// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: coupon/v1/coupon.proto

package campaignv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateCampaignRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	AvailableCoupons int32                  `protobuf:"varint,1,opt,name=available_coupons,json=availableCoupons,proto3" json:"available_coupons,omitempty"`
	StartTime        int64                  `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateCampaignRequest) Reset() {
	*x = CreateCampaignRequest{}
	mi := &file_coupon_v1_coupon_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCampaignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCampaignRequest) ProtoMessage() {}

func (x *CreateCampaignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coupon_v1_coupon_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCampaignRequest.ProtoReflect.Descriptor instead.
func (*CreateCampaignRequest) Descriptor() ([]byte, []int) {
	return file_coupon_v1_coupon_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCampaignRequest) GetAvailableCoupons() int32 {
	if x != nil {
		return x.AvailableCoupons
	}
	return 0
}

func (x *CreateCampaignRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

type CreateCampaignResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Campaign      *Campaign              `protobuf:"bytes,1,opt,name=campaign,proto3" json:"campaign,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCampaignResponse) Reset() {
	*x = CreateCampaignResponse{}
	mi := &file_coupon_v1_coupon_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCampaignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCampaignResponse) ProtoMessage() {}

func (x *CreateCampaignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coupon_v1_coupon_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCampaignResponse.ProtoReflect.Descriptor instead.
func (*CreateCampaignResponse) Descriptor() ([]byte, []int) {
	return file_coupon_v1_coupon_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCampaignResponse) GetCampaign() *Campaign {
	if x != nil {
		return x.Campaign
	}
	return nil
}

type GetCampaignRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CampaignId    string                 `protobuf:"bytes,1,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCampaignRequest) Reset() {
	*x = GetCampaignRequest{}
	mi := &file_coupon_v1_coupon_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCampaignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCampaignRequest) ProtoMessage() {}

func (x *GetCampaignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coupon_v1_coupon_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCampaignRequest.ProtoReflect.Descriptor instead.
func (*GetCampaignRequest) Descriptor() ([]byte, []int) {
	return file_coupon_v1_coupon_proto_rawDescGZIP(), []int{2}
}

func (x *GetCampaignRequest) GetCampaignId() string {
	if x != nil {
		return x.CampaignId
	}
	return ""
}

type GetCampaignResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Campaign          *Campaign              `protobuf:"bytes,1,opt,name=campaign,proto3" json:"campaign,omitempty"`
	IssuedCouponCodes []string               `protobuf:"bytes,2,rep,name=issued_coupon_codes,json=issuedCouponCodes,proto3" json:"issued_coupon_codes,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetCampaignResponse) Reset() {
	*x = GetCampaignResponse{}
	mi := &file_coupon_v1_coupon_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCampaignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCampaignResponse) ProtoMessage() {}

func (x *GetCampaignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coupon_v1_coupon_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCampaignResponse.ProtoReflect.Descriptor instead.
func (*GetCampaignResponse) Descriptor() ([]byte, []int) {
	return file_coupon_v1_coupon_proto_rawDescGZIP(), []int{3}
}

func (x *GetCampaignResponse) GetCampaign() *Campaign {
	if x != nil {
		return x.Campaign
	}
	return nil
}

func (x *GetCampaignResponse) GetIssuedCouponCodes() []string {
	if x != nil {
		return x.IssuedCouponCodes
	}
	return nil
}

type IssueCouponRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CampaignId    string                 `protobuf:"bytes,1,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IssueCouponRequest) Reset() {
	*x = IssueCouponRequest{}
	mi := &file_coupon_v1_coupon_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueCouponRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCouponRequest) ProtoMessage() {}

func (x *IssueCouponRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coupon_v1_coupon_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCouponRequest.ProtoReflect.Descriptor instead.
func (*IssueCouponRequest) Descriptor() ([]byte, []int) {
	return file_coupon_v1_coupon_proto_rawDescGZIP(), []int{4}
}

func (x *IssueCouponRequest) GetCampaignId() string {
	if x != nil {
		return x.CampaignId
	}
	return ""
}

type IssueCouponResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Coupon        *Coupon                `protobuf:"bytes,1,opt,name=coupon,proto3" json:"coupon,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IssueCouponResponse) Reset() {
	*x = IssueCouponResponse{}
	mi := &file_coupon_v1_coupon_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueCouponResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCouponResponse) ProtoMessage() {}

