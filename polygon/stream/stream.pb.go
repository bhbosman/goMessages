// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: polygon/stream.proto

package stream

import (
	context "context"
	constants "github.com/bhbosman/gocommon/constants"
	stream "github.com/bhbosman/gocommon/stream"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// https://polygon.io/sockets for forex
type PolygonMessageReceived struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType string `protobuf:"bytes,1,opt,name=EventType,json=ev,proto3" json:"EventType,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=Status,json=status,proto3" json:"Status,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=Message,json=message,proto3" json:"Message,omitempty"`
}

func (x *PolygonMessageReceived) Reset() {
	*x = PolygonMessageReceived{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polygon_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolygonMessageReceived) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolygonMessageReceived) ProtoMessage() {}

func (x *PolygonMessageReceived) ProtoReflect() protoreflect.Message {
	mi := &file_polygon_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolygonMessageReceived.ProtoReflect.Descriptor instead.
func (*PolygonMessageReceived) Descriptor() ([]byte, []int) {
	return file_polygon_stream_proto_rawDescGZIP(), []int{0}
}

func (x *PolygonMessageReceived) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *PolygonMessageReceived) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PolygonMessageReceived) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (self *PolygonMessageReceived) TypeCode() uint32 {
	return PolygonMessageReceivedTypeCode
}

type PolygonMessageSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action string `protobuf:"bytes,1,opt,name=Action,json=action,proto3" json:"Action,omitempty"`
	Params string `protobuf:"bytes,2,opt,name=Params,json=params,proto3" json:"Params,omitempty"`
}

func (x *PolygonMessageSend) Reset() {
	*x = PolygonMessageSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polygon_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolygonMessageSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolygonMessageSend) ProtoMessage() {}

func (x *PolygonMessageSend) ProtoReflect() protoreflect.Message {
	mi := &file_polygon_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolygonMessageSend.ProtoReflect.Descriptor instead.
func (*PolygonMessageSend) Descriptor() ([]byte, []int) {
	return file_polygon_stream_proto_rawDescGZIP(), []int{1}
}

func (x *PolygonMessageSend) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *PolygonMessageSend) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (self *PolygonMessageSend) TypeCode() uint32 {
	return PolygonMessageSendTypeCode
}

// https://polygon.io/sockets for forex
type PolygonForexQuote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType      string  `protobuf:"bytes,1,opt,name=EventType,json=ev,proto3" json:"EventType,omitempty"`
	CurrencyPair   string  `protobuf:"bytes,2,opt,name=CurrencyPair,json=p,proto3" json:"CurrencyPair,omitempty"`
	FXExchangeId   string  `protobuf:"bytes,3,opt,name=FXExchangeId,json=x,proto3" json:"FXExchangeId,omitempty"`
	AskPrice       float64 `protobuf:"fixed64,4,opt,name=AskPrice,json=a,proto3" json:"AskPrice,omitempty"`
	BidPrice       float64 `protobuf:"fixed64,5,opt,name=BidPrice,json=b,proto3" json:"BidPrice,omitempty"`
	QuoteTimestamp int64   `protobuf:"varint,6,opt,name=QuoteTimestamp,json=t,proto3" json:"QuoteTimestamp,omitempty"`
}

func (x *PolygonForexQuote) Reset() {
	*x = PolygonForexQuote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polygon_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolygonForexQuote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolygonForexQuote) ProtoMessage() {}

func (x *PolygonForexQuote) ProtoReflect() protoreflect.Message {
	mi := &file_polygon_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolygonForexQuote.ProtoReflect.Descriptor instead.
func (*PolygonForexQuote) Descriptor() ([]byte, []int) {
	return file_polygon_stream_proto_rawDescGZIP(), []int{2}
}

func (x *PolygonForexQuote) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *PolygonForexQuote) GetCurrencyPair() string {
	if x != nil {
		return x.CurrencyPair
	}
	return ""
}

func (x *PolygonForexQuote) GetFXExchangeId() string {
	if x != nil {
		return x.FXExchangeId
	}
	return ""
}

func (x *PolygonForexQuote) GetAskPrice() float64 {
	if x != nil {
		return x.AskPrice
	}
	return 0
}

