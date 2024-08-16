//*
// Defines the Wi-Fi Bandwidth values used in the 802.11 Survey message.
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
// source: com/craxiom/messaging/wifi/bandwidth/wifi_bandwidth.proto

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

type WifiBandwidth int32

const (
	WifiBandwidth_UNKNOWN     WifiBandwidth = 0
	WifiBandwidth_MHZ_20      WifiBandwidth = 1
	WifiBandwidth_MHZ_40      WifiBandwidth = 2
	WifiBandwidth_MHZ_80      WifiBandwidth = 3
	WifiBandwidth_MHZ_80_PLUS WifiBandwidth = 4
	WifiBandwidth_MHZ_160     WifiBandwidth = 5
	WifiBandwidth_MHZ_320     WifiBandwidth = 6
)

// Enum value maps for WifiBandwidth.
var (
	WifiBandwidth_name = map[int32]string{
		0: "UNKNOWN",
		1: "MHZ_20",
		2: "MHZ_40",
		3: "MHZ_80",
		4: "MHZ_80_PLUS",
		5: "MHZ_160",
		6: "MHZ_320",
	}
	WifiBandwidth_value = map[string]int32{
		"UNKNOWN":     0,
		"MHZ_20":      1,
		"MHZ_40":      2,
		"MHZ_80":      3,
		"MHZ_80_PLUS": 4,
		"MHZ_160":     5,
		"MHZ_320":     6,
	}
)

func (x WifiBandwidth) Enum() *WifiBandwidth {
	p := new(WifiBandwidth)
	*p = x
	return p
}

func (x WifiBandwidth) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WifiBandwidth) Descriptor() protoreflect.EnumDescriptor {
	return file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_enumTypes[0].Descriptor()
}

func (WifiBandwidth) Type() protoreflect.EnumType {
	return &file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_enumTypes[0]
}

func (x WifiBandwidth) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WifiBandwidth.Descriptor instead.
func (WifiBandwidth) EnumDescriptor() ([]byte, []int) {
	return file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_rawDescGZIP(), []int{0}
}

var File_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto protoreflect.FileDescriptor

var file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_rawDesc = []byte{
	0x0a, 0x39, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x77, 0x69, 0x66, 0x69, 0x2f, 0x62, 0x61, 0x6e,
	0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x2f, 0x77, 0x69, 0x66, 0x69, 0x5f, 0x62, 0x61, 0x6e, 0x64,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x77, 0x69, 0x66, 0x69, 0x2e, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x2a, 0x6b, 0x0a, 0x0d, 0x57, 0x69, 0x66, 0x69, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4d, 0x48, 0x5a, 0x5f, 0x32, 0x30, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x48, 0x5a, 0x5f, 0x34, 0x30, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x48, 0x5a, 0x5f, 0x38,
	0x30, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x48, 0x5a, 0x5f, 0x38, 0x30, 0x5f, 0x50, 0x4c,
	0x55, 0x53, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x48, 0x5a, 0x5f, 0x31, 0x36, 0x30, 0x10,
	0x05, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x48, 0x5a, 0x5f, 0x33, 0x32, 0x30, 0x10, 0x06, 0x42, 0x35,
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x77, 0x69, 0x66, 0x69, 0x50, 0x01, 0x5a, 0x15,
	0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_rawDescOnce sync.Once
	file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_rawDescData = file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_rawDesc
)

func file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_rawDescGZIP() []byte {
	file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_rawDescOnce.Do(func() {
		file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_rawDescData)
	})
	return file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_rawDescData
}

var file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_goTypes = []any{
	(WifiBandwidth)(0), // 0: com.craxiom.messaging.wifi.bandwidth.WifiBandwidth
}
var file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_init() }
func file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_init() {
	if File_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_goTypes,
		DependencyIndexes: file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_depIdxs,
		EnumInfos:         file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_enumTypes,
	}.Build()
	File_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto = out.File
	file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_rawDesc = nil
	file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_goTypes = nil
	file_com_craxiom_messaging_wifi_bandwidth_wifi_bandwidth_proto_depIdxs = nil
}
