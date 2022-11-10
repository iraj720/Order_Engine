// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: userpb/enums.proto

package userpb

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

type UserRole int32

const (
	UserRole_Guest UserRole = 0
	UserRole_Plain UserRole = 1
	UserRole_Admin UserRole = 2
)

// Enum value maps for UserRole.
var (
	UserRole_name = map[int32]string{
		0: "Guest",
		1: "Plain",
		2: "Admin",
	}
	UserRole_value = map[string]int32{
		"Guest": 0,
		"Plain": 1,
		"Admin": 2,
	}
)

func (x UserRole) Enum() *UserRole {
	p := new(UserRole)
	*p = x
	return p
}

func (x UserRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRole) Descriptor() protoreflect.EnumDescriptor {
	return file_userpb_enums_proto_enumTypes[0].Descriptor()
}

func (UserRole) Type() protoreflect.EnumType {
	return &file_userpb_enums_proto_enumTypes[0]
}

func (x UserRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRole.Descriptor instead.
func (UserRole) EnumDescriptor() ([]byte, []int) {
	return file_userpb_enums_proto_rawDescGZIP(), []int{0}
}

type TOTPChannel int32

const (
	TOTPChannel_SMS              TOTPChannel = 0
	TOTPChannel_EMAIL            TOTPChannel = 1
	TOTPChannel_AuthenticatorApp TOTPChannel = 2
	TOTPChannel_VoiceCall        TOTPChannel = 3
)

// Enum value maps for TOTPChannel.
var (
	TOTPChannel_name = map[int32]string{
		0: "SMS",
		1: "EMAIL",
		2: "AuthenticatorApp",
		3: "VoiceCall",
	}
	TOTPChannel_value = map[string]int32{
		"SMS":              0,
		"EMAIL":            1,
		"AuthenticatorApp": 2,
		"VoiceCall":        3,
	}
)

func (x TOTPChannel) Enum() *TOTPChannel {
	p := new(TOTPChannel)
	*p = x
	return p
}

func (x TOTPChannel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TOTPChannel) Descriptor() protoreflect.EnumDescriptor {
	return file_userpb_enums_proto_enumTypes[1].Descriptor()
}

func (TOTPChannel) Type() protoreflect.EnumType {
	return &file_userpb_enums_proto_enumTypes[1]
}

func (x TOTPChannel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TOTPChannel.Descriptor instead.
func (TOTPChannel) EnumDescriptor() ([]byte, []int) {
	return file_userpb_enums_proto_rawDescGZIP(), []int{1}
}

type TOTPChannelStatus int32

const (
	TOTPChannelStatus_TOTPCreated     TOTPChannelStatus = 0
	TOTPChannelStatus_TOTPBlocked     TOTPChannelStatus = 1
	TOTPChannelStatus_TOTPWaiting     TOTPChannelStatus = 2
	TOTPChannelStatus_TOTPUnreachable TOTPChannelStatus = 3
	TOTPChannelStatus_TOTPVerified    TOTPChannelStatus = 4
)

// Enum value maps for TOTPChannelStatus.
var (
	TOTPChannelStatus_name = map[int32]string{
		0: "TOTPCreated",
		1: "TOTPBlocked",
		2: "TOTPWaiting",
		3: "TOTPUnreachable",
		4: "TOTPVerified",
	}
	TOTPChannelStatus_value = map[string]int32{
		"TOTPCreated":     0,
		"TOTPBlocked":     1,
		"TOTPWaiting":     2,
		"TOTPUnreachable": 3,
		"TOTPVerified":    4,
	}
)

func (x TOTPChannelStatus) Enum() *TOTPChannelStatus {
	p := new(TOTPChannelStatus)
	*p = x
	return p
}

func (x TOTPChannelStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TOTPChannelStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_userpb_enums_proto_enumTypes[2].Descriptor()
}

func (TOTPChannelStatus) Type() protoreflect.EnumType {
	return &file_userpb_enums_proto_enumTypes[2]
}

func (x TOTPChannelStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TOTPChannelStatus.Descriptor instead.
func (TOTPChannelStatus) EnumDescriptor() ([]byte, []int) {
	return file_userpb_enums_proto_rawDescGZIP(), []int{2}
}

type UserPermission int32

const (
	UserPermission_NoPermission   UserPermission = 0
	UserPermission_Account        UserPermission = 1
	UserPermission_FiatWithdraw   UserPermission = 2
	UserPermission_CryptoWithdraw UserPermission = 4
	UserPermission_FiatDeposit    UserPermission = 8
	UserPermission_CryptoDeposit  UserPermission = 16
	UserPermission_Exchange       UserPermission = 32
)

// Enum value maps for UserPermission.
var (
	UserPermission_name = map[int32]string{
		0:  "NoPermission",
		1:  "Account",
		2:  "FiatWithdraw",
		4:  "CryptoWithdraw",
		8:  "FiatDeposit",
		16: "CryptoDeposit",
		32: "Exchange",
	}
	UserPermission_value = map[string]int32{
		"NoPermission":   0,
		"Account":        1,
		"FiatWithdraw":   2,
		"CryptoWithdraw": 4,
		"FiatDeposit":    8,
		"CryptoDeposit":  16,
		"Exchange":       32,
	}
)

func (x UserPermission) Enum() *UserPermission {
	p := new(UserPermission)
	*p = x
	return p
}

func (x UserPermission) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserPermission) Descriptor() protoreflect.EnumDescriptor {
	return file_userpb_enums_proto_enumTypes[3].Descriptor()
}

func (UserPermission) Type() protoreflect.EnumType {
	return &file_userpb_enums_proto_enumTypes[3]
}

func (x UserPermission) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserPermission.Descriptor instead.
func (UserPermission) EnumDescriptor() ([]byte, []int) {
	return file_userpb_enums_proto_rawDescGZIP(), []int{3}
}

type TaskType int32

const (
	TaskType_TOTP TaskType = 0
)

// Enum value maps for TaskType.
var (
	TaskType_name = map[int32]string{
		0: "TOTP",
	}
	TaskType_value = map[string]int32{
		"TOTP": 0,
	}
)

func (x TaskType) Enum() *TaskType {
	p := new(TaskType)
	*p = x
	return p
}

func (x TaskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskType) Descriptor() protoreflect.EnumDescriptor {
	return file_userpb_enums_proto_enumTypes[4].Descriptor()
}

func (TaskType) Type() protoreflect.EnumType {
	return &file_userpb_enums_proto_enumTypes[4]
}

func (x TaskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskType.Descriptor instead.
func (TaskType) EnumDescriptor() ([]byte, []int) {
	return file_userpb_enums_proto_rawDescGZIP(), []int{4}
}

type TaskStatus int32

const (
	TaskStatus_Todo  TaskStatus = 0
	TaskStatus_Doing TaskStatus = 1
	TaskStatus_Done  TaskStatus = 2
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "Todo",
		1: "Doing",
		2: "Done",
	}
	TaskStatus_value = map[string]int32{
		"Todo":  0,
		"Doing": 1,
		"Done":  2,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_userpb_enums_proto_enumTypes[5].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_userpb_enums_proto_enumTypes[5]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_userpb_enums_proto_rawDescGZIP(), []int{5}
}

type GroupType int32

const (
	GroupType_Nobody              GroupType = 0
	GroupType_KYCAgents           GroupType = 1
	GroupType_MoneyTransferAgents GroupType = 2
	GroupType_TechnicalAgents     GroupType = 3
	GroupType_MoneyMonitoring     GroupType = 4
	GroupType_SuperAdmin          GroupType = 5
)

// Enum value maps for GroupType.
var (
	GroupType_name = map[int32]string{
		0: "Nobody",
		1: "KYCAgents",
		2: "MoneyTransferAgents",
		3: "TechnicalAgents",
		4: "MoneyMonitoring",
		5: "SuperAdmin",
	}
	GroupType_value = map[string]int32{
		"Nobody":              0,
		"KYCAgents":           1,
		"MoneyTransferAgents": 2,
		"TechnicalAgents":     3,
		"MoneyMonitoring":     4,
		"SuperAdmin":          5,
	}
)

func (x GroupType) Enum() *GroupType {
	p := new(GroupType)
	*p = x
	return p
}

func (x GroupType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupType) Descriptor() protoreflect.EnumDescriptor {
	return file_userpb_enums_proto_enumTypes[6].Descriptor()
}

func (GroupType) Type() protoreflect.EnumType {
	return &file_userpb_enums_proto_enumTypes[6]
}

func (x GroupType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupType.Descriptor instead.
func (GroupType) EnumDescriptor() ([]byte, []int) {
	return file_userpb_enums_proto_rawDescGZIP(), []int{6}
}

type UserAnomaly int32

const (
	UserAnomaly_WhiteListed UserAnomaly = 0
	UserAnomaly_GrayListed  UserAnomaly = 1
	UserAnomaly_BlackListed UserAnomaly = 2
)

// Enum value maps for UserAnomaly.
var (
	UserAnomaly_name = map[int32]string{
		0: "WhiteListed",
		1: "GrayListed",
		2: "BlackListed",
	}
	UserAnomaly_value = map[string]int32{
		"WhiteListed": 0,
		"GrayListed":  1,
		"BlackListed": 2,
	}
)

func (x UserAnomaly) Enum() *UserAnomaly {
	p := new(UserAnomaly)
	*p = x
	return p
}

func (x UserAnomaly) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserAnomaly) Descriptor() protoreflect.EnumDescriptor {
	return file_userpb_enums_proto_enumTypes[7].Descriptor()
}

func (UserAnomaly) Type() protoreflect.EnumType {
	return &file_userpb_enums_proto_enumTypes[7]
}

func (x UserAnomaly) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserAnomaly.Descriptor instead.
func (UserAnomaly) EnumDescriptor() ([]byte, []int) {
	return file_userpb_enums_proto_rawDescGZIP(), []int{7}
}

var File_userpb_enums_proto protoreflect.FileDescriptor

var file_userpb_enums_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2a, 0x2b, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x09, 0x0a, 0x05, 0x47, 0x75, 0x65, 0x73, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x50,
	0x6c, 0x61, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x10,
	0x02, 0x2a, 0x46, 0x0a, 0x0b, 0x54, 0x4f, 0x54, 0x50, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x07, 0x0a, 0x03, 0x53, 0x4d, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41,
	0x49, 0x4c, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x70, 0x70, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x10, 0x03, 0x2a, 0x6d, 0x0a, 0x11, 0x54, 0x4f, 0x54,
	0x50, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f,
	0x0a, 0x0b, 0x54, 0x4f, 0x54, 0x50, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x54, 0x4f, 0x54, 0x50, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x4f, 0x54, 0x50, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x10,
	0x02, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x4f, 0x54, 0x50, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x63, 0x68,
	0x61, 0x62, 0x6c, 0x65, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x4f, 0x54, 0x50, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x04, 0x2a, 0x87, 0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x0c, 0x4e,
	0x6f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x69,
	0x61, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x10, 0x04,
	0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x69, 0x61, 0x74, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x10,
	0x08, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x10, 0x20, 0x2a, 0x14, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x54, 0x4f, 0x54, 0x50, 0x10, 0x00, 0x2a, 0x2b, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x44, 0x6f, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44,
	0x6f, 0x6e, 0x65, 0x10, 0x02, 0x2a, 0x79, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f, 0x62, 0x6f, 0x64, 0x79, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x4b, 0x59, 0x43, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x73, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x69,
	0x63, 0x61, 0x6c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x04,
	0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x75, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x05,
	0x2a, 0x3f, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x12,
	0x0f, 0x0a, 0x0b, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x72, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x10,
	0x02, 0x42, 0x1c, 0x5a, 0x1a, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userpb_enums_proto_rawDescOnce sync.Once
	file_userpb_enums_proto_rawDescData = file_userpb_enums_proto_rawDesc
)

func file_userpb_enums_proto_rawDescGZIP() []byte {
	file_userpb_enums_proto_rawDescOnce.Do(func() {
		file_userpb_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_userpb_enums_proto_rawDescData)
	})
	return file_userpb_enums_proto_rawDescData
}

var file_userpb_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 8)
var file_userpb_enums_proto_goTypes = []interface{}{
	(UserRole)(0),          // 0: rabex.api.user.UserRole
	(TOTPChannel)(0),       // 1: rabex.api.user.TOTPChannel
	(TOTPChannelStatus)(0), // 2: rabex.api.user.TOTPChannelStatus
	(UserPermission)(0),    // 3: rabex.api.user.UserPermission
	(TaskType)(0),          // 4: rabex.api.user.TaskType
	(TaskStatus)(0),        // 5: rabex.api.user.TaskStatus
	(GroupType)(0),         // 6: rabex.api.user.GroupType
	(UserAnomaly)(0),       // 7: rabex.api.user.UserAnomaly
}
var file_userpb_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_userpb_enums_proto_init() }
func file_userpb_enums_proto_init() {
	if File_userpb_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_userpb_enums_proto_rawDesc,
			NumEnums:      8,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_userpb_enums_proto_goTypes,
		DependencyIndexes: file_userpb_enums_proto_depIdxs,
		EnumInfos:         file_userpb_enums_proto_enumTypes,
	}.Build()
	File_userpb_enums_proto = out.File
	file_userpb_enums_proto_rawDesc = nil
	file_userpb_enums_proto_goTypes = nil
	file_userpb_enums_proto_depIdxs = nil
}