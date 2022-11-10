// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: commons/enums.proto

package commons

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type Clients int32

const (
	Clients_FiatGateway    Clients = 0
	Clients_CryptoGateway  Clients = 1
	Clients_Accounting     Clients = 2
	Clients_UserManagement Clients = 3
	Clients_Ticket         Clients = 4
	Clients_Exchange       Clients = 5
)

// Enum value maps for Clients.
var (
	Clients_name = map[int32]string{
		0: "FiatGateway",
		1: "CryptoGateway",
		2: "Accounting",
		3: "UserManagement",
		4: "Ticket",
		5: "Exchange",
	}
	Clients_value = map[string]int32{
		"FiatGateway":    0,
		"CryptoGateway":  1,
		"Accounting":     2,
		"UserManagement": 3,
		"Ticket":         4,
		"Exchange":       5,
	}
)

func (x Clients) Enum() *Clients {
	p := new(Clients)
	*p = x
	return p
}

func (x Clients) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Clients) Descriptor() protoreflect.EnumDescriptor {
	return file_commons_enums_proto_enumTypes[0].Descriptor()
}

func (Clients) Type() protoreflect.EnumType {
	return &file_commons_enums_proto_enumTypes[0]
}

func (x Clients) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Clients.Descriptor instead.
func (Clients) EnumDescriptor() ([]byte, []int) {
	return file_commons_enums_proto_rawDescGZIP(), []int{0}
}

type Side int32

const (
	Side_UNSPECIFIED Side = 0
	Side_Sell        Side = 1
	Side_Buy         Side = 2
)

// Enum value maps for Side.
var (
	Side_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "Sell",
		2: "Buy",
	}
	Side_value = map[string]int32{
		"UNSPECIFIED": 0,
		"Sell":        1,
		"Buy":         2,
	}
)

func (x Side) Enum() *Side {
	p := new(Side)
	*p = x
	return p
}

func (x Side) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Side) Descriptor() protoreflect.EnumDescriptor {
	return file_commons_enums_proto_enumTypes[1].Descriptor()
}

func (Side) Type() protoreflect.EnumType {
	return &file_commons_enums_proto_enumTypes[1]
}

func (x Side) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Side.Descriptor instead.
func (Side) EnumDescriptor() ([]byte, []int) {
	return file_commons_enums_proto_rawDescGZIP(), []int{1}
}

type Assets int32

const (
	Assets_UnspecifiedAsset Assets = 0
	Assets_BTC              Assets = 1
	Assets_ETH              Assets = 2
	Assets_XRP              Assets = 3
	Assets_BCH              Assets = 4
	Assets_LTC              Assets = 5
	Assets_LINK             Assets = 6
	Assets_EOS              Assets = 7
	Assets_TRX              Assets = 8
	Assets_XLM              Assets = 9
	Assets_XMR              Assets = 10
	Assets_IOTA             Assets = 11
	Assets_ZEC              Assets = 12
	Assets_USDC             Assets = 13
	Assets_DASH             Assets = 14
	Assets_DOGE             Assets = 15
	Assets_PAX              Assets = 16
	Assets_SC               Assets = 17
	Assets_BAT              Assets = 18
	Assets_BCHSV            Assets = 19
	Assets_BTT              Assets = 20
	Assets_ETC              Assets = 21
	Assets_NEO              Assets = 22
	Assets_ATOM             Assets = 23
	Assets_ONG              Assets = 24
	Assets_VET              Assets = 25
	Assets_DCR              Assets = 26
	Assets_QTUM             Assets = 27
	Assets_IOST             Assets = 28
	Assets_XTZ              Assets = 29
	Assets_WAVES            Assets = 30
	Assets_ALGO             Assets = 31
	Assets_TUSD             Assets = 32
	Assets_HOT              Assets = 33
	Assets_KAVA             Assets = 34
	Assets_ZRX              Assets = 35
	Assets_ICX              Assets = 36
	Assets_NANO             Assets = 37
	Assets_ONE              Assets = 38
	Assets_NPXS             Assets = 39
	Assets_FTM              Assets = 40
	Assets_DENT             Assets = 41
	Assets_THETA            Assets = 42
	Assets_MITH             Assets = 43
	Assets_WAN              Assets = 44
	Assets_DOCK             Assets = 45
	Assets_USDS             Assets = 46
	Assets_ENJ              Assets = 47
	Assets_COCOS            Assets = 48
	Assets_RVN              Assets = 49
	Assets_TOMO             Assets = 50
	Assets_ZIL              Assets = 51
	Assets_BEAM             Assets = 52
	Assets_HC               Assets = 53
	Assets_NULS             Assets = 54
	Assets_STORM            Assets = 55
	Assets_MTL              Assets = 56
	Assets_CVC              Assets = 57
	Assets_GTO              Assets = 58
	Assets_FUN              Assets = 59
	Assets_PERL             Assets = 60
	Assets_DUSK             Assets = 61
	Assets_DGB              Assets = 62
	Assets_EGLD             Assets = 63
	Assets_AAVE             Assets = 64
	Assets_VTHO             Assets = 65
	Assets_JST              Assets = 66
	Assets_PAXG             Assets = 67
	Assets_DOT              Assets = 68
	Assets_UNI              Assets = 69
	Assets_BAND             Assets = 70
	Assets_FIL              Assets = 71
	Assets_SUN              Assets = 72
	Assets_OCEAN            Assets = 73
	Assets_SOL              Assets = 74
	Assets_CHZ              Assets = 75
	Assets_SUSHI            Assets = 76
	Assets_TFUEL            Assets = 77
	Assets_BNB              Assets = 78
	Assets_USDT             Assets = 79
	Assets_ONEINCH          Assets = 80
	Assets_CAKE             Assets = 81
	Assets_BUSD             Assets = 82
	Assets_WIN              Assets = 83
	Assets_ONT              Assets = 84
	Assets_SHIB             Assets = 85
	Assets_MATIC            Assets = 86
	Assets_YFI              Assets = 87
	Assets_YFII             Assets = 88
	Assets_AVAX             Assets = 89
	Assets_ADA              Assets = 90
	Assets_IRR              Assets = 301
)

