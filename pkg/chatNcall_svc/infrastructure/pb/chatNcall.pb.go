// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/chatNcall_svc/infrastructure/pb/chatNcall.proto

package pb

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

type RequestUserChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestUserChat) Reset() {
	*x = RequestUserChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestUserChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestUserChat) ProtoMessage() {}

func (x *RequestUserChat) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestUserChat.ProtoReflect.Descriptor instead.
func (*RequestUserChat) Descriptor() ([]byte, []int) {
	return file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_rawDescGZIP(), []int{0}
}

type ResponseUserChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResponseUserChat) Reset() {
	*x = ResponseUserChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseUserChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseUserChat) ProtoMessage() {}

func (x *ResponseUserChat) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseUserChat.ProtoReflect.Descriptor instead.
func (*ResponseUserChat) Descriptor() ([]byte, []int) {
	return file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_rawDescGZIP(), []int{1}
}

var File_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto protoreflect.FileDescriptor

var file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_rawDesc = []byte{
	0x0a, 0x33, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x4e, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x73, 0x76, 0x63, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x4e, 0x63, 0x61, 0x6c, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x4e, 0x63, 0x61, 0x6c, 0x6c,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x32, 0x67, 0x0a,
	0x10, 0x43, 0x68, 0x61, 0x74, 0x4e, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x53, 0x0a, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61,
	0x74, 0x12, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x4e, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x68, 0x61, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x4e, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x42, 0x27, 0x5a, 0x25, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x4e, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x76, 0x63, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_rawDescOnce sync.Once
	file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_rawDescData = file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_rawDesc
)

func file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_rawDescGZIP() []byte {
	file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_rawDescOnce.Do(func() {
		file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_rawDescData)
	})
	return file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_rawDescData
}

var file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_goTypes = []interface{}{
	(*RequestUserChat)(nil),  // 0: chatNcall_proto.RequestUserChat
	(*ResponseUserChat)(nil), // 1: chatNcall_proto.ResponseUserChat
}
var file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_depIdxs = []int32{
	0, // 0: chatNcall_proto.ChatNCallService.PassUserChat:input_type -> chatNcall_proto.RequestUserChat
	1, // 1: chatNcall_proto.ChatNCallService.PassUserChat:output_type -> chatNcall_proto.ResponseUserChat
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_init() }
func file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_init() {
	if File_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestUserChat); i {
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
		file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseUserChat); i {
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
			RawDescriptor: file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_goTypes,
		DependencyIndexes: file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_depIdxs,
		MessageInfos:      file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_msgTypes,
	}.Build()
	File_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto = out.File
	file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_rawDesc = nil
	file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_goTypes = nil
	file_pkg_chatNcall_svc_infrastructure_pb_chatNcall_proto_depIdxs = nil
}