func (x *PolygonForexQuote) GetBidPrice() float64 {
	if x != nil {
		return x.BidPrice
	}
	return 0
}

func (x *PolygonForexQuote) GetQuoteTimestamp() int64 {
	if x != nil {
		return x.QuoteTimestamp
	}
	return 0
}

func (self *PolygonForexQuote) TypeCode() uint32 {
	return PolygonForexQuoteTypeCode
}

type PolygonForexAggregate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType          string  `protobuf:"bytes,1,opt,name=EventType,json=ev,proto3" json:"EventType,omitempty"`
	CurrencyPair       string  `protobuf:"bytes,2,opt,name=CurrencyPair,json=pair,proto3" json:"CurrencyPair,omitempty"`
	OpenPrice          float64 `protobuf:"fixed64,3,opt,name=OpenPrice,json=o,proto3" json:"OpenPrice,omitempty"`
	ClosePrice         float64 `protobuf:"fixed64,4,opt,name=ClosePrice,json=c,proto3" json:"ClosePrice,omitempty"`
	HighPrice          float64 `protobuf:"fixed64,5,opt,name=HighPrice,json=h,proto3" json:"HighPrice,omitempty"`
	LowPrice           float64 `protobuf:"fixed64,6,opt,name=LowPrice,json=l,proto3" json:"LowPrice,omitempty"`
	Volume             int64   `protobuf:"varint,7,opt,name=Volume,json=v,proto3" json:"Volume,omitempty"`
	TickStartTimestamp int64   `protobuf:"varint,8,opt,name=TickStartTimestamp,json=s,proto3" json:"TickStartTimestamp,omitempty"`
}

func (x *PolygonForexAggregate) Reset() {
	*x = PolygonForexAggregate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polygon_stream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolygonForexAggregate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolygonForexAggregate) ProtoMessage() {}

func (x *PolygonForexAggregate) ProtoReflect() protoreflect.Message {
	mi := &file_polygon_stream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolygonForexAggregate.ProtoReflect.Descriptor instead.
func (*PolygonForexAggregate) Descriptor() ([]byte, []int) {
	return file_polygon_stream_proto_rawDescGZIP(), []int{3}
}

func (x *PolygonForexAggregate) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *PolygonForexAggregate) GetCurrencyPair() string {
	if x != nil {
		return x.CurrencyPair
	}
	return ""
}

func (x *PolygonForexAggregate) GetOpenPrice() float64 {
	if x != nil {
		return x.OpenPrice
	}
	return 0
}

func (x *PolygonForexAggregate) GetClosePrice() float64 {
	if x != nil {
		return x.ClosePrice
	}
	return 0
}

func (x *PolygonForexAggregate) GetHighPrice() float64 {
	if x != nil {
		return x.HighPrice
	}
	return 0
}

func (x *PolygonForexAggregate) GetLowPrice() float64 {
	if x != nil {
		return x.LowPrice
	}
	return 0
}

func (x *PolygonForexAggregate) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *PolygonForexAggregate) GetTickStartTimestamp() int64 {
	if x != nil {
		return x.TickStartTimestamp
	}
	return 0
}

func (self *PolygonForexAggregate) TypeCode() uint32 {
	return PolygonForexAggregateTypeCode
}