// Enum value maps for Assets.
var (
	Assets_name = map[int32]string{
		0:   "UnspecifiedAsset",
		1:   "BTC",
		2:   "ETH",
		3:   "XRP",
		4:   "BCH",
		5:   "LTC",
		6:   "LINK",
		7:   "EOS",
		8:   "TRX",
		9:   "XLM",
		10:  "XMR",
		11:  "IOTA",
		12:  "ZEC",
		13:  "USDC",
		14:  "DASH",
		15:  "DOGE",
		16:  "PAX",
		17:  "SC",
		18:  "BAT",
		19:  "BCHSV",
		20:  "BTT",
		21:  "ETC",
		22:  "NEO",
		23:  "ATOM",
		24:  "ONG",
		25:  "VET",
		26:  "DCR",
		27:  "QTUM",
		28:  "IOST",
		29:  "XTZ",
		30:  "WAVES",
		31:  "ALGO",
		32:  "TUSD",
		33:  "HOT",
		34:  "KAVA",
		35:  "ZRX",
		36:  "ICX",
		37:  "NANO",
		38:  "ONE",
		39:  "NPXS",
		40:  "FTM",
		41:  "DENT",
		42:  "THETA",
		43:  "MITH",
		44:  "WAN",
		45:  "DOCK",
		46:  "USDS",
		47:  "ENJ",
		48:  "COCOS",
		49:  "RVN",
		50:  "TOMO",
		51:  "ZIL",
		52:  "BEAM",
		53:  "HC",
		54:  "NULS",
		55:  "STORM",
		56:  "MTL",
		57:  "CVC",
		58:  "GTO",
		59:  "FUN",
		60:  "PERL",
		61:  "DUSK",
		62:  "DGB",
		63:  "EGLD",
		64:  "AAVE",
		65:  "VTHO",
		66:  "JST",
		67:  "PAXG",
		68:  "DOT",
		69:  "UNI",
		70:  "BAND",
		71:  "FIL",
		72:  "SUN",
		73:  "OCEAN",
		74:  "SOL",
		75:  "CHZ",
		76:  "SUSHI",
		77:  "TFUEL",
		78:  "BNB",
		79:  "USDT",
		80:  "ONEINCH",
		81:  "CAKE",
		82:  "BUSD",
		83:  "WIN",
		84:  "ONT",
		85:  "SHIB",
		86:  "MATIC",
		87:  "YFI",
		88:  "YFII",
		89:  "AVAX",
		90:  "ADA",
		301: "IRR",
	}
	Assets_value = map[string]int32{
		"UnspecifiedAsset": 0,
		"BTC":              1,
		"ETH":              2,
		"XRP":              3,
		"BCH":              4,
		"LTC":              5,
		"LINK":             6,
		"EOS":              7,
		"TRX":              8,
		"XLM":              9,
		"XMR":              10,
		"IOTA":             11,
		"ZEC":              12,
		"USDC":             13,
		"DASH":             14,
		"DOGE":             15,
		"PAX":              16,
		"SC":               17,
		"BAT":              18,
		"BCHSV":            19,
		"BTT":              20,
		"ETC":              21,
		"NEO":              22,
		"ATOM":             23,
		"ONG":              24,
		"VET":              25,
		"DCR":              26,
		"QTUM":             27,
		"IOST":             28,
		"XTZ":              29,
		"WAVES":            30,
		"ALGO":             31,
		"TUSD":             32,
		"HOT":              33,
		"KAVA":             34,
		"ZRX":              35,
		"ICX":              36,
		"NANO":             37,
		"ONE":              38,
		"NPXS":             39,
		"FTM":              40,
		"DENT":             41,
		"THETA":            42,
		"MITH":             43,
		"WAN":              44,
		"DOCK":             45,
		"USDS":             46,
		"ENJ":              47,
		"COCOS":            48,
		"RVN":              49,
		"TOMO":             50,
		"ZIL":              51,
		"BEAM":             52,
		"HC":               53,
		"NULS":             54,
		"STORM":            55,
		"MTL":              56,
		"CVC":              57,
		"GTO":              58,
		"FUN":              59,
		"PERL":             60,
		"DUSK":             61,
		"DGB":              62,
		"EGLD":             63,
		"AAVE":             64,
		"VTHO":             65,
		"JST":              66,
		"PAXG":             67,
		"DOT":              68,
		"UNI":              69,
		"BAND":             70,
		"FIL":              71,
		"SUN":              72,
		"OCEAN":            73,
		"SOL":              74,
		"CHZ":              75,
		"SUSHI":            76,
		"TFUEL":            77,
		"BNB":              78,
		"USDT":             79,
		"ONEINCH":          80,
		"CAKE":             81,
		"BUSD":             82,
		"WIN":              83,
		"ONT":              84,
		"SHIB":             85,
		"MATIC":            86,
		"YFI":              87,
		"YFII":             88,
		"AVAX":             89,
		"ADA":              90,
		"IRR":              301,
	}
)

