//*
// A protobuf definition of the Device Status message defined in the Network Survey Messaging API.
//
// This protobuf definition is provided as a convenience only. See the official API documentation for the true Device
// Status message schema.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: com/craxiom/messaging/device_status.proto

package messaging

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeviceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     string            `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	MessageType string            `protobuf:"bytes,2,opt,name=messageType,proto3" json:"messageType,omitempty"`
	Data        *DeviceStatusData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DeviceStatus) Reset() {
	*x = DeviceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_craxiom_messaging_device_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceStatus) ProtoMessage() {}

func (x *DeviceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_com_craxiom_messaging_device_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceStatus.ProtoReflect.Descriptor instead.
func (*DeviceStatus) Descriptor() ([]byte, []int) {
	return file_com_craxiom_messaging_device_status_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceStatus) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *DeviceStatus) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *DeviceStatus) GetData() *DeviceStatusData {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeviceStatusData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceSerialNumber  string                 `protobuf:"bytes,1,opt,name=deviceSerialNumber,proto3" json:"deviceSerialNumber,omitempty"`
	DeviceName          string                 `protobuf:"bytes,2,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	DeviceTime          string                 `protobuf:"bytes,3,opt,name=deviceTime,proto3" json:"deviceTime,omitempty"`
	Latitude            float64                `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude           float64                `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Altitude            float32                `protobuf:"fixed32,6,opt,name=altitude,proto3" json:"altitude,omitempty"`
	DeviceModel         string                 `protobuf:"bytes,8,opt,name=deviceModel,proto3" json:"deviceModel,omitempty"`
	Accuracy            int32                  `protobuf:"varint,9,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	Heading             float32                `protobuf:"fixed32,50,opt,name=heading,proto3" json:"heading,omitempty"`
	Pitch               float32                `protobuf:"fixed32,51,opt,name=pitch,proto3" json:"pitch,omitempty"`
	Roll                float32                `protobuf:"fixed32,52,opt,name=roll,proto3" json:"roll,omitempty"`
	FieldOfView         float32                `protobuf:"fixed32,53,opt,name=fieldOfView,proto3" json:"fieldOfView,omitempty"`
	ReceiverSensitivity float32                `protobuf:"fixed32,54,opt,name=receiverSensitivity,proto3" json:"receiverSensitivity,omitempty"`
	Speed               float32                `protobuf:"fixed32,55,opt,name=speed,proto3" json:"speed,omitempty"`
	BatteryLevelPercent *wrapperspb.Int32Value `protobuf:"bytes,7,opt,name=batteryLevelPercent,proto3" json:"batteryLevelPercent,omitempty"`
	Error               *Error                 `protobuf:"bytes,10,opt,name=error,proto3" json:"error,omitempty"`
	MdmOverride         *wrapperspb.BoolValue  `protobuf:"bytes,11,opt,name=mdmOverride,proto3" json:"mdmOverride,omitempty"`
	AppVersion          string                 `protobuf:"bytes,12,opt,name=appVersion,proto3" json:"appVersion,omitempty"`
	GnssLatitude        float64                `protobuf:"fixed64,15,opt,name=gnssLatitude,proto3" json:"gnssLatitude,omitempty"`
	GnssLongitude       float64                `protobuf:"fixed64,16,opt,name=gnssLongitude,proto3" json:"gnssLongitude,omitempty"`
	GnssAltitude        float32                `protobuf:"fixed32,17,opt,name=gnssAltitude,proto3" json:"gnssAltitude,omitempty"`
	GnssAccuracy        int32                  `protobuf:"varint,18,opt,name=gnssAccuracy,proto3" json:"gnssAccuracy,omitempty"`
	NetworkLatitude     float64                `protobuf:"fixed64,19,opt,name=networkLatitude,proto3" json:"networkLatitude,omitempty"`
	NetworkLongitude    float64                `protobuf:"fixed64,20,opt,name=networkLongitude,proto3" json:"networkLongitude,omitempty"`
	NetworkAltitude     float32                `protobuf:"fixed32,21,opt,name=networkAltitude,proto3" json:"networkAltitude,omitempty"`
	NetworkAccuracy     int32                  `protobuf:"varint,22,opt,name=networkAccuracy,proto3" json:"networkAccuracy,omitempty"`
}

func (x *DeviceStatusData) Reset() {
	*x = DeviceStatusData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_craxiom_messaging_device_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceStatusData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceStatusData) ProtoMessage() {}

func (x *DeviceStatusData) ProtoReflect() protoreflect.Message {
	mi := &file_com_craxiom_messaging_device_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceStatusData.ProtoReflect.Descriptor instead.
func (*DeviceStatusData) Descriptor() ([]byte, []int) {
	return file_com_craxiom_messaging_device_status_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceStatusData) GetDeviceSerialNumber() string {
	if x != nil {
		return x.DeviceSerialNumber
	}
	return ""
}

func (x *DeviceStatusData) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *DeviceStatusData) GetDeviceTime() string {
	if x != nil {
		return x.DeviceTime
	}
	return ""
}

func (x *DeviceStatusData) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *DeviceStatusData) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *DeviceStatusData) GetAltitude() float32 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *DeviceStatusData) GetDeviceModel() string {
	if x != nil {
		return x.DeviceModel
	}
	return ""
}

func (x *DeviceStatusData) GetAccuracy() int32 {
	if x != nil {
		return x.Accuracy
	}
	return 0
}

func (x *DeviceStatusData) GetHeading() float32 {
	if x != nil {
		return x.Heading
	}
	return 0
}

func (x *DeviceStatusData) GetPitch() float32 {
	if x != nil {
		return x.Pitch
	}
	return 0
}

func (x *DeviceStatusData) GetRoll() float32 {
	if x != nil {
		return x.Roll
	}
	return 0
}

func (x *DeviceStatusData) GetFieldOfView() float32 {
	if x != nil {
		return x.FieldOfView
	}
	return 0
}

func (x *DeviceStatusData) GetReceiverSensitivity() float32 {
	if x != nil {
		return x.ReceiverSensitivity
	}
	return 0
}

func (x *DeviceStatusData) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *DeviceStatusData) GetBatteryLevelPercent() *wrapperspb.Int32Value {
	if x != nil {
		return x.BatteryLevelPercent
	}
	return nil
}

func (x *DeviceStatusData) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *DeviceStatusData) GetMdmOverride() *wrapperspb.BoolValue {
	if x != nil {
		return x.MdmOverride
	}
	return nil
}

func (x *DeviceStatusData) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

func (x *DeviceStatusData) GetGnssLatitude() float64 {
	if x != nil {
		return x.GnssLatitude
	}
	return 0
}

func (x *DeviceStatusData) GetGnssLongitude() float64 {
	if x != nil {
		return x.GnssLongitude
	}
	return 0
}

func (x *DeviceStatusData) GetGnssAltitude() float32 {
	if x != nil {
		return x.GnssAltitude
	}
	return 0
}

func (x *DeviceStatusData) GetGnssAccuracy() int32 {
	if x != nil {
		return x.GnssAccuracy
	}
	return 0
}

func (x *DeviceStatusData) GetNetworkLatitude() float64 {
	if x != nil {
		return x.NetworkLatitude
	}
	return 0
}

func (x *DeviceStatusData) GetNetworkLongitude() float64 {
	if x != nil {
		return x.NetworkLongitude
	}
	return 0
}

func (x *DeviceStatusData) GetNetworkAltitude() float32 {
	if x != nil {
		return x.NetworkAltitude
	}
	return 0
}

func (x *DeviceStatusData) GetNetworkAccuracy() int32 {
	if x != nil {
		return x.NetworkAccuracy
	}
	return 0
}

// Used to report an error with the device
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMessage string `protobuf:"bytes,1,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_craxiom_messaging_device_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_com_craxiom_messaging_device_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_com_craxiom_messaging_device_status_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_com_craxiom_messaging_device_status_proto protoreflect.FileDescriptor