type PolygonForexCombined struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType string `protobuf:"bytes,1,opt,name=EventType,json=ev,proto3" json:"EventType,omitempty"`
	// Status Message
	StatusStatus  string `protobuf:"bytes,2,opt,name=StatusStatus,json=status,proto3" json:"StatusStatus,omitempty"`
	StatusMessage string `protobuf:"bytes,3,opt,name=StatusMessage,json=message,proto3" json:"StatusMessage,omitempty"`
	// Quote Message
	QuoteCurrencyPair   string  `protobuf:"bytes,4,opt,name=QuoteCurrencyPair,json=p,proto3" json:"QuoteCurrencyPair,omitempty"`
	QuoteFXExchangeId   string  `protobuf:"bytes,5,opt,name=QuoteFXExchangeId,json=x,proto3" json:"QuoteFXExchangeId,omitempty"`
	QuoteAskPrice       float64 `protobuf:"fixed64,6,opt,name=QuoteAskPrice,json=a,proto3" json:"QuoteAskPrice,omitempty"`
	QuoteBidPrice       float64 `protobuf:"fixed64,7,opt,name=QuoteBidPrice,json=b,proto3" json:"QuoteBidPrice,omitempty"`
	QuoteQuoteTimestamp int64   `protobuf:"varint,8,opt,name=QuoteQuoteTimestamp,json=t,proto3" json:"QuoteQuoteTimestamp,omitempty"`
	// Aggregate Message
	AggregateCurrencyPair       string  `protobuf:"bytes,9,opt,name=AggregateCurrencyPair,json=pair,proto3" json:"AggregateCurrencyPair,omitempty"`
	AggregateOpenPrice          float64 `protobuf:"fixed64,10,opt,name=AggregateOpenPrice,json=o,proto3" json:"AggregateOpenPrice,omitempty"`
	AggregateClosePrice         float64 `protobuf:"fixed64,11,opt,name=AggregateClosePrice,json=c,proto3" json:"AggregateClosePrice,omitempty"`
	AggregateHighPrice          float64 `protobuf:"fixed64,12,opt,name=AggregateHighPrice,json=h,proto3" json:"AggregateHighPrice,omitempty"`
	AggregateLowPrice           float64 `protobuf:"fixed64,13,opt,name=AggregateLowPrice,json=l,proto3" json:"AggregateLowPrice,omitempty"`
	AggregateVolume             int64   `protobuf:"varint,14,opt,name=AggregateVolume,json=v,proto3" json:"AggregateVolume,omitempty"`
	AggregateTickStartTimestamp int64   `protobuf:"varint,15,opt,name=AggregateTickStartTimestamp,json=s,proto3" json:"AggregateTickStartTimestamp,omitempty"`
}

func (x *PolygonForexCombined) Reset() {
	*x = PolygonForexCombined{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polygon_stream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolygonForexCombined) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolygonForexCombined) ProtoMessage() {}

func (x *PolygonForexCombined) ProtoReflect() protoreflect.Message {
	mi := &file_polygon_stream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolygonForexCombined.ProtoReflect.Descriptor instead.
func (*PolygonForexCombined) Descriptor() ([]byte, []int) {
	return file_polygon_stream_proto_rawDescGZIP(), []int{4}
}

func (x *PolygonForexCombined) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *PolygonForexCombined) GetStatusStatus() string {
	if x != nil {
		return x.StatusStatus
	}
	return ""
}

func (x *PolygonForexCombined) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *PolygonForexCombined) GetQuoteCurrencyPair() string {
	if x != nil {
		return x.QuoteCurrencyPair
	}
	return ""
}

func (x *PolygonForexCombined) GetQuoteFXExchangeId() string {
	if x != nil {
		return x.QuoteFXExchangeId
	}
	return ""
}

func (x *PolygonForexCombined) GetQuoteAskPrice() float64 {
	if x != nil {
		return x.QuoteAskPrice
	}
	return 0
}

func (x *PolygonForexCombined) GetQuoteBidPrice() float64 {
	if x != nil {
		return x.QuoteBidPrice
	}
	return 0
}

func (x *PolygonForexCombined) GetQuoteQuoteTimestamp() int64 {
	if x != nil {
		return x.QuoteQuoteTimestamp
	}
	return 0
}

func (x *PolygonForexCombined) GetAggregateCurrencyPair() string {
	if x != nil {
		return x.AggregateCurrencyPair
	}
	return ""
}

func (x *PolygonForexCombined) GetAggregateOpenPrice() float64 {
	if x != nil {
		return x.AggregateOpenPrice
	}
	return 0
}

func (x *PolygonForexCombined) GetAggregateClosePrice() float64 {
	if x != nil {
		return x.AggregateClosePrice
	}
	return 0
}

func (x *PolygonForexCombined) GetAggregateHighPrice() float64 {
	if x != nil {
		return x.AggregateHighPrice
	}
	return 0
}

func (x *PolygonForexCombined) GetAggregateLowPrice() float64 {
	if x != nil {
		return x.AggregateLowPrice
	}
	return 0
}

