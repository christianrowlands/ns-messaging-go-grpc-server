//*
// A protobuf definition of the GNSS survey record defined in the Network Survey Messaging API.
//
// This protobuf definition is provided as a convenience only. See the official API documentation for the true GNSS
// Record message schema.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: com/craxiom/messaging/gnss_record.proto

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

type GnssRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     string          `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	MessageType string          `protobuf:"bytes,2,opt,name=messageType,proto3" json:"messageType,omitempty"`
	Data        *GnssRecordData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GnssRecord) Reset() {
	*x = GnssRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_craxiom_messaging_gnss_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GnssRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GnssRecord) ProtoMessage() {}

func (x *GnssRecord) ProtoReflect() protoreflect.Message {
	mi := &file_com_craxiom_messaging_gnss_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GnssRecord.ProtoReflect.Descriptor instead.
func (*GnssRecord) Descriptor() ([]byte, []int) {
	return file_com_craxiom_messaging_gnss_record_proto_rawDescGZIP(), []int{0}
}

func (x *GnssRecord) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GnssRecord) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *GnssRecord) GetData() *GnssRecordData {
	if x != nil {
		return x.Data
	}
	return nil
}

type GnssRecordData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceSerialNumber  string                  `protobuf:"bytes,1,opt,name=deviceSerialNumber,proto3" json:"deviceSerialNumber,omitempty"`
	DeviceName          string                  `protobuf:"bytes,2,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	DeviceTime          string                  `protobuf:"bytes,3,opt,name=deviceTime,proto3" json:"deviceTime,omitempty"`
	Latitude            float64                 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude           float64                 `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Altitude            float32                 `protobuf:"fixed32,6,opt,name=altitude,proto3" json:"altitude,omitempty"`
	MissionId           string                  `protobuf:"bytes,7,opt,name=missionId,proto3" json:"missionId,omitempty"`
	RecordNumber        int32                   `protobuf:"varint,8,opt,name=recordNumber,proto3" json:"recordNumber,omitempty"`
	GroupNumber         int32                   `protobuf:"varint,9,opt,name=groupNumber,proto3" json:"groupNumber,omitempty"`
	DeviceModel         string                  `protobuf:"bytes,10,opt,name=deviceModel,proto3" json:"deviceModel,omitempty"`
	Accuracy            int32                   `protobuf:"varint,11,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	Heading             float32                 `protobuf:"fixed32,50,opt,name=heading,proto3" json:"heading,omitempty"`
	Pitch               float32                 `protobuf:"fixed32,51,opt,name=pitch,proto3" json:"pitch,omitempty"`
	Roll                float32                 `protobuf:"fixed32,52,opt,name=roll,proto3" json:"roll,omitempty"`
	FieldOfView         float32                 `protobuf:"fixed32,53,opt,name=fieldOfView,proto3" json:"fieldOfView,omitempty"`
	ReceiverSensitivity float32                 `protobuf:"fixed32,54,opt,name=receiverSensitivity,proto3" json:"receiverSensitivity,omitempty"`
	Speed               float32                 `protobuf:"fixed32,55,opt,name=speed,proto3" json:"speed,omitempty"`
	Constellation       Constellation           `protobuf:"varint,12,opt,name=constellation,proto3,enum=com.craxiom.messaging.gnss.constellation.Constellation" json:"constellation,omitempty"`
	SpaceVehicleId      *wrapperspb.UInt32Value `protobuf:"bytes,13,opt,name=spaceVehicleId,proto3" json:"spaceVehicleId,omitempty"`
	CarrierFreqHz       *wrapperspb.UInt64Value `protobuf:"bytes,14,opt,name=carrierFreqHz,proto3" json:"carrierFreqHz,omitempty"`
	ClockOffset         *wrapperspb.DoubleValue `protobuf:"bytes,15,opt,name=clockOffset,proto3" json:"clockOffset,omitempty"`
	UsedInSolution      *wrapperspb.BoolValue   `protobuf:"bytes,16,opt,name=usedInSolution,proto3" json:"usedInSolution,omitempty"`
	UndulationM         *wrapperspb.FloatValue  `protobuf:"bytes,17,opt,name=undulationM,proto3" json:"undulationM,omitempty"`
	LatitudeStdDevM     *wrapperspb.FloatValue  `protobuf:"bytes,18,opt,name=latitudeStdDevM,proto3" json:"latitudeStdDevM,omitempty"`
	LongitudeStdDevM    *wrapperspb.FloatValue  `protobuf:"bytes,19,opt,name=longitudeStdDevM,proto3" json:"longitudeStdDevM,omitempty"`
	AltitudeStdDevM     *wrapperspb.FloatValue  `protobuf:"bytes,20,opt,name=altitudeStdDevM,proto3" json:"altitudeStdDevM,omitempty"`
	AgcDb               *wrapperspb.FloatValue  `protobuf:"bytes,21,opt,name=agcDb,proto3" json:"agcDb,omitempty"`
	Cn0DbHz             *wrapperspb.FloatValue  `protobuf:"bytes,22,opt,name=cn0DbHz,proto3" json:"cn0DbHz,omitempty"`
	Hdop                *wrapperspb.FloatValue  `protobuf:"bytes,23,opt,name=hdop,proto3" json:"hdop,omitempty"`
	Vdop                *wrapperspb.FloatValue  `protobuf:"bytes,24,opt,name=vdop,proto3" json:"vdop,omitempty"`
}

