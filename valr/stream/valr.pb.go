// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: valr/valr.proto

package stream

import (
	context "context"
	stream "github.com/bhbosman/gocommon/stream"
	goerrors "github.com/bhbosman/goerrors"
	goprotoextra "github.com/bhbosman/goprotoextra"
	proto "github.com/golang/protobuf/proto"
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
		mi := &file_valr_valr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_valr_valr_proto_msgTypes[0]
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
	return file_valr_valr_proto_rawDescGZIP(), []int{0}
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

var File_valr_valr_proto protoreflect.FileDescriptor

var file_valr_valr_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x35, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_valr_valr_proto_rawDescOnce sync.Once
	file_valr_valr_proto_rawDescData = file_valr_valr_proto_rawDesc
)

func file_valr_valr_proto_rawDescGZIP() []byte {
	file_valr_valr_proto_rawDescOnce.Do(func() {
		file_valr_valr_proto_rawDescData = protoimpl.X.CompressGZIP(file_valr_valr_proto_rawDescData)
	})
	return file_valr_valr_proto_rawDescData
}

var file_valr_valr_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_valr_valr_proto_goTypes = []interface{}{
	(*Point)(nil), // 0: Point
}
var file_valr_valr_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_valr_valr_proto_init() }
func file_valr_valr_proto_init() {
	if File_valr_valr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_valr_valr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_valr_valr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_valr_valr_proto_goTypes,
		DependencyIndexes: file_valr_valr_proto_depIdxs,
		MessageInfos:      file_valr_valr_proto_msgTypes,
	}.Build()
	File_valr_valr_proto = out.File
	file_valr_valr_proto_rawDesc = nil
	file_valr_valr_proto_goTypes = nil
	file_valr_valr_proto_depIdxs = nil
}

const PointTypeCode uint32 = 2882114659

//true
//true
//false
//false
type PointWrapper struct {
	goprotoextra.BaseMessageWrapper
	Data *Point
}

func (self *PointWrapper) Message() interface{} {
	return self.Data
}

func (self *PointWrapper) messageWrapper() interface{} {
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

var _ = stream.Register(
	PointTypeCode,
	stream.TypeCodeData{
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
			return nil, goerrors.InvalidParam
		}})