// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: usermnm/enums.proto

package usermnm

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "rabex/api/pb/commons"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ActionTypePermission int32

const (
	ActionTypePermission_CREATE ActionTypePermission = 0
	ActionTypePermission_READ   ActionTypePermission = 1
	ActionTypePermission_UPDATE ActionTypePermission = 2
	ActionTypePermission_DELETE ActionTypePermission = 3
)

// Enum value maps for ActionTypePermission.
var (
	ActionTypePermission_name = map[int32]string{
		0: "CREATE",
		1: "READ",
		2: "UPDATE",
		3: "DELETE",
	}
	ActionTypePermission_value = map[string]int32{
		"CREATE": 0,
		"READ":   1,
		"UPDATE": 2,
		"DELETE": 3,
	}
)

func (x ActionTypePermission) Enum() *ActionTypePermission {
	p := new(ActionTypePermission)
	*p = x
	return p
}

func (x ActionTypePermission) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionTypePermission) Descriptor() protoreflect.EnumDescriptor {
	return file_usermnm_enums_proto_enumTypes[0].Descriptor()
}

func (ActionTypePermission) Type() protoreflect.EnumType {
	return &file_usermnm_enums_proto_enumTypes[0]
}

func (x ActionTypePermission) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionTypePermission.Descriptor instead.
func (ActionTypePermission) EnumDescriptor() ([]byte, []int) {
	return file_usermnm_enums_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_Nothing         Status = 0
	Status_WaitingForAdmin Status = 1
	Status_Rejected        Status = 2
	Status_Confirmed       Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Nothing",
		1: "WaitingForAdmin",
		2: "Rejected",
		3: "Confirmed",
	}
	Status_value = map[string]int32{
		"Nothing":         0,
		"WaitingForAdmin": 1,
		"Rejected":        2,
		"Confirmed":       3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_usermnm_enums_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_usermnm_enums_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_usermnm_enums_proto_rawDescGZIP(), []int{1}
}

var File_usermnm_enums_proto protoreflect.FileDescriptor

var file_usermnm_enums_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6e, 0x6d, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6e, 0x6d, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x44, 0x0a,
	0x14, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x45, 0x10, 0x03, 0x2a, 0x47, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x61,
	0x69, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x10, 0x03, 0x42, 0x30, 0x5a, 0x1c,
	0x72, 0x61, 0x62, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x6d, 0x6e, 0x6d, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x6e, 0x6d, 0x92, 0x41, 0x0f, 0x1a,
	0x0d, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x36, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usermnm_enums_proto_rawDescOnce sync.Once
	file_usermnm_enums_proto_rawDescData = file_usermnm_enums_proto_rawDesc
)

func file_usermnm_enums_proto_rawDescGZIP() []byte {
	file_usermnm_enums_proto_rawDescOnce.Do(func() {
		file_usermnm_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_usermnm_enums_proto_rawDescData)
	})
	return file_usermnm_enums_proto_rawDescData
}

var file_usermnm_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_usermnm_enums_proto_goTypes = []interface{}{
	(ActionTypePermission)(0), // 0: rabex.api.usermnm.ActionTypePermission
	(Status)(0),               // 1: rabex.api.usermnm.Status
}
var file_usermnm_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_usermnm_enums_proto_init() }
func file_usermnm_enums_proto_init() {
	if File_usermnm_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_usermnm_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_usermnm_enums_proto_goTypes,
		DependencyIndexes: file_usermnm_enums_proto_depIdxs,
		EnumInfos:         file_usermnm_enums_proto_enumTypes,
	}.Build()
	File_usermnm_enums_proto = out.File
	file_usermnm_enums_proto_rawDesc = nil
	file_usermnm_enums_proto_goTypes = nil
	file_usermnm_enums_proto_depIdxs = nil
}
