// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        v5.28.0
// source: valr/valr.proto

package stream

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

type Subscriptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event string   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Pairs []string `protobuf:"bytes,2,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *Subscriptions) Reset() {
	*x = Subscriptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valr_valr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscriptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscriptions) ProtoMessage() {}

func (x *Subscriptions) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Subscriptions.ProtoReflect.Descriptor instead.
func (*Subscriptions) Descriptor() ([]byte, []int) {
	return file_valr_valr_proto_rawDescGZIP(), []int{0}
}

func (x *Subscriptions) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *Subscriptions) GetPairs() []string {
	if x != nil {
		return x.Pairs
	}
	return nil
}

type Subscribe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Subscribe) Reset() {
	*x = Subscribe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valr_valr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscribe) ProtoMessage() {}

func (x *Subscribe) ProtoReflect() protoreflect.Message {
	mi := &file_valr_valr_proto_msgTypes[1]
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
	return file_valr_valr_proto_rawDescGZIP(), []int{1}
}

func (x *Subscribe) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_valr_valr_proto protoreflect.FileDescriptor

var file_valr_valr_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3b, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x69, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x22, 0x1f,
	0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42,
	0x09, 0x5a, 0x07, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_valr_valr_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_valr_valr_proto_goTypes = []interface{}{
	(*Subscriptions)(nil), // 0: Subscriptions
	(*Subscribe)(nil),     // 1: Subscribe
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
			switch v := v.(*Subscriptions); i {
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
		file_valr_valr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_valr_valr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
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
