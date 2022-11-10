// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: fiatpb/enums.proto

package fiatpb

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

type PaymentService int32

const (
	PaymentService_VANDAR  PaymentService = 0
	PaymentService_AYANDEH PaymentService = 1
)

// Enum value maps for PaymentService.
var (
	PaymentService_name = map[int32]string{
		0: "VANDAR",
		1: "AYANDEH",
	}
	PaymentService_value = map[string]int32{
		"VANDAR":  0,
		"AYANDEH": 1,
	}
)

func (x PaymentService) Enum() *PaymentService {
	p := new(PaymentService)
	*p = x
	return p
}

func (x PaymentService) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentService) Descriptor() protoreflect.EnumDescriptor {
	return file_fiatpb_enums_proto_enumTypes[0].Descriptor()
}

func (PaymentService) Type() protoreflect.EnumType {
	return &file_fiatpb_enums_proto_enumTypes[0]
}

func (x PaymentService) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentService.Descriptor instead.
func (PaymentService) EnumDescriptor() ([]byte, []int) {
	return file_fiatpb_enums_proto_rawDescGZIP(), []int{0}
}

type VandarBusinesses int32

const (
	VandarBusinesses_RABIN    VandarBusinesses = 0
	VandarBusinesses_RABEXPRO VandarBusinesses = 1
)

// Enum value maps for VandarBusinesses.
var (
	VandarBusinesses_name = map[int32]string{
		0: "RABIN",
		1: "RABEXPRO",
	}
	VandarBusinesses_value = map[string]int32{
		"RABIN":    0,
		"RABEXPRO": 1,
	}
)

func (x VandarBusinesses) Enum() *VandarBusinesses {
	p := new(VandarBusinesses)
	*p = x
	return p
}

func (x VandarBusinesses) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VandarBusinesses) Descriptor() protoreflect.EnumDescriptor {
	return file_fiatpb_enums_proto_enumTypes[1].Descriptor()
}

func (VandarBusinesses) Type() protoreflect.EnumType {
	return &file_fiatpb_enums_proto_enumTypes[1]
}

func (x VandarBusinesses) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VandarBusinesses.Descriptor instead.
func (VandarBusinesses) EnumDescriptor() ([]byte, []int) {
	return file_fiatpb_enums_proto_rawDescGZIP(), []int{1}
}

type DepositDeliveryStatus int32

const (
	DepositDeliveryStatus_Unspecified     DepositDeliveryStatus = 0
	DepositDeliveryStatus_SuccessFinished DepositDeliveryStatus = 1
	DepositDeliveryStatus_SuccessWait     DepositDeliveryStatus = 2
	DepositDeliveryStatus_Failed          DepositDeliveryStatus = 4
)

// Enum value maps for DepositDeliveryStatus.
var (
	DepositDeliveryStatus_name = map[int32]string{
		0: "Unspecified",
		1: "SuccessFinished",
		2: "SuccessWait",
		4: "Failed",
	}
	DepositDeliveryStatus_value = map[string]int32{
		"Unspecified":     0,
		"SuccessFinished": 1,
		"SuccessWait":     2,
		"Failed":          4,
	}
)

func (x DepositDeliveryStatus) Enum() *DepositDeliveryStatus {
	p := new(DepositDeliveryStatus)
	*p = x
	return p
}

func (x DepositDeliveryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DepositDeliveryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_fiatpb_enums_proto_enumTypes[2].Descriptor()
}

func (DepositDeliveryStatus) Type() protoreflect.EnumType {
	return &file_fiatpb_enums_proto_enumTypes[2]
}

func (x DepositDeliveryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DepositDeliveryStatus.Descriptor instead.
func (DepositDeliveryStatus) EnumDescriptor() ([]byte, []int) {
	return file_fiatpb_enums_proto_rawDescGZIP(), []int{2}
}

var File_fiatpb_enums_proto protoreflect.FileDescriptor

var file_fiatpb_enums_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x69, 0x61, 0x74, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x66, 0x69, 0x61, 0x74, 0x2a, 0x29, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x41, 0x4e, 0x44, 0x41, 0x52,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x59, 0x41, 0x4e, 0x44, 0x45, 0x48, 0x10, 0x01, 0x2a,
	0x2b, 0x0a, 0x10, 0x56, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x41, 0x42, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x52, 0x41, 0x42, 0x45, 0x58, 0x50, 0x52, 0x4f, 0x10, 0x01, 0x2a, 0x5a, 0x0a, 0x15,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x61, 0x69, 0x74, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x04, 0x42, 0x1c, 0x5a, 0x1a, 0x72, 0x61, 0x62, 0x65,
	0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x70, 0x62, 0x3b,
	0x66, 0x69, 0x61, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fiatpb_enums_proto_rawDescOnce sync.Once
	file_fiatpb_enums_proto_rawDescData = file_fiatpb_enums_proto_rawDesc
)

func file_fiatpb_enums_proto_rawDescGZIP() []byte {
	file_fiatpb_enums_proto_rawDescOnce.Do(func() {
		file_fiatpb_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_fiatpb_enums_proto_rawDescData)
	})
	return file_fiatpb_enums_proto_rawDescData
}

var file_fiatpb_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_fiatpb_enums_proto_goTypes = []interface{}{
	(PaymentService)(0),        // 0: rabex.api.fiat.PaymentService
	(VandarBusinesses)(0),      // 1: rabex.api.fiat.VandarBusinesses
	(DepositDeliveryStatus)(0), // 2: rabex.api.fiat.DepositDeliveryStatus
}
var file_fiatpb_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fiatpb_enums_proto_init() }
func file_fiatpb_enums_proto_init() {
	if File_fiatpb_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fiatpb_enums_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fiatpb_enums_proto_goTypes,
		DependencyIndexes: file_fiatpb_enums_proto_depIdxs,
		EnumInfos:         file_fiatpb_enums_proto_enumTypes,
	}.Build()
	File_fiatpb_enums_proto = out.File
	file_fiatpb_enums_proto_rawDesc = nil
	file_fiatpb_enums_proto_goTypes = nil
	file_fiatpb_enums_proto_depIdxs = nil
}
