// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: accounting/exchange/messages.proto

package exchange

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Exchange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId            string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	MakerAsset      string `protobuf:"bytes,2,opt,name=maker_asset,json=makerAsset,proto3" json:"maker_asset,omitempty"`
	TakerAsset      string `protobuf:"bytes,3,opt,name=taker_asset,json=takerAsset,proto3" json:"taker_asset,omitempty"`
	MakerAmount     string `protobuf:"bytes,4,opt,name=maker_amount,json=makerAmount,proto3" json:"maker_amount,omitempty"`
	TakerAmount     string `protobuf:"bytes,5,opt,name=taker_amount,json=takerAmount,proto3" json:"taker_amount,omitempty"`
	MakerWarehouse  string `protobuf:"bytes,6,opt,name=maker_warehouse,json=makerWarehouse,proto3" json:"maker_warehouse,omitempty"`
	TakerWarehouse  string `protobuf:"bytes,7,opt,name=taker_warehouse,json=takerWarehouse,proto3" json:"taker_warehouse,omitempty"`
	Broker          string `protobuf:"bytes,8,opt,name=broker,proto3" json:"broker,omitempty"`
	BrokerWarehouse string `protobuf:"bytes,9,opt,name=broker_warehouse,json=brokerWarehouse,proto3" json:"broker_warehouse,omitempty"`
	ReferralAsset   string `protobuf:"bytes,10,opt,name=referral_asset,json=referralAsset,proto3" json:"referral_asset,omitempty"`
	ReferralAmount  string `protobuf:"bytes,11,opt,name=referral_amount,json=referralAmount,proto3" json:"referral_amount,omitempty"`
	Maker           string `protobuf:"bytes,12,opt,name=maker,proto3" json:"maker,omitempty"`
	Taker           string `protobuf:"bytes,13,opt,name=taker,proto3" json:"taker,omitempty"`
}

func (x *Exchange) Reset() {
	*x = Exchange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_exchange_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exchange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exchange) ProtoMessage() {}

func (x *Exchange) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_exchange_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exchange.ProtoReflect.Descriptor instead.
func (*Exchange) Descriptor() ([]byte, []int) {
	return file_accounting_exchange_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Exchange) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *Exchange) GetMakerAsset() string {
	if x != nil {
		return x.MakerAsset
	}
	return ""
}

func (x *Exchange) GetTakerAsset() string {
	if x != nil {
		return x.TakerAsset
	}
	return ""
}

func (x *Exchange) GetMakerAmount() string {
	if x != nil {
		return x.MakerAmount
	}
	return ""
}

func (x *Exchange) GetTakerAmount() string {
	if x != nil {
		return x.TakerAmount
	}
	return ""
}

func (x *Exchange) GetMakerWarehouse() string {
	if x != nil {
		return x.MakerWarehouse
	}
	return ""
}

func (x *Exchange) GetTakerWarehouse() string {
	if x != nil {
		return x.TakerWarehouse
	}
	return ""
}

func (x *Exchange) GetBroker() string {
	if x != nil {
		return x.Broker
	}
	return ""
}

func (x *Exchange) GetBrokerWarehouse() string {
	if x != nil {
		return x.BrokerWarehouse
	}
	return ""
}

func (x *Exchange) GetReferralAsset() string {
	if x != nil {
		return x.ReferralAsset
	}
	return ""
}

func (x *Exchange) GetReferralAmount() string {
	if x != nil {
		return x.ReferralAmount
	}
	return ""
}

func (x *Exchange) GetMaker() string {
	if x != nil {
		return x.Maker
	}
	return ""
}

func (x *Exchange) GetTaker() string {
	if x != nil {
		return x.Taker
	}
	return ""
}