func (x *PolygonForexCombined) GetAggregateVolume() int64 {
	if x != nil {
		return x.AggregateVolume
	}
	return 0
}

func (x *PolygonForexCombined) GetAggregateTickStartTimestamp() int64 {
	if x != nil {
		return x.AggregateTickStartTimestamp
	}
	return 0
}

func (self *PolygonForexCombined) TypeCode() uint32 {
	return PolygonForexCombinedTypeCode
}

var File_polygon_stream_proto protoreflect.FileDescriptor

var file_polygon_stream_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x61, 0x0a, 0x16, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x15, 0x0a, 0x09, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x76,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x44, 0x0a, 0x12, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xa1, 0x01, 0x0a, 0x11, 0x50, 0x6f, 0x6c,
	0x79, 0x67, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x65, 0x78, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x15,
	0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x65, 0x76, 0x12, 0x17, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x50, 0x61, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x70, 0x12, 0x17,
	0x0a, 0x0c, 0x46, 0x58, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x78, 0x12, 0x13, 0x0a, 0x08, 0x41, 0x73, 0x6b, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x61, 0x12, 0x13, 0x0a, 0x08,
	0x42, 0x69, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01,
	0x62, 0x12, 0x19, 0x0a, 0x0e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x74, 0x22, 0xd4, 0x01, 0x0a,
	0x15, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x65, 0x78, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x76, 0x12, 0x1a, 0x0a,
	0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x09, 0x4f, 0x70, 0x65,
	0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x6f, 0x12,
	0x15, 0x0a, 0x0a, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x01, 0x63, 0x12, 0x14, 0x0a, 0x09, 0x48, 0x69, 0x67, 0x68, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x68, 0x12, 0x13, 0x0a, 0x08,
	0x4c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01,
	0x6c, 0x12, 0x11, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x01, 0x76, 0x12, 0x1d, 0x0a, 0x12, 0x54, 0x69, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x73, 0x22, 0xe0, 0x03, 0x0a, 0x14, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x46,
	0x6f, 0x72, 0x65, 0x78, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x15, 0x0a, 0x09,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x65, 0x76, 0x12, 0x1c, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1e, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x11, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x70, 0x12,
	0x1c, 0x0a, 0x11, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x46, 0x58, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x78, 0x12, 0x18, 0x0a,
	0x0d, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x61, 0x12, 0x18, 0x0a, 0x0d, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x42, 0x69, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01,
	0x62, 0x12, 0x1e, 0x0a, 0x13, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x74, 0x12, 0x23, 0x0a, 0x15, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x69, 0x72, 0x12, 0x1d, 0x0a, 0x12, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x01, 0x6f, 0x12, 0x1e, 0x0a, 0x13, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x01, 0x63, 0x12, 0x1d, 0x0a, 0x12, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x48, 0x69, 0x67, 0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x01, 0x68, 0x12, 0x1c, 0x0a, 0x11, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x01, 0x6c, 0x12, 0x1a, 0x0a, 0x0f, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x76, 0x12, 0x26,
	0x0a, 0x1b, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x01, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_polygon_stream_proto_rawDescOnce sync.Once
	file_polygon_stream_proto_rawDescData = file_polygon_stream_proto_rawDesc
)

func file_polygon_stream_proto_rawDescGZIP() []byte {
	file_polygon_stream_proto_rawDescOnce.Do(func() {
		file_polygon_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_polygon_stream_proto_rawDescData)
	})
	return file_polygon_stream_proto_rawDescData
}

