// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: accounting/exchange/services.proto

package exchange

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_accounting_exchange_services_proto protoreflect.FileDescriptor

var file_accounting_exchange_services_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x22,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xec, 0x01, 0x0a, 0x0b, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x50, 0x49,
	0x12, 0x69, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x72, 0x0a, 0x08, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x12, 0x19, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x22, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x3a, 0x01, 0x2a, 0x42,
	0x2b, 0x5a, 0x29, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x3b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_accounting_exchange_services_proto_goTypes = []interface{}{
	(*ExchangeRequest)(nil),  // 0: exchange.ExchangeRequest
	(*ReferralRequest)(nil),  // 1: exchange.ReferralRequest
	(*ExchangeResponse)(nil), // 2: exchange.ExchangeResponse
	(*ReferralResponse)(nil), // 3: exchange.ReferralResponse
}
var file_accounting_exchange_services_proto_depIdxs = []int32{
	0, // 0: exchange.ExchangeAPI.Exchange:input_type -> exchange.ExchangeRequest
	1, // 1: exchange.ExchangeAPI.Referral:input_type -> exchange.ReferralRequest
	2, // 2: exchange.ExchangeAPI.Exchange:output_type -> exchange.ExchangeResponse
	3, // 3: exchange.ExchangeAPI.Referral:output_type -> exchange.ReferralResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_accounting_exchange_services_proto_init() }
func file_accounting_exchange_services_proto_init() {
	if File_accounting_exchange_services_proto != nil {
		return
	}
	file_accounting_exchange_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_accounting_exchange_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accounting_exchange_services_proto_goTypes,
		DependencyIndexes: file_accounting_exchange_services_proto_depIdxs,
	}.Build()
	File_accounting_exchange_services_proto = out.File
	file_accounting_exchange_services_proto_rawDesc = nil
	file_accounting_exchange_services_proto_goTypes = nil
	file_accounting_exchange_services_proto_depIdxs = nil
}
