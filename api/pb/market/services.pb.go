// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: market/services.proto

package market

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_market_services_proto protoreflect.FileDescriptor

var file_market_services_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x1a, 0x15, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa7, 0x12, 0x0a, 0x06,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x9b, 0x01, 0x0a, 0x11, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x2e, 0x72,
	0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x61, 0x6e, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x22, 0x22, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x61, 0x6e, 0x69, 0x74,
	0x79, 0x3a, 0x01, 0x2a, 0x12, 0x86, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x72,
	0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x82, 0x01,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x24, 0x2e,
	0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x20, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0x7a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12,
	0x22, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x80,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f, 0x74, 0x63, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2c, 0x2e,
	0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x74, 0x63, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0xa8, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x74, 0x63, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x2e, 0x72, 0x61,
	0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x74, 0x63, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x72, 0x61,
	0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x74, 0x63, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7a, 0x0a, 0x09,
	0x41, 0x64, 0x64, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x22, 0x2e, 0x72, 0x61, 0x62, 0x65,
	0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x41, 0x64, 0x64,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2e, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x61, 0x64, 0x64, 0x2d, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x8c, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x72,
	0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x8b, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x12, 0x28, 0x2e, 0x72, 0x61,
	0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x63, 0x68, 0x61, 0x72,
	0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x8b, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x73, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2b, 0x2e, 0x72,
	0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x61, 0x62, 0x65,
	0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x25, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x3a, 0x01, 0x2a, 0x12, 0x8e, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2c, 0x2e, 0x72, 0x61,
	0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x61, 0x62, 0x65,
	0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x26, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x79, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x2c, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x8f, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x27, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x72, 0x61,
	0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x98, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2a, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x62, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x22, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x63, 0x0a, 0x05,
	0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e,
	0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4b, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x6b, 0x6c, 0x69, 0x6e, 0x65, 0x3a, 0x01, 0x2a, 0x30,
	0x01, 0x12, 0x81, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x66, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x27, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x66, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x66, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12,
	0x14, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x72, 0x65, 0x66, 0x2f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a, 0x11, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x2b, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x42, 0x2e, 0x5a, 0x1a, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x3b, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x92, 0x41, 0x0f, 0x1a, 0x0d, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e,
	0x31, 0x30, 0x30, 0x2e, 0x36, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_market_services_proto_goTypes = []interface{}{
	(*MarketPriceSanityRequest)(nil),      // 0: rabex.api.market.MarketPriceSanityRequest
	(*UpdateMarketRequest)(nil),           // 1: rabex.api.market.UpdateMarketRequest
	(*ListMarketsRequest)(nil),            // 2: rabex.api.market.ListMarketsRequest
	(*GetMarketRequest)(nil),              // 3: rabex.api.market.GetMarketRequest
	(*emptypb.Empty)(nil),                 // 4: google.protobuf.Empty
	(*UpdateOtcMarketStatusRequest)(nil),  // 5: rabex.api.market.UpdateOtcMarketStatusRequest
	(*AddMarketRequest)(nil),              // 6: rabex.api.market.AddMarketRequest
	(*GetMarketsPriceRequest)(nil),        // 7: rabex.api.market.GetMarketsPriceRequest
	(*GetMarketChartsRequest)(nil),        // 8: rabex.api.market.GetMarketChartsRequest
	(*UpdateBaseSelectorRequest)(nil),     // 9: rabex.api.market.UpdateBaseSelectorRequest
	(*UpdateQuoteSelectorRequest)(nil),    // 10: rabex.api.market.UpdateQuoteSelectorRequest
	(*UpdateSelectorRequest)(nil),         // 11: rabex.api.market.UpdateSelectorRequest
	(*UpdateSubSelectorRequest)(nil),      // 12: rabex.api.market.UpdateSubSelectorRequest
	(*MarketRefPriceRequest)(nil),         // 13: rabex.api.market.MarketRefPriceRequest
	(*MarketPriceSanityResponse)(nil),     // 14: rabex.api.market.MarketPriceSanityResponse
	(*UpdateMarketResponse)(nil),          // 15: rabex.api.market.UpdateMarketResponse
	(*ListMarketsResponse)(nil),           // 16: rabex.api.market.ListMarketsResponse
	(*GetMarketResponse)(nil),             // 17: rabex.api.market.GetMarketResponse
	(*GetOtcMarketStatusResponse)(nil),    // 18: rabex.api.market.GetOtcMarketStatusResponse
	(*UpdateOtcMarketStatusResponse)(nil), // 19: rabex.api.market.UpdateOtcMarketStatusResponse
	(*AddMarketResponse)(nil),             // 20: rabex.api.market.AddMarketResponse
	(*GetMarketsPriceResponse)(nil),       // 21: rabex.api.market.GetMarketsPriceResponse
	(*GetMarketChartsResponse)(nil),       // 22: rabex.api.market.GetMarketChartsResponse
	(*MarketSelectors)(nil),               // 23: rabex.api.market.MarketSelectors
	(*MarketSelectorsGetResponse)(nil),    // 24: rabex.api.market.MarketSelectorsGetResponse
	(*MarketSelectorsResponse)(nil),       // 25: rabex.api.market.MarketSelectorsResponse
	(*MarketKlineStream)(nil),             // 26: rabex.api.market.MarketKlineStream
	(*MarketRefPriceResponse)(nil),        // 27: rabex.api.market.MarketRefPriceResponse
	(*TickerHealthCheckResponse)(nil),     // 28: rabex.api.market.TickerHealthCheckResponse
}
var file_market_services_proto_depIdxs = []int32{
	0,  // 0: rabex.api.market.Market.MarketPriceSanity:input_type -> rabex.api.market.MarketPriceSanityRequest
	1,  // 1: rabex.api.market.Market.UpdateMarket:input_type -> rabex.api.market.UpdateMarketRequest
	2,  // 2: rabex.api.market.Market.ListMarkets:input_type -> rabex.api.market.ListMarketsRequest
	3,  // 3: rabex.api.market.Market.GetMarket:input_type -> rabex.api.market.GetMarketRequest
	4,  // 4: rabex.api.market.Market.GetOtcMarketStatus:input_type -> google.protobuf.Empty
	5,  // 5: rabex.api.market.Market.UpdateOtcMarketStatus:input_type -> rabex.api.market.UpdateOtcMarketStatusRequest
	6,  // 6: rabex.api.market.Market.AddMarket:input_type -> rabex.api.market.AddMarketRequest
	7,  // 7: rabex.api.market.Market.GetMarketsPrice:input_type -> rabex.api.market.GetMarketsPriceRequest
	8,  // 8: rabex.api.market.Market.GetMarketCharts:input_type -> rabex.api.market.GetMarketChartsRequest
	9,  // 9: rabex.api.market.Market.UpdateBaseSelector:input_type -> rabex.api.market.UpdateBaseSelectorRequest
	10, // 10: rabex.api.market.Market.UpdateQuoteSelector:input_type -> rabex.api.market.UpdateQuoteSelectorRequest
	4,  // 11: rabex.api.market.Market.GetAllSelectors:input_type -> google.protobuf.Empty
	11, // 12: rabex.api.market.Market.UpdateSelector:input_type -> rabex.api.market.UpdateSelectorRequest
	12, // 13: rabex.api.market.Market.UpdateSubSelector:input_type -> rabex.api.market.UpdateSubSelectorRequest
	4,  // 14: rabex.api.market.Market.Kline:input_type -> google.protobuf.Empty
	13, // 15: rabex.api.market.Market.MarketRefPrice:input_type -> rabex.api.market.MarketRefPriceRequest
	4,  // 16: rabex.api.market.Market.TickerHealthCheck:input_type -> google.protobuf.Empty
	14, // 17: rabex.api.market.Market.MarketPriceSanity:output_type -> rabex.api.market.MarketPriceSanityResponse
	15, // 18: rabex.api.market.Market.UpdateMarket:output_type -> rabex.api.market.UpdateMarketResponse
	16, // 19: rabex.api.market.Market.ListMarkets:output_type -> rabex.api.market.ListMarketsResponse
	17, // 20: rabex.api.market.Market.GetMarket:output_type -> rabex.api.market.GetMarketResponse
	18, // 21: rabex.api.market.Market.GetOtcMarketStatus:output_type -> rabex.api.market.GetOtcMarketStatusResponse
	19, // 22: rabex.api.market.Market.UpdateOtcMarketStatus:output_type -> rabex.api.market.UpdateOtcMarketStatusResponse
	20, // 23: rabex.api.market.Market.AddMarket:output_type -> rabex.api.market.AddMarketResponse
	21, // 24: rabex.api.market.Market.GetMarketsPrice:output_type -> rabex.api.market.GetMarketsPriceResponse
	22, // 25: rabex.api.market.Market.GetMarketCharts:output_type -> rabex.api.market.GetMarketChartsResponse
	23, // 26: rabex.api.market.Market.UpdateBaseSelector:output_type -> rabex.api.market.MarketSelectors
	23, // 27: rabex.api.market.Market.UpdateQuoteSelector:output_type -> rabex.api.market.MarketSelectors
	24, // 28: rabex.api.market.Market.GetAllSelectors:output_type -> rabex.api.market.MarketSelectorsGetResponse
	25, // 29: rabex.api.market.Market.UpdateSelector:output_type -> rabex.api.market.MarketSelectorsResponse
	25, // 30: rabex.api.market.Market.UpdateSubSelector:output_type -> rabex.api.market.MarketSelectorsResponse
	26, // 31: rabex.api.market.Market.Kline:output_type -> rabex.api.market.MarketKlineStream
	27, // 32: rabex.api.market.Market.MarketRefPrice:output_type -> rabex.api.market.MarketRefPriceResponse
	28, // 33: rabex.api.market.Market.TickerHealthCheck:output_type -> rabex.api.market.TickerHealthCheckResponse
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_market_services_proto_init() }
func file_market_services_proto_init() {
	if File_market_services_proto != nil {
		return
	}
	file_market_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_market_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_market_services_proto_goTypes,
		DependencyIndexes: file_market_services_proto_depIdxs,
	}.Build()
	File_market_services_proto = out.File
	file_market_services_proto_rawDesc = nil
	file_market_services_proto_goTypes = nil
	file_market_services_proto_depIdxs = nil
}