var file_com_craxiom_messaging_device_status_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xe1, 0x07, 0x0a,
	0x10, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x2e, 0x0a, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x61,
	0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63,
	0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x63, 0x63,
	0x75, 0x72, 0x61, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x18, 0x33, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x69, 0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x6c, 0x18, 0x34, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x66, 0x56, 0x69, 0x65, 0x77, 0x18, 0x35, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x66, 0x56, 0x69, 0x65, 0x77, 0x12, 0x30, 0x0a, 0x13, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x18, 0x36, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x37, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x12, 0x4d, 0x0a, 0x13, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x13, 0x62,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x0b, 0x6d, 0x64, 0x6d, 0x4f, 0x76, 0x65,
	0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x6d, 0x64, 0x6d, 0x4f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x6e, 0x73, 0x73, 0x4c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x67, 0x6e, 0x73, 0x73,
	0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x67, 0x6e, 0x73, 0x73,
	0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x67, 0x6e, 0x73, 0x73, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x67, 0x6e, 0x73, 0x73, 0x41, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x67, 0x6e, 0x73, 0x73, 0x41, 0x6c, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x6e, 0x73, 0x73, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61,
	0x63, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x67, 0x6e, 0x73, 0x73, 0x41, 0x63,
	0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x6c,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79,
	0x22, 0x2b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x30, 0x0a,
	0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x01, 0x5a, 0x15, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_craxiom_messaging_device_status_proto_rawDescOnce sync.Once
	file_com_craxiom_messaging_device_status_proto_rawDescData = file_com_craxiom_messaging_device_status_proto_rawDesc
)

