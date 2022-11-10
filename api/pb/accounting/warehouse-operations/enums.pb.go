// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: accounting/warehouse-operations/enums.proto

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

type Action int32

const (
	Action_ACTION_UNSPECIFIED Action = 0
	Action_ACTION_CREDIT      Action = 1
	Action_ACTION_CONFISCATE  Action = 2
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "ACTION_UNSPECIFIED",
		1: "ACTION_CREDIT",
		2: "ACTION_CONFISCATE",
	}
	Action_value = map[string]int32{
		"ACTION_UNSPECIFIED": 0,
		"ACTION_CREDIT":      1,
		"ACTION_CONFISCATE":  2,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_accounting_warehouse_operations_enums_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_accounting_warehouse_operations_enums_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_accounting_warehouse_operations_enums_proto_rawDescGZIP(), []int{0}
}

var File_accounting_warehouse_operations_enums_proto protoreflect.FileDescriptor

var file_accounting_warehouse_operations_enums_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x77, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4f, 0x70, 0x73, 0x2a, 0x4a, 0x0a, 0x06, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x10, 0x01,
	0x12, 0x15, 0x0a, 0x11, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x53, 0x43, 0x41, 0x54, 0x45, 0x10, 0x02, 0x42, 0x48, 0x5a, 0x46, 0x72, 0x61, 0x62, 0x65, 0x78,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x4f, 0x70, 0x73, 0x3b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4f, 0x70,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounting_warehouse_operations_enums_proto_rawDescOnce sync.Once
	file_accounting_warehouse_operations_enums_proto_rawDescData = file_accounting_warehouse_operations_enums_proto_rawDesc
)

func file_accounting_warehouse_operations_enums_proto_rawDescGZIP() []byte {
	file_accounting_warehouse_operations_enums_proto_rawDescOnce.Do(func() {
		file_accounting_warehouse_operations_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounting_warehouse_operations_enums_proto_rawDescData)
	})
	return file_accounting_warehouse_operations_enums_proto_rawDescData
}

var file_accounting_warehouse_operations_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_accounting_warehouse_operations_enums_proto_goTypes = []interface{}{
	(Action)(0), // 0: warehouseOps.Action
}
var file_accounting_warehouse_operations_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_accounting_warehouse_operations_enums_proto_init() }
func file_accounting_warehouse_operations_enums_proto_init() {
	if File_accounting_warehouse_operations_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_accounting_warehouse_operations_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accounting_warehouse_operations_enums_proto_goTypes,
		DependencyIndexes: file_accounting_warehouse_operations_enums_proto_depIdxs,
		EnumInfos:         file_accounting_warehouse_operations_enums_proto_enumTypes,
	}.Build()
	File_accounting_warehouse_operations_enums_proto = out.File
	file_accounting_warehouse_operations_enums_proto_rawDesc = nil
	file_accounting_warehouse_operations_enums_proto_goTypes = nil
	file_accounting_warehouse_operations_enums_proto_depIdxs = nil
}