func (x *IssueCouponResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coupon_v1_coupon_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCouponResponse.ProtoReflect.Descriptor instead.
func (*IssueCouponResponse) Descriptor() ([]byte, []int) {
	return file_coupon_v1_coupon_proto_rawDescGZIP(), []int{5}
}

func (x *IssueCouponResponse) GetCoupon() *Coupon {
	if x != nil {
		return x.Coupon
	}
	return nil
}

type Campaign struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AvailableCoupons int32                  `protobuf:"varint,2,opt,name=available_coupons,json=availableCoupons,proto3" json:"available_coupons,omitempty"`
	StartTime        int64                  `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Campaign) Reset() {
	*x = Campaign{}
	mi := &file_coupon_v1_coupon_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Campaign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Campaign) ProtoMessage() {}

func (x *Campaign) ProtoReflect() protoreflect.Message {
	mi := &file_coupon_v1_coupon_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Campaign.ProtoReflect.Descriptor instead.
func (*Campaign) Descriptor() ([]byte, []int) {
	return file_coupon_v1_coupon_proto_rawDescGZIP(), []int{6}
}

func (x *Campaign) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Campaign) GetAvailableCoupons() int32 {
	if x != nil {
		return x.AvailableCoupons
	}
	return 0
}

func (x *Campaign) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

type Coupon struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CampaignId    string                 `protobuf:"bytes,2,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	CouponCode    string                 `protobuf:"bytes,3,opt,name=coupon_code,json=couponCode,proto3" json:"coupon_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Coupon) Reset() {
	*x = Coupon{}
	mi := &file_coupon_v1_coupon_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Coupon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coupon) ProtoMessage() {}

func (x *Coupon) ProtoReflect() protoreflect.Message {
	mi := &file_coupon_v1_coupon_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coupon.ProtoReflect.Descriptor instead.
func (*Coupon) Descriptor() ([]byte, []int) {
	return file_coupon_v1_coupon_proto_rawDescGZIP(), []int{7}
}

func (x *Coupon) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Coupon) GetCampaignId() string {
	if x != nil {
		return x.CampaignId
	}
	return ""
}

func (x *Coupon) GetCouponCode() string {
	if x != nil {
		return x.CouponCode
	}
	return ""
}

var File_coupon_v1_coupon_proto protoreflect.FileDescriptor

var file_coupon_v1_coupon_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x63, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x08, 0x63,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x22, 0x78,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x08,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x12, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x22,
	0x42, 0x0a, 0x13, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x22, 0x66, 0x0a, 0x08, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2b, 0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x06, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x96, 0x02, 0x0a, 0x0f, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x22, 0x2e,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0b,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x63, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x27, 0x5a, 0x25, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_coupon_v1_coupon_proto_rawDescOnce sync.Once
	file_coupon_v1_coupon_proto_rawDescData []byte
)

func file_coupon_v1_coupon_proto_rawDescGZIP() []byte {
	file_coupon_v1_coupon_proto_rawDescOnce.Do(func() {
		file_coupon_v1_coupon_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_coupon_v1_coupon_proto_rawDesc), len(file_coupon_v1_coupon_proto_rawDesc)))
	})
	return file_coupon_v1_coupon_proto_rawDescData
}

var file_coupon_v1_coupon_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_coupon_v1_coupon_proto_goTypes = []any{
	(*CreateCampaignRequest)(nil),  // 0: campaign.v1.CreateCampaignRequest
	(*CreateCampaignResponse)(nil), // 1: campaign.v1.CreateCampaignResponse
	(*GetCampaignRequest)(nil),     // 2: campaign.v1.GetCampaignRequest
	(*GetCampaignResponse)(nil),    // 3: campaign.v1.GetCampaignResponse
	(*IssueCouponRequest)(nil),     // 4: campaign.v1.IssueCouponRequest
	(*IssueCouponResponse)(nil),    // 5: campaign.v1.IssueCouponResponse
	(*Campaign)(nil),               // 6: campaign.v1.Campaign
	(*Coupon)(nil),                 // 7: campaign.v1.Coupon
}
var file_coupon_v1_coupon_proto_depIdxs = []int32{
	6, // 0: campaign.v1.CreateCampaignResponse.campaign:type_name -> campaign.v1.Campaign
	6, // 1: campaign.v1.GetCampaignResponse.campaign:type_name -> campaign.v1.Campaign
	7, // 2: campaign.v1.IssueCouponResponse.coupon:type_name -> campaign.v1.Coupon
	0, // 3: campaign.v1.CampaignService.CreateCampaign:input_type -> campaign.v1.CreateCampaignRequest
	2, // 4: campaign.v1.CampaignService.GetCampaign:input_type -> campaign.v1.GetCampaignRequest
	4, // 5: campaign.v1.CampaignService.IssueCoupon:input_type -> campaign.v1.IssueCouponRequest
	1, // 6: campaign.v1.CampaignService.CreateCampaign:output_type -> campaign.v1.CreateCampaignResponse
	3, // 7: campaign.v1.CampaignService.GetCampaign:output_type -> campaign.v1.GetCampaignResponse
	5, // 8: campaign.v1.CampaignService.IssueCoupon:output_type -> campaign.v1.IssueCouponResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_coupon_v1_coupon_proto_init() }
func file_coupon_v1_coupon_proto_init() {
	if File_coupon_v1_coupon_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_coupon_v1_coupon_proto_rawDesc), len(file_coupon_v1_coupon_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coupon_v1_coupon_proto_goTypes,
		DependencyIndexes: file_coupon_v1_coupon_proto_depIdxs,
		MessageInfos:      file_coupon_v1_coupon_proto_msgTypes,
	}.Build()
	File_coupon_v1_coupon_proto = out.File
	file_coupon_v1_coupon_proto_goTypes = nil
	file_coupon_v1_coupon_proto_depIdxs = nil
}