type Referral struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId            string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	BrokerWarehouse string `protobuf:"bytes,2,opt,name=broker_warehouse,json=brokerWarehouse,proto3" json:"broker_warehouse,omitempty"`
	ReferralAsset   string `protobuf:"bytes,3,opt,name=referral_asset,json=referralAsset,proto3" json:"referral_asset,omitempty"`
	ReferralAmount  string `protobuf:"bytes,4,opt,name=referral_amount,json=referralAmount,proto3" json:"referral_amount,omitempty"`
	UserWarehouse   string `protobuf:"bytes,5,opt,name=user_warehouse,json=userWarehouse,proto3" json:"user_warehouse,omitempty"`
	Broker          string `protobuf:"bytes,6,opt,name=broker,proto3" json:"broker,omitempty"`
	User            string `protobuf:"bytes,7,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Referral) Reset() {
	*x = Referral{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_exchange_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Referral) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Referral) ProtoMessage() {}

func (x *Referral) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_exchange_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Referral.ProtoReflect.Descriptor instead.
func (*Referral) Descriptor() ([]byte, []int) {
	return file_accounting_exchange_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Referral) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *Referral) GetBrokerWarehouse() string {
	if x != nil {
		return x.BrokerWarehouse
	}
	return ""
}

func (x *Referral) GetReferralAsset() string {
	if x != nil {
		return x.ReferralAsset
	}
	return ""
}

func (x *Referral) GetReferralAmount() string {
	if x != nil {
		return x.ReferralAmount
	}
	return ""
}

func (x *Referral) GetUserWarehouse() string {
	if x != nil {
		return x.UserWarehouse
	}
	return ""
}

func (x *Referral) GetBroker() string {
	if x != nil {
		return x.Broker
	}
	return ""
}

func (x *Referral) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type ExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExchangeSingle *Exchange      `protobuf:"bytes,1,opt,name=exchange_single,json=exchangeSingle,proto3" json:"exchange_single,omitempty"`
	Action         ExchangeAction `protobuf:"varint,2,opt,name=action,proto3,enum=exchange.ExchangeAction" json:"action,omitempty"`
}

func (x *ExchangeRequest) Reset() {
	*x = ExchangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_exchange_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRequest) ProtoMessage() {}

func (x *ExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_exchange_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRequest.ProtoReflect.Descriptor instead.
func (*ExchangeRequest) Descriptor() ([]byte, []int) {
	return file_accounting_exchange_messages_proto_rawDescGZIP(), []int{2}
}

func (x *ExchangeRequest) GetExchangeSingle() *Exchange {
	if x != nil {
		return x.ExchangeSingle
	}
	return nil
}

func (x *ExchangeRequest) GetAction() ExchangeAction {
	if x != nil {
		return x.Action
	}
	return ExchangeAction_EXCHANGE_ACTION_UNSPECIFIED
}

type ExchangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExchangeResponse) Reset() {
	*x = ExchangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_exchange_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeResponse) ProtoMessage() {}

func (x *ExchangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_exchange_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeResponse.ProtoReflect.Descriptor instead.
func (*ExchangeResponse) Descriptor() ([]byte, []int) {
	return file_accounting_exchange_messages_proto_rawDescGZIP(), []int{3}
}

type ReferralRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferralMultiple []*Referral    `protobuf:"bytes,1,rep,name=referral_multiple,json=referralMultiple,proto3" json:"referral_multiple,omitempty"`
	Action           ReferralAction `protobuf:"varint,2,opt,name=action,proto3,enum=exchange.ReferralAction" json:"action,omitempty"`
}

func (x *ReferralRequest) Reset() {
	*x = ReferralRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_exchange_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferralRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferralRequest) ProtoMessage() {}

func (x *ReferralRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_exchange_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferralRequest.ProtoReflect.Descriptor instead.
func (*ReferralRequest) Descriptor() ([]byte, []int) {
	return file_accounting_exchange_messages_proto_rawDescGZIP(), []int{4}
}

func (x *ReferralRequest) GetReferralMultiple() []*Referral {
	if x != nil {
		return x.ReferralMultiple
	}
	return nil
}

func (x *ReferralRequest) GetAction() ReferralAction {
	if x != nil {
		return x.Action
	}
	return ReferralAction_REFERRAL_ACTION_UNSPECIFIED
}

type ReferralResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReferralResponse) Reset() {
	*x = ReferralResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_exchange_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferralResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferralResponse) ProtoMessage() {}

func (x *ReferralResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_exchange_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferralResponse.ProtoReflect.Descriptor instead.
func (*ReferralResponse) Descriptor() ([]byte, []int) {
	return file_accounting_exchange_messages_proto_rawDescGZIP(), []int{5}
}

var File_accounting_exchange_messages_proto protoreflect.FileDescriptor

var file_accounting_exchange_messages_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x1f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb8, 0x03, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x13, 0x0a, 0x05,
	0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x6b, 0x65, 0x72,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61,
	0x6b, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x6b,
	0x65, 0x72, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61, 0x6b,
	0x65, 0x72, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61,
	0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x61, 0x6b, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x6b, 0x65, 0x72, 0x22, 0xed, 0x01, 0x0a, 0x08, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x57, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x61, 0x6c, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61,
	0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x75, 0x73, 0x65, 0x72, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x80, 0x01, 0x0a, 0x0f, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b,
	0x0a, 0x0f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0e, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x12, 0x0a,
	0x10, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x84, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x11, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61,
	0x6c, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x72, 0x61, 0x6c, 0x52, 0x10, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29,
	0x72, 0x61, 0x62, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x3b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_accounting_exchange_messages_proto_rawDescOnce sync.Once
	file_accounting_exchange_messages_proto_rawDescData = file_accounting_exchange_messages_proto_rawDesc
)

func file_accounting_exchange_messages_proto_rawDescGZIP() []byte {
	file_accounting_exchange_messages_proto_rawDescOnce.Do(func() {
		file_accounting_exchange_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounting_exchange_messages_proto_rawDescData)
	})
	return file_accounting_exchange_messages_proto_rawDescData
}

var file_accounting_exchange_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_accounting_exchange_messages_proto_goTypes = []interface{}{
	(*Exchange)(nil),         // 0: exchange.Exchange
	(*Referral)(nil),         // 1: exchange.Referral
	(*ExchangeRequest)(nil),  // 2: exchange.ExchangeRequest
	(*ExchangeResponse)(nil), // 3: exchange.ExchangeResponse
	(*ReferralRequest)(nil),  // 4: exchange.ReferralRequest
	(*ReferralResponse)(nil), // 5: exchange.ReferralResponse
	(ExchangeAction)(0),      // 6: exchange.ExchangeAction
	(ReferralAction)(0),      // 7: exchange.ReferralAction
}
var file_accounting_exchange_messages_proto_depIdxs = []int32{
	0, // 0: exchange.ExchangeRequest.exchange_single:type_name -> exchange.Exchange
	6, // 1: exchange.ExchangeRequest.action:type_name -> exchange.ExchangeAction
	1, // 2: exchange.ReferralRequest.referral_multiple:type_name -> exchange.Referral
	7, // 3: exchange.ReferralRequest.action:type_name -> exchange.ReferralAction
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_accounting_exchange_messages_proto_init() }
func file_accounting_exchange_messages_proto_init() {
	if File_accounting_exchange_messages_proto != nil {
		return
	}
	file_accounting_exchange_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_accounting_exchange_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exchange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_accounting_exchange_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Referral); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_accounting_exchange_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_accounting_exchange_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_accounting_exchange_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferralRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_accounting_exchange_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferralResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_accounting_exchange_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accounting_exchange_messages_proto_goTypes,
		DependencyIndexes: file_accounting_exchange_messages_proto_depIdxs,
		MessageInfos:      file_accounting_exchange_messages_proto_msgTypes,
	}.Build()
	File_accounting_exchange_messages_proto = out.File
	file_accounting_exchange_messages_proto_rawDesc = nil
	file_accounting_exchange_messages_proto_goTypes = nil
	file_accounting_exchange_messages_proto_depIdxs = nil
}
