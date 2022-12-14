// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: accounting/treasury-statement/messages.proto

package tStatement

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TreasuryStatement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TreasuryId   string                 `protobuf:"bytes,1,opt,name=treasury_id,json=treasuryId,proto3" json:"treasury_id,omitempty"`
	Broker       string                 `protobuf:"bytes,2,opt,name=broker,proto3" json:"broker,omitempty"`
	AssetNetwork string                 `protobuf:"bytes,3,opt,name=asset_network,json=assetNetwork,proto3" json:"asset_network,omitempty"`
	LastUpdate   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	Balance      string                 `protobuf:"bytes,5,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *TreasuryStatement) Reset() {
	*x = TreasuryStatement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_treasury_statement_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasuryStatement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasuryStatement) ProtoMessage() {}

func (x *TreasuryStatement) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_treasury_statement_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasuryStatement.ProtoReflect.Descriptor instead.
func (*TreasuryStatement) Descriptor() ([]byte, []int) {
	return file_accounting_treasury_statement_messages_proto_rawDescGZIP(), []int{0}
}

func (x *TreasuryStatement) GetTreasuryId() string {
	if x != nil {
		return x.TreasuryId
	}
	return ""
}

func (x *TreasuryStatement) GetBroker() string {
	if x != nil {
		return x.Broker
	}
	return ""
}

func (x *TreasuryStatement) GetAssetNetwork() string {
	if x != nil {
		return x.AssetNetwork
	}
	return ""
}

func (x *TreasuryStatement) GetLastUpdate() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdate
	}
	return nil
}

func (x *TreasuryStatement) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

type GetBalanceAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TreasuryId string `protobuf:"bytes,1,opt,name=treasury_id,json=treasuryId,proto3" json:"treasury_id,omitempty"`
	Page       uint32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`                            // for pagination - determines what page
	PageLimit  uint32 `protobuf:"varint,3,opt,name=page_limit,json=pageLimit,proto3" json:"page_limit,omitempty"` // for pagination - determines number of records in each page
}

func (x *GetBalanceAllRequest) Reset() {
	*x = GetBalanceAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_treasury_statement_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceAllRequest) ProtoMessage() {}

func (x *GetBalanceAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_treasury_statement_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceAllRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceAllRequest) Descriptor() ([]byte, []int) {
	return file_accounting_treasury_statement_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GetBalanceAllRequest) GetTreasuryId() string {
	if x != nil {
		return x.TreasuryId
	}
	return ""
}

func (x *GetBalanceAllRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetBalanceAllRequest) GetPageLimit() uint32 {
	if x != nil {
		return x.PageLimit
	}
	return 0
}

type GetBalanceAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*TreasuryStatement `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetBalanceAllResponse) Reset() {
	*x = GetBalanceAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_treasury_statement_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceAllResponse) ProtoMessage() {}

func (x *GetBalanceAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_treasury_statement_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceAllResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceAllResponse) Descriptor() ([]byte, []int) {
	return file_accounting_treasury_statement_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetBalanceAllResponse) GetData() []*TreasuryStatement {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetBalanceByAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TreasuryId   string `protobuf:"bytes,1,opt,name=treasury_id,json=treasuryId,proto3" json:"treasury_id,omitempty"`
	AssetNetwork string `protobuf:"bytes,2,opt,name=asset_network,json=assetNetwork,proto3" json:"asset_network,omitempty"`
}

func (x *GetBalanceByAssetRequest) Reset() {
	*x = GetBalanceByAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_treasury_statement_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceByAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceByAssetRequest) ProtoMessage() {}

func (x *GetBalanceByAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_treasury_statement_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceByAssetRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceByAssetRequest) Descriptor() ([]byte, []int) {
	return file_accounting_treasury_statement_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GetBalanceByAssetRequest) GetTreasuryId() string {
	if x != nil {
		return x.TreasuryId
	}
	return ""
}

func (x *GetBalanceByAssetRequest) GetAssetNetwork() string {
	if x != nil {
		return x.AssetNetwork
	}
	return ""
}

type GetBalanceByAssetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*TreasuryStatement `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetBalanceByAssetResponse) Reset() {
	*x = GetBalanceByAssetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_treasury_statement_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceByAssetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceByAssetResponse) ProtoMessage() {}

func (x *GetBalanceByAssetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_treasury_statement_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceByAssetResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceByAssetResponse) Descriptor() ([]byte, []int) {
	return file_accounting_treasury_statement_messages_proto_rawDescGZIP(), []int{4}
}

func (x *GetBalanceByAssetResponse) GetData() []*TreasuryStatement {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetBalancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TreasuryId string `protobuf:"bytes,1,opt,name=treasury_id,json=treasuryId,proto3" json:"treasury_id,omitempty"`
}

