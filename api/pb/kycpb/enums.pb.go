// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: kycpb/enums.proto

package kycpb

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

type RendezvousMechanisms int32

const (
	RendezvousMechanisms_DefaultValue     RendezvousMechanisms = 0
	RendezvousMechanisms_PhoneCall        RendezvousMechanisms = 1
	RendezvousMechanisms_PhysicalPresence RendezvousMechanisms = 2
	RendezvousMechanisms_TrusteeEntity    RendezvousMechanisms = 3
)

// Enum value maps for RendezvousMechanisms.
var (
	RendezvousMechanisms_name = map[int32]string{
		0: "DefaultValue",
		1: "PhoneCall",
		2: "PhysicalPresence",
		3: "TrusteeEntity",
	}
	RendezvousMechanisms_value = map[string]int32{
		"DefaultValue":     0,
		"PhoneCall":        1,
		"PhysicalPresence": 2,
		"TrusteeEntity":    3,
	}
)

func (x RendezvousMechanisms) Enum() *RendezvousMechanisms {
	p := new(RendezvousMechanisms)
	*p = x
	return p
}

func (x RendezvousMechanisms) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RendezvousMechanisms) Descriptor() protoreflect.EnumDescriptor {
	return file_kycpb_enums_proto_enumTypes[0].Descriptor()
}

func (RendezvousMechanisms) Type() protoreflect.EnumType {
	return &file_kycpb_enums_proto_enumTypes[0]
}

func (x RendezvousMechanisms) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RendezvousMechanisms.Descriptor instead.
func (RendezvousMechanisms) EnumDescriptor() ([]byte, []int) {
	return file_kycpb_enums_proto_rawDescGZIP(), []int{0}
}

type RendezvousStatus int32

const (
	RendezvousStatus_NoAction           RendezvousStatus = 0
	RendezvousStatus_Created            RendezvousStatus = 1
	RendezvousStatus_RejectedByUser     RendezvousStatus = 2
	RendezvousStatus_RejectedByAdmin    RendezvousStatus = 3
	RendezvousStatus_CancelledByUser    RendezvousStatus = 4
	RendezvousStatus_CancelledByAdmin   RendezvousStatus = 5
	RendezvousStatus_MovedToAnotherTime RendezvousStatus = 6
	RendezvousStatus_InProgress         RendezvousStatus = 7
	RendezvousStatus_NoShow             RendezvousStatus = 8
	RendezvousStatus_Completed          RendezvousStatus = 9
)

// Enum value maps for RendezvousStatus.
var (
	RendezvousStatus_name = map[int32]string{
		0: "NoAction",
		1: "Created",
		2: "RejectedByUser",
		3: "RejectedByAdmin",
		4: "CancelledByUser",
		5: "CancelledByAdmin",
		6: "MovedToAnotherTime",
		7: "InProgress",
		8: "NoShow",
		9: "Completed",
	}
	RendezvousStatus_value = map[string]int32{
		"NoAction":           0,
		"Created":            1,
		"RejectedByUser":     2,
		"RejectedByAdmin":    3,
		"CancelledByUser":    4,
		"CancelledByAdmin":   5,
		"MovedToAnotherTime": 6,
		"InProgress":         7,
		"NoShow":             8,
		"Completed":          9,
	}
)

func (x RendezvousStatus) Enum() *RendezvousStatus {
	p := new(RendezvousStatus)
	*p = x
	return p
}

func (x RendezvousStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RendezvousStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_kycpb_enums_proto_enumTypes[1].Descriptor()
}

func (RendezvousStatus) Type() protoreflect.EnumType {
	return &file_kycpb_enums_proto_enumTypes[1]
}

