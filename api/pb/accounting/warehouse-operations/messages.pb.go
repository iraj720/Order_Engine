// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: accounting/warehouse-operations/messages.proto

package warehouseOps

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

type WarehouseOperationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List   []*WarehouseOperationsRequest_Item `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Action Action                             `protobuf:"varint,2,opt,name=action,proto3,enum=warehouseOps.Action" json:"action,omitempty"`
}

func (x *WarehouseOperationsRequest) Reset() {
	*x = WarehouseOperationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_warehouse_operations_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseOperationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseOperationsRequest) ProtoMessage() {}

func (x *WarehouseOperationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_warehouse_operations_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseOperationsRequest.ProtoReflect.Descriptor instead.
func (*WarehouseOperationsRequest) Descriptor() ([]byte, []int) {
	return file_accounting_warehouse_operations_messages_proto_rawDescGZIP(), []int{0}
}

func (x *WarehouseOperationsRequest) GetList() []*WarehouseOperationsRequest_Item {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *WarehouseOperationsRequest) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ACTION_UNSPECIFIED
}

type WarehouseOperationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WarehouseOperationsResponse) Reset() {
	*x = WarehouseOperationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_warehouse_operations_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseOperationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseOperationsResponse) ProtoMessage() {}

func (x *WarehouseOperationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_warehouse_operations_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseOperationsResponse.ProtoReflect.Descriptor instead.
func (*WarehouseOperationsResponse) Descriptor() ([]byte, []int) {
	return file_accounting_warehouse_operations_messages_proto_rawDescGZIP(), []int{1}
}

type WarehouseTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*WarehouseTransferRequest_Item `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *WarehouseTransferRequest) Reset() {
	*x = WarehouseTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_warehouse_operations_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseTransferRequest) ProtoMessage() {}

func (x *WarehouseTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_warehouse_operations_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseTransferRequest.ProtoReflect.Descriptor instead.
func (*WarehouseTransferRequest) Descriptor() ([]byte, []int) {
	return file_accounting_warehouse_operations_messages_proto_rawDescGZIP(), []int{2}
}

func (x *WarehouseTransferRequest) GetList() []*WarehouseTransferRequest_Item {
	if x != nil {
		return x.List
	}
	return nil
}

type WarehouseTransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WarehouseTransferResponse) Reset() {
	*x = WarehouseTransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_warehouse_operations_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseTransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseTransferResponse) ProtoMessage() {}

func (x *WarehouseTransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_warehouse_operations_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseTransferResponse.ProtoReflect.Descriptor instead.
func (*WarehouseTransferResponse) Descriptor() ([]byte, []int) {
	return file_accounting_warehouse_operations_messages_proto_rawDescGZIP(), []int{3}
}

type WarehouseOperationsRequest_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId            string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Asset           string `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Amount          string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	UserWarehouse   string `protobuf:"bytes,6,opt,name=user_warehouse,json=userWarehouse,proto3" json:"user_warehouse,omitempty"`
	BrokerWarehouse string `protobuf:"bytes,7,opt,name=broker_warehouse,json=brokerWarehouse,proto3" json:"broker_warehouse,omitempty"`
	Broker          string `protobuf:"bytes,8,opt,name=broker,proto3" json:"broker,omitempty"`
	User            string `protobuf:"bytes,9,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *WarehouseOperationsRequest_Item) Reset() {
	*x = WarehouseOperationsRequest_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_warehouse_operations_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseOperationsRequest_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseOperationsRequest_Item) ProtoMessage() {}

func (x *WarehouseOperationsRequest_Item) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_warehouse_operations_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseOperationsRequest_Item.ProtoReflect.Descriptor instead.
func (*WarehouseOperationsRequest_Item) Descriptor() ([]byte, []int) {
	return file_accounting_warehouse_operations_messages_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WarehouseOperationsRequest_Item) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *WarehouseOperationsRequest_Item) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *WarehouseOperationsRequest_Item) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *WarehouseOperationsRequest_Item) GetUserWarehouse() string {
	if x != nil {
		return x.UserWarehouse
	}
	return ""
}

func (x *WarehouseOperationsRequest_Item) GetBrokerWarehouse() string {
	if x != nil {
		return x.BrokerWarehouse
	}
	return ""
}

func (x *WarehouseOperationsRequest_Item) GetBroker() string {
	if x != nil {
		return x.Broker
	}
	return ""
}

