// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: accounting/externalexchange/messages.proto

package externalexchange

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

type ExternalExchangeCost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Asset  string `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ExternalExchangeCost) Reset() {
	*x = ExternalExchangeCost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_externalexchange_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalExchangeCost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalExchangeCost) ProtoMessage() {}

func (x *ExternalExchangeCost) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_externalexchange_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalExchangeCost.ProtoReflect.Descriptor instead.
func (*ExternalExchangeCost) Descriptor() ([]byte, []int) {
	return file_accounting_externalexchange_messages_proto_rawDescGZIP(), []int{0}
}

func (x *ExternalExchangeCost) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *ExternalExchangeCost) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type ExternalExchange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId                 string                  `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	EgressAsset          string                  `protobuf:"bytes,2,opt,name=egress_asset,json=egressAsset,proto3" json:"egress_asset,omitempty"`
	EgressAmount         string                  `protobuf:"bytes,3,opt,name=egress_amount,json=egressAmount,proto3" json:"egress_amount,omitempty"`
	EgressExecutedAmount string                  `protobuf:"bytes,4,opt,name=egress_executed_amount,json=egressExecutedAmount,proto3" json:"egress_executed_amount,omitempty"`
	IngressAsset         string                  `protobuf:"bytes,5,opt,name=ingress_asset,json=ingressAsset,proto3" json:"ingress_asset,omitempty"`
	IngressAmount        string                  `protobuf:"bytes,6,opt,name=ingress_amount,json=ingressAmount,proto3" json:"ingress_amount,omitempty"`
	BrokerWarehouse      string                  `protobuf:"bytes,7,opt,name=broker_warehouse,json=brokerWarehouse,proto3" json:"broker_warehouse,omitempty"`
	Broker               string                  `protobuf:"bytes,8,opt,name=broker,proto3" json:"broker,omitempty"`
	BrokerTreasury       string                  `protobuf:"bytes,9,opt,name=broker_treasury,json=brokerTreasury,proto3" json:"broker_treasury,omitempty"`
	Costs                []*ExternalExchangeCost `protobuf:"bytes,10,rep,name=costs,proto3" json:"costs,omitempty"`
}

func (x *ExternalExchange) Reset() {
	*x = ExternalExchange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_externalexchange_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalExchange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalExchange) ProtoMessage() {}

func (x *ExternalExchange) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_externalexchange_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalExchange.ProtoReflect.Descriptor instead.
func (*ExternalExchange) Descriptor() ([]byte, []int) {
	return file_accounting_externalexchange_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ExternalExchange) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *ExternalExchange) GetEgressAsset() string {
	if x != nil {
		return x.EgressAsset
	}
	return ""
}

func (x *ExternalExchange) GetEgressAmount() string {
	if x != nil {
		return x.EgressAmount
	}
	return ""
}

func (x *ExternalExchange) GetEgressExecutedAmount() string {
	if x != nil {
		return x.EgressExecutedAmount
	}
	return ""
}

func (x *ExternalExchange) GetIngressAsset() string {
	if x != nil {
		return x.IngressAsset
	}
	return ""
}

func (x *ExternalExchange) GetIngressAmount() string {
	if x != nil {
		return x.IngressAmount
	}
	return ""
}

func (x *ExternalExchange) GetBrokerWarehouse() string {
	if x != nil {
		return x.BrokerWarehouse
	}
	return ""
}

func (x *ExternalExchange) GetBroker() string {
	if x != nil {
		return x.Broker
	}
	return ""
}

func (x *ExternalExchange) GetBrokerTreasury() string {
	if x != nil {
		return x.BrokerTreasury
	}
	return ""
}

func (x *ExternalExchange) GetCosts() []*ExternalExchangeCost {
	if x != nil {
		return x.Costs
	}
	return nil
}

type ExternalExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExternalExchangeSingle *ExternalExchange `protobuf:"bytes,1,opt,name=external_exchange_single,json=externalExchangeSingle,proto3" json:"external_exchange_single,omitempty"`
	Action                 Action            `protobuf:"varint,2,opt,name=action,proto3,enum=externalexchange.Action" json:"action,omitempty"`
}

func (x *ExternalExchangeRequest) Reset() {
	*x = ExternalExchangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_externalexchange_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalExchangeRequest) ProtoMessage() {}

