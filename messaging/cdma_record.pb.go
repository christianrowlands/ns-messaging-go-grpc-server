//*
// A protobuf definition of the CDMA survey record defined in the Network Survey Messaging API.
//
// This protobuf definition is provided as a convenience only. See the official API documentation for the true CDMA
// Record message schema.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: com/craxiom/messaging/cdma_record.proto

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

type CdmaRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     string          `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	MessageType string          `protobuf:"bytes,2,opt,name=messageType,proto3" json:"messageType,omitempty"`
	Data        *CdmaRecordData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CdmaRecord) Reset() {
	*x = CdmaRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_craxiom_messaging_cdma_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CdmaRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CdmaRecord) ProtoMessage() {}

func (x *CdmaRecord) ProtoReflect() protoreflect.Message {
	mi := &file_com_craxiom_messaging_cdma_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CdmaRecord.ProtoReflect.Descriptor instead.
func (*CdmaRecord) Descriptor() ([]byte, []int) {
	return file_com_craxiom_messaging_cdma_record_proto_rawDescGZIP(), []int{0}
}

func (x *CdmaRecord) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CdmaRecord) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *CdmaRecord) GetData() *CdmaRecordData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CdmaRecordData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceSerialNumber  string                 `protobuf:"bytes,1,opt,name=deviceSerialNumber,proto3" json:"deviceSerialNumber,omitempty"`
	DeviceName          string                 `protobuf:"bytes,2,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	DeviceTime          string                 `protobuf:"bytes,3,opt,name=deviceTime,proto3" json:"deviceTime,omitempty"`
	Latitude            float64                `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude           float64                `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Altitude            float32                `protobuf:"fixed32,6,opt,name=altitude,proto3" json:"altitude,omitempty"`
	MissionId           string                 `protobuf:"bytes,7,opt,name=missionId,proto3" json:"missionId,omitempty"`
	RecordNumber        int32                  `protobuf:"varint,8,opt,name=recordNumber,proto3" json:"recordNumber,omitempty"`
	GroupNumber         int32                  `protobuf:"varint,9,opt,name=groupNumber,proto3" json:"groupNumber,omitempty"`
	Accuracy            int32                  `protobuf:"varint,10,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	Heading             float32                `protobuf:"fixed32,50,opt,name=heading,proto3" json:"heading,omitempty"`
	Pitch               float32                `protobuf:"fixed32,51,opt,name=pitch,proto3" json:"pitch,omitempty"`
	Roll                float32                `protobuf:"fixed32,52,opt,name=roll,proto3" json:"roll,omitempty"`
	FieldOfView         float32                `protobuf:"fixed32,53,opt,name=fieldOfView,proto3" json:"fieldOfView,omitempty"`
	ReceiverSensitivity float32                `protobuf:"fixed32,54,opt,name=receiverSensitivity,proto3" json:"receiverSensitivity,omitempty"`
	Speed               float32                `protobuf:"fixed32,55,opt,name=speed,proto3" json:"speed,omitempty"`
	Sid                 *wrapperspb.Int32Value `protobuf:"bytes,16,opt,name=sid,proto3" json:"sid,omitempty"`
	Nid                 *wrapperspb.Int32Value `protobuf:"bytes,17,opt,name=nid,proto3" json:"nid,omitempty"`
	Zone                *wrapperspb.Int32Value `protobuf:"bytes,18,opt,name=zone,proto3" json:"zone,omitempty"`
	Bsid                *wrapperspb.Int32Value `protobuf:"bytes,19,opt,name=bsid,proto3" json:"bsid,omitempty"`
	Channel             *wrapperspb.Int32Value `protobuf:"bytes,20,opt,name=channel,proto3" json:"channel,omitempty"`
	PnOffset            *wrapperspb.Int32Value `protobuf:"bytes,21,opt,name=pnOffset,proto3" json:"pnOffset,omitempty"`
	SignalStrength      *wrapperspb.FloatValue `protobuf:"bytes,22,opt,name=signalStrength,proto3" json:"signalStrength,omitempty"`
	Ecio                *wrapperspb.FloatValue `protobuf:"bytes,23,opt,name=ecio,proto3" json:"ecio,omitempty"`
	ServingCell         *wrapperspb.BoolValue  `protobuf:"bytes,25,opt,name=servingCell,proto3" json:"servingCell,omitempty"`
	Provider            string                 `protobuf:"bytes,27,opt,name=provider,proto3" json:"provider,omitempty"`
	Slot                *wrapperspb.Int32Value `protobuf:"bytes,28,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (x *CdmaRecordData) Reset() {
	*x = CdmaRecordData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_craxiom_messaging_cdma_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CdmaRecordData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CdmaRecordData) ProtoMessage() {}