func (x *WarehouseOperationsRequest_Item) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type WarehouseTransferRequest_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId          string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Asset         string `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Amount        string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	FromWarehouse string `protobuf:"bytes,6,opt,name=from_warehouse,json=fromWarehouse,proto3" json:"from_warehouse,omitempty"`
	ToWarehouse   string `protobuf:"bytes,7,opt,name=to_warehouse,json=toWarehouse,proto3" json:"to_warehouse,omitempty"`
	Broker        string `protobuf:"bytes,8,opt,name=broker,proto3" json:"broker,omitempty"`
	FromUser      string `protobuf:"bytes,9,opt,name=from_user,json=fromUser,proto3" json:"from_user,omitempty"`
	ToUser        string `protobuf:"bytes,10,opt,name=to_user,json=toUser,proto3" json:"to_user,omitempty"`
}

func (x *WarehouseTransferRequest_Item) Reset() {
	*x = WarehouseTransferRequest_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_warehouse_operations_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseTransferRequest_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseTransferRequest_Item) ProtoMessage() {}

func (x *WarehouseTransferRequest_Item) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_warehouse_operations_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseTransferRequest_Item.ProtoReflect.Descriptor instead.
func (*WarehouseTransferRequest_Item) Descriptor() ([]byte, []int) {
	return file_accounting_warehouse_operations_messages_proto_rawDescGZIP(), []int{2, 0}
}

func (x *WarehouseTransferRequest_Item) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *WarehouseTransferRequest_Item) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *WarehouseTransferRequest_Item) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *WarehouseTransferRequest_Item) GetFromWarehouse() string {
	if x != nil {
		return x.FromWarehouse
	}
	return ""
}

func (x *WarehouseTransferRequest_Item) GetToWarehouse() string {
	if x != nil {
		return x.ToWarehouse
	}
	return ""
}

func (x *WarehouseTransferRequest_Item) GetBroker() string {
	if x != nil {
		return x.Broker
	}
	return ""
}

func (x *WarehouseTransferRequest_Item) GetFromUser() string {
	if x != nil {
		return x.FromUser
	}
	return ""
}

func (x *WarehouseTransferRequest_Item) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

var File_accounting_warehouse_operations_messages_proto protoreflect.FileDescriptor

var file_accounting_warehouse_operations_messages_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x77, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4f, 0x70, 0x73, 0x1a, 0x2b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x2d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x02, 0x0a, 0x1a,
	0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x4f, 0x70, 0x73, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4f, 0x70, 0x73, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xc7, 0x01, 0x0a, 0x04,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x75, 0x73, 0x65, 0x72, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x1d, 0x0a, 0x1b, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xbf, 0x02, 0x0a, 0x18, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4f, 0x70, 0x73, 0x2e, 0x57,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x1a, 0xe1, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x13, 0x0a, 0x05, 0x74,
	0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x5f, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x57,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x22, 0x1b, 0x0a, 0x19, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x48, 0x5a, 0x46, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x62, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4f, 0x70, 0x73,
	0x3b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4f, 0x70, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounting_warehouse_operations_messages_proto_rawDescOnce sync.Once
	file_accounting_warehouse_operations_messages_proto_rawDescData = file_accounting_warehouse_operations_messages_proto_rawDesc
)

func file_accounting_warehouse_operations_messages_proto_rawDescGZIP() []byte {
	file_accounting_warehouse_operations_messages_proto_rawDescOnce.Do(func() {
		file_accounting_warehouse_operations_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounting_warehouse_operations_messages_proto_rawDescData)
	})
	return file_accounting_warehouse_operations_messages_proto_rawDescData
}

var file_accounting_warehouse_operations_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_accounting_warehouse_operations_messages_proto_goTypes = []interface{}{
	(*WarehouseOperationsRequest)(nil),      // 0: warehouseOps.WarehouseOperationsRequest
	(*WarehouseOperationsResponse)(nil),     // 1: warehouseOps.WarehouseOperationsResponse
	(*WarehouseTransferRequest)(nil),        // 2: warehouseOps.WarehouseTransferRequest
	(*WarehouseTransferResponse)(nil),       // 3: warehouseOps.WarehouseTransferResponse
	(*WarehouseOperationsRequest_Item)(nil), // 4: warehouseOps.WarehouseOperationsRequest.Item
	(*WarehouseTransferRequest_Item)(nil),   // 5: warehouseOps.WarehouseTransferRequest.Item
	(Action)(0),                             // 6: warehouseOps.Action
}
var file_accounting_warehouse_operations_messages_proto_depIdxs = []int32{
	4, // 0: warehouseOps.WarehouseOperationsRequest.list:type_name -> warehouseOps.WarehouseOperationsRequest.Item
	6, // 1: warehouseOps.WarehouseOperationsRequest.action:type_name -> warehouseOps.Action
	5, // 2: warehouseOps.WarehouseTransferRequest.list:type_name -> warehouseOps.WarehouseTransferRequest.Item
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_accounting_warehouse_operations_messages_proto_init() }
func file_accounting_warehouse_operations_messages_proto_init() {
	if File_accounting_warehouse_operations_messages_proto != nil {
		return
	}
	file_accounting_warehouse_operations_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_accounting_warehouse_operations_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseOperationsRequest); i {
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
		file_accounting_warehouse_operations_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseOperationsResponse); i {
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
		file_accounting_warehouse_operations_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseTransferRequest); i {
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
		file_accounting_warehouse_operations_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseTransferResponse); i {
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
		file_accounting_warehouse_operations_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseOperationsRequest_Item); i {
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
		file_accounting_warehouse_operations_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WarehouseTransferRequest_Item); i {
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
			RawDescriptor: file_accounting_warehouse_operations_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accounting_warehouse_operations_messages_proto_goTypes,
		DependencyIndexes: file_accounting_warehouse_operations_messages_proto_depIdxs,
		MessageInfos:      file_accounting_warehouse_operations_messages_proto_msgTypes,
	}.Build()
	File_accounting_warehouse_operations_messages_proto = out.File
	file_accounting_warehouse_operations_messages_proto_rawDesc = nil
	file_accounting_warehouse_operations_messages_proto_goTypes = nil
	file_accounting_warehouse_operations_messages_proto_depIdxs = nil
}