func (x *ExternalExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_externalexchange_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalExchangeRequest.ProtoReflect.Descriptor instead.
func (*ExternalExchangeRequest) Descriptor() ([]byte, []int) {
	return file_accounting_externalexchange_messages_proto_rawDescGZIP(), []int{2}
}

func (x *ExternalExchangeRequest) GetExternalExchangeSingle() *ExternalExchange {
	if x != nil {
		return x.ExternalExchangeSingle
	}
	return nil
}

func (x *ExternalExchangeRequest) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ACTION_UNSPECIFIED
}

type ExternalExchangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error   []string `protobuf:"bytes,3,rep,name=error,proto3" json:"error,omitempty"`
	Data    string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ExternalExchangeResponse) Reset() {
	*x = ExternalExchangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounting_externalexchange_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalExchangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalExchangeResponse) ProtoMessage() {}

func (x *ExternalExchangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_externalexchange_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalExchangeResponse.ProtoReflect.Descriptor instead.
func (*ExternalExchangeResponse) Descriptor() ([]byte, []int) {
	return file_accounting_externalexchange_messages_proto_rawDescGZIP(), []int{3}
}

func (x *ExternalExchangeResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ExternalExchangeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ExternalExchangeResponse) GetError() []string {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ExternalExchangeResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_accounting_externalexchange_messages_proto protoreflect.FileDescriptor

var file_accounting_externalexchange_messages_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x27,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x14, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x9b, 0x03,
	0x0a, 0x10, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x34, 0x0a, 0x16, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x74,
	0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x12, 0x3c, 0x0a,
	0x05, 0x63, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x43, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x63, 0x6f, 0x73, 0x74, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x17,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5c, 0x0a, 0x18, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x16, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x76, 0x0a, 0x18, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x3b, 0x5a, 0x39, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x3b, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounting_externalexchange_messages_proto_rawDescOnce sync.Once
	file_accounting_externalexchange_messages_proto_rawDescData = file_accounting_externalexchange_messages_proto_rawDesc
)

func file_accounting_externalexchange_messages_proto_rawDescGZIP() []byte {
	file_accounting_externalexchange_messages_proto_rawDescOnce.Do(func() {
		file_accounting_externalexchange_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounting_externalexchange_messages_proto_rawDescData)
	})
	return file_accounting_externalexchange_messages_proto_rawDescData
}

var file_accounting_externalexchange_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_accounting_externalexchange_messages_proto_goTypes = []interface{}{
	(*ExternalExchangeCost)(nil),     // 0: externalexchange.ExternalExchangeCost
	(*ExternalExchange)(nil),         // 1: externalexchange.ExternalExchange
	(*ExternalExchangeRequest)(nil),  // 2: externalexchange.ExternalExchangeRequest
	(*ExternalExchangeResponse)(nil), // 3: externalexchange.ExternalExchangeResponse
	(Action)(0),                      // 4: externalexchange.Action
}
var file_accounting_externalexchange_messages_proto_depIdxs = []int32{
	0, // 0: externalexchange.ExternalExchange.costs:type_name -> externalexchange.ExternalExchangeCost
	1, // 1: externalexchange.ExternalExchangeRequest.external_exchange_single:type_name -> externalexchange.ExternalExchange
	4, // 2: externalexchange.ExternalExchangeRequest.action:type_name -> externalexchange.Action
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_accounting_externalexchange_messages_proto_init() }
func file_accounting_externalexchange_messages_proto_init() {
	if File_accounting_externalexchange_messages_proto != nil {
		return
	}
	file_accounting_externalexchange_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_accounting_externalexchange_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalExchangeCost); i {
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
		file_accounting_externalexchange_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalExchange); i {
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
		file_accounting_externalexchange_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalExchangeRequest); i {
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
		file_accounting_externalexchange_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalExchangeResponse); i {
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
			RawDescriptor: file_accounting_externalexchange_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accounting_externalexchange_messages_proto_goTypes,
		DependencyIndexes: file_accounting_externalexchange_messages_proto_depIdxs,
		MessageInfos:      file_accounting_externalexchange_messages_proto_msgTypes,
	}.Build()
	File_accounting_externalexchange_messages_proto = out.File
	file_accounting_externalexchange_messages_proto_rawDesc = nil
	file_accounting_externalexchange_messages_proto_goTypes = nil
	file_accounting_externalexchange_messages_proto_depIdxs = nil
}
