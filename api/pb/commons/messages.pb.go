// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: commons/messages.proto

package commons

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Null int32

const (
	Null_NULL_VALUE Null = 0
)

// Enum value maps for Null.
var (
	Null_name = map[int32]string{
		0: "NULL_VALUE",
	}
	Null_value = map[string]int32{
		"NULL_VALUE": 0,
	}
)

func (x Null) Enum() *Null {
	p := new(Null)
	*p = x
	return p
}

func (x Null) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Null) Descriptor() protoreflect.EnumDescriptor {
	return file_commons_messages_proto_enumTypes[0].Descriptor()
}

func (Null) Type() protoreflect.EnumType {
	return &file_commons_messages_proto_enumTypes[0]
}

func (x Null) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Null.Descriptor instead.
func (Null) EnumDescriptor() ([]byte, []int) {
	return file_commons_messages_proto_rawDescGZIP(), []int{0}
}

type TargetUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub string `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
}

func (x *TargetUser) Reset() {
	*x = TargetUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetUser) ProtoMessage() {}

func (x *TargetUser) ProtoReflect() protoreflect.Message {
	mi := &file_commons_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetUser.ProtoReflect.Descriptor instead.
func (*TargetUser) Descriptor() ([]byte, []int) {
	return file_commons_messages_proto_rawDescGZIP(), []int{0}
}

func (x *TargetUser) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

type InternalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin      string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// reference :
	// some messages could be a response to a previous one
	// thus, we can set reference for them
	Reference int64   `protobuf:"varint,3,opt,name=reference,proto3" json:"reference,omitempty"`
	Topic     string  `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"` // works like message type
	Owner     []byte  `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Payload   *Struct `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *InternalRequest) Reset() {
	*x = InternalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalRequest) ProtoMessage() {}

func (x *InternalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commons_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalRequest.ProtoReflect.Descriptor instead.
func (*InternalRequest) Descriptor() ([]byte, []int) {
	return file_commons_messages_proto_rawDescGZIP(), []int{1}
}

func (x *InternalRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *InternalRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *InternalRequest) GetReference() int64 {
	if x != nil {
		return x.Reference
	}
	return 0
}

func (x *InternalRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *InternalRequest) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *InternalRequest) GetPayload() *Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

type InternalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin      string  `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination string  `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Reference   int64   `protobuf:"varint,3,opt,name=reference,proto3" json:"reference,omitempty"`
	Status      int32   `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Topic       string  `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
	Owner       []byte  `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Payload     *Struct `protobuf:"bytes,7,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *InternalResponse) Reset() {
	*x = InternalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalResponse) ProtoMessage() {}

func (x *InternalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commons_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalResponse.ProtoReflect.Descriptor instead.
func (*InternalResponse) Descriptor() ([]byte, []int) {
	return file_commons_messages_proto_rawDescGZIP(), []int{2}
}

func (x *InternalResponse) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *InternalResponse) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *InternalResponse) GetReference() int64 {
	if x != nil {
		return x.Reference
	}
	return 0
}

func (x *InternalResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *InternalResponse) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *InternalResponse) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *InternalResponse) GetPayload() *Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

type GeneralResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Array `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GeneralResponse) Reset() {
	*x = GeneralResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralResponse) ProtoMessage() {}

func (x *GeneralResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commons_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralResponse.ProtoReflect.Descriptor instead.
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return file_commons_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GeneralResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GeneralResponse) GetData() []*Array {
	if x != nil {
		return x.Data
	}
	return nil
}

type Array struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*Value `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *Array) Reset() {
	*x = Array{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Array) ProtoMessage() {}

func (x *Array) ProtoReflect() protoreflect.Message {
	mi := &file_commons_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Array.ProtoReflect.Descriptor instead.
func (*Array) Descriptor() ([]byte, []int) {
	return file_commons_messages_proto_rawDescGZIP(), []int{4}
}

func (x *Array) GetValue() []*Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//	*Value_NullValue
	//	*Value_NumberValue
	//	*Value_StringValue
	//	*Value_BoolValue
	//	*Value_StructValue
	//	*Value_IntValue
	//	*Value_Int32Value
	//	*Value_Float32Value
	//	*Value_Float64Value
	//	*Value_AnyValue
	//	*Value_ArrayValue
	Kind isValue_Kind `protobuf_oneof:"kind"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_commons_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_commons_messages_proto_rawDescGZIP(), []int{5}
}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Value) GetNullValue() Null {
	if x, ok := x.GetKind().(*Value_NullValue); ok {
		return x.NullValue
	}
	return Null_NULL_VALUE
}

func (x *Value) GetNumberValue() float64 {
	if x, ok := x.GetKind().(*Value_NumberValue); ok {
		return x.NumberValue
	}
	return 0
}

func (x *Value) GetStringValue() string {
	if x, ok := x.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *Value) GetBoolValue() bool {
	if x, ok := x.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *Value) GetStructValue() *Struct {
	if x, ok := x.GetKind().(*Value_StructValue); ok {
		return x.StructValue
	}
	return nil
}

func (x *Value) GetIntValue() int64 {
	if x, ok := x.GetKind().(*Value_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *Value) GetInt32Value() int32 {
	if x, ok := x.GetKind().(*Value_Int32Value); ok {
		return x.Int32Value
	}
	return 0
}

func (x *Value) GetFloat32Value() float32 {
	if x, ok := x.GetKind().(*Value_Float32Value); ok {
		return x.Float32Value
	}
	return 0
}

func (x *Value) GetFloat64Value() float64 {
	if x, ok := x.GetKind().(*Value_Float64Value); ok {
		return x.Float64Value
	}
	return 0
}

func (x *Value) GetAnyValue() *anypb.Any {
	if x, ok := x.GetKind().(*Value_AnyValue); ok {
		return x.AnyValue
	}
	return nil
}

func (x *Value) GetArrayValue() *Array {
	if x, ok := x.GetKind().(*Value_ArrayValue); ok {
		return x.ArrayValue
	}
	return nil
}

type isValue_Kind interface {
	isValue_Kind()
}

type Value_NullValue struct {
	NullValue Null `protobuf:"varint,1,opt,name=null_value,json=nullValue,proto3,enum=rabex.api.commons.Null,oneof"`
}

type Value_NumberValue struct {
	NumberValue float64 `protobuf:"fixed64,2,opt,name=number_value,json=numberValue,proto3,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Value_StructValue struct {
	StructValue *Struct `protobuf:"bytes,5,opt,name=struct_value,json=structValue,proto3,oneof"`
}

type Value_IntValue struct {
	IntValue int64 `protobuf:"varint,6,opt,name=int_value,json=intValue,proto3,oneof"`
}

type Value_Int32Value struct {
	Int32Value int32 `protobuf:"varint,7,opt,name=int32_value,json=int32Value,proto3,oneof"`
}

type Value_Float32Value struct {
	Float32Value float32 `protobuf:"fixed32,8,opt,name=float32_value,json=float32Value,proto3,oneof"`
}

type Value_Float64Value struct {
	Float64Value float64 `protobuf:"fixed64,9,opt,name=float64_value,json=float64Value,proto3,oneof"`
}

type Value_AnyValue struct {
	AnyValue *anypb.Any `protobuf:"bytes,10,opt,name=any_value,json=anyValue,proto3,oneof"`
}

type Value_ArrayValue struct {
	ArrayValue *Array `protobuf:"bytes,11,opt,name=array_value,json=arrayValue,proto3,oneof"`
}

func (*Value_NullValue) isValue_Kind() {}

func (*Value_NumberValue) isValue_Kind() {}

func (*Value_StringValue) isValue_Kind() {}

func (*Value_BoolValue) isValue_Kind() {}

func (*Value_StructValue) isValue_Kind() {}

func (*Value_IntValue) isValue_Kind() {}

func (*Value_Int32Value) isValue_Kind() {}

func (*Value_Float32Value) isValue_Kind() {}

func (*Value_Float64Value) isValue_Kind() {}

func (*Value_AnyValue) isValue_Kind() {}

func (*Value_ArrayValue) isValue_Kind() {}

type Struct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields map[string]*Value `protobuf:"bytes,1,rep,name=Fields,proto3" json:"Fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Struct) Reset() {
	*x = Struct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Struct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Struct) ProtoMessage() {}

func (x *Struct) ProtoReflect() protoreflect.Message {
	mi := &file_commons_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Struct.ProtoReflect.Descriptor instead.
func (*Struct) Descriptor() ([]byte, []int) {
	return file_commons_messages_proto_rawDescGZIP(), []int{6}
}

func (x *Struct) GetFields() map[string]*Value {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Validation []string `protobuf:"bytes,2,rep,name=validation,proto3" json:"validation,omitempty"`
	Message    string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Status     string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_commons_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_commons_messages_proto_rawDescGZIP(), []int{7}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetValidation() []string {
	if x != nil {
		return x.Validation
	}
	return nil
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type AssetList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetList []Assets `protobuf:"varint,1,rep,packed,name=assetList,proto3,enum=rabex.api.commons.Assets" json:"assetList,omitempty"`
}

func (x *AssetList) Reset() {
	*x = AssetList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetList) ProtoMessage() {}

func (x *AssetList) ProtoReflect() protoreflect.Message {
	mi := &file_commons_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetList.ProtoReflect.Descriptor instead.
func (*AssetList) Descriptor() ([]byte, []int) {
	return file_commons_messages_proto_rawDescGZIP(), []int{8}
}

func (x *AssetList) GetAssetList() []Assets {
	if x != nil {
		return x.AssetList
	}
	return nil
}

var File_commons_messages_proto protoreflect.FileDescriptor

var file_commons_messages_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0a, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x62, 0x22, 0xca, 0x01, 0x0a, 0x0f,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xe3, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x61, 0x62,
	0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x5f,
	0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x61, 0x62, 0x65,
	0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22,
	0x37, 0x0a, 0x05, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xf6, 0x03, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x48,
	0x00, 0x52, 0x09, 0x6e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02,
	0x48, 0x00, 0x52, 0x0c, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x25, 0x0a, 0x0d, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0c, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x61, 0x6e, 0x79, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x48, 0x00, 0x52, 0x08, 0x61, 0x6e, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3b, 0x0a, 0x0b,
	0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x0a, 0x61,
	0x72, 0x72, 0x61, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x22, 0x9c, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x3d, 0x0a, 0x06,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72,
	0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x53, 0x0a, 0x0b, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x61,
	0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x6d, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x44, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x09,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x2a, 0x16, 0x0a, 0x04, 0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x0e, 0x0a,
	0x0a, 0x4e, 0x55, 0x4c, 0x4c, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x00, 0x42, 0x1e, 0x5a,
	0x1c, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commons_messages_proto_rawDescOnce sync.Once
	file_commons_messages_proto_rawDescData = file_commons_messages_proto_rawDesc
)

func file_commons_messages_proto_rawDescGZIP() []byte {
	file_commons_messages_proto_rawDescOnce.Do(func() {
		file_commons_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_commons_messages_proto_rawDescData)
	})
	return file_commons_messages_proto_rawDescData
}

var file_commons_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_commons_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_commons_messages_proto_goTypes = []interface{}{
	(Null)(0),                // 0: rabex.api.commons.Null
	(*TargetUser)(nil),       // 1: rabex.api.commons.TargetUser
	(*InternalRequest)(nil),  // 2: rabex.api.commons.InternalRequest
	(*InternalResponse)(nil), // 3: rabex.api.commons.InternalResponse
	(*GeneralResponse)(nil),  // 4: rabex.api.commons.GeneralResponse
	(*Array)(nil),            // 5: rabex.api.commons.Array
	(*Value)(nil),            // 6: rabex.api.commons.Value
	(*Struct)(nil),           // 7: rabex.api.commons.Struct
	(*Error)(nil),            // 8: rabex.api.commons.Error
	(*AssetList)(nil),        // 9: rabex.api.commons.AssetList
	nil,                      // 10: rabex.api.commons.Struct.FieldsEntry
	(*anypb.Any)(nil),        // 11: google.protobuf.Any
	(Assets)(0),              // 12: rabex.api.commons.Assets
}
var file_commons_messages_proto_depIdxs = []int32{
	7,  // 0: rabex.api.commons.InternalRequest.payload:type_name -> rabex.api.commons.Struct
	7,  // 1: rabex.api.commons.InternalResponse.payload:type_name -> rabex.api.commons.Struct
	5,  // 2: rabex.api.commons.GeneralResponse.data:type_name -> rabex.api.commons.Array
	6,  // 3: rabex.api.commons.Array.value:type_name -> rabex.api.commons.Value
	0,  // 4: rabex.api.commons.Value.null_value:type_name -> rabex.api.commons.Null
	7,  // 5: rabex.api.commons.Value.struct_value:type_name -> rabex.api.commons.Struct
	11, // 6: rabex.api.commons.Value.any_value:type_name -> google.protobuf.Any
	5,  // 7: rabex.api.commons.Value.array_value:type_name -> rabex.api.commons.Array
	10, // 8: rabex.api.commons.Struct.Fields:type_name -> rabex.api.commons.Struct.FieldsEntry
	12, // 9: rabex.api.commons.AssetList.assetList:type_name -> rabex.api.commons.Assets
	6,  // 10: rabex.api.commons.Struct.FieldsEntry.value:type_name -> rabex.api.commons.Value
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_commons_messages_proto_init() }
func file_commons_messages_proto_init() {
	if File_commons_messages_proto != nil {
		return
	}
	file_commons_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commons_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetUser); i {
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
		file_commons_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalRequest); i {
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
		file_commons_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalResponse); i {
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
		file_commons_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralResponse); i {
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
		file_commons_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Array); i {
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
		file_commons_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_commons_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Struct); i {
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
		file_commons_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_commons_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetList); i {
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
	file_commons_messages_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Value_NullValue)(nil),
		(*Value_NumberValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_StructValue)(nil),
		(*Value_IntValue)(nil),
		(*Value_Int32Value)(nil),
		(*Value_Float32Value)(nil),
		(*Value_Float64Value)(nil),
		(*Value_AnyValue)(nil),
		(*Value_ArrayValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commons_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commons_messages_proto_goTypes,
		DependencyIndexes: file_commons_messages_proto_depIdxs,
		EnumInfos:         file_commons_messages_proto_enumTypes,
		MessageInfos:      file_commons_messages_proto_msgTypes,
	}.Build()
	File_commons_messages_proto = out.File
	file_commons_messages_proto_rawDesc = nil
	file_commons_messages_proto_goTypes = nil
	file_commons_messages_proto_depIdxs = nil
}
