// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: treasury/messages.proto

package treasury

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

type WarehouseType int32

const (
	WarehouseType_Main    WarehouseType = 0
	WarehouseType_Reserve WarehouseType = 1
	WarehouseType_Pending WarehouseType = 2
	WarehouseType_Gas     WarehouseType = 3
)

// Enum value maps for WarehouseType.
var (
	WarehouseType_name = map[int32]string{
		0: "Main",
		1: "Reserve",
		2: "Pending",
		3: "Gas",
	}
	WarehouseType_value = map[string]int32{
		"Main":    0,
		"Reserve": 1,
		"Pending": 2,
		"Gas":     3,
	}
)

func (x WarehouseType) Enum() *WarehouseType {
	p := new(WarehouseType)
	*p = x
	return p
}

func (x WarehouseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WarehouseType) Descriptor() protoreflect.EnumDescriptor {
	return file_treasury_messages_proto_enumTypes[0].Descriptor()
}

func (WarehouseType) Type() protoreflect.EnumType {
	return &file_treasury_messages_proto_enumTypes[0]
}

func (x WarehouseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WarehouseType.Descriptor instead.
func (WarehouseType) EnumDescriptor() ([]byte, []int) {
	return file_treasury_messages_proto_rawDescGZIP(), []int{0}
}

var File_treasury_messages_proto protoreflect.FileDescriptor

var file_treasury_messages_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x61, 0x62, 0x65, 0x78,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x2a, 0x3c, 0x0a,
	0x0d, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x61, 0x73, 0x10, 0x03, 0x42, 0x20, 0x5a, 0x1e, 0x72,
	0x61, 0x62, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x79, 0x3b, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_treasury_messages_proto_rawDescOnce sync.Once
	file_treasury_messages_proto_rawDescData = file_treasury_messages_proto_rawDesc
)

func file_treasury_messages_proto_rawDescGZIP() []byte {
	file_treasury_messages_proto_rawDescOnce.Do(func() {
		file_treasury_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_treasury_messages_proto_rawDescData)
	})
	return file_treasury_messages_proto_rawDescData
}

var file_treasury_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_treasury_messages_proto_goTypes = []interface{}{
	(WarehouseType)(0), // 0: rabex.api.treasury.WarehouseType
}
var file_treasury_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_treasury_messages_proto_init() }
func file_treasury_messages_proto_init() {
	if File_treasury_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_treasury_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_treasury_messages_proto_goTypes,
		DependencyIndexes: file_treasury_messages_proto_depIdxs,
		EnumInfos:         file_treasury_messages_proto_enumTypes,
	}.Build()
	File_treasury_messages_proto = out.File
	file_treasury_messages_proto_rawDesc = nil
	file_treasury_messages_proto_goTypes = nil
	file_treasury_messages_proto_depIdxs = nil
}
