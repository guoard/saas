// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: protobuf/queue.proto

package protobuf

import (
	proto "github.com/golang/protobuf/proto"
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

type QueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *QueueRequest) Reset() {
	*x = QueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_queue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueRequest) ProtoMessage() {}

func (x *QueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_queue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueRequest.ProtoReflect.Descriptor instead.
func (*QueueRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_queue_proto_rawDescGZIP(), []int{0}
}

func (x *QueueRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type QueueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *QueueResponse) Reset() {
	*x = QueueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_queue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueResponse) ProtoMessage() {}

func (x *QueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_queue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueResponse.ProtoReflect.Descriptor instead.
func (*QueueResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_queue_proto_rawDescGZIP(), []int{1}
}

func (x *QueueResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_protobuf_queue_proto protoreflect.FileDescriptor

var file_protobuf_queue_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x23, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x32, 0x33, 0x0a,
	0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x0d, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x32, 0x68, 0x61, 0x6d, 0x65, 0x64, 0x2f, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_queue_proto_rawDescOnce sync.Once
	file_protobuf_queue_proto_rawDescData = file_protobuf_queue_proto_rawDesc
)

func file_protobuf_queue_proto_rawDescGZIP() []byte {
	file_protobuf_queue_proto_rawDescOnce.Do(func() {
		file_protobuf_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_queue_proto_rawDescData)
	})
	return file_protobuf_queue_proto_rawDescData
}

var file_protobuf_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_queue_proto_goTypes = []interface{}{
	(*QueueRequest)(nil),  // 0: QueueRequest
	(*QueueResponse)(nil), // 1: QueueResponse
}
var file_protobuf_queue_proto_depIdxs = []int32{
	0, // 0: Queue.Capture:input_type -> QueueRequest
	1, // 1: Queue.Capture:output_type -> QueueResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_queue_proto_init() }
func file_protobuf_queue_proto_init() {
	if File_protobuf_queue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_queue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueRequest); i {
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
		file_protobuf_queue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueResponse); i {
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
			RawDescriptor: file_protobuf_queue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_queue_proto_goTypes,
		DependencyIndexes: file_protobuf_queue_proto_depIdxs,
		MessageInfos:      file_protobuf_queue_proto_msgTypes,
	}.Build()
	File_protobuf_queue_proto = out.File
	file_protobuf_queue_proto_rawDesc = nil
	file_protobuf_queue_proto_goTypes = nil
	file_protobuf_queue_proto_depIdxs = nil
}
