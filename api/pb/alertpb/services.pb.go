// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: alertpb/services.proto

package alertpb

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

var File_alertpb_services_proto protoreflect.FileDescriptor

var file_alertpb_services_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x1a, 0x16, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xba, 0x06, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x81, 0x01, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x25, 0x2e, 0x72, 0x61,
	0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x74, 0x63, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x90, 0x01, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x77, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x42, 0x79,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x2b, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x42, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x21, 0x22, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x74, 0x63, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0x81, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x12, 0x25, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x61, 0x62, 0x65,
	0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x74, 0x63,
	0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x81, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x25, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70,
	0x62, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x74, 0x63, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x88, 0x01, 0x0a, 0x10, 0x53,
	0x68, 0x6f, 0x77, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x2a, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x61,
	0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x73, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x74, 0x63, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x86, 0x01, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x77, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x24, 0x2e, 0x72, 0x61, 0x62, 0x65, 0x78,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x74, 0x63, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x1e,
	0x5a, 0x1c, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x3b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_alertpb_services_proto_goTypes = []interface{}{
	(*CreateAlertRequest)(nil),       // 0: rabex.api.alertpb.CreateAlertRequest
	(*ListAlertsByAdminRequest)(nil), // 1: rabex.api.alertpb.ListAlertsByAdminRequest
	(*RemoveAlertRequest)(nil),       // 2: rabex.api.alertpb.RemoveAlertRequest
	(*UpdateAlertRequest)(nil),       // 3: rabex.api.alertpb.UpdateAlertRequest
	(*ListAlertsByUserRequest)(nil),  // 4: rabex.api.alertpb.ListAlertsByUserRequest
	(*UserGroupsRequest)(nil),        // 5: rabex.api.alertpb.UserGroupsRequest
	(*AlertResponse)(nil),            // 6: rabex.api.alertpb.AlertResponse
	(*ListAlertResponses)(nil),       // 7: rabex.api.alertpb.ListAlertResponses
	(*UserGroupsResponse)(nil),       // 8: rabex.api.alertpb.UserGroupsResponse
}
var file_alertpb_services_proto_depIdxs = []int32{
	0, // 0: rabex.api.alertpb.Alert.CreateAlert:input_type -> rabex.api.alertpb.CreateAlertRequest
	1, // 1: rabex.api.alertpb.Alert.ShowAlertsByAdmin:input_type -> rabex.api.alertpb.ListAlertsByAdminRequest
	2, // 2: rabex.api.alertpb.Alert.RemoveAlert:input_type -> rabex.api.alertpb.RemoveAlertRequest
	3, // 3: rabex.api.alertpb.Alert.UpdateAlert:input_type -> rabex.api.alertpb.UpdateAlertRequest
	4, // 4: rabex.api.alertpb.Alert.ShowAlertsByUser:input_type -> rabex.api.alertpb.ListAlertsByUserRequest
	5, // 5: rabex.api.alertpb.Alert.ShowUserGroups:input_type -> rabex.api.alertpb.UserGroupsRequest
	6, // 6: rabex.api.alertpb.Alert.CreateAlert:output_type -> rabex.api.alertpb.AlertResponse
	7, // 7: rabex.api.alertpb.Alert.ShowAlertsByAdmin:output_type -> rabex.api.alertpb.ListAlertResponses
	6, // 8: rabex.api.alertpb.Alert.RemoveAlert:output_type -> rabex.api.alertpb.AlertResponse
	6, // 9: rabex.api.alertpb.Alert.UpdateAlert:output_type -> rabex.api.alertpb.AlertResponse
	7, // 10: rabex.api.alertpb.Alert.ShowAlertsByUser:output_type -> rabex.api.alertpb.ListAlertResponses
	8, // 11: rabex.api.alertpb.Alert.ShowUserGroups:output_type -> rabex.api.alertpb.UserGroupsResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_alertpb_services_proto_init() }
func file_alertpb_services_proto_init() {
	if File_alertpb_services_proto != nil {
		return
	}
	file_alertpb_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_alertpb_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alertpb_services_proto_goTypes,
		DependencyIndexes: file_alertpb_services_proto_depIdxs,
	}.Build()
	File_alertpb_services_proto = out.File
	file_alertpb_services_proto_rawDesc = nil
	file_alertpb_services_proto_goTypes = nil
	file_alertpb_services_proto_depIdxs = nil
}
