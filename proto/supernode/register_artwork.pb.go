// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: supernode/register_artwork.proto

package supernode

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

type SessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeID string `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
}

func (x *SessionRequest) Reset() {
	*x = SessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supernode_register_artwork_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionRequest) ProtoMessage() {}

func (x *SessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supernode_register_artwork_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionRequest.ProtoReflect.Descriptor instead.
func (*SessionRequest) Descriptor() ([]byte, []int) {
	return file_supernode_register_artwork_proto_rawDescGZIP(), []int{0}
}

func (x *SessionRequest) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

type SessionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessID string `protobuf:"bytes,1,opt,name=sessID,proto3" json:"sessID,omitempty"`
}

func (x *SessionReply) Reset() {
	*x = SessionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supernode_register_artwork_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionReply) ProtoMessage() {}

func (x *SessionReply) ProtoReflect() protoreflect.Message {
	mi := &file_supernode_register_artwork_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionReply.ProtoReflect.Descriptor instead.
func (*SessionReply) Descriptor() ([]byte, []int) {
	return file_supernode_register_artwork_proto_rawDescGZIP(), []int{1}
}

func (x *SessionReply) GetSessID() string {
	if x != nil {
		return x.SessID
	}
	return ""
}

var File_supernode_register_artwork_proto protoreflect.FileDescriptor

var file_supernode_register_artwork_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x73, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x28, 0x0a,
	0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x22, 0x26, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x73, 0x73, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x73, 0x73, 0x49, 0x44, 0x32,
	0x54, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x72, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x41, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e,
	0x73, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x28, 0x01, 0x30, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x73, 0x74, 0x65, 0x6c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x67, 0x6f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_supernode_register_artwork_proto_rawDescOnce sync.Once
	file_supernode_register_artwork_proto_rawDescData = file_supernode_register_artwork_proto_rawDesc
)

func file_supernode_register_artwork_proto_rawDescGZIP() []byte {
	file_supernode_register_artwork_proto_rawDescOnce.Do(func() {
		file_supernode_register_artwork_proto_rawDescData = protoimpl.X.CompressGZIP(file_supernode_register_artwork_proto_rawDescData)
	})
	return file_supernode_register_artwork_proto_rawDescData
}

var file_supernode_register_artwork_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_supernode_register_artwork_proto_goTypes = []interface{}{
	(*SessionRequest)(nil), // 0: supernode.SessionRequest
	(*SessionReply)(nil),   // 1: supernode.SessionReply
}
var file_supernode_register_artwork_proto_depIdxs = []int32{
	0, // 0: supernode.RegisterArtwork.Session:input_type -> supernode.SessionRequest
	1, // 1: supernode.RegisterArtwork.Session:output_type -> supernode.SessionReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_supernode_register_artwork_proto_init() }
func file_supernode_register_artwork_proto_init() {
	if File_supernode_register_artwork_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_supernode_register_artwork_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionRequest); i {
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
		file_supernode_register_artwork_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionReply); i {
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
			RawDescriptor: file_supernode_register_artwork_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_supernode_register_artwork_proto_goTypes,
		DependencyIndexes: file_supernode_register_artwork_proto_depIdxs,
		MessageInfos:      file_supernode_register_artwork_proto_msgTypes,
	}.Build()
	File_supernode_register_artwork_proto = out.File
	file_supernode_register_artwork_proto_rawDesc = nil
	file_supernode_register_artwork_proto_goTypes = nil
	file_supernode_register_artwork_proto_depIdxs = nil
}
