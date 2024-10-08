//*
// Defines the AKM Suite values used in the 802.11 Survey message.
//
// This protobuf definition is provided as a convenience only. See the official API documentation for the true field
// schema.
//
// It is necessary to define this enum in an individual file so that duplicate enum values can be used. This file
// specifies a different protobuf package than other enums so that the enum value scope is different than all other
// enums.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: com/craxiom/messaging/wifi/akmsuite/akm_suite.proto

package messaging

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

// This enum defines the AKM Suites field of the 802.11 messages.
//
// Note that a single Beacon, Probe Response, Association Request, or Reassociation Request Frame can contain
// multiple Cipher Suites, and multiple AKM suites.  It is necessary to list all of the supported suites so that a
// proper security assessment can be performed on the Access Point.
//
// Following are how the enum values map to IEEE Std 802.11-2012
// 00-0F-AC:1 - IEEE_8021X
// 00-0F-AC:2 - PSK
// 00-0F-AC:3 - FT_IEEE_8021X
// 00-0F-AC:4 - FT_PSK
// 00-0F-AC:5 - IEEE_8021X_SHA256
// 00-0F-AC:6 - PSK_SHA256
// 00-0F-AC:7 - TDLS
// 00-0F-AC:8 - SAE
// 00-0F-AC:9 - FT_SAE
// <p>
// And OPEN is no AKM.
type AkmSuite int32

const (
	AkmSuite_UNKNOWN AkmSuite = 0
	// EAP AKM Methods
	AkmSuite_IEEE_8021X        AkmSuite = 1
	AkmSuite_FT_IEEE_8021X     AkmSuite = 3
	AkmSuite_IEEE_8021X_SHA256 AkmSuite = 5
	// PSK AKM Methods
	AkmSuite_PSK        AkmSuite = 2
	AkmSuite_FT_PSK     AkmSuite = 4
	AkmSuite_PSK_SHA256 AkmSuite = 6
	AkmSuite_SAE        AkmSuite = 8
	AkmSuite_FT_SAE     AkmSuite = 9
	// TDLS
	AkmSuite_TDLS AkmSuite = 7
	AkmSuite_OPEN AkmSuite = 13
)

// Enum value maps for AkmSuite.
var (
	AkmSuite_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "IEEE_8021X",
		3:  "FT_IEEE_8021X",
		5:  "IEEE_8021X_SHA256",
		2:  "PSK",
		4:  "FT_PSK",
		6:  "PSK_SHA256",
		8:  "SAE",
		9:  "FT_SAE",
		7:  "TDLS",
		13: "OPEN",
	}
	AkmSuite_value = map[string]int32{
		"UNKNOWN":           0,
		"IEEE_8021X":        1,
		"FT_IEEE_8021X":     3,
		"IEEE_8021X_SHA256": 5,
		"PSK":               2,
		"FT_PSK":            4,
		"PSK_SHA256":        6,
		"SAE":               8,
		"FT_SAE":            9,
		"TDLS":              7,
		"OPEN":              13,
	}
)

func (x AkmSuite) Enum() *AkmSuite {
	p := new(AkmSuite)
	*p = x
	return p
}

func (x AkmSuite) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AkmSuite) Descriptor() protoreflect.EnumDescriptor {
	return file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_enumTypes[0].Descriptor()
}

func (AkmSuite) Type() protoreflect.EnumType {
	return &file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_enumTypes[0]
}

func (x AkmSuite) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AkmSuite.Descriptor instead.
func (AkmSuite) EnumDescriptor() ([]byte, []int) {
	return file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_rawDescGZIP(), []int{0}
}

var File_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto protoreflect.FileDescriptor

var file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x77, 0x69, 0x66, 0x69, 0x2f, 0x61, 0x6b, 0x6d,
	0x73, 0x75, 0x69, 0x74, 0x65, 0x2f, 0x61, 0x6b, 0x6d, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x61, 0x78, 0x69,
	0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x77, 0x69, 0x66,
	0x69, 0x2e, 0x61, 0x6b, 0x6d, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2a, 0x9f, 0x01, 0x0a, 0x08, 0x41,
	0x6b, 0x6d, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x45, 0x45, 0x45, 0x5f, 0x38, 0x30, 0x32,
	0x31, 0x58, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x54, 0x5f, 0x49, 0x45, 0x45, 0x45, 0x5f,
	0x38, 0x30, 0x32, 0x31, 0x58, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x45, 0x45, 0x45, 0x5f,
	0x38, 0x30, 0x32, 0x31, 0x58, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x05, 0x12, 0x07,
	0x0a, 0x03, 0x50, 0x53, 0x4b, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x54, 0x5f, 0x50, 0x53,
	0x4b, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x53, 0x4b, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35,
	0x36, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x41, 0x45, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x54, 0x5f, 0x53, 0x41, 0x45, 0x10, 0x09, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x44, 0x4c, 0x53,
	0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x0d, 0x42, 0x35, 0x0a, 0x1a,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x77, 0x69, 0x66, 0x69, 0x50, 0x01, 0x5a, 0x15, 0x63, 0x72,
	0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_rawDescOnce sync.Once
	file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_rawDescData = file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_rawDesc
)

func file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_rawDescGZIP() []byte {
	file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_rawDescOnce.Do(func() {
		file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_rawDescData)
	})
	return file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_rawDescData
}

var file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_goTypes = []any{
	(AkmSuite)(0), // 0: com.craxiom.messaging.wifi.akmsuite.AkmSuite
}
var file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_init() }
func file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_init() {
	if File_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_goTypes,
		DependencyIndexes: file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_depIdxs,
		EnumInfos:         file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_enumTypes,
	}.Build()
	File_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto = out.File
	file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_rawDesc = nil
	file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_goTypes = nil
	file_com_craxiom_messaging_wifi_akmsuite_akm_suite_proto_depIdxs = nil
}
