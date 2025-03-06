package campaign

import (
	"context"
	"fmt"
	"testing"
	"time"

	campaignv1 "go-project/gen/campaign/v1"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestCreateCampaign(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	ctx := context.Background()

	rdb.FlushAll(ctx)
	service := NewCampaignService()

	t.Run("[success] create campaign", func(t *testing.T) {
		startTime := time.Now().Add(-time.Hour).Unix()
		req := &campaignv1.CreateCampaignRequest{
			AvailableCoupons: 1000,
			StartTime:        startTime,
		}

		campaign, err := service.CreateCampaign(ctx, req)
		assert.NoError(t, err)
		assert.NotEmpty(t, campaign.Id)
		assert.Equal(t, req.AvailableCoupons, campaign.AvailableCoupons)
		assert.Equal(t, req.StartTime, campaign.StartTime)

		key := "campaign:" + campaign.Id
		values := rdb.HGetAll(ctx, key).Val()
		assert.Equal(t, campaign.Id, values["id"])
		assert.Equal(t, "1000", values["available_coupons"])
		assert.Equal(t, fmt.Sprintf("%d", startTime), values["start_time"])
	})
}

func TestGetCampaign(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	ctx := context.Background()

	rdb.FlushAll(ctx)
	service := NewCampaignService()

	t.Run("[success] get campaign", func(t *testing.T) {
		campaign := &campaignv1.Campaign{
			Id:               "campaign1",
			AvailableCoupons: 1000,
			StartTime:        time.Now().Add(-time.Hour).Unix(),
		}

		key := "campaign:" + campaign.Id
		rdb.HSet(ctx, key,
			"id", campaign.Id,
			"available_coupons", campaign.AvailableCoupons,
			"start_time", campaign.StartTime,
		)

		rdb.SAdd(ctx, "campaign:"+campaign.Id+":codes", "coupon1", "coupon2")

		result, codes, err := service.GetCampaign(ctx, campaign.Id)
		assert.NoError(t, err)
		assert.Equal(t, campaign.Id, result.Id)
		assert.Equal(t, campaign.AvailableCoupons, result.AvailableCoupons)
		assert.Equal(t, campaign.StartTime, result.StartTime)
		assert.ElementsMatch(t, []string{"coupon1", "coupon2"}, codes)
	})

	t.Run("[error] campaign not found", func(t *testing.T) {
		result, codes, err := service.GetCampaign(ctx, "campaign not found")
		assert.Error(t, err)
		assert.Equal(t, "campaign not found", err.Error())
		assert.Nil(t, result)
		assert.Nil(t, codes)
	})
}

func TestIssueCoupon(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	ctx := context.Background()

	rdb.FlushAll(ctx)
	service := NewCampaignService()

	t.Run("[success] issue coupon", func(t *testing.T) {
		campaign := &campaignv1.Campaign{
			Id:               "campaign1",
			AvailableCoupons: 1000,
			StartTime:        time.Now().Add(-time.Hour).Unix(),
		}

		key := "campaign:" + campaign.Id
		rdb.HSet(ctx, key,
			"id", campaign.Id,
			"available_coupons", campaign.AvailableCoupons,
			"start_time", campaign.StartTime,
		)

		coupon, err := service.IssueCoupon(ctx, campaign.Id)
		assert.NoError(t, err)
		assert.NotEmpty(t, coupon.Id)
		assert.Equal(t, campaign.Id, coupon.CampaignId)
		assert.NotEmpty(t, coupon.CouponCode)

		values := rdb.HGetAll(ctx, key).Val()
		assert.Equal(t, "999", values["available_coupons"])

		codes := rdb.SMembers(ctx, "campaign:"+campaign.Id+":codes").Val()
		assert.Contains(t, codes, coupon.CouponCode)

		couponValues := rdb.HGetAll(ctx, "coupon:"+coupon.Id).Val()
		assert.Equal(t, coupon.Id, couponValues["id"])
		assert.Equal(t, coupon.CampaignId, couponValues["campaign_id"])
		assert.Equal(t, coupon.CouponCode, couponValues["coupon_code"])
	})

	t.Run("[error] campaign not found", func(t *testing.T) {
		coupon, err := service.IssueCoupon(ctx, "campaign2")
		assert.Error(t, err)
		assert.Equal(t, "campaign not found", err.Error())
		assert.Nil(t, coupon)
	})

	t.Run("[error] campaign not started", func(t *testing.T) {
		campaign := &campaignv1.Campaign{
			Id:               "campaign1",
			AvailableCoupons: 1000,
			StartTime:        time.Now().Add(time.Hour).Unix(),
		}

		key := "campaign:" + campaign.Id
		rdb.HSet(ctx, key,
			"id", campaign.Id,
			"available_coupons", campaign.AvailableCoupons,
			"start_time", campaign.StartTime,
		)

		coupon, err := service.IssueCoupon(ctx, campaign.Id)
		assert.Error(t, err)
		assert.Equal(t, "campaign has not started yet", err.Error())
		assert.Nil(t, coupon)
	})

	t.Run("[error] no available coupons", func(t *testing.T) {
		campaign := &campaignv1.Campaign{
			Id:               "campaign1",
			AvailableCoupons: 0,
			StartTime:        time.Now().Add(-time.Hour).Unix(),
		}

		key := "campaign:" + campaign.Id
		rdb.HSet(ctx, key,
			"id", campaign.Id,
			"available_coupons", campaign.AvailableCoupons,
			"start_time", campaign.StartTime,
		)

		coupon, err := service.IssueCoupon(ctx, campaign.Id)
		assert.Error(t, err)
		assert.Equal(t, "no available coupons", err.Error())
		assert.Nil(t, coupon)
	})
}

func TestIssueMultipleCoupons(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	ctx := context.Background()

	rdb.FlushAll(ctx)
	service := NewCampaignService()

	t.Run("[success] issue 1000 coupons", func(t *testing.T) {
		campaign := &campaignv1.Campaign{
			Id:               "campaign_1000",
			AvailableCoupons: 1000,
			StartTime:        time.Now().Add(-time.Hour).Unix(),
		}

		key := "campaign:" + campaign.Id
		rdb.HSet(ctx, key,
			"id", campaign.Id,
			"available_coupons", campaign.AvailableCoupons,
			"start_time", campaign.StartTime,
		)

		coupons := make([]*campaignv1.Coupon, 1000)
		for i := 0; i < 1000; i++ {
			coupon, err := service.IssueCoupon(ctx, campaign.Id)
			assert.NoError(t, err)
			coupons[i] = coupon
		}

		codes := make(map[string]bool)
		for _, coupon := range coupons {
			assert.False(t, codes[coupon.CouponCode], "duplicate coupon code", coupon.CouponCode)
			codes[coupon.CouponCode] = true
		}

		values := rdb.HGetAll(ctx, key).Val()
		assert.Equal(t, "0", values["available_coupons"])
	})

	t.Run("[success] issue 100 coupons", func(t *testing.T) {
		campaign := &campaignv1.Campaign{
			Id:               "campaign_100",
			AvailableCoupons: 1000,
			StartTime:        time.Now().Add(-time.Hour).Unix(),
		}

		key := "campaign:" + campaign.Id
		rdb.HSet(ctx, key,
			"id", campaign.Id,
			"available_coupons", campaign.AvailableCoupons,
			"start_time", campaign.StartTime,
		)

		coupons := make([]*campaignv1.Coupon, 100)
		for i := 0; i < 100; i++ {
			coupon, err := service.IssueCoupon(ctx, campaign.Id)
			assert.NoError(t, err)
			coupons[i] = coupon
		}

		codes := make(map[string]bool)
		for _, coupon := range coupons {
			assert.False(t, codes[coupon.CouponCode], "duplicate coupon code", coupon.CouponCode)
			codes[coupon.CouponCode] = true
		}

		values := rdb.HGetAll(ctx, key).Val()
		assert.Equal(t, "900", values["available_coupons"])
	})

	t.Run("[error] issue over 1000 coupons", func(t *testing.T) {
		campaign := &campaignv1.Campaign{
			Id:               "campaign_over",
			AvailableCoupons: 1000,
			StartTime:        time.Now().Add(-time.Hour).Unix(),
		}

		key := "campaign:" + campaign.Id
		rdb.HSet(ctx, key,
			"id", campaign.Id,
			"available_coupons", campaign.AvailableCoupons,
			"start_time", campaign.StartTime,
		)

		coupons := make([]*campaignv1.Coupon, 1000)
		for i := 0; i < 1000; i++ {
			coupon, err := service.IssueCoupon(ctx, campaign.Id)
			assert.NoError(t, err)
			coupons[i] = coupon
		}

		codes := make(map[string]bool)
		for _, coupon := range coupons {
			assert.False(t, codes[coupon.CouponCode], "duplicate coupon code", coupon.CouponCode)
			codes[coupon.CouponCode] = true
		}

		values := rdb.HGetAll(ctx, key).Val()
		assert.Equal(t, "0", values["available_coupons"])

		coupon, err := service.IssueCoupon(ctx, campaign.Id)
		assert.Error(t, err)
		assert.Equal(t, "no available coupons", err.Error())
		assert.Nil(t, coupon)
	})
}