func (x Assets) Enum() *Assets {
	p := new(Assets)
	*p = x
	return p
}

func (x Assets) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Assets) Descriptor() protoreflect.EnumDescriptor {
	return file_commons_enums_proto_enumTypes[2].Descriptor()
}

func (Assets) Type() protoreflect.EnumType {
	return &file_commons_enums_proto_enumTypes[2]
}

func (x Assets) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Assets.Descriptor instead.
func (Assets) EnumDescriptor() ([]byte, []int) {
	return file_commons_enums_proto_rawDescGZIP(), []int{2}
}

type Networks int32

const (
	Networks_UnspecifiedNetwork Networks = 0
	Networks_Shetab             Networks = 1
	Networks_Visa               Networks = 2
	Networks_BTCMain            Networks = 101
	Networks_ETHMain            Networks = 102
	Networks_Tron               Networks = 103
	Networks_BNBMain            Networks = 104
)

// Enum value maps for Networks.
var (
	Networks_name = map[int32]string{
		0:   "UnspecifiedNetwork",
		1:   "Shetab",
		2:   "Visa",
		101: "BTCMain",
		102: "ETHMain",
		103: "Tron",
		104: "BNBMain",
	}
	Networks_value = map[string]int32{
		"UnspecifiedNetwork": 0,
		"Shetab":             1,
		"Visa":               2,
		"BTCMain":            101,
		"ETHMain":            102,
		"Tron":               103,
		"BNBMain":            104,
	}
)

func (x Networks) Enum() *Networks {
	p := new(Networks)
	*p = x
	return p
}

func (x Networks) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Networks) Descriptor() protoreflect.EnumDescriptor {
	return file_commons_enums_proto_enumTypes[3].Descriptor()
}

func (Networks) Type() protoreflect.EnumType {
	return &file_commons_enums_proto_enumTypes[3]
}

func (x Networks) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Networks.Descriptor instead.
func (Networks) EnumDescriptor() ([]byte, []int) {
	return file_commons_enums_proto_rawDescGZIP(), []int{3}
}