func (x *CdmaRecordData) ProtoReflect() protoreflect.Message {
	mi := &file_com_craxiom_messaging_cdma_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CdmaRecordData.ProtoReflect.Descriptor instead.
func (*CdmaRecordData) Descriptor() ([]byte, []int) {
	return file_com_craxiom_messaging_cdma_record_proto_rawDescGZIP(), []int{1}
}

func (x *CdmaRecordData) GetDeviceSerialNumber() string {
	if x != nil {
		return x.DeviceSerialNumber
	}
	return ""
}

func (x *CdmaRecordData) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *CdmaRecordData) GetDeviceTime() string {
	if x != nil {
		return x.DeviceTime
	}
	return ""
}

func (x *CdmaRecordData) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *CdmaRecordData) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *CdmaRecordData) GetAltitude() float32 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *CdmaRecordData) GetMissionId() string {
	if x != nil {
		return x.MissionId
	}
	return ""
}

func (x *CdmaRecordData) GetRecordNumber() int32 {
	if x != nil {
		return x.RecordNumber
	}
	return 0
}

func (x *CdmaRecordData) GetGroupNumber() int32 {
	if x != nil {
		return x.GroupNumber
	}
	return 0
}

func (x *CdmaRecordData) GetAccuracy() int32 {
	if x != nil {
		return x.Accuracy
	}
	return 0
}

func (x *CdmaRecordData) GetHeading() float32 {
	if x != nil {
		return x.Heading
	}
	return 0
}

func (x *CdmaRecordData) GetPitch() float32 {
	if x != nil {
		return x.Pitch
	}
	return 0
}

func (x *CdmaRecordData) GetRoll() float32 {
	if x != nil {
		return x.Roll
	}
	return 0
}

func (x *CdmaRecordData) GetFieldOfView() float32 {
	if x != nil {
		return x.FieldOfView
	}
	return 0
}

func (x *CdmaRecordData) GetReceiverSensitivity() float32 {
	if x != nil {
		return x.ReceiverSensitivity
	}
	return 0
}

func (x *CdmaRecordData) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *CdmaRecordData) GetSid() *wrapperspb.Int32Value {
	if x != nil {
		return x.Sid
	}
	return nil
}

func (x *CdmaRecordData) GetNid() *wrapperspb.Int32Value {
	if x != nil {
		return x.Nid
	}
	return nil
}

func (x *CdmaRecordData) GetZone() *wrapperspb.Int32Value {
	if x != nil {
		return x.Zone
	}
	return nil
}

func (x *CdmaRecordData) GetBsid() *wrapperspb.Int32Value {
	if x != nil {
		return x.Bsid
	}
	return nil
}

func (x *CdmaRecordData) GetChannel() *wrapperspb.Int32Value {
	if x != nil {
		return x.Channel
	}
	return nil
}

func (x *CdmaRecordData) GetPnOffset() *wrapperspb.Int32Value {
	if x != nil {
		return x.PnOffset
	}
	return nil
}

func (x *CdmaRecordData) GetSignalStrength() *wrapperspb.FloatValue {
	if x != nil {
		return x.SignalStrength
	}
	return nil
}

func (x *CdmaRecordData) GetEcio() *wrapperspb.FloatValue {
	if x != nil {
		return x.Ecio
	}
	return nil
}

func (x *CdmaRecordData) GetServingCell() *wrapperspb.BoolValue {
	if x != nil {
		return x.ServingCell
	}
	return nil
}

func (x *CdmaRecordData) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CdmaRecordData) GetSlot() *wrapperspb.Int32Value {
	if x != nil {
		return x.Slot
	}
	return nil
}

var File_com_craxiom_messaging_cdma_record_proto protoreflect.FileDescriptor

