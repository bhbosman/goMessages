// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.10
// 	protoc        v3.14.0
// source: kraken/kraken.proto

package stream

import (
	stream "github.com/bhbosman/gocommon/stream"
	goerrors "github.com/bhbosman/goerrors"
	goprotoextra "github.com/bhbosman/goprotoextra"
	v2 "github.com/reactivex/rxgo/v2"
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Subscribe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reqid uint32 `protobuf:"varint,1,opt,name=reqid,proto3" json:"reqid,omitempty"`
	Pair  string `protobuf:"bytes,2,opt,name=pair,proto3" json:"pair,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Subscribe) Reset() {
	*x = Subscribe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kraken_kraken_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscribe) ProtoMessage() {}

func (x *Subscribe) ProtoReflect() protoreflect.Message {
	mi := &file_kraken_kraken_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscribe.ProtoReflect.Descriptor instead.
func (*Subscribe) Descriptor() ([]byte, []int) {
	return file_kraken_kraken_proto_rawDescGZIP(), []int{0}
}

func (x *Subscribe) GetReqid() uint32 {
	if x != nil {
		return x.Reqid
	}
	return 0
}

func (x *Subscribe) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *Subscribe) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (self *Subscribe) TypeCode() uint32 {
	return SubscribeTypeCode
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Open   float64 `protobuf:"fixed64,1,opt,name=Open,proto3" json:"Open,omitempty"`
	Close  float64 `protobuf:"fixed64,2,opt,name=Close,proto3" json:"Close,omitempty"`
	Volume float64 `protobuf:"fixed64,3,opt,name=Volume,proto3" json:"Volume,omitempty"`
	High   float64 `protobuf:"fixed64,4,opt,name=High,proto3" json:"High,omitempty"`
	Low    float64 `protobuf:"fixed64,5,opt,name=Low,proto3" json:"Low,omitempty"`
	Bid    float64 `protobuf:"fixed64,6,opt,name=Bid,proto3" json:"Bid,omitempty"`
	Ask    float64 `protobuf:"fixed64,7,opt,name=Ask,proto3" json:"Ask,omitempty"`
	Pair   string  `protobuf:"bytes,8,opt,name=Pair,proto3" json:"Pair,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kraken_kraken_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_kraken_kraken_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_kraken_kraken_proto_rawDescGZIP(), []int{1}
}

func (x *Price) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *Price) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *Price) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Price) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Price) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Price) GetBid() float64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *Price) GetAsk() float64 {
	if x != nil {
		return x.Ask
	}
	return 0
}

func (x *Price) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (self *Price) TypeCode() uint32 {
	return PriceTypeCode
}

var File_kraken_kraken_proto protoreflect.FileDescriptor

var file_kraken_kraken_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6b, 0x72, 0x61, 0x6b, 0x65, 0x6e, 0x2f, 0x6b, 0x72, 0x61, 0x6b, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49,
	0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x71, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x65, 0x71, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x69, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x69, 0x67, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x04, 0x48, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f, 0x77,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x4c, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x42,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x42, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x41, 0x73, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x41, 0x73, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50,
	0x61, 0x69, 0x72, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kraken_kraken_proto_rawDescOnce sync.Once
	file_kraken_kraken_proto_rawDescData = file_kraken_kraken_proto_rawDesc
)

func file_kraken_kraken_proto_rawDescGZIP() []byte {
	file_kraken_kraken_proto_rawDescOnce.Do(func() {
		file_kraken_kraken_proto_rawDescData = protoimpl.X.CompressGZIP(file_kraken_kraken_proto_rawDescData)
	})
	return file_kraken_kraken_proto_rawDescData
}

var file_kraken_kraken_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kraken_kraken_proto_goTypes = []interface{}{
	(*Subscribe)(nil), // 0: golang.example.policy.Subscribe
	(*Price)(nil),     // 1: golang.example.policy.Price
}
var file_kraken_kraken_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kraken_kraken_proto_init() }
func file_kraken_kraken_proto_init() {
	if File_kraken_kraken_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kraken_kraken_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscribe); i {
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
		file_kraken_kraken_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
			RawDescriptor: file_kraken_kraken_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kraken_kraken_proto_goTypes,
		DependencyIndexes: file_kraken_kraken_proto_depIdxs,
		MessageInfos:      file_kraken_kraken_proto_msgTypes,
	}.Build()
	File_kraken_kraken_proto = out.File
	file_kraken_kraken_proto_rawDesc = nil
	file_kraken_kraken_proto_goTypes = nil
	file_kraken_kraken_proto_depIdxs = nil
}

// Typecode generated from: "Subscribe"
const SubscribeTypeCode uint32 = 3930643869

// Typecode generated from: "Price"
const PriceTypeCode uint32 = 185142749

//true
//true
//false
//false
type SubscribeWrapper struct {
	goprotoextra.BaseMessageWrapper
	Data *Subscribe
}

func (self *SubscribeWrapper) Message() interface{} {
	return self.Data
}

func (self *SubscribeWrapper) messageWrapper() interface{} {
	return self
}

func NewSubscribeWrapper(
	toReactor v2.NextFunc,
	toConnection v2.NextFunc,
	data *Subscribe) *SubscribeWrapper {
	return &SubscribeWrapper{
		BaseMessageWrapper: goprotoextra.NewBaseMessageWrapper(
			toReactor,
			toConnection),
		Data: data,
	}
}

var _ = stream.Register(
	SubscribeTypeCode,
	stream.TypeCodeData{
		CreateMessage: func() proto.Message {
			return &Subscribe{}
		},
		CreateWrapper: func(
			toReactor v2.NextFunc,
			toConnection v2.NextFunc,
			data proto.Message) (goprotoextra.IMessageWrapper, error) {
			if msg, ok := data.(*Subscribe); ok {
				return NewSubscribeWrapper(
					toReactor,
					toConnection,
					msg), nil
			}
			return nil, goerrors.InvalidParam
		}})

//true
//true
//false
//false
type PriceWrapper struct {
	goprotoextra.BaseMessageWrapper
	Data *Price
}

func (self *PriceWrapper) Message() interface{} {
	return self.Data
}

func (self *PriceWrapper) messageWrapper() interface{} {
	return self
}

func NewPriceWrapper(
	toReactor v2.NextFunc,
	toConnection v2.NextFunc,
	data *Price) *PriceWrapper {
	return &PriceWrapper{
		BaseMessageWrapper: goprotoextra.NewBaseMessageWrapper(
			toReactor,
			toConnection),
		Data: data,
	}
}

var _ = stream.Register(
	PriceTypeCode,
	stream.TypeCodeData{
		CreateMessage: func() proto.Message {
			return &Price{}
		},
		CreateWrapper: func(
			toReactor v2.NextFunc,
			toConnection v2.NextFunc,
			data proto.Message) (goprotoextra.IMessageWrapper, error) {
			if msg, ok := data.(*Price); ok {
				return NewPriceWrapper(
					toReactor,
					toConnection,
					msg), nil
			}
			return nil, goerrors.InvalidParam
		}})