var file_polygon_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_polygon_stream_proto_goTypes = []interface{}{
	(*PolygonMessageReceived)(nil), // 0: golang.example.policy.PolygonMessageReceived
	(*PolygonMessageSend)(nil),     // 1: golang.example.policy.PolygonMessageSend
	(*PolygonForexQuote)(nil),      // 2: golang.example.policy.PolygonForexQuote
	(*PolygonForexAggregate)(nil),  // 3: golang.example.policy.PolygonForexAggregate
	(*PolygonForexCombined)(nil),   // 4: golang.example.policy.PolygonForexCombined
}
var file_polygon_stream_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_polygon_stream_proto_init() }
func file_polygon_stream_proto_init() {
	if File_polygon_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_polygon_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolygonMessageReceived); i {
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
		file_polygon_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolygonMessageSend); i {
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
		file_polygon_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolygonForexQuote); i {
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
		file_polygon_stream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolygonForexAggregate); i {
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
		file_polygon_stream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolygonForexCombined); i {
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
			RawDescriptor: file_polygon_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_polygon_stream_proto_goTypes,
		DependencyIndexes: file_polygon_stream_proto_depIdxs,
		MessageInfos:      file_polygon_stream_proto_msgTypes,
	}.Build()
	File_polygon_stream_proto = out.File
	file_polygon_stream_proto_rawDesc = nil
	file_polygon_stream_proto_goTypes = nil
	file_polygon_stream_proto_depIdxs = nil
}

const PolygonMessageReceivedTypeCode uint32 = 764086940
const PolygonMessageSendTypeCode uint32 = 2766463228
const PolygonForexQuoteTypeCode uint32 = 3646923243
const PolygonForexAggregateTypeCode uint32 = 1453767431
const PolygonForexCombinedTypeCode uint32 = 2532026489

//true
//true
//false
//false
type PolygonMessageReceivedWrapper struct {
	stream.BaseMessageWrapper
	Data *PolygonMessageReceived
}

func (self *PolygonMessageReceivedWrapper) Message() proto1.Message {
	return self.Data
}

func (self *PolygonMessageReceivedWrapper) messageWrapper() stream.IMessageWrapper {
	return self
}

func NewPolygonMessageReceivedWrapper(
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	toReactor stream.ToReactorFunc,
	toConnection stream.ToConnectionFunc,
	data *PolygonMessageReceived) *PolygonMessageReceivedWrapper {
	return &PolygonMessageReceivedWrapper{
		BaseMessageWrapper: stream.NewBaseMessageWrapper(
			cancelCtx,
			cancelFunc,
			toReactor,
			toConnection),
		Data: data,
	}
}

