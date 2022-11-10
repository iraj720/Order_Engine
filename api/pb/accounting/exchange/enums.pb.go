// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: accounting/exchange/enums.proto

package exchange

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

type ExchangeAction int32

const (
	ExchangeAction_EXCHANGE_ACTION_UNSPECIFIED   ExchangeAction = 0
	ExchangeAction_EXCHANGE_ACTION_INIT          ExchangeAction = 1
	ExchangeAction_EXCHANGE_ACTION_FINALIZE      ExchangeAction = 2
	ExchangeAction_EXCHANGE_ACTION_CANCEL        ExchangeAction = 3
	ExchangeAction_EXCHANGE_ACTION_INIT_FINALIZE ExchangeAction = 4
)

// Enum value maps for ExchangeAction.
var (
	ExchangeAction_name = map[int32]string{
		0: "EXCHANGE_ACTION_UNSPECIFIED",
		1: "EXCHANGE_ACTION_INIT",
		2: "EXCHANGE_ACTION_FINALIZE",
		3: "EXCHANGE_ACTION_CANCEL",
		4: "EXCHANGE_ACTION_INIT_FINALIZE",
	}
	ExchangeAction_value = map[string]int32{
		"EXCHANGE_ACTION_UNSPECIFIED":   0,
		"EXCHANGE_ACTION_INIT":          1,
		"EXCHANGE_ACTION_FINALIZE":      2,
		"EXCHANGE_ACTION_CANCEL":        3,
		"EXCHANGE_ACTION_INIT_FINALIZE": 4,
	}
)

func (x ExchangeAction) Enum() *ExchangeAction {
	p := new(ExchangeAction)
	*p = x
	return p
}

func (x ExchangeAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExchangeAction) Descriptor() protoreflect.EnumDescriptor {
	return file_accounting_exchange_enums_proto_enumTypes[0].Descriptor()
}

func (ExchangeAction) Type() protoreflect.EnumType {
	return &file_accounting_exchange_enums_proto_enumTypes[0]
}

func (x ExchangeAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExchangeAction.Descriptor instead.
func (ExchangeAction) EnumDescriptor() ([]byte, []int) {
	return file_accounting_exchange_enums_proto_rawDescGZIP(), []int{0}
}

type ReferralAction int32

const (
	ReferralAction_REFERRAL_ACTION_UNSPECIFIED ReferralAction = 0
	ReferralAction_REFERRAL_ACTION_FINALIZE    ReferralAction = 1
)

// Enum value maps for ReferralAction.
var (
	ReferralAction_name = map[int32]string{
		0: "REFERRAL_ACTION_UNSPECIFIED",
		1: "REFERRAL_ACTION_FINALIZE",
	}
	ReferralAction_value = map[string]int32{
		"REFERRAL_ACTION_UNSPECIFIED": 0,
		"REFERRAL_ACTION_FINALIZE":    1,
	}
)

func (x ReferralAction) Enum() *ReferralAction {
	p := new(ReferralAction)
	*p = x
	return p
}

func (x ReferralAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReferralAction) Descriptor() protoreflect.EnumDescriptor {
	return file_accounting_exchange_enums_proto_enumTypes[1].Descriptor()
}

func (ReferralAction) Type() protoreflect.EnumType {
	return &file_accounting_exchange_enums_proto_enumTypes[1]
}

func (x ReferralAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReferralAction.Descriptor instead.
func (ReferralAction) EnumDescriptor() ([]byte, []int) {
	return file_accounting_exchange_enums_proto_rawDescGZIP(), []int{1}
}

var File_accounting_exchange_enums_proto protoreflect.FileDescriptor

var file_accounting_exchange_enums_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2a, 0xa8, 0x01, 0x0a, 0x0e,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x1b, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x18, 0x0a, 0x14, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x58, 0x43,
	0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x4e,
	0x41, 0x4c, 0x49, 0x5a, 0x45, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x58, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x5f, 0x46, 0x49, 0x4e, 0x41,
	0x4c, 0x49, 0x5a, 0x45, 0x10, 0x04, 0x2a, 0x4f, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x46, 0x45,
	0x52, 0x52, 0x41, 0x4c, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x46,
	0x45, 0x52, 0x52, 0x41, 0x4c, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x4e,
	0x41, 0x4c, 0x49, 0x5a, 0x45, 0x10, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x72, 0x61, 0x62, 0x65, 0x78,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x3b, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounting_exchange_enums_proto_rawDescOnce sync.Once
	file_accounting_exchange_enums_proto_rawDescData = file_accounting_exchange_enums_proto_rawDesc
)

func file_accounting_exchange_enums_proto_rawDescGZIP() []byte {
	file_accounting_exchange_enums_proto_rawDescOnce.Do(func() {
		file_accounting_exchange_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounting_exchange_enums_proto_rawDescData)
	})
	return file_accounting_exchange_enums_proto_rawDescData
}

var file_accounting_exchange_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_accounting_exchange_enums_proto_goTypes = []interface{}{
	(ExchangeAction)(0), // 0: exchange.ExchangeAction
	(ReferralAction)(0), // 1: exchange.ReferralAction
}
var file_accounting_exchange_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_accounting_exchange_enums_proto_init() }
func file_accounting_exchange_enums_proto_init() {
	if File_accounting_exchange_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_accounting_exchange_enums_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accounting_exchange_enums_proto_goTypes,
		DependencyIndexes: file_accounting_exchange_enums_proto_depIdxs,
		EnumInfos:         file_accounting_exchange_enums_proto_enumTypes,
	}.Build()
	File_accounting_exchange_enums_proto = out.File
	file_accounting_exchange_enums_proto_rawDesc = nil
	file_accounting_exchange_enums_proto_goTypes = nil
	file_accounting_exchange_enums_proto_depIdxs = nil
}