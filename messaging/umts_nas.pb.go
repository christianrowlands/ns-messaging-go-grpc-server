//*
// A protobuf definition of the UMTS NAS and UMTS NAS DSDS message defined in the Network Survey Messaging API.
//
// This protobuf definition is provided as a convenience only. See the official API documentation for the true UMTS
// Record message schema.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: com/craxiom/messaging/umts_nas.proto

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

type UmtsNas struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     string       `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	MessageType string       `protobuf:"bytes,2,opt,name=messageType,proto3" json:"messageType,omitempty"`
	Data        *UmtsNasData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UmtsNas) Reset() {
	*x = UmtsNas{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_craxiom_messaging_umts_nas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UmtsNas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UmtsNas) ProtoMessage() {}

func (x *UmtsNas) ProtoReflect() protoreflect.Message {
	mi := &file_com_craxiom_messaging_umts_nas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UmtsNas.ProtoReflect.Descriptor instead.
func (*UmtsNas) Descriptor() ([]byte, []int) {
	return file_com_craxiom_messaging_umts_nas_proto_rawDescGZIP(), []int{0}
}

func (x *UmtsNas) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *UmtsNas) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *UmtsNas) GetData() *UmtsNasData {
	if x != nil {
		return x.Data
	}
	return nil
}

type UmtsNasData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceSerialNumber string  `protobuf:"bytes,1,opt,name=deviceSerialNumber,proto3" json:"deviceSerialNumber,omitempty"`
	DeviceName         string  `protobuf:"bytes,2,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	DeviceTime         string  `protobuf:"bytes,3,opt,name=deviceTime,proto3" json:"deviceTime,omitempty"`
	Latitude           float64 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude          float64 `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Altitude           float32 `protobuf:"fixed32,6,opt,name=altitude,proto3" json:"altitude,omitempty"`
	MissionId          string  `protobuf:"bytes,7,opt,name=missionId,proto3" json:"missionId,omitempty"`
	Accuracy           int32   `protobuf:"varint,8,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	// Orientation of sensor and sensor parameters. optional.
	Heading             float32 `protobuf:"fixed32,50,opt,name=heading,proto3" json:"heading,omitempty"`
	Pitch               float32 `protobuf:"fixed32,51,opt,name=pitch,proto3" json:"pitch,omitempty"`
	Roll                float32 `protobuf:"fixed32,52,opt,name=roll,proto3" json:"roll,omitempty"`
	FieldOfView         float32 `protobuf:"fixed32,53,opt,name=fieldOfView,proto3" json:"fieldOfView,omitempty"`
	ReceiverSensitivity float32 `protobuf:"fixed32,54,opt,name=receiverSensitivity,proto3" json:"receiverSensitivity,omitempty"`
	Speed               float32 `protobuf:"fixed32,55,opt,name=speed,proto3" json:"speed,omitempty"`
	PcapRecord          []byte  `protobuf:"bytes,11,opt,name=pcapRecord,proto3" json:"pcapRecord,omitempty"`
}

func (x *UmtsNasData) Reset() {
	*x = UmtsNasData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_craxiom_messaging_umts_nas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UmtsNasData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UmtsNasData) ProtoMessage() {}

func (x *UmtsNasData) ProtoReflect() protoreflect.Message {
	mi := &file_com_craxiom_messaging_umts_nas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UmtsNasData.ProtoReflect.Descriptor instead.
func (*UmtsNasData) Descriptor() ([]byte, []int) {
	return file_com_craxiom_messaging_umts_nas_proto_rawDescGZIP(), []int{1}
}

func (x *UmtsNasData) GetDeviceSerialNumber() string {
	if x != nil {
		return x.DeviceSerialNumber
	}
	return ""
}

func (x *UmtsNasData) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *UmtsNasData) GetDeviceTime() string {
	if x != nil {
		return x.DeviceTime
	}
	return ""
}

func (x *UmtsNasData) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *UmtsNasData) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *UmtsNasData) GetAltitude() float32 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *UmtsNasData) GetMissionId() string {
	if x != nil {
		return x.MissionId
	}
	return ""
}