func (x *GnssRecordData) Reset() {
	*x = GnssRecordData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_craxiom_messaging_gnss_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GnssRecordData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GnssRecordData) ProtoMessage() {}

func (x *GnssRecordData) ProtoReflect() protoreflect.Message {
	mi := &file_com_craxiom_messaging_gnss_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GnssRecordData.ProtoReflect.Descriptor instead.
func (*GnssRecordData) Descriptor() ([]byte, []int) {
	return file_com_craxiom_messaging_gnss_record_proto_rawDescGZIP(), []int{1}
}

func (x *GnssRecordData) GetDeviceSerialNumber() string {
	if x != nil {
		return x.DeviceSerialNumber
	}
	return ""
}

func (x *GnssRecordData) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *GnssRecordData) GetDeviceTime() string {
	if x != nil {
		return x.DeviceTime
	}
	return ""
}

func (x *GnssRecordData) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GnssRecordData) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *GnssRecordData) GetAltitude() float32 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *GnssRecordData) GetMissionId() string {
	if x != nil {
		return x.MissionId
	}
	return ""
}

func (x *GnssRecordData) GetRecordNumber() int32 {
	if x != nil {
		return x.RecordNumber
	}
	return 0
}

func (x *GnssRecordData) GetGroupNumber() int32 {
	if x != nil {
		return x.GroupNumber
	}
	return 0
}

func (x *GnssRecordData) GetDeviceModel() string {
	if x != nil {
		return x.DeviceModel
	}
	return ""
}

func (x *GnssRecordData) GetAccuracy() int32 {
	if x != nil {
		return x.Accuracy
	}
	return 0
}

func (x *GnssRecordData) GetHeading() float32 {
	if x != nil {
		return x.Heading
	}
	return 0
}

func (x *GnssRecordData) GetPitch() float32 {
	if x != nil {
		return x.Pitch
	}
	return 0
}

func (x *GnssRecordData) GetRoll() float32 {
	if x != nil {
		return x.Roll
	}
	return 0
}

func (x *GnssRecordData) GetFieldOfView() float32 {
	if x != nil {
		return x.FieldOfView
	}
	return 0
}

func (x *GnssRecordData) GetReceiverSensitivity() float32 {
	if x != nil {
		return x.ReceiverSensitivity
	}
	return 0
}

func (x *GnssRecordData) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *GnssRecordData) GetConstellation() Constellation {
	if x != nil {
		return x.Constellation
	}
	return Constellation_UNKNOWN
}

func (x *GnssRecordData) GetSpaceVehicleId() *wrapperspb.UInt32Value {
	if x != nil {
		return x.SpaceVehicleId
	}
	return nil
}

func (x *GnssRecordData) GetCarrierFreqHz() *wrapperspb.UInt64Value {
	if x != nil {
		return x.CarrierFreqHz
	}
	return nil
}

func (x *GnssRecordData) GetClockOffset() *wrapperspb.DoubleValue {
	if x != nil {
		return x.ClockOffset
	}
	return nil
}

func (x *GnssRecordData) GetUsedInSolution() *wrapperspb.BoolValue {
	if x != nil {
		return x.UsedInSolution
	}
	return nil
}

func (x *GnssRecordData) GetUndulationM() *wrapperspb.FloatValue {
	if x != nil {
		return x.UndulationM
	}
	return nil
}