type Brokers int32

const (
	Brokers_Rabex      Brokers = 0
	Brokers_AmirMasoud Brokers = 1
	Brokers_Mohsen     Brokers = 2
)

// Enum value maps for Brokers.
var (
	Brokers_name = map[int32]string{
		0: "Rabex",
		1: "AmirMasoud",
		2: "Mohsen",
	}
	Brokers_value = map[string]int32{
		"Rabex":      0,
		"AmirMasoud": 1,
		"Mohsen":     2,
	}
)

func (x Brokers) Enum() *Brokers {
	p := new(Brokers)
	*p = x
	return p
}

func (x Brokers) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Brokers) Descriptor() protoreflect.EnumDescriptor {
	return file_commons_enums_proto_enumTypes[4].Descriptor()
}

func (Brokers) Type() protoreflect.EnumType {
	return &file_commons_enums_proto_enumTypes[4]
}

func (x Brokers) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Brokers.Descriptor instead.
func (Brokers) EnumDescriptor() ([]byte, []int) {
	return file_commons_enums_proto_rawDescGZIP(), []int{4}
}

type TransactionType int32

const (
	TransactionType_Withdraw TransactionType = 0
	TransactionType_Deposit  TransactionType = 1
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "Withdraw",
		1: "Deposit",
	}
	TransactionType_value = map[string]int32{
		"Withdraw": 0,
		"Deposit":  1,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_commons_enums_proto_enumTypes[5].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_commons_enums_proto_enumTypes[5]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_commons_enums_proto_rawDescGZIP(), []int{5}
}

var file_commons_enums_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51234,
		Name:          "rabex.api.commons.custom_name",
		Tag:           "bytes,51234,opt,name=custom_name",
		Filename:      "commons/enums.proto",
	},
}

// Extension fields to descriptor.EnumValueOptions.
var (
	// optional string custom_name = 51234;
	E_CustomName = &file_commons_enums_proto_extTypes[0]
)

var File_commons_enums_proto protoreflect.FileDescriptor