func (x RendezvousStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RendezvousStatus.Descriptor instead.
func (RendezvousStatus) EnumDescriptor() ([]byte, []int) {
	return file_kycpb_enums_proto_rawDescGZIP(), []int{1}
}

type WorkdayLoadType int32

const (
	WorkdayLoadType_FullLoad              WorkdayLoadType = 0
	WorkdayLoadType_HalfLoad              WorkdayLoadType = 1
	WorkdayLoadType_NoLoad                WorkdayLoadType = 2
	WorkdayLoadType_ExtraLoad             WorkdayLoadType = 3
	WorkdayLoadType_MorningDoubleCapacity WorkdayLoadType = 4
)

// Enum value maps for WorkdayLoadType.
var (
	WorkdayLoadType_name = map[int32]string{
		0: "FullLoad",
		1: "HalfLoad",
		2: "NoLoad",
		3: "ExtraLoad",
		4: "MorningDoubleCapacity",
	}
	WorkdayLoadType_value = map[string]int32{
		"FullLoad":              0,
		"HalfLoad":              1,
		"NoLoad":                2,
		"ExtraLoad":             3,
		"MorningDoubleCapacity": 4,
	}
)

func (x WorkdayLoadType) Enum() *WorkdayLoadType {
	p := new(WorkdayLoadType)
	*p = x
	return p
}

func (x WorkdayLoadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkdayLoadType) Descriptor() protoreflect.EnumDescriptor {
	return file_kycpb_enums_proto_enumTypes[2].Descriptor()
}

func (WorkdayLoadType) Type() protoreflect.EnumType {
	return &file_kycpb_enums_proto_enumTypes[2]
}

func (x WorkdayLoadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkdayLoadType.Descriptor instead.
func (WorkdayLoadType) EnumDescriptor() ([]byte, []int) {
	return file_kycpb_enums_proto_rawDescGZIP(), []int{2}
}

type UserVerificationStatus int32

const (
	UserVerificationStatus_UVPending        UserVerificationStatus = 0
	UserVerificationStatus_UVVerified       UserVerificationStatus = 1
	UserVerificationStatus_UVRejected       UserVerificationStatus = 2
	UserVerificationStatus_UVReset          UserVerificationStatus = 3
	UserVerificationStatus_UVSendToOperator UserVerificationStatus = 4
)

// Enum value maps for UserVerificationStatus.
var (
	UserVerificationStatus_name = map[int32]string{
		0: "UVPending",
		1: "UVVerified",
		2: "UVRejected",
		3: "UVReset",
		4: "UVSendToOperator",
	}
	UserVerificationStatus_value = map[string]int32{
		"UVPending":        0,
		"UVVerified":       1,
		"UVRejected":       2,
		"UVReset":          3,
		"UVSendToOperator": 4,
	}
)

func (x UserVerificationStatus) Enum() *UserVerificationStatus {
	p := new(UserVerificationStatus)
	*p = x
	return p
}

func (x UserVerificationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserVerificationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_kycpb_enums_proto_enumTypes[3].Descriptor()
}

func (UserVerificationStatus) Type() protoreflect.EnumType {
	return &file_kycpb_enums_proto_enumTypes[3]
}

func (x UserVerificationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserVerificationStatus.Descriptor instead.
func (UserVerificationStatus) EnumDescriptor() ([]byte, []int) {
	return file_kycpb_enums_proto_rawDescGZIP(), []int{3}
}

type UserBankCardVerificationStatus int32

const (
	UserBankCardVerificationStatus_UBCPending  UserBankCardVerificationStatus = 0
	UserBankCardVerificationStatus_UBCVerified UserBankCardVerificationStatus = 1
	UserBankCardVerificationStatus_UBCRejected UserBankCardVerificationStatus = 2
)

// Enum value maps for UserBankCardVerificationStatus.
var (
	UserBankCardVerificationStatus_name = map[int32]string{
		0: "UBCPending",
		1: "UBCVerified",
		2: "UBCRejected",
	}
	UserBankCardVerificationStatus_value = map[string]int32{
		"UBCPending":  0,
		"UBCVerified": 1,
		"UBCRejected": 2,
	}
)

func (x UserBankCardVerificationStatus) Enum() *UserBankCardVerificationStatus {
	p := new(UserBankCardVerificationStatus)
	*p = x
	return p
}

func (x UserBankCardVerificationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserBankCardVerificationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_kycpb_enums_proto_enumTypes[4].Descriptor()
}

func (UserBankCardVerificationStatus) Type() protoreflect.EnumType {
	return &file_kycpb_enums_proto_enumTypes[4]
}

func (x UserBankCardVerificationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserBankCardVerificationStatus.Descriptor instead.
func (UserBankCardVerificationStatus) EnumDescriptor() ([]byte, []int) {
	return file_kycpb_enums_proto_rawDescGZIP(), []int{4}
}

type UserClaimVerificationStatus int32

const (
	UserClaimVerificationStatus_UCPending  UserClaimVerificationStatus = 0
	UserClaimVerificationStatus_UCVerified UserClaimVerificationStatus = 1
	UserClaimVerificationStatus_UCRejected UserClaimVerificationStatus = 2
	UserClaimVerificationStatus_UCReset    UserClaimVerificationStatus = 3
)

// Enum value maps for UserClaimVerificationStatus.
var (
	UserClaimVerificationStatus_name = map[int32]string{
		0: "UCPending",
		1: "UCVerified",
		2: "UCRejected",
		3: "UCReset",
	}
	UserClaimVerificationStatus_value = map[string]int32{
		"UCPending":  0,
		"UCVerified": 1,
		"UCRejected": 2,
		"UCReset":    3,
	}
)

func (x UserClaimVerificationStatus) Enum() *UserClaimVerificationStatus {
	p := new(UserClaimVerificationStatus)
	*p = x
	return p
}

func (x UserClaimVerificationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserClaimVerificationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_kycpb_enums_proto_enumTypes[5].Descriptor()
}

func (UserClaimVerificationStatus) Type() protoreflect.EnumType {
	return &file_kycpb_enums_proto_enumTypes[5]
}

func (x UserClaimVerificationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserClaimVerificationStatus.Descriptor instead.
func (UserClaimVerificationStatus) EnumDescriptor() ([]byte, []int) {
	return file_kycpb_enums_proto_rawDescGZIP(), []int{5}
}

type UserIBANVerificationStatus int32

const (
	UserIBANVerificationStatus_UIBANPending  UserIBANVerificationStatus = 0
	UserIBANVerificationStatus_UIBANVerified UserIBANVerificationStatus = 1
	UserIBANVerificationStatus_UIBANRejected UserIBANVerificationStatus = 2
)

// Enum value maps for UserIBANVerificationStatus.
var (
	UserIBANVerificationStatus_name = map[int32]string{
		0: "UIBANPending",
		1: "UIBANVerified",
		2: "UIBANRejected",
	}
	UserIBANVerificationStatus_value = map[string]int32{
		"UIBANPending":  0,
		"UIBANVerified": 1,
		"UIBANRejected": 2,
	}
)

func (x UserIBANVerificationStatus) Enum() *UserIBANVerificationStatus {
	p := new(UserIBANVerificationStatus)
	*p = x
	return p
}

func (x UserIBANVerificationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserIBANVerificationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_kycpb_enums_proto_enumTypes[6].Descriptor()
}

func (UserIBANVerificationStatus) Type() protoreflect.EnumType {
	return &file_kycpb_enums_proto_enumTypes[6]
}

func (x UserIBANVerificationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserIBANVerificationStatus.Descriptor instead.
func (UserIBANVerificationStatus) EnumDescriptor() ([]byte, []int) {
	return file_kycpb_enums_proto_rawDescGZIP(), []int{6}
}

type UserClaimType int32

const (
	UserClaimType_NID                    UserClaimType = 0
	UserClaimType_BookletID              UserClaimType = 1
	UserClaimType_FatherName             UserClaimType = 2
	UserClaimType_AdvertisementChannel   UserClaimType = 3
	UserClaimType_AdvertisementComment   UserClaimType = 4
	UserClaimType_BuyExchangeAssetSymbol UserClaimType = 5
	UserClaimType_BirthDate              UserClaimType = 6
	UserClaimType_Address                UserClaimType = 7
	UserClaimType_State                  UserClaimType = 8
	UserClaimType_PostalCode             UserClaimType = 9
	UserClaimType_LandlineNumber         UserClaimType = 10
	UserClaimType_LandlineTime           UserClaimType = 11
	UserClaimType_SelfieFile             UserClaimType = 12
	UserClaimType_DocsFile               UserClaimType = 13
	UserClaimType_IDDescFile             UserClaimType = 14
	UserClaimType_FirstName              UserClaimType = 15
	UserClaimType_LastName               UserClaimType = 16
	UserClaimType_MiddleName             UserClaimType = 17
	UserClaimType_Email                  UserClaimType = 18
)

// Enum value maps for UserClaimType.
var (
	UserClaimType_name = map[int32]string{
		0:  "NID",
		1:  "BookletID",
		2:  "FatherName",
		3:  "AdvertisementChannel",
		4:  "AdvertisementComment",
		5:  "BuyExchangeAssetSymbol",
		6:  "BirthDate",
		7:  "Address",
		8:  "State",
		9:  "PostalCode",
		10: "LandlineNumber",
		11: "LandlineTime",
		12: "SelfieFile",
		13: "DocsFile",
		14: "IDDescFile",
		15: "FirstName",
		16: "LastName",
		17: "MiddleName",
		18: "Email",
	}
	UserClaimType_value = map[string]int32{
		"NID":                    0,
		"BookletID":              1,
		"FatherName":             2,
		"AdvertisementChannel":   3,
		"AdvertisementComment":   4,
		"BuyExchangeAssetSymbol": 5,
		"BirthDate":              6,
		"Address":                7,
		"State":                  8,
		"PostalCode":             9,
		"LandlineNumber":         10,
		"LandlineTime":           11,
		"SelfieFile":             12,
		"DocsFile":               13,
		"IDDescFile":             14,
		"FirstName":              15,
		"LastName":               16,
		"MiddleName":             17,
		"Email":                  18,
	}
)

func (x UserClaimType) Enum() *UserClaimType {
	p := new(UserClaimType)
	*p = x
	return p
}

func (x UserClaimType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserClaimType) Descriptor() protoreflect.EnumDescriptor {
	return file_kycpb_enums_proto_enumTypes[7].Descriptor()
}

func (UserClaimType) Type() protoreflect.EnumType {
	return &file_kycpb_enums_proto_enumTypes[7]
}

func (x UserClaimType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserClaimType.Descriptor instead.
func (UserClaimType) EnumDescriptor() ([]byte, []int) {
	return file_kycpb_enums_proto_rawDescGZIP(), []int{7}
}

type UserClaimBatchType int32

const (
	UserClaimBatchType_BaseInfoBatch  UserClaimBatchType = 0
	UserClaimBatchType_DocumentsBatch UserClaimBatchType = 1
	UserClaimBatchType_LandlineBatch  UserClaimBatchType = 2
)

// Enum value maps for UserClaimBatchType.
var (
	UserClaimBatchType_name = map[int32]string{
		0: "BaseInfoBatch",
		1: "DocumentsBatch",
		2: "LandlineBatch",
	}
	UserClaimBatchType_value = map[string]int32{
		"BaseInfoBatch":  0,
		"DocumentsBatch": 1,
		"LandlineBatch":  2,
	}
)

func (x UserClaimBatchType) Enum() *UserClaimBatchType {
	p := new(UserClaimBatchType)
	*p = x
	return p
}

func (x UserClaimBatchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserClaimBatchType) Descriptor() protoreflect.EnumDescriptor {
	return file_kycpb_enums_proto_enumTypes[8].Descriptor()
}

func (UserClaimBatchType) Type() protoreflect.EnumType {
	return &file_kycpb_enums_proto_enumTypes[8]
}

func (x UserClaimBatchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserClaimBatchType.Descriptor instead.
func (UserClaimBatchType) EnumDescriptor() ([]byte, []int) {
	return file_kycpb_enums_proto_rawDescGZIP(), []int{8}
}

type RejectionReasonType int32

const (
	RejectionReasonType_BaseInfoReason  RejectionReasonType = 0
	RejectionReasonType_DocumentsReason RejectionReasonType = 1
	RejectionReasonType_LandlineReason  RejectionReasonType = 2
	RejectionReasonType_UserReason      RejectionReasonType = 3
)

// Enum value maps for RejectionReasonType.
var (
	RejectionReasonType_name = map[int32]string{
		0: "BaseInfoReason",
		1: "DocumentsReason",
		2: "LandlineReason",
		3: "UserReason",
	}
	RejectionReasonType_value = map[string]int32{
		"BaseInfoReason":  0,
		"DocumentsReason": 1,
		"LandlineReason":  2,
		"UserReason":      3,
	}
)

func (x RejectionReasonType) Enum() *RejectionReasonType {
	p := new(RejectionReasonType)
	*p = x
	return p
}

func (x RejectionReasonType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RejectionReasonType) Descriptor() protoreflect.EnumDescriptor {
	return file_kycpb_enums_proto_enumTypes[9].Descriptor()
}

func (RejectionReasonType) Type() protoreflect.EnumType {
	return &file_kycpb_enums_proto_enumTypes[9]
}

func (x RejectionReasonType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RejectionReasonType.Descriptor instead.
func (RejectionReasonType) EnumDescriptor() ([]byte, []int) {
	return file_kycpb_enums_proto_rawDescGZIP(), []int{9}
}

var File_kycpb_enums_proto protoreflect.FileDescriptor

var file_kycpb_enums_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6b, 0x79, 0x63, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x79, 0x63, 0x2a, 0x60, 0x0a, 0x14, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73,
	0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x50,
	0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x10,
	0x02, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x65, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x10, 0x03, 0x2a, 0xc4, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76,
	0x6f, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x6f, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x42, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x03, 0x12, 0x13, 0x0a,
	0x0f, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x42,
	0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x6f, 0x76, 0x65,
	0x64, 0x54, 0x6f, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x10, 0x06,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x07,
	0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f, 0x53, 0x68, 0x6f, 0x77, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x09, 0x2a, 0x63, 0x0a, 0x0f, 0x57,
	0x6f, 0x72, 0x6b, 0x64, 0x61, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c,
	0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4c, 0x6f, 0x61, 0x64, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x48, 0x61, 0x6c, 0x66, 0x4c, 0x6f, 0x61, 0x64, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f,
	0x4c, 0x6f, 0x61, 0x64, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c,
	0x6f, 0x61, 0x64, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x6f, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x10, 0x04,
	0x2a, 0x6a, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x56,
	0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x56, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x56, 0x52,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x56, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x56, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x6f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x10, 0x04, 0x2a, 0x52, 0x0a, 0x1e,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e,
	0x0a, 0x0a, 0x55, 0x42, 0x43, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x42, 0x43, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x42, 0x43, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x02,
	0x2a, 0x59, 0x0a, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0d, 0x0a, 0x09, 0x55, 0x43, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x55, 0x43, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x55, 0x43, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x43, 0x52, 0x65, 0x73, 0x65, 0x74, 0x10, 0x03, 0x2a, 0x54, 0x0a, 0x1a, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x42, 0x41, 0x4e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x49, 0x42,
	0x41, 0x4e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x55,
	0x49, 0x42, 0x41, 0x4e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x55, 0x49, 0x42, 0x41, 0x4e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10,
	0x02, 0x2a, 0xca, 0x02, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x42, 0x6f, 0x6f, 0x6b, 0x6c, 0x65, 0x74, 0x49, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x46,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x41,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x04, 0x12,
	0x1a, 0x0a, 0x16, 0x42, 0x75, 0x79, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x42,
	0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65,
	0x10, 0x09, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x61, 0x6e, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x10, 0x0b, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x66,
	0x69, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x73,
	0x46, 0x69, 0x6c, 0x65, 0x10, 0x0d, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x44, 0x44, 0x65, 0x73, 0x63,
	0x46, 0x69, 0x6c, 0x65, 0x10, 0x0e, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x10, 0x0f, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x10, 0x10, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x10, 0x11, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0x12, 0x2a, 0x4e,
	0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4c,
	0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x10, 0x02, 0x2a, 0x62,
	0x0a, 0x13, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x10, 0x01, 0x12, 0x12,
	0x0a, 0x0e, 0x4c, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x10, 0x03, 0x42, 0x1a, 0x5a, 0x18, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x62, 0x2f, 0x6b, 0x79, 0x63, 0x70, 0x62, 0x3b, 0x6b, 0x79, 0x63, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kycpb_enums_proto_rawDescOnce sync.Once
	file_kycpb_enums_proto_rawDescData = file_kycpb_enums_proto_rawDesc
)

func file_kycpb_enums_proto_rawDescGZIP() []byte {
	file_kycpb_enums_proto_rawDescOnce.Do(func() {
		file_kycpb_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_kycpb_enums_proto_rawDescData)
	})
	return file_kycpb_enums_proto_rawDescData
}

var file_kycpb_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 10)
var file_kycpb_enums_proto_goTypes = []interface{}{
	(RendezvousMechanisms)(0),           // 0: rabex.api.kyc.RendezvousMechanisms
	(RendezvousStatus)(0),               // 1: rabex.api.kyc.RendezvousStatus
	(WorkdayLoadType)(0),                // 2: rabex.api.kyc.WorkdayLoadType
	(UserVerificationStatus)(0),         // 3: rabex.api.kyc.UserVerificationStatus
	(UserBankCardVerificationStatus)(0), // 4: rabex.api.kyc.UserBankCardVerificationStatus
	(UserClaimVerificationStatus)(0),    // 5: rabex.api.kyc.UserClaimVerificationStatus
	(UserIBANVerificationStatus)(0),     // 6: rabex.api.kyc.UserIBANVerificationStatus
	(UserClaimType)(0),                  // 7: rabex.api.kyc.UserClaimType
	(UserClaimBatchType)(0),             // 8: rabex.api.kyc.UserClaimBatchType
	(RejectionReasonType)(0),            // 9: rabex.api.kyc.RejectionReasonType
}
var file_kycpb_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kycpb_enums_proto_init() }
func file_kycpb_enums_proto_init() {
	if File_kycpb_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kycpb_enums_proto_rawDesc,
			NumEnums:      10,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kycpb_enums_proto_goTypes,
		DependencyIndexes: file_kycpb_enums_proto_depIdxs,
		EnumInfos:         file_kycpb_enums_proto_enumTypes,
	}.Build()
	File_kycpb_enums_proto = out.File
	file_kycpb_enums_proto_rawDesc = nil
	file_kycpb_enums_proto_goTypes = nil
	file_kycpb_enums_proto_depIdxs = nil
}