func (x *GnssRecordData) GetLatitudeStdDevM() *wrapperspb.FloatValue {
	if x != nil {
		return x.LatitudeStdDevM
	}
	return nil
}

func (x *GnssRecordData) GetLongitudeStdDevM() *wrapperspb.FloatValue {
	if x != nil {
		return x.LongitudeStdDevM
	}
	return nil
}

func (x *GnssRecordData) GetAltitudeStdDevM() *wrapperspb.FloatValue {
	if x != nil {
		return x.AltitudeStdDevM
	}
	return nil
}

func (x *GnssRecordData) GetAgcDb() *wrapperspb.FloatValue {
	if x != nil {
		return x.AgcDb
	}
	return nil
}

func (x *GnssRecordData) GetCn0DbHz() *wrapperspb.FloatValue {
	if x != nil {
		return x.Cn0DbHz
	}
	return nil
}

func (x *GnssRecordData) GetHdop() *wrapperspb.FloatValue {
	if x != nil {
		return x.Hdop
	}
	return nil
}

func (x *GnssRecordData) GetVdop() *wrapperspb.FloatValue {
	if x != nil {
		return x.Vdop
	}
	return nil
}

var File_com_craxiom_messaging_gnss_record_proto protoreflect.FileDescriptor

var file_com_craxiom_messaging_gnss_record_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x6e, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x6e, 0x73, 0x73, 0x2f, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x83, 0x01, 0x0a, 0x0a, 0x47, 0x6e, 0x73, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x47, 0x6e, 0x73, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xf5, 0x0a, 0x0a, 0x0e, 0x47, 0x6e, 0x73, 0x73, 0x52,
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
	0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x32, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x69, 0x74, 0x63, 0x68, 0x18, 0x33, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x69, 0x74, 0x63,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x6c, 0x18, 0x34, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x66,
	0x56, 0x69, 0x65, 0x77, 0x18, 0x35, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x66, 0x56, 0x69, 0x65, 0x77, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x36,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x18, 0x37, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12,
	0x5d, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x61,
	0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x67,
	0x6e, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44,
	0x0a, 0x0e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x0d, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x46,
	0x72, 0x65, 0x71, 0x48, 0x7a, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x63, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x46, 0x72, 0x65, 0x71, 0x48, 0x7a, 0x12, 0x3e, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x63, 0x6c, 0x6f,
	0x63, 0x6b, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x42, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x64,
	0x49, 0x6e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x75, 0x73,
	0x65, 0x64, 0x49, 0x6e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0b,
	0x75, 0x6e, 0x64, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b,
	0x75, 0x6e, 0x64, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x12, 0x45, 0x0a, 0x0f, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x53, 0x74, 0x64, 0x44, 0x65, 0x76, 0x4d, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0f, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x53, 0x74, 0x64, 0x44, 0x65,
	0x76, 0x4d, 0x12, 0x47, 0x0a, 0x10, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x53,
	0x74, 0x64, 0x44, 0x65, 0x76, 0x4d, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x53, 0x74, 0x64, 0x44, 0x65, 0x76, 0x4d, 0x12, 0x45, 0x0a, 0x0f, 0x61,
	0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x53, 0x74, 0x64, 0x44, 0x65, 0x76, 0x4d, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0f, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x53, 0x74, 0x64, 0x44, 0x65,
	0x76, 0x4d, 0x12, 0x31, 0x0a, 0x05, 0x61, 0x67, 0x63, 0x44, 0x62, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x61, 0x67, 0x63, 0x44, 0x62, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x6e, 0x30, 0x44, 0x62, 0x48, 0x7a,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x6e, 0x30, 0x44, 0x62, 0x48, 0x7a, 0x12, 0x2f, 0x0a, 0x04,
	0x68, 0x64, 0x6f, 0x70, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x68, 0x64, 0x6f, 0x70, 0x12, 0x2f, 0x0a,
	0x04, 0x76, 0x64, 0x6f, 0x70, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x76, 0x64, 0x6f, 0x70, 0x42, 0x30,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x72, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x01, 0x5a, 0x15, 0x63, 0x72, 0x61, 0x78, 0x69,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_craxiom_messaging_gnss_record_proto_rawDescOnce sync.Once
	file_com_craxiom_messaging_gnss_record_proto_rawDescData = file_com_craxiom_messaging_gnss_record_proto_rawDesc
)