var _ = stream.Register(
	PolygonMessageReceivedTypeCode,
	stream.TypeCodeData{
		CreateMessage: func() proto1.Message {
			return &PolygonMessageReceived{}
		},
		CreateWrapper: func(
			cancelCtx context.Context,
			cancelFunc context.CancelFunc,
			toReactor stream.ToReactorFunc,
			toConnection stream.ToConnectionFunc,
			data proto1.Message) (stream.IMessageWrapper, error) {
			if msg, ok := data.(*PolygonMessageReceived); ok {
				return NewPolygonMessageReceivedWrapper(
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
type PolygonMessageSendWrapper struct {
	stream.BaseMessageWrapper
	Data *PolygonMessageSend
}

func (self *PolygonMessageSendWrapper) Message() proto1.Message {
	return self.Data
}

func (self *PolygonMessageSendWrapper) messageWrapper() stream.IMessageWrapper {
	return self
}

func NewPolygonMessageSendWrapper(
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	toReactor stream.ToReactorFunc,
	toConnection stream.ToConnectionFunc,
	data *PolygonMessageSend) *PolygonMessageSendWrapper {
	return &PolygonMessageSendWrapper{
		BaseMessageWrapper: stream.NewBaseMessageWrapper(
			cancelCtx,
			cancelFunc,
			toReactor,
			toConnection),
		Data: data,
	}
}

var _ = stream.Register(
	PolygonMessageSendTypeCode,
	stream.TypeCodeData{
		CreateMessage: func() proto1.Message {
			return &PolygonMessageSend{}
		},
		CreateWrapper: func(
			cancelCtx context.Context,
			cancelFunc context.CancelFunc,
			toReactor stream.ToReactorFunc,
			toConnection stream.ToConnectionFunc,
			data proto1.Message) (stream.IMessageWrapper, error) {
			if msg, ok := data.(*PolygonMessageSend); ok {
				return NewPolygonMessageSendWrapper(
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
type PolygonForexQuoteWrapper struct {
	stream.BaseMessageWrapper
	Data *PolygonForexQuote
}

func (self *PolygonForexQuoteWrapper) Message() proto1.Message {
	return self.Data
}

func (self *PolygonForexQuoteWrapper) messageWrapper() stream.IMessageWrapper {
	return self
}

func NewPolygonForexQuoteWrapper(
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	toReactor stream.ToReactorFunc,
	toConnection stream.ToConnectionFunc,
	data *PolygonForexQuote) *PolygonForexQuoteWrapper {
	return &PolygonForexQuoteWrapper{
		BaseMessageWrapper: stream.NewBaseMessageWrapper(
			cancelCtx,
			cancelFunc,
			toReactor,
			toConnection),
		Data: data,
	}
}

var _ = stream.Register(
	PolygonForexQuoteTypeCode,
	stream.TypeCodeData{
		CreateMessage: func() proto1.Message {
			return &PolygonForexQuote{}
		},
		CreateWrapper: func(
			cancelCtx context.Context,
			cancelFunc context.CancelFunc,
			toReactor stream.ToReactorFunc,
			toConnection stream.ToConnectionFunc,
			data proto1.Message) (stream.IMessageWrapper, error) {
			if msg, ok := data.(*PolygonForexQuote); ok {
				return NewPolygonForexQuoteWrapper(
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
type PolygonForexAggregateWrapper struct {
	stream.BaseMessageWrapper
	Data *PolygonForexAggregate
}

func (self *PolygonForexAggregateWrapper) Message() proto1.Message {
	return self.Data
}

func (self *PolygonForexAggregateWrapper) messageWrapper() stream.IMessageWrapper {
	return self
}

func NewPolygonForexAggregateWrapper(
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	toReactor stream.ToReactorFunc,
	toConnection stream.ToConnectionFunc,
	data *PolygonForexAggregate) *PolygonForexAggregateWrapper {
	return &PolygonForexAggregateWrapper{
		BaseMessageWrapper: stream.NewBaseMessageWrapper(
			cancelCtx,
			cancelFunc,
			toReactor,
			toConnection),
		Data: data,
	}
}

var _ = stream.Register(
	PolygonForexAggregateTypeCode,
	stream.TypeCodeData{
		CreateMessage: func() proto1.Message {
			return &PolygonForexAggregate{}
		},
		CreateWrapper: func(
			cancelCtx context.Context,
			cancelFunc context.CancelFunc,
			toReactor stream.ToReactorFunc,
			toConnection stream.ToConnectionFunc,
			data proto1.Message) (stream.IMessageWrapper, error) {
			if msg, ok := data.(*PolygonForexAggregate); ok {
				return NewPolygonForexAggregateWrapper(
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
type PolygonForexCombinedWrapper struct {
	stream.BaseMessageWrapper
	Data *PolygonForexCombined
}

func (self *PolygonForexCombinedWrapper) Message() proto1.Message {
	return self.Data
}

func (self *PolygonForexCombinedWrapper) messageWrapper() stream.IMessageWrapper {
	return self
}

func NewPolygonForexCombinedWrapper(
	cancelCtx context.Context,
	cancelFunc context.CancelFunc,
	toReactor stream.ToReactorFunc,
	toConnection stream.ToConnectionFunc,
	data *PolygonForexCombined) *PolygonForexCombinedWrapper {
	return &PolygonForexCombinedWrapper{
		BaseMessageWrapper: stream.NewBaseMessageWrapper(
			cancelCtx,
			cancelFunc,
			toReactor,
			toConnection),
		Data: data,
	}
}

var _ = stream.Register(
	PolygonForexCombinedTypeCode,
	stream.TypeCodeData{
		CreateMessage: func() proto1.Message {
			return &PolygonForexCombined{}
		},
		CreateWrapper: func(
			cancelCtx context.Context,
			cancelFunc context.CancelFunc,
			toReactor stream.ToReactorFunc,
			toConnection stream.ToConnectionFunc,
			data proto1.Message) (stream.IMessageWrapper, error) {
			if msg, ok := data.(*PolygonForexCombined); ok {
				return NewPolygonForexCombinedWrapper(
					cancelCtx,
					cancelFunc,
					toReactor,
					toConnection,
					msg), nil
			}
			return nil, constants.InvalidParam
		}})