func file_com_craxiom_messaging_device_status_proto_rawDescGZIP() []byte {
	file_com_craxiom_messaging_device_status_proto_rawDescOnce.Do(func() {
		file_com_craxiom_messaging_device_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_craxiom_messaging_device_status_proto_rawDescData)
	})
	return file_com_craxiom_messaging_device_status_proto_rawDescData
}

var file_com_craxiom_messaging_device_status_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_craxiom_messaging_device_status_proto_goTypes = []any{
	(*DeviceStatus)(nil),          // 0: com.craxiom.messaging.DeviceStatus
	(*DeviceStatusData)(nil),      // 1: com.craxiom.messaging.DeviceStatusData
	(*Error)(nil),                 // 2: com.craxiom.messaging.Error
	(*wrapperspb.Int32Value)(nil), // 3: google.protobuf.Int32Value
	(*wrapperspb.BoolValue)(nil),  // 4: google.protobuf.BoolValue
}
var file_com_craxiom_messaging_device_status_proto_depIdxs = []int32{
	1, // 0: com.craxiom.messaging.DeviceStatus.data:type_name -> com.craxiom.messaging.DeviceStatusData
	3, // 1: com.craxiom.messaging.DeviceStatusData.batteryLevelPercent:type_name -> google.protobuf.Int32Value
	2, // 2: com.craxiom.messaging.DeviceStatusData.error:type_name -> com.craxiom.messaging.Error
	4, // 3: com.craxiom.messaging.DeviceStatusData.mdmOverride:type_name -> google.protobuf.BoolValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_craxiom_messaging_device_status_proto_init() }
func file_com_craxiom_messaging_device_status_proto_init() {
	if File_com_craxiom_messaging_device_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_craxiom_messaging_device_status_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DeviceStatus); i {
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
		file_com_craxiom_messaging_device_status_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DeviceStatusData); i {
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
		file_com_craxiom_messaging_device_status_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_com_craxiom_messaging_device_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_craxiom_messaging_device_status_proto_goTypes,
		DependencyIndexes: file_com_craxiom_messaging_device_status_proto_depIdxs,
		MessageInfos:      file_com_craxiom_messaging_device_status_proto_msgTypes,
	}.Build()
	File_com_craxiom_messaging_device_status_proto = out.File
	file_com_craxiom_messaging_device_status_proto_rawDesc = nil
	file_com_craxiom_messaging_device_status_proto_goTypes = nil
	file_com_craxiom_messaging_device_status_proto_depIdxs = nil
}