var file_commons_enums_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x72, 0x61, 0x62, 0x65, 0x78, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x6b, 0x0a, 0x07, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x69, 0x61, 0x74, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x03, 0x12, 0x0a, 0x0a,
	0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x10, 0x05, 0x2a, 0x2a, 0x0a, 0x04, 0x53, 0x69, 0x64, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x65, 0x6c, 0x6c, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x75,
	0x79, 0x10, 0x02, 0x2a, 0x8e, 0x07, 0x0a, 0x06, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x14,
	0x0a, 0x10, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x54, 0x43, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x45, 0x54, 0x48, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x52, 0x50, 0x10, 0x03, 0x12,
	0x07, 0x0a, 0x03, 0x42, 0x43, 0x48, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x54, 0x43, 0x10,
	0x05, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x45,
	0x4f, 0x53, 0x10, 0x07, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x52, 0x58, 0x10, 0x08, 0x12, 0x07, 0x0a,
	0x03, 0x58, 0x4c, 0x4d, 0x10, 0x09, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x4d, 0x52, 0x10, 0x0a, 0x12,
	0x08, 0x0a, 0x04, 0x49, 0x4f, 0x54, 0x41, 0x10, 0x0b, 0x12, 0x07, 0x0a, 0x03, 0x5a, 0x45, 0x43,
	0x10, 0x0c, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x44, 0x43, 0x10, 0x0d, 0x12, 0x08, 0x0a, 0x04,
	0x44, 0x41, 0x53, 0x48, 0x10, 0x0e, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x47, 0x45, 0x10, 0x0f,
	0x12, 0x07, 0x0a, 0x03, 0x50, 0x41, 0x58, 0x10, 0x10, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x43, 0x10,
	0x11, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x41, 0x54, 0x10, 0x12, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x43,
	0x48, 0x53, 0x56, 0x10, 0x13, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x54, 0x54, 0x10, 0x14, 0x12, 0x07,
	0x0a, 0x03, 0x45, 0x54, 0x43, 0x10, 0x15, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x4f, 0x10, 0x16,
	0x12, 0x08, 0x0a, 0x04, 0x41, 0x54, 0x4f, 0x4d, 0x10, 0x17, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4e,
	0x47, 0x10, 0x18, 0x12, 0x07, 0x0a, 0x03, 0x56, 0x45, 0x54, 0x10, 0x19, 0x12, 0x07, 0x0a, 0x03,
	0x44, 0x43, 0x52, 0x10, 0x1a, 0x12, 0x08, 0x0a, 0x04, 0x51, 0x54, 0x55, 0x4d, 0x10, 0x1b, 0x12,
	0x08, 0x0a, 0x04, 0x49, 0x4f, 0x53, 0x54, 0x10, 0x1c, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x54, 0x5a,
	0x10, 0x1d, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x41, 0x56, 0x45, 0x53, 0x10, 0x1e, 0x12, 0x08, 0x0a,
	0x04, 0x41, 0x4c, 0x47, 0x4f, 0x10, 0x1f, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x55, 0x53, 0x44, 0x10,
	0x20, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x4f, 0x54, 0x10, 0x21, 0x12, 0x08, 0x0a, 0x04, 0x4b, 0x41,
	0x56, 0x41, 0x10, 0x22, 0x12, 0x07, 0x0a, 0x03, 0x5a, 0x52, 0x58, 0x10, 0x23, 0x12, 0x07, 0x0a,
	0x03, 0x49, 0x43, 0x58, 0x10, 0x24, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x4e, 0x4f, 0x10, 0x25,
	0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10, 0x26, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x50, 0x58,
	0x53, 0x10, 0x27, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x54, 0x4d, 0x10, 0x28, 0x12, 0x08, 0x0a, 0x04,
	0x44, 0x45, 0x4e, 0x54, 0x10, 0x29, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x48, 0x45, 0x54, 0x41, 0x10,
	0x2a, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x49, 0x54, 0x48, 0x10, 0x2b, 0x12, 0x07, 0x0a, 0x03, 0x57,
	0x41, 0x4e, 0x10, 0x2c, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x43, 0x4b, 0x10, 0x2d, 0x12, 0x08,
	0x0a, 0x04, 0x55, 0x53, 0x44, 0x53, 0x10, 0x2e, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x4e, 0x4a, 0x10,
	0x2f, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4f, 0x43, 0x4f, 0x53, 0x10, 0x30, 0x12, 0x07, 0x0a, 0x03,
	0x52, 0x56, 0x4e, 0x10, 0x31, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x4f, 0x4d, 0x4f, 0x10, 0x32, 0x12,
	0x07, 0x0a, 0x03, 0x5a, 0x49, 0x4c, 0x10, 0x33, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x45, 0x41, 0x4d,
	0x10, 0x34, 0x12, 0x06, 0x0a, 0x02, 0x48, 0x43, 0x10, 0x35, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55,
	0x4c, 0x53, 0x10, 0x36, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x4f, 0x52, 0x4d, 0x10, 0x37, 0x12,
	0x07, 0x0a, 0x03, 0x4d, 0x54, 0x4c, 0x10, 0x38, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x56, 0x43, 0x10,
	0x39, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x54, 0x4f, 0x10, 0x3a, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x55,
	0x4e, 0x10, 0x3b, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x45, 0x52, 0x4c, 0x10, 0x3c, 0x12, 0x08, 0x0a,
	0x04, 0x44, 0x55, 0x53, 0x4b, 0x10, 0x3d, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x47, 0x42, 0x10, 0x3e,
	0x12, 0x08, 0x0a, 0x04, 0x45, 0x47, 0x4c, 0x44, 0x10, 0x3f, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x41,
	0x56, 0x45, 0x10, 0x40, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x54, 0x48, 0x4f, 0x10, 0x41, 0x12, 0x07,
	0x0a, 0x03, 0x4a, 0x53, 0x54, 0x10, 0x42, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x58, 0x47, 0x10,
	0x43, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x4f, 0x54, 0x10, 0x44, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x4e,
	0x49, 0x10, 0x45, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x41, 0x4e, 0x44, 0x10, 0x46, 0x12, 0x07, 0x0a,
	0x03, 0x46, 0x49, 0x4c, 0x10, 0x47, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x55, 0x4e, 0x10, 0x48, 0x12,
	0x09, 0x0a, 0x05, 0x4f, 0x43, 0x45, 0x41, 0x4e, 0x10, 0x49, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x4f,
	0x4c, 0x10, 0x4a, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x48, 0x5a, 0x10, 0x4b, 0x12, 0x09, 0x0a, 0x05,
	0x53, 0x55, 0x53, 0x48, 0x49, 0x10, 0x4c, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x46, 0x55, 0x45, 0x4c,
	0x10, 0x4d, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x4e, 0x42, 0x10, 0x4e, 0x12, 0x08, 0x0a, 0x04, 0x55,
	0x53, 0x44, 0x54, 0x10, 0x4f, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x4e, 0x45, 0x49, 0x4e, 0x43, 0x48,
	0x10, 0x50, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x41, 0x4b, 0x45, 0x10, 0x51, 0x12, 0x08, 0x0a, 0x04,
	0x42, 0x55, 0x53, 0x44, 0x10, 0x52, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x49, 0x4e, 0x10, 0x53, 0x12,
	0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x54, 0x10, 0x54, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x48, 0x49, 0x42,
	0x10, 0x55, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x54, 0x49, 0x43, 0x10, 0x56, 0x12, 0x07, 0x0a,
	0x03, 0x59, 0x46, 0x49, 0x10, 0x57, 0x12, 0x08, 0x0a, 0x04, 0x59, 0x46, 0x49, 0x49, 0x10, 0x58,
	0x12, 0x08, 0x0a, 0x04, 0x41, 0x56, 0x41, 0x58, 0x10, 0x59, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44,
	0x41, 0x10, 0x5a, 0x12, 0x08, 0x0a, 0x03, 0x49, 0x52, 0x52, 0x10, 0xad, 0x02, 0x22, 0x05, 0x08,
	0x5b, 0x10, 0xac, 0x02, 0x2a, 0x6f, 0x0a, 0x08, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x68, 0x65, 0x74,
	0x61, 0x62, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x69, 0x73, 0x61, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x42, 0x54, 0x43, 0x4d, 0x61, 0x69, 0x6e, 0x10, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x45,
	0x54, 0x48, 0x4d, 0x61, 0x69, 0x6e, 0x10, 0x66, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x72, 0x6f, 0x6e,
	0x10, 0x67, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4e, 0x42, 0x4d, 0x61, 0x69, 0x6e, 0x10, 0x68, 0x22,
	0x04, 0x08, 0x03, 0x10, 0x64, 0x2a, 0x30, 0x0a, 0x07, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x73,
	0x12, 0x09, 0x0a, 0x05, 0x52, 0x61, 0x62, 0x65, 0x78, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x41,
	0x6d, 0x69, 0x72, 0x4d, 0x61, 0x73, 0x6f, 0x75, 0x64, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x6f, 0x68, 0x73, 0x65, 0x6e, 0x10, 0x02, 0x2a, 0x2c, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x10, 0x01, 0x3a, 0x44, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa2, 0x90, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x72,
	0x61, 0x62, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_commons_enums_proto_rawDescOnce sync.Once
	file_commons_enums_proto_rawDescData = file_commons_enums_proto_rawDesc
)

