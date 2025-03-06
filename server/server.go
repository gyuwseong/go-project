package server

import (
	"context"

	campaignv1 "go-project/gen/campaign/v1"
	"go-project/internal/campaign"

	"connectrpc.com/connect"
)

type CampaignServer struct {
	campaignService *campaign.CampaignService
}

func NewCampaignServer() *CampaignServer {
	return &CampaignServer{
		campaignService: campaign.NewCampaignService(),
	}
}

func (server *CampaignServer) CreateCampaign(
	ctx context.Context,
	req *connect.Request[campaignv1.CreateCampaignRequest],
) (*connect.Response[campaignv1.CreateCampaignResponse], error) {
	campaign, err := server.campaignService.CreateCampaign(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&campaignv1.CreateCampaignResponse{
		Campaign: campaign,
	}), nil
}

func (server *CampaignServer) GetCampaign(
	ctx context.Context,
	req *connect.Request[campaignv1.GetCampaignRequest],
) (*connect.Response[campaignv1.GetCampaignResponse], error) {
	campaign, issuedCouponCodes, err := server.campaignService.GetCampaign(ctx, req.Msg.CampaignId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	response := &campaignv1.GetCampaignResponse{
		Campaign:          campaign,
		IssuedCouponCodes: issuedCouponCodes,
	}
	return connect.NewResponse(response), nil
}

func (server *CampaignServer) IssueCoupon(
	ctx context.Context,
	req *connect.Request[campaignv1.IssueCouponRequest],
) (*connect.Response[campaignv1.IssueCouponResponse], error) {
	coupon, err := server.campaignService.IssueCoupon(ctx, req.Msg.CampaignId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&campaignv1.IssueCouponResponse{
		Coupon: coupon,
	}), nil
}
