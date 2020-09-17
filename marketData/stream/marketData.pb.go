// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: marketData/marketData.proto

package stream

import (
	context "context"
	constants "github.com/bhbosman/gocommon/constants"
	goprotoextra "github.com/bhbosman/goprotoextra"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	proto1 "google.golang.org/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price  float64 `protobuf:"fixed64,1,opt,name=Price,proto3" json:"Price,omitempty"`
	Volume float64 `protobuf:"fixed64,2,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketData_marketData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_marketData_marketData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_marketData_marketData_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Point) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (self *Point) TypeCode() uint32 {
	return PointTypeCode
}

type PublishTop5 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instrument         string   `protobuf:"bytes,1,opt,name=Instrument,proto3" json:"Instrument,omitempty"`
	Spread             float64  `protobuf:"fixed64,2,opt,name=Spread,proto3" json:"Spread,omitempty"`
	SourceTimeStamp    int64    `protobuf:"varint,3,opt,name=SourceTimeStamp,proto3" json:"SourceTimeStamp,omitempty"`
	SourceMessageCount int64    `protobuf:"varint,4,opt,name=SourceMessageCount,proto3" json:"SourceMessageCount,omitempty"`
	UpdateCount        int64    `protobuf:"varint,5,opt,name=UpdateCount,proto3" json:"UpdateCount,omitempty"`
	Bid                []*Point `protobuf:"bytes,6,rep,name=Bid,proto3" json:"Bid,omitempty"`
	Ask                []*Point `protobuf:"bytes,7,rep,name=Ask,proto3" json:"Ask,omitempty"`
}

func (x *PublishTop5) Reset() {
	*x = PublishTop5{}
	if protoimpl.UnsafeEnabled {
		mi := &file_marketData_marketData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishTop5) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishTop5) ProtoMessage() {}

func (x *PublishTop5) ProtoReflect() protoreflect.Message {
	mi := &file_marketData_marketData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishTop5.ProtoReflect.Descriptor instead.
func (*PublishTop5) Descriptor() ([]byte, []int) {
	return file_marketData_marketData_proto_rawDescGZIP(), []int{1}
}

func (x *PublishTop5) GetInstrument() string {
	if x != nil {
		return x.Instrument
	}
	return ""
}

func (x *PublishTop5) GetSpread() float64 {
	if x != nil {
		return x.Spread
	}
	return 0
}

func (x *PublishTop5) GetSourceTimeStamp() int64 {
	if x != nil {
		return x.SourceTimeStamp
	}
	return 0
}

func (x *PublishTop5) GetSourceMessageCount() int64 {
	if x != nil {
		return x.SourceMessageCount
	}
	return 0
}

func (x *PublishTop5) GetUpdateCount() int64 {
	if x != nil {
		return x.UpdateCount
	}
	return 0
}

func (x *PublishTop5) GetBid() []*Point {
	if x != nil {
		return x.Bid
	}
	return nil
}

func (x *PublishTop5) GetAsk() []*Point {
	if x != nil {
		return x.Ask
	}
	return nil
}

func (self *PublishTop5) TypeCode() uint32 {
	return PublishTop5TypeCode
}

var File_marketData_marketData_proto protoreflect.FileDescriptor

var file_marketData_marketData_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x05, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x22, 0xf5, 0x01, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x6f, 0x70,
	0x35, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x70, 0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x53, 0x70, 0x72, 0x65, 0x61, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x2e, 0x0a, 0x12, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x12, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x03, 0x42, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x03, 0x41, 0x73, 0x6b, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x03, 0x41, 0x73, 0x6b, 0x32, 0x3a, 0x0a, 0x09, 0x4d, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x0c, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x6f, 0x70, 0x35, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_marketData_marketData_proto_rawDescOnce sync.Once
	file_marketData_marketData_proto_rawDescData = file_marketData_marketData_proto_rawDesc
)

func file_marketData_marketData_proto_rawDescGZIP() []byte {
	file_marketData_marketData_proto_rawDescOnce.Do(func() {
		file_marketData_marketData_proto_rawDescData = protoimpl.X.CompressGZIP(file_marketData_marketData_proto_rawDescData)
	})
	return file_marketData_marketData_proto_rawDescData
}

var file_marketData_marketData_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_marketData_marketData_proto_goTypes = []interface{}{
	(*Point)(nil),       // 0: Point
	(*PublishTop5)(nil), // 1: PublishTop5
	(*empty.Empty)(nil), // 2: google.protobuf.Empty
}
var file_marketData_marketData_proto_depIdxs = []int32{
	0, // 0: PublishTop5.Bid:type_name -> Point
	0, // 1: PublishTop5.Ask:type_name -> Point
	1, // 2: MyService.Check:input_type -> PublishTop5
	2, // 3: MyService.Check:output_type -> google.protobuf.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_marketData_marketData_proto_init() }
func file_marketData_marketData_proto_init() {
	if File_marketData_marketData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_marketData_marketData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_marketData_marketData_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishTop5); i {
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
			RawDescriptor: file_marketData_marketData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_marketData_marketData_proto_goTypes,
		DependencyIndexes: file_marketData_marketData_proto_depIdxs,
		MessageInfos:      file_marketData_marketData_proto_msgTypes,
	}.Build()
	File_marketData_marketData_proto = out.File
	file_marketData_marketData_proto_rawDesc = nil
	file_marketData_marketData_proto_goTypes = nil
	file_marketData_marketData_proto_depIdxs = nil
}

const PointTypeCode uint32 = 2882114659
const PublishTop5TypeCode uint32 = 2016047094

//true
//true
//false
//false
type PointWrapper struct {
	goprotoextra.BaseMessageWrapper
	Data *Point
}

func (self *PointWrapper) Message() proto1.Message {
	return self.Data
}

func (self *PointWrapper) messageWrapper() goprotoextra.IMessageWrapper {
	return self
}

func NewPointWrapper(
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	toReactor goprotoextra.ToReactorFunc,
	toConnection goprotoextra.ToConnectionFunc,
	data *Point) *PointWrapper {
	return &PointWrapper{
		BaseMessageWrapper: goprotoextra.NewBaseMessageWrapper(
			cancelCtx,
			cancelFunc,
			toReactor,
			toConnection),
		Data: data,
	}
}

var _ = goprotoextra.Register(
	PointTypeCode,
	goprotoextra.TypeCodeData{
		CreateMessage: func() proto1.Message {
			return &Point{}
		},
		CreateWrapper: func(
			cancelCtx context.Context,
			cancelFunc context.CancelFunc,
			toReactor goprotoextra.ToReactorFunc,
			toConnection goprotoextra.ToConnectionFunc,
			data proto1.Message) (goprotoextra.IMessageWrapper, error) {
			if msg, ok := data.(*Point); ok {
				return NewPointWrapper(
					cancelCtx,
					cancelFunc,
					toReactor,
					toConnection,
					msg), nil
			}
			return nil, constants.InvalidParam
		}})

//true
//true
//false
//false
type PublishTop5Wrapper struct {
	goprotoextra.BaseMessageWrapper
	Data *PublishTop5
}

func (self *PublishTop5Wrapper) Message() proto1.Message {
	return self.Data
}

func (self *PublishTop5Wrapper) messageWrapper() goprotoextra.IMessageWrapper {
	return self
}

func NewPublishTop5Wrapper(
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	toReactor goprotoextra.ToReactorFunc,
	toConnection goprotoextra.ToConnectionFunc,
	data *PublishTop5) *PublishTop5Wrapper {
	return &PublishTop5Wrapper{
		BaseMessageWrapper: goprotoextra.NewBaseMessageWrapper(
			cancelCtx,
			cancelFunc,
			toReactor,
			toConnection),
		Data: data,
	}
}

var _ = goprotoextra.Register(
	PublishTop5TypeCode,
	goprotoextra.TypeCodeData{
		CreateMessage: func() proto1.Message {
			return &PublishTop5{}
		},
		CreateWrapper: func(
			cancelCtx context.Context,
			cancelFunc context.CancelFunc,
			toReactor goprotoextra.ToReactorFunc,
			toConnection goprotoextra.ToConnectionFunc,
			data proto1.Message) (goprotoextra.IMessageWrapper, error) {
			if msg, ok := data.(*PublishTop5); ok {
				return NewPublishTop5Wrapper(
					cancelCtx,
					cancelFunc,
					toReactor,
					toConnection,
					msg), nil
			}
			return nil, constants.InvalidParam
		}})
