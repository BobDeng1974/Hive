// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: models.proto

package grpc

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Alarm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Severity string `protobuf:"bytes,3,opt,name=Severity,proto3" json:"Severity,omitempty"`
	DeviceID int64  `protobuf:"varint,4,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
}

func (x *Alarm) Reset() {
	*x = Alarm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alarm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alarm) ProtoMessage() {}

func (x *Alarm) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alarm.ProtoReflect.Descriptor instead.
func (*Alarm) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{0}
}

func (x *Alarm) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Alarm) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Alarm) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *Alarm) GetDeviceID() int64 {
	if x != nil {
		return x.DeviceID
	}
	return 0
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string  `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	IMEI      string  `protobuf:"bytes,2,opt,name=IMEI,proto3" json:"IMEI,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	Latitude  float64 `protobuf:"fixed64,4,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{1}
}

func (x *Device) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Device) GetIMEI() string {
	if x != nil {
		return x.IMEI
	}
	return ""
}

func (x *Device) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Device) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type Measurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string  `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Value    float64 `protobuf:"fixed64,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Unit     string  `protobuf:"bytes,3,opt,name=Unit,proto3" json:"Unit,omitempty"`
	DeviceID int64   `protobuf:"varint,4,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
}

func (x *Measurement) Reset() {
	*x = Measurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Measurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measurement) ProtoMessage() {}

func (x *Measurement) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measurement.ProtoReflect.Descriptor instead.
func (*Measurement) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{2}
}

func (x *Measurement) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Measurement) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Measurement) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Measurement) GetDeviceID() int64 {
	if x != nil {
		return x.DeviceID
	}
	return 0
}

type Confirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply int64 `protobuf:"varint,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *Confirmation) Reset() {
	*x = Confirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmation) ProtoMessage() {}

func (x *Confirmation) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmation.ProtoReflect.Descriptor instead.
func (*Confirmation) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{3}
}

func (x *Confirmation) GetReply() int64 {
	if x != nil {
		return x.Reply
	}
	return 0
}

var File_models_proto protoreflect.FileDescriptor

var file_models_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b,
	0x0a, 0x05, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x22, 0x6a, 0x0a, 0x06, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x4d, 0x45,
	0x49, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x4d, 0x45, 0x49, 0x12, 0x1c, 0x0a,
	0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x4c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x67, 0x0a, 0x0b, 0x4d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44,
	0x22, 0x24, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x7f, 0x0a, 0x12, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x0c, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a,
	0x0d, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0d, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x00, 0x28, 0x01, 0x32, 0x66, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x07, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x0d, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x07, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0d, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x28, 0x01, 0x32,
	0x61, 0x0a, 0x0c, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x26, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x12, 0x06,
	0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x1a, 0x0d, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x73, 0x12, 0x06, 0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x1a,
	0x0d, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00,
	0x28, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_proto_rawDescOnce sync.Once
	file_models_proto_rawDescData = file_models_proto_rawDesc
)

func file_models_proto_rawDescGZIP() []byte {
	file_models_proto_rawDescOnce.Do(func() {
		file_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_proto_rawDescData)
	})
	return file_models_proto_rawDescData
}

var file_models_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_models_proto_goTypes = []interface{}{
	(*Alarm)(nil),        // 0: Alarm
	(*Device)(nil),       // 1: Device
	(*Measurement)(nil),  // 2: Measurement
	(*Confirmation)(nil), // 3: Confirmation
}
var file_models_proto_depIdxs = []int32{
	2, // 0: MeasurementService.CreateMeasurement:input_type -> Measurement
	2, // 1: MeasurementService.CreateMeasurements:input_type -> Measurement
	1, // 2: DeviceService.CreateDevice:input_type -> Device
	1, // 3: DeviceService.CreateDevices:input_type -> Device
	0, // 4: AlarmService.CreateAlarm:input_type -> Alarm
	0, // 5: AlarmService.CreateAlarms:input_type -> Alarm
	3, // 6: MeasurementService.CreateMeasurement:output_type -> Confirmation
	3, // 7: MeasurementService.CreateMeasurements:output_type -> Confirmation
	3, // 8: DeviceService.CreateDevice:output_type -> Confirmation
	3, // 9: DeviceService.CreateDevices:output_type -> Confirmation
	3, // 10: AlarmService.CreateAlarm:output_type -> Confirmation
	3, // 11: AlarmService.CreateAlarms:output_type -> Confirmation
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_proto_init() }
func file_models_proto_init() {
	if File_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alarm); i {
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
		file_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Measurement); i {
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
		file_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirmation); i {
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
			RawDescriptor: file_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_models_proto_goTypes,
		DependencyIndexes: file_models_proto_depIdxs,
		MessageInfos:      file_models_proto_msgTypes,
	}.Build()
	File_models_proto = out.File
	file_models_proto_rawDesc = nil
	file_models_proto_goTypes = nil
	file_models_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MeasurementServiceClient is the client API for MeasurementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeasurementServiceClient interface {
	CreateMeasurement(ctx context.Context, in *Measurement, opts ...grpc.CallOption) (*Confirmation, error)
	CreateMeasurements(ctx context.Context, opts ...grpc.CallOption) (MeasurementService_CreateMeasurementsClient, error)
}

type measurementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeasurementServiceClient(cc grpc.ClientConnInterface) MeasurementServiceClient {
	return &measurementServiceClient{cc}
}

func (c *measurementServiceClient) CreateMeasurement(ctx context.Context, in *Measurement, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, "/MeasurementService/CreateMeasurement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measurementServiceClient) CreateMeasurements(ctx context.Context, opts ...grpc.CallOption) (MeasurementService_CreateMeasurementsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MeasurementService_serviceDesc.Streams[0], "/MeasurementService/CreateMeasurements", opts...)
	if err != nil {
		return nil, err
	}
	x := &measurementServiceCreateMeasurementsClient{stream}
	return x, nil
}

type MeasurementService_CreateMeasurementsClient interface {
	Send(*Measurement) error
	CloseAndRecv() (*Confirmation, error)
	grpc.ClientStream
}

type measurementServiceCreateMeasurementsClient struct {
	grpc.ClientStream
}

func (x *measurementServiceCreateMeasurementsClient) Send(m *Measurement) error {
	return x.ClientStream.SendMsg(m)
}

func (x *measurementServiceCreateMeasurementsClient) CloseAndRecv() (*Confirmation, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Confirmation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MeasurementServiceServer is the server API for MeasurementService service.
type MeasurementServiceServer interface {
	CreateMeasurement(context.Context, *Measurement) (*Confirmation, error)
	CreateMeasurements(MeasurementService_CreateMeasurementsServer) error
}

// UnimplementedMeasurementServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMeasurementServiceServer struct {
}

func (*UnimplementedMeasurementServiceServer) CreateMeasurement(context.Context, *Measurement) (*Confirmation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMeasurement not implemented")
}
func (*UnimplementedMeasurementServiceServer) CreateMeasurements(MeasurementService_CreateMeasurementsServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateMeasurements not implemented")
}

func RegisterMeasurementServiceServer(s *grpc.Server, srv MeasurementServiceServer) {
	s.RegisterService(&_MeasurementService_serviceDesc, srv)
}

func _MeasurementService_CreateMeasurement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Measurement)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasurementServiceServer).CreateMeasurement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MeasurementService/CreateMeasurement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasurementServiceServer).CreateMeasurement(ctx, req.(*Measurement))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasurementService_CreateMeasurements_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MeasurementServiceServer).CreateMeasurements(&measurementServiceCreateMeasurementsServer{stream})
}

type MeasurementService_CreateMeasurementsServer interface {
	SendAndClose(*Confirmation) error
	Recv() (*Measurement, error)
	grpc.ServerStream
}

type measurementServiceCreateMeasurementsServer struct {
	grpc.ServerStream
}

func (x *measurementServiceCreateMeasurementsServer) SendAndClose(m *Confirmation) error {
	return x.ServerStream.SendMsg(m)
}

func (x *measurementServiceCreateMeasurementsServer) Recv() (*Measurement, error) {
	m := new(Measurement)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MeasurementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MeasurementService",
	HandlerType: (*MeasurementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMeasurement",
			Handler:    _MeasurementService_CreateMeasurement_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateMeasurements",
			Handler:       _MeasurementService_CreateMeasurements_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "models.proto",
}

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceServiceClient interface {
	CreateDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Confirmation, error)
	CreateDevices(ctx context.Context, opts ...grpc.CallOption) (DeviceService_CreateDevicesClient, error)
}

type deviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceServiceClient(cc grpc.ClientConnInterface) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) CreateDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, "/DeviceService/CreateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) CreateDevices(ctx context.Context, opts ...grpc.CallOption) (DeviceService_CreateDevicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DeviceService_serviceDesc.Streams[0], "/DeviceService/CreateDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &deviceServiceCreateDevicesClient{stream}
	return x, nil
}

type DeviceService_CreateDevicesClient interface {
	Send(*Device) error
	CloseAndRecv() (*Confirmation, error)
	grpc.ClientStream
}

type deviceServiceCreateDevicesClient struct {
	grpc.ClientStream
}

func (x *deviceServiceCreateDevicesClient) Send(m *Device) error {
	return x.ClientStream.SendMsg(m)
}

