// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        v5.28.0
// source: pingpong/pingpong.proto

package stream

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId        int64                  `protobuf:"varint,1,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	RequestTimeStamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=RequestTimeStamp,proto3" json:"RequestTimeStamp,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pingpong_pingpong_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_pingpong_pingpong_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_pingpong_pingpong_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *Ping) GetRequestTimeStamp() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestTimeStamp
	}
	return nil
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId         int64                  `protobuf:"varint,1,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	RequestTimeStamp  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=RequestTimeStamp,proto3" json:"RequestTimeStamp,omitempty"`
	ResponseTimeStamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=ResponseTimeStamp,proto3" json:"ResponseTimeStamp,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pingpong_pingpong_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_pingpong_pingpong_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_pingpong_pingpong_proto_rawDescGZIP(), []int{1}
}

func (x *Pong) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *Pong) GetRequestTimeStamp() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestTimeStamp
	}
	return nil
}

func (x *Pong) GetResponseTimeStamp() *timestamppb.Timestamp {
	if x != nil {
		return x.ResponseTimeStamp
	}
	return nil
}

type TestMessageForPingPong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AString string `protobuf:"bytes,1,opt,name=AString,proto3" json:"AString,omitempty"`
}

func (x *TestMessageForPingPong) Reset() {
	*x = TestMessageForPingPong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pingpong_pingpong_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessageForPingPong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessageForPingPong) ProtoMessage() {}

func (x *TestMessageForPingPong) ProtoReflect() protoreflect.Message {
	mi := &file_pingpong_pingpong_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessageForPingPong.ProtoReflect.Descriptor instead.
func (*TestMessageForPingPong) Descriptor() ([]byte, []int) {
	return file_pingpong_pingpong_proto_rawDescGZIP(), []int{2}
}

func (x *TestMessageForPingPong) GetAString() string {
	if x != nil {
		return x.AString
	}
	return ""
}

var File_pingpong_pingpong_proto protoreflect.FileDescriptor

var file_pingpong_pingpong_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6e, 0x67, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x70,
	0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x46, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xb6, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x6e,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x46, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x48, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x32, 0x0a, 0x16, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x46, 0x6f, 0x72, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x41,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pingpong_pingpong_proto_rawDescOnce sync.Once
	file_pingpong_pingpong_proto_rawDescData = file_pingpong_pingpong_proto_rawDesc
)

func file_pingpong_pingpong_proto_rawDescGZIP() []byte {
	file_pingpong_pingpong_proto_rawDescOnce.Do(func() {
		file_pingpong_pingpong_proto_rawDescData = protoimpl.X.CompressGZIP(file_pingpong_pingpong_proto_rawDescData)
	})
	return file_pingpong_pingpong_proto_rawDescData
}

var file_pingpong_pingpong_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pingpong_pingpong_proto_goTypes = []interface{}{
	(*Ping)(nil),                   // 0: Ping
	(*Pong)(nil),                   // 1: Pong
	(*TestMessageForPingPong)(nil), // 2: TestMessageForPingPong
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
}
var file_pingpong_pingpong_proto_depIdxs = []int32{
	3, // 0: Ping.RequestTimeStamp:type_name -> google.protobuf.Timestamp
	3, // 1: Pong.RequestTimeStamp:type_name -> google.protobuf.Timestamp
	3, // 2: Pong.ResponseTimeStamp:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pingpong_pingpong_proto_init() }
func file_pingpong_pingpong_proto_init() {
	if File_pingpong_pingpong_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pingpong_pingpong_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_pingpong_pingpong_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
		file_pingpong_pingpong_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessageForPingPong); i {
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
			RawDescriptor: file_pingpong_pingpong_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pingpong_pingpong_proto_goTypes,
		DependencyIndexes: file_pingpong_pingpong_proto_depIdxs,
		MessageInfos:      file_pingpong_pingpong_proto_msgTypes,
	}.Build()
	File_pingpong_pingpong_proto = out.File
	file_pingpong_pingpong_proto_rawDesc = nil
	file_pingpong_pingpong_proto_goTypes = nil
	file_pingpong_pingpong_proto_depIdxs = nil
}