var file_com_craxiom_messaging_cdma_record_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x64, 0x6d, 0x61, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x83, 0x01, 0x0a, 0x0a, 0x43, 0x64, 0x6d, 0x61, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x64, 0x6d, 0x61, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb5, 0x08, 0x0a, 0x0e, 0x43, 0x64, 0x6d, 0x61, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x12, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x32, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x69, 0x74,
	0x63, 0x68, 0x18, 0x33, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x69, 0x74, 0x63, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x6c, 0x18, 0x34, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x66, 0x56, 0x69,
	0x65, 0x77, 0x18, 0x35, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x66, 0x56, 0x69, 0x65, 0x77, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x36, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x18, 0x37, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x2d, 0x0a,
	0x03, 0x73, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x03,
	0x6e, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x6e, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x7a,
	0x6f, 0x6e, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x2f, 0x0a, 0x04,
	0x62, 0x73, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x62, 0x73, 0x69, 0x64, 0x12, 0x35, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x6e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x08, 0x70, 0x6e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x43, 0x0a,
	0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x12, 0x2f, 0x0a, 0x04, 0x65, 0x63, 0x69, 0x6f, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x65,
	0x63, 0x69, 0x6f, 0x12, 0x3c, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x65,
	0x6c, 0x6c, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6c,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x1b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a,
	0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x42, 0x30,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x01, 0x5a, 0x15, 0x63, 0x72, 0x61, 0x78, 0x69,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_craxiom_messaging_cdma_record_proto_rawDescOnce sync.Once
	file_com_craxiom_messaging_cdma_record_proto_rawDescData = file_com_craxiom_messaging_cdma_record_proto_rawDesc
)

func file_com_craxiom_messaging_cdma_record_proto_rawDescGZIP() []byte {
	file_com_craxiom_messaging_cdma_record_proto_rawDescOnce.Do(func() {
		file_com_craxiom_messaging_cdma_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_craxiom_messaging_cdma_record_proto_rawDescData)
	})
	return file_com_craxiom_messaging_cdma_record_proto_rawDescData
}

var file_com_craxiom_messaging_cdma_record_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_craxiom_messaging_cdma_record_proto_goTypes = []any{
	(*CdmaRecord)(nil),            // 0: com.craxiom.messaging.CdmaRecord
	(*CdmaRecordData)(nil),        // 1: com.craxiom.messaging.CdmaRecordData
	(*wrapperspb.Int32Value)(nil), // 2: google.protobuf.Int32Value
	(*wrapperspb.FloatValue)(nil), // 3: google.protobuf.FloatValue
	(*wrapperspb.BoolValue)(nil),  // 4: google.protobuf.BoolValue
}
var file_com_craxiom_messaging_cdma_record_proto_depIdxs = []int32{
	1,  // 0: com.craxiom.messaging.CdmaRecord.data:type_name -> com.craxiom.messaging.CdmaRecordData
	2,  // 1: com.craxiom.messaging.CdmaRecordData.sid:type_name -> google.protobuf.Int32Value
	2,  // 2: com.craxiom.messaging.CdmaRecordData.nid:type_name -> google.protobuf.Int32Value
	2,  // 3: com.craxiom.messaging.CdmaRecordData.zone:type_name -> google.protobuf.Int32Value
	2,  // 4: com.craxiom.messaging.CdmaRecordData.bsid:type_name -> google.protobuf.Int32Value
	2,  // 5: com.craxiom.messaging.CdmaRecordData.channel:type_name -> google.protobuf.Int32Value
	2,  // 6: com.craxiom.messaging.CdmaRecordData.pnOffset:type_name -> google.protobuf.Int32Value
	3,  // 7: com.craxiom.messaging.CdmaRecordData.signalStrength:type_name -> google.protobuf.FloatValue
	3,  // 8: com.craxiom.messaging.CdmaRecordData.ecio:type_name -> google.protobuf.FloatValue
	4,  // 9: com.craxiom.messaging.CdmaRecordData.servingCell:type_name -> google.protobuf.BoolValue
	2,  // 10: com.craxiom.messaging.CdmaRecordData.slot:type_name -> google.protobuf.Int32Value
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_com_craxiom_messaging_cdma_record_proto_init() }
func file_com_craxiom_messaging_cdma_record_proto_init() {
	if File_com_craxiom_messaging_cdma_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_craxiom_messaging_cdma_record_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CdmaRecord); i {
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
		file_com_craxiom_messaging_cdma_record_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CdmaRecordData); i {
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
			RawDescriptor: file_com_craxiom_messaging_cdma_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_craxiom_messaging_cdma_record_proto_goTypes,
		DependencyIndexes: file_com_craxiom_messaging_cdma_record_proto_depIdxs,
		MessageInfos:      file_com_craxiom_messaging_cdma_record_proto_msgTypes,
	}.Build()
	File_com_craxiom_messaging_cdma_record_proto = out.File
	file_com_craxiom_messaging_cdma_record_proto_rawDesc = nil
	file_com_craxiom_messaging_cdma_record_proto_goTypes = nil
	file_com_craxiom_messaging_cdma_record_proto_depIdxs = nil
}
