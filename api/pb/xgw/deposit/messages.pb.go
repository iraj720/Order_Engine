// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: xgw/deposit/messages.proto

package deposit

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

type DepositHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel   string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Asset     string `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"` // binance statuses (0:pending,6: credited but cannot withdraw, 1:success) if not provided all statuses will be returned
	Offset    int32  `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit     int32  `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	StartTime int64  `protobuf:"varint,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64  `protobuf:"varint,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *DepositHistoryRequest) Reset() {
	*x = DepositHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xgw_deposit_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositHistoryRequest) ProtoMessage() {}

func (x *DepositHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xgw_deposit_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositHistoryRequest.ProtoReflect.Descriptor instead.
func (*DepositHistoryRequest) Descriptor() ([]byte, []int) {
	return file_xgw_deposit_messages_proto_rawDescGZIP(), []int{0}
}

func (x *DepositHistoryRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *DepositHistoryRequest) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *DepositHistoryRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DepositHistoryRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *DepositHistoryRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *DepositHistoryRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *DepositHistoryRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type DepositHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string status = 1;
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// repeated string error = 3;
	Data []*DepositHistoryResponse_Data `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *DepositHistoryResponse) Reset() {
	*x = DepositHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xgw_deposit_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositHistoryResponse) ProtoMessage() {}

func (x *DepositHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xgw_deposit_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositHistoryResponse.ProtoReflect.Descriptor instead.
func (*DepositHistoryResponse) Descriptor() ([]byte, []int) {
	return file_xgw_deposit_messages_proto_rawDescGZIP(), []int{1}
}

func (x *DepositHistoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DepositHistoryResponse) GetData() []*DepositHistoryResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type DepositHistoryResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address       string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Amount        string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	InsertTime    int64  `protobuf:"varint,3,opt,name=insert_time,json=insertTime,proto3" json:"insert_time,omitempty"`
	Asset         string `protobuf:"bytes,4,opt,name=asset,proto3" json:"asset,omitempty"`
	TxId          string `protobuf:"bytes,5,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	AddressTag    string `protobuf:"bytes,6,opt,name=address_tag,json=addressTag,proto3" json:"address_tag,omitempty"`
	Network       string `protobuf:"bytes,7,opt,name=network,proto3" json:"network,omitempty"`
	Status        int32  `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	TransferType  int32  `protobuf:"varint,9,opt,name=transfer_type,json=transferType,proto3" json:"transfer_type,omitempty"`
	UnlockConfirm int32  `protobuf:"varint,10,opt,name=unlock_confirm,json=unlockConfirm,proto3" json:"unlock_confirm,omitempty"`
	ConfirmTimes  string `protobuf:"bytes,11,opt,name=confirm_times,json=confirmTimes,proto3" json:"confirm_times,omitempty"`
	WalletType    int32  `protobuf:"varint,12,opt,name=wallet_type,json=walletType,proto3" json:"wallet_type,omitempty"`
}

func (x *DepositHistoryResponse_Data) Reset() {
	*x = DepositHistoryResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xgw_deposit_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositHistoryResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositHistoryResponse_Data) ProtoMessage() {}

func (x *DepositHistoryResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_xgw_deposit_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositHistoryResponse_Data.ProtoReflect.Descriptor instead.
func (*DepositHistoryResponse_Data) Descriptor() ([]byte, []int) {
	return file_xgw_deposit_messages_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DepositHistoryResponse_Data) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DepositHistoryResponse_Data) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *DepositHistoryResponse_Data) GetInsertTime() int64 {
	if x != nil {
		return x.InsertTime
	}
	return 0
}

func (x *DepositHistoryResponse_Data) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *DepositHistoryResponse_Data) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *DepositHistoryResponse_Data) GetAddressTag() string {
	if x != nil {
		return x.AddressTag
	}
	return ""
}

func (x *DepositHistoryResponse_Data) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *DepositHistoryResponse_Data) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DepositHistoryResponse_Data) GetTransferType() int32 {
	if x != nil {
		return x.TransferType
	}
	return 0
}

func (x *DepositHistoryResponse_Data) GetUnlockConfirm() int32 {
	if x != nil {
		return x.UnlockConfirm
	}
	return 0
}

func (x *DepositHistoryResponse_Data) GetConfirmTimes() string {
	if x != nil {
		return x.ConfirmTimes
	}
	return ""
}

func (x *DepositHistoryResponse_Data) GetWalletType() int32 {
	if x != nil {
		return x.WalletType
	}
	return 0
}

var File_xgw_deposit_messages_proto protoreflect.FileDescriptor

var file_xgw_deposit_messages_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x78, 0x67, 0x77, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x22, 0xc7, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0xd8, 0x03, 0x0a, 0x16, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xe9,
	0x02, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x78, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x54, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x72, 0x61,
	0x62, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x78, 0x67, 0x77, 0x2f, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x3b, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xgw_deposit_messages_proto_rawDescOnce sync.Once
	file_xgw_deposit_messages_proto_rawDescData = file_xgw_deposit_messages_proto_rawDesc
)

func file_xgw_deposit_messages_proto_rawDescGZIP() []byte {
	file_xgw_deposit_messages_proto_rawDescOnce.Do(func() {
		file_xgw_deposit_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_xgw_deposit_messages_proto_rawDescData)
	})
	return file_xgw_deposit_messages_proto_rawDescData
}

var file_xgw_deposit_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_xgw_deposit_messages_proto_goTypes = []interface{}{
	(*DepositHistoryRequest)(nil),       // 0: deposit.DepositHistoryRequest
	(*DepositHistoryResponse)(nil),      // 1: deposit.DepositHistoryResponse
	(*DepositHistoryResponse_Data)(nil), // 2: deposit.DepositHistoryResponse.Data
}
var file_xgw_deposit_messages_proto_depIdxs = []int32{
	2, // 0: deposit.DepositHistoryResponse.data:type_name -> deposit.DepositHistoryResponse.Data
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_xgw_deposit_messages_proto_init() }
func file_xgw_deposit_messages_proto_init() {
	if File_xgw_deposit_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xgw_deposit_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositHistoryRequest); i {
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
		file_xgw_deposit_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositHistoryResponse); i {
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
		file_xgw_deposit_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositHistoryResponse_Data); i {
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
			RawDescriptor: file_xgw_deposit_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xgw_deposit_messages_proto_goTypes,
		DependencyIndexes: file_xgw_deposit_messages_proto_depIdxs,
		MessageInfos:      file_xgw_deposit_messages_proto_msgTypes,
	}.Build()
	File_xgw_deposit_messages_proto = out.File
	file_xgw_deposit_messages_proto_rawDesc = nil
	file_xgw_deposit_messages_proto_goTypes = nil
	file_xgw_deposit_messages_proto_depIdxs = nil
}