func (x *GetBalancesRequest) Reset() {
	*x = GetBalancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_treasury_statement_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalancesRequest) ProtoMessage() {}

func (x *GetBalancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_treasury_statement_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalancesRequest.ProtoReflect.Descriptor instead.
func (*GetBalancesRequest) Descriptor() ([]byte, []int) {
	return file_accounting_treasury_statement_messages_proto_rawDescGZIP(), []int{5}
}

func (x *GetBalancesRequest) GetTreasuryId() string {
	if x != nil {
		return x.TreasuryId
	}
	return ""
}

type GetBalancesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*TreasuryStatement `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetBalancesResponse) Reset() {
	*x = GetBalancesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_treasury_statement_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalancesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalancesResponse) ProtoMessage() {}

func (x *GetBalancesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_treasury_statement_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalancesResponse.ProtoReflect.Descriptor instead.
func (*GetBalancesResponse) Descriptor() ([]byte, []int) {
	return file_accounting_treasury_statement_messages_proto_rawDescGZIP(), []int{6}
}

func (x *GetBalancesResponse) GetData() []*TreasuryStatement {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_accounting_treasury_statement_messages_proto protoreflect.FileDescriptor

var file_accounting_treasury_statement_messages_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x72, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x79, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x11,
	0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x3b, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x6a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x4a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x60,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x22, 0x4e, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x79, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x42, 0x5a, 0x40, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x62, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x72, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x79, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3b, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounting_treasury_statement_messages_proto_rawDescOnce sync.Once
	file_accounting_treasury_statement_messages_proto_rawDescData = file_accounting_treasury_statement_messages_proto_rawDesc
)

func file_accounting_treasury_statement_messages_proto_rawDescGZIP() []byte {
	file_accounting_treasury_statement_messages_proto_rawDescOnce.Do(func() {
		file_accounting_treasury_statement_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounting_treasury_statement_messages_proto_rawDescData)
	})
	return file_accounting_treasury_statement_messages_proto_rawDescData
}

var file_accounting_treasury_statement_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_accounting_treasury_statement_messages_proto_goTypes = []interface{}{
	(*TreasuryStatement)(nil),         // 0: tStatement.TreasuryStatement
	(*GetBalanceAllRequest)(nil),      // 1: tStatement.GetBalanceAllRequest
	(*GetBalanceAllResponse)(nil),     // 2: tStatement.GetBalanceAllResponse
	(*GetBalanceByAssetRequest)(nil),  // 3: tStatement.GetBalanceByAssetRequest
	(*GetBalanceByAssetResponse)(nil), // 4: tStatement.GetBalanceByAssetResponse
	(*GetBalancesRequest)(nil),        // 5: tStatement.GetBalancesRequest
	(*GetBalancesResponse)(nil),       // 6: tStatement.GetBalancesResponse
	(*timestamppb.Timestamp)(nil),     // 7: google.protobuf.Timestamp
}
var file_accounting_treasury_statement_messages_proto_depIdxs = []int32{
	7, // 0: tStatement.TreasuryStatement.last_update:type_name -> google.protobuf.Timestamp
	0, // 1: tStatement.GetBalanceAllResponse.data:type_name -> tStatement.TreasuryStatement
	0, // 2: tStatement.GetBalanceByAssetResponse.data:type_name -> tStatement.TreasuryStatement
	0, // 3: tStatement.GetBalancesResponse.data:type_name -> tStatement.TreasuryStatement
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_accounting_treasury_statement_messages_proto_init() }
func file_accounting_treasury_statement_messages_proto_init() {
	if File_accounting_treasury_statement_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accounting_treasury_statement_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasuryStatement); i {
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
		file_accounting_treasury_statement_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceAllRequest); i {
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
		file_accounting_treasury_statement_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceAllResponse); i {
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
		file_accounting_treasury_statement_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceByAssetRequest); i {
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
		file_accounting_treasury_statement_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceByAssetResponse); i {
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
		file_accounting_treasury_statement_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalancesRequest); i {
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
		file_accounting_treasury_statement_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalancesResponse); i {
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
			RawDescriptor: file_accounting_treasury_statement_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accounting_treasury_statement_messages_proto_goTypes,
		DependencyIndexes: file_accounting_treasury_statement_messages_proto_depIdxs,
		MessageInfos:      file_accounting_treasury_statement_messages_proto_msgTypes,
	}.Build()
	File_accounting_treasury_statement_messages_proto = out.File
	file_accounting_treasury_statement_messages_proto_rawDesc = nil
	file_accounting_treasury_statement_messages_proto_goTypes = nil
	file_accounting_treasury_statement_messages_proto_depIdxs = nil
}