func file_commons_enums_proto_rawDescGZIP() []byte {
	file_commons_enums_proto_rawDescOnce.Do(func() {
		file_commons_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_commons_enums_proto_rawDescData)
	})
	return file_commons_enums_proto_rawDescData
}

var file_commons_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_commons_enums_proto_goTypes = []interface{}{
	(Clients)(0),                        // 0: rabex.api.commons.Clients
	(Side)(0),                           // 1: rabex.api.commons.Side
	(Assets)(0),                         // 2: rabex.api.commons.Assets
	(Networks)(0),                       // 3: rabex.api.commons.Networks
	(Brokers)(0),                        // 4: rabex.api.commons.Brokers
	(TransactionType)(0),                // 5: rabex.api.commons.TransactionType
	(*descriptor.EnumValueOptions)(nil), // 6: google.protobuf.EnumValueOptions
}
var file_commons_enums_proto_depIdxs = []int32{
	6, // 0: rabex.api.commons.custom_name:extendee -> google.protobuf.EnumValueOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commons_enums_proto_init() }
func file_commons_enums_proto_init() {
	if File_commons_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commons_enums_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_commons_enums_proto_goTypes,
		DependencyIndexes: file_commons_enums_proto_depIdxs,
		EnumInfos:         file_commons_enums_proto_enumTypes,
		ExtensionInfos:    file_commons_enums_proto_extTypes,
	}.Build()
	File_commons_enums_proto = out.File
	file_commons_enums_proto_rawDesc = nil
	file_commons_enums_proto_goTypes = nil
	file_commons_enums_proto_depIdxs = nil
}