func (x *UmtsNasData) GetAccuracy() int32 {
	if x != nil {
		return x.Accuracy
	}
	return 0
}

func (x *UmtsNasData) GetHeading() float32 {
	if x != nil {
		return x.Heading
	}
	return 0
}

func (x *UmtsNasData) GetPitch() float32 {
	if x != nil {
		return x.Pitch
	}
	return 0
}

func (x *UmtsNasData) GetRoll() float32 {
	if x != nil {
		return x.Roll
	}
	return 0
}

func (x *UmtsNasData) GetFieldOfView() float32 {
	if x != nil {
		return x.FieldOfView
	}
	return 0
}

func (x *UmtsNasData) GetReceiverSensitivity() float32 {
	if x != nil {
		return x.ReceiverSensitivity
	}
	return 0
}

func (x *UmtsNasData) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *UmtsNasData) GetPcapRecord() []byte {
	if x != nil {
		return x.PcapRecord
	}
	return nil
}

var File_com_craxiom_messaging_umts_nas_proto protoreflect.FileDescriptor

var file_com_craxiom_messaging_umts_nas_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x75, 0x6d, 0x74, 0x73, 0x5f, 0x6e, 0x61, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x61, 0x78,
	0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x0a,
	0x07, 0x55, 0x6d, 0x74, 0x73, 0x4e, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x6d, 0x74, 0x73, 0x4e,
	0x61, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xdb, 0x03, 0x0a,
	0x0b, 0x55, 0x6d, 0x74, 0x73, 0x4e, 0x61, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x12,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x32, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x18,
	0x33, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x6c, 0x18, 0x34, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x6c,
	0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x66, 0x56, 0x69, 0x65, 0x77, 0x18,
	0x35, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x66, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x36, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x13, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x37, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x63,
	0x61, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x70, 0x63, 0x61, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x30, 0x0a, 0x15, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x50, 0x01, 0x5a, 0x15, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_craxiom_messaging_umts_nas_proto_rawDescOnce sync.Once
	file_com_craxiom_messaging_umts_nas_proto_rawDescData = file_com_craxiom_messaging_umts_nas_proto_rawDesc
)

func file_com_craxiom_messaging_umts_nas_proto_rawDescGZIP() []byte {
	file_com_craxiom_messaging_umts_nas_proto_rawDescOnce.Do(func() {
		file_com_craxiom_messaging_umts_nas_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_craxiom_messaging_umts_nas_proto_rawDescData)
	})
	return file_com_craxiom_messaging_umts_nas_proto_rawDescData
}

var file_com_craxiom_messaging_umts_nas_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_craxiom_messaging_umts_nas_proto_goTypes = []any{
	(*UmtsNas)(nil),     // 0: com.craxiom.messaging.UmtsNas
	(*UmtsNasData)(nil), // 1: com.craxiom.messaging.UmtsNasData
}
var file_com_craxiom_messaging_umts_nas_proto_depIdxs = []int32{
	1, // 0: com.craxiom.messaging.UmtsNas.data:type_name -> com.craxiom.messaging.UmtsNasData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_craxiom_messaging_umts_nas_proto_init() }
func file_com_craxiom_messaging_umts_nas_proto_init() {
	if File_com_craxiom_messaging_umts_nas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_craxiom_messaging_umts_nas_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UmtsNas); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_com_craxiom_messaging_umts_nas_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UmtsNasData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_craxiom_messaging_umts_nas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_craxiom_messaging_umts_nas_proto_goTypes,
		DependencyIndexes: file_com_craxiom_messaging_umts_nas_proto_depIdxs,
		MessageInfos:      file_com_craxiom_messaging_umts_nas_proto_msgTypes,
	}.Build()
	File_com_craxiom_messaging_umts_nas_proto = out.File
	file_com_craxiom_messaging_umts_nas_proto_rawDesc = nil
	file_com_craxiom_messaging_umts_nas_proto_goTypes = nil
	file_com_craxiom_messaging_umts_nas_proto_depIdxs = nil
}