func file_com_craxiom_messaging_gnss_record_proto_rawDescGZIP() []byte {
	file_com_craxiom_messaging_gnss_record_proto_rawDescOnce.Do(func() {
		file_com_craxiom_messaging_gnss_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_craxiom_messaging_gnss_record_proto_rawDescData)
	})
	return file_com_craxiom_messaging_gnss_record_proto_rawDescData
}

var file_com_craxiom_messaging_gnss_record_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_craxiom_messaging_gnss_record_proto_goTypes = []any{
	(*GnssRecord)(nil),             // 0: com.craxiom.messaging.GnssRecord
	(*GnssRecordData)(nil),         // 1: com.craxiom.messaging.GnssRecordData
	(Constellation)(0),             // 2: com.craxiom.messaging.gnss.constellation.Constellation
	(*wrapperspb.UInt32Value)(nil), // 3: google.protobuf.UInt32Value
	(*wrapperspb.UInt64Value)(nil), // 4: google.protobuf.UInt64Value
	(*wrapperspb.DoubleValue)(nil), // 5: google.protobuf.DoubleValue
	(*wrapperspb.BoolValue)(nil),   // 6: google.protobuf.BoolValue
	(*wrapperspb.FloatValue)(nil),  // 7: google.protobuf.FloatValue
}
var file_com_craxiom_messaging_gnss_record_proto_depIdxs = []int32{
	1,  // 0: com.craxiom.messaging.GnssRecord.data:type_name -> com.craxiom.messaging.GnssRecordData
	2,  // 1: com.craxiom.messaging.GnssRecordData.constellation:type_name -> com.craxiom.messaging.gnss.constellation.Constellation
	3,  // 2: com.craxiom.messaging.GnssRecordData.spaceVehicleId:type_name -> google.protobuf.UInt32Value
	4,  // 3: com.craxiom.messaging.GnssRecordData.carrierFreqHz:type_name -> google.protobuf.UInt64Value
	5,  // 4: com.craxiom.messaging.GnssRecordData.clockOffset:type_name -> google.protobuf.DoubleValue
	6,  // 5: com.craxiom.messaging.GnssRecordData.usedInSolution:type_name -> google.protobuf.BoolValue
	7,  // 6: com.craxiom.messaging.GnssRecordData.undulationM:type_name -> google.protobuf.FloatValue
	7,  // 7: com.craxiom.messaging.GnssRecordData.latitudeStdDevM:type_name -> google.protobuf.FloatValue
	7,  // 8: com.craxiom.messaging.GnssRecordData.longitudeStdDevM:type_name -> google.protobuf.FloatValue
	7,  // 9: com.craxiom.messaging.GnssRecordData.altitudeStdDevM:type_name -> google.protobuf.FloatValue
	7,  // 10: com.craxiom.messaging.GnssRecordData.agcDb:type_name -> google.protobuf.FloatValue
	7,  // 11: com.craxiom.messaging.GnssRecordData.cn0DbHz:type_name -> google.protobuf.FloatValue
	7,  // 12: com.craxiom.messaging.GnssRecordData.hdop:type_name -> google.protobuf.FloatValue
	7,  // 13: com.craxiom.messaging.GnssRecordData.vdop:type_name -> google.protobuf.FloatValue
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_com_craxiom_messaging_gnss_record_proto_init() }
func file_com_craxiom_messaging_gnss_record_proto_init() {
	if File_com_craxiom_messaging_gnss_record_proto != nil {
		return
	}
	file_com_craxiom_messaging_gnss_constellation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_craxiom_messaging_gnss_record_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GnssRecord); i {
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
		file_com_craxiom_messaging_gnss_record_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GnssRecordData); i {
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
			RawDescriptor: file_com_craxiom_messaging_gnss_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_craxiom_messaging_gnss_record_proto_goTypes,
		DependencyIndexes: file_com_craxiom_messaging_gnss_record_proto_depIdxs,
		MessageInfos:      file_com_craxiom_messaging_gnss_record_proto_msgTypes,
	}.Build()
	File_com_craxiom_messaging_gnss_record_proto = out.File
	file_com_craxiom_messaging_gnss_record_proto_rawDesc = nil
	file_com_craxiom_messaging_gnss_record_proto_goTypes = nil
	file_com_craxiom_messaging_gnss_record_proto_depIdxs = nil
}
