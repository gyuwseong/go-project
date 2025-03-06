package campaign

import (
	"context"
	"errors"
	"fmt"
	"time"

	campaignv1 "go-project/gen/campaign/v1"
	"go-project/internal/coupon"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/google/uuid"
)

type CampaignService struct {
	couponGenerator *coupon.CouponGenerator
	redis           *redis.Client
	rs              *redsync.Redsync
}

func NewCampaignService() *CampaignService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0, 
	})

	pool := goredis.NewPool(rdb)
	rs := redsync.New(pool)

	return &CampaignService{
		couponGenerator: coupon.NewCouponGenerator(),
		redis:           rdb,
		rs:              rs,
	}
}

func (service *CampaignService) CreateCampaign(ctx context.Context, req *campaignv1.CreateCampaignRequest) (*campaignv1.Campaign, error) {
	campaign := &campaignv1.Campaign{
		Id:               uuid.New().String(),
		AvailableCoupons: req.AvailableCoupons,
		StartTime:        req.StartTime,
	}

	key := "campaign:" + campaign.Id
	service.redis.HSet(ctx, key,
		"id", campaign.Id,
		"available_coupons", campaign.AvailableCoupons,
		"start_time", campaign.StartTime,
	)

	return campaign, nil
}

func (service *CampaignService) GetCampaign(ctx context.Context, campaignID string) (*campaignv1.Campaign, []string, error) {
	key := "campaign:" + campaignID
	values := service.redis.HGetAll(ctx, key).Val()
	if len(values) == 0 {
		return nil, nil, errors.New("campaign not found")
	}

	campaign := &campaignv1.Campaign{
		Id:               values["id"],
		AvailableCoupons: parseInt32(values["available_coupons"]),
		StartTime:        parseInt64(values["start_time"]),
	}

	codes := service.redis.SMembers(ctx, "campaign:"+campaignID+":codes").Val()

	return campaign, codes, nil
}

func (service *CampaignService) IssueCoupon(ctx context.Context, campaignID string) (*campaignv1.Coupon, error) {
	mutex := service.rs.NewMutex("campaign:"+campaignID+":lock", redsync.WithExpiry(time.Second*10))
	if err := mutex.Lock(); err != nil {
		return nil, errors.New("failed to acquire lock")
	}
	defer mutex.Unlock()

	key := "campaign:" + campaignID
	values := service.redis.HGetAll(ctx, key).Val()
	if len(values) == 0 {
		return nil, errors.New("campaign not found")
	}

	campaign := &campaignv1.Campaign{
		Id:               values["id"],
		AvailableCoupons: parseInt32(values["available_coupons"]),
		StartTime:        parseInt64(values["start_time"]),
	}

	now := time.Now().Unix()
	if now < campaign.StartTime {
		return nil, errors.New("campaign has not started yet")
	}

	available := service.redis.HIncrBy(ctx, key, "available_coupons", -1).Val()
	if available < 0 {
		service.redis.HIncrBy(ctx, key, "available_coupons", 1)
		return nil, errors.New("no available coupons")
	}

	couponCode := service.couponGenerator.GenerateCode()
	exists := service.redis.SIsMember(ctx, "campaign:"+campaignID+":codes", couponCode).Val()
	if exists {
		service.redis.HIncrBy(ctx, key, "available_coupons", 1)
		return nil, errors.New("duplicate coupon code")
	}

	service.redis.SAdd(ctx, "campaign:"+campaignID+":codes", couponCode)

	newCoupon := &campaignv1.Coupon{
		Id:          uuid.New().String(),
		CampaignId:  campaignID,
		CouponCode:  couponCode,
	}

	couponKey := "coupon:" + newCoupon.Id
	service.redis.HSet(ctx, couponKey,
		"id", newCoupon.Id,
		"campaign_id", newCoupon.CampaignId,
		"coupon_code", newCoupon.CouponCode,
	)

	return newCoupon, nil
}

func parseInt32(s string) int32 {
	var v int32
	fmt.Sscanf(s, "%d", &v)
	return v
}

func parseInt64(s string) int64 {
	var v int64
	fmt.Sscanf(s, "%d", &v)
	return v
}