func (x *deviceServiceCreateDevicesClient) CloseAndRecv() (*Confirmation, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Confirmation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DeviceServiceServer is the server API for DeviceService service.
type DeviceServiceServer interface {
	CreateDevice(context.Context, *Device) (*Confirmation, error)
	CreateDevices(DeviceService_CreateDevicesServer) error
}

// UnimplementedDeviceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceServiceServer struct {
}

func (*UnimplementedDeviceServiceServer) CreateDevice(context.Context, *Device) (*Confirmation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (*UnimplementedDeviceServiceServer) CreateDevices(DeviceService_CreateDevicesServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateDevices not implemented")
}

func RegisterDeviceServiceServer(s *grpc.Server, srv DeviceServiceServer) {
	s.RegisterService(&_DeviceService_serviceDesc, srv)
}

func _DeviceService_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeviceService/CreateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).CreateDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_CreateDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DeviceServiceServer).CreateDevices(&deviceServiceCreateDevicesServer{stream})
}

type DeviceService_CreateDevicesServer interface {
	SendAndClose(*Confirmation) error
	Recv() (*Device, error)
	grpc.ServerStream
}

type deviceServiceCreateDevicesServer struct {
	grpc.ServerStream
}

func (x *deviceServiceCreateDevicesServer) SendAndClose(m *Confirmation) error {
	return x.ServerStream.SendMsg(m)
}

func (x *deviceServiceCreateDevicesServer) Recv() (*Device, error) {
	m := new(Device)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DeviceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDevice",
			Handler:    _DeviceService_CreateDevice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateDevices",
			Handler:       _DeviceService_CreateDevices_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "models.proto",
}

// AlarmServiceClient is the client API for AlarmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlarmServiceClient interface {
	CreateAlarm(ctx context.Context, in *Alarm, opts ...grpc.CallOption) (*Confirmation, error)
	CreateAlarms(ctx context.Context, opts ...grpc.CallOption) (AlarmService_CreateAlarmsClient, error)
}

type alarmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlarmServiceClient(cc grpc.ClientConnInterface) AlarmServiceClient {
	return &alarmServiceClient{cc}
}

func (c *alarmServiceClient) CreateAlarm(ctx context.Context, in *Alarm, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, "/AlarmService/CreateAlarm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmServiceClient) CreateAlarms(ctx context.Context, opts ...grpc.CallOption) (AlarmService_CreateAlarmsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AlarmService_serviceDesc.Streams[0], "/AlarmService/CreateAlarms", opts...)
	if err != nil {
		return nil, err
	}
	x := &alarmServiceCreateAlarmsClient{stream}
	return x, nil
}

type AlarmService_CreateAlarmsClient interface {
	Send(*Alarm) error
	CloseAndRecv() (*Confirmation, error)
	grpc.ClientStream
}

type alarmServiceCreateAlarmsClient struct {
	grpc.ClientStream
}

func (x *alarmServiceCreateAlarmsClient) Send(m *Alarm) error {
	return x.ClientStream.SendMsg(m)
}

func (x *alarmServiceCreateAlarmsClient) CloseAndRecv() (*Confirmation, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Confirmation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AlarmServiceServer is the server API for AlarmService service.
type AlarmServiceServer interface {
	CreateAlarm(context.Context, *Alarm) (*Confirmation, error)
	CreateAlarms(AlarmService_CreateAlarmsServer) error
}

// UnimplementedAlarmServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAlarmServiceServer struct {
}

func (*UnimplementedAlarmServiceServer) CreateAlarm(context.Context, *Alarm) (*Confirmation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlarm not implemented")
}
func (*UnimplementedAlarmServiceServer) CreateAlarms(AlarmService_CreateAlarmsServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateAlarms not implemented")
}

func RegisterAlarmServiceServer(s *grpc.Server, srv AlarmServiceServer) {
	s.RegisterService(&_AlarmService_serviceDesc, srv)
}

func _AlarmService_CreateAlarm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Alarm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmServiceServer).CreateAlarm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AlarmService/CreateAlarm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmServiceServer).CreateAlarm(ctx, req.(*Alarm))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmService_CreateAlarms_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AlarmServiceServer).CreateAlarms(&alarmServiceCreateAlarmsServer{stream})
}

type AlarmService_CreateAlarmsServer interface {
	SendAndClose(*Confirmation) error
	Recv() (*Alarm, error)
	grpc.ServerStream
}

type alarmServiceCreateAlarmsServer struct {
	grpc.ServerStream
}

func (x *alarmServiceCreateAlarmsServer) SendAndClose(m *Confirmation) error {
	return x.ServerStream.SendMsg(m)
}

func (x *alarmServiceCreateAlarmsServer) Recv() (*Alarm, error) {
	m := new(Alarm)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AlarmService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AlarmService",
	HandlerType: (*AlarmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAlarm",
			Handler:    _AlarmService_CreateAlarm_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateAlarms",
			Handler:       _AlarmService_CreateAlarms_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "models.proto",
}