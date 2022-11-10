// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: warehouse/transaction/messages.proto

package transaction

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	commons "rabex/api/pb/commons"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WithdrawStatus commons.WithdrawStatus `protobuf:"varint,1,opt,name=WithdrawStatus,proto3,enum=rabex.api.commons.withdraw.WithdrawStatus" json:"WithdrawStatus,omitempty"`
	DepositStatus  commons.DepositStatus  `protobuf:"varint,2,opt,name=DepositStatus,proto3,enum=rabex.api.commons.deposit.DepositStatus" json:"DepositStatus,omitempty"`
}

func (x *TransactionStatus) Reset() {
	*x = TransactionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_transaction_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionStatus) ProtoMessage() {}

func (x *TransactionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_transaction_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionStatus.ProtoReflect.Descriptor instead.
func (*TransactionStatus) Descriptor() ([]byte, []int) {
	return file_warehouse_transaction_messages_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionStatus) GetWithdrawStatus() commons.WithdrawStatus {
	if x != nil {
		return x.WithdrawStatus
	}
	return commons.WithdrawStatus(0)
}

func (x *TransactionStatus) GetDepositStatus() commons.DepositStatus {
	if x != nil {
		return x.DepositStatus
	}
	return commons.DepositStatus(0)
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_transaction_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_transaction_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_warehouse_transaction_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type TransactionFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID           string                  `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	TxType         commons.TransactionType `protobuf:"varint,2,opt,name=txType,proto3,enum=rabex.api.commons.TransactionType" json:"txType,omitempty"`
	Owners         []string                `protobuf:"bytes,3,rep,name=owners,proto3" json:"owners,omitempty"`
	AccountType    AccountType             `protobuf:"varint,4,opt,name=accountType,proto3,enum=rabex.api.warehouse.transaction.AccountType" json:"accountType,omitempty"`
	WithdrawStatus commons.WithdrawStatus  `protobuf:"varint,5,opt,name=withdrawStatus,proto3,enum=rabex.api.commons.withdraw.WithdrawStatus" json:"withdrawStatus,omitempty"`
	DepositStatus  commons.DepositStatus   `protobuf:"varint,6,opt,name=depositStatus,proto3,enum=rabex.api.commons.deposit.DepositStatus" json:"depositStatus,omitempty"`
}

func (x *TransactionFilter) Reset() {
	*x = TransactionFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_transaction_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionFilter) ProtoMessage() {}

func (x *TransactionFilter) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_transaction_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionFilter.ProtoReflect.Descriptor instead.
func (*TransactionFilter) Descriptor() ([]byte, []int) {
	return file_warehouse_transaction_messages_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionFilter) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *TransactionFilter) GetTxType() commons.TransactionType {
	if x != nil {
		return x.TxType
	}
	return commons.TransactionType(0)
}

func (x *TransactionFilter) GetOwners() []string {
	if x != nil {
		return x.Owners
	}
	return nil
}

func (x *TransactionFilter) GetAccountType() AccountType {
	if x != nil {
		return x.AccountType
	}
	return AccountType_Cost
}

func (x *TransactionFilter) GetWithdrawStatus() commons.WithdrawStatus {
	if x != nil {
		return x.WithdrawStatus
	}
	return commons.WithdrawStatus(0)
}

func (x *TransactionFilter) GetDepositStatus() commons.DepositStatus {
	if x != nil {
		return x.DepositStatus
	}
	return commons.DepositStatus(0)
}

var File_warehouse_transaction_messages_proto protoreflect.FileDescriptor

var file_warehouse_transaction_messages_proto_rawDesc = []byte{
	0x0a, 0x24, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x11, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x52, 0x0a, 0x0e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x4e, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x72, 0x61,
	0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x22, 0xef, 0x02, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x3a, 0x0a,
	0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x73, 0x12, 0x4e, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x52, 0x0a, 0x0e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x72, 0x61, 0x62, 0x65,
	0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4e, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x72,
	0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x30, 0x5a, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_warehouse_transaction_messages_proto_rawDescOnce sync.Once
	file_warehouse_transaction_messages_proto_rawDescData = file_warehouse_transaction_messages_proto_rawDesc
)

func file_warehouse_transaction_messages_proto_rawDescGZIP() []byte {
	file_warehouse_transaction_messages_proto_rawDescOnce.Do(func() {
		file_warehouse_transaction_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_warehouse_transaction_messages_proto_rawDescData)
	})
	return file_warehouse_transaction_messages_proto_rawDescData
}

var file_warehouse_transaction_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_warehouse_transaction_messages_proto_goTypes = []interface{}{
	(*TransactionStatus)(nil),    // 0: rabex.api.warehouse.transaction.TransactionStatus
	(*Request)(nil),              // 1: rabex.api.warehouse.transaction.Request
	(*TransactionFilter)(nil),    // 2: rabex.api.warehouse.transaction.TransactionFilter
	(commons.WithdrawStatus)(0),  // 3: rabex.api.commons.withdraw.WithdrawStatus
	(commons.DepositStatus)(0),   // 4: rabex.api.commons.deposit.DepositStatus
	(commons.TransactionType)(0), // 5: rabex.api.commons.TransactionType
	(AccountType)(0),             // 6: rabex.api.warehouse.transaction.AccountType
}
var file_warehouse_transaction_messages_proto_depIdxs = []int32{
	3, // 0: rabex.api.warehouse.transaction.TransactionStatus.WithdrawStatus:type_name -> rabex.api.commons.withdraw.WithdrawStatus
	4, // 1: rabex.api.warehouse.transaction.TransactionStatus.DepositStatus:type_name -> rabex.api.commons.deposit.DepositStatus
	5, // 2: rabex.api.warehouse.transaction.TransactionFilter.txType:type_name -> rabex.api.commons.TransactionType
	6, // 3: rabex.api.warehouse.transaction.TransactionFilter.accountType:type_name -> rabex.api.warehouse.transaction.AccountType
	3, // 4: rabex.api.warehouse.transaction.TransactionFilter.withdrawStatus:type_name -> rabex.api.commons.withdraw.WithdrawStatus
	4, // 5: rabex.api.warehouse.transaction.TransactionFilter.depositStatus:type_name -> rabex.api.commons.deposit.DepositStatus
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_warehouse_transaction_messages_proto_init() }
func file_warehouse_transaction_messages_proto_init() {
	if File_warehouse_transaction_messages_proto != nil {
		return
	}
	file_warehouse_transaction_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_warehouse_transaction_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionStatus); i {
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
		file_warehouse_transaction_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_warehouse_transaction_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionFilter); i {
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
			RawDescriptor: file_warehouse_transaction_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_warehouse_transaction_messages_proto_goTypes,
		DependencyIndexes: file_warehouse_transaction_messages_proto_depIdxs,
		MessageInfos:      file_warehouse_transaction_messages_proto_msgTypes,
	}.Build()
	File_warehouse_transaction_messages_proto = out.File
	file_warehouse_transaction_messages_proto_rawDesc = nil
	file_warehouse_transaction_messages_proto_goTypes = nil
	file_warehouse_transaction_messages_proto_depIdxs = nil
}