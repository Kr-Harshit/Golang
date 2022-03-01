// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pkg/gopher/gopher.proto

package gopher_GRPC

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The Request message containing the user's name.
type GopherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GopherRequest) Reset() {
	*x = GopherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_gopher_gopher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GopherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GopherRequest) ProtoMessage() {}

func (x *GopherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_gopher_gopher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GopherRequest.ProtoReflect.Descriptor instead.
func (*GopherRequest) Descriptor() ([]byte, []int) {
	return file_pkg_gopher_gopher_proto_rawDescGZIP(), []int{0}
}

func (x *GopherRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type GopherReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GopherReply) Reset() {
	*x = GopherReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_gopher_gopher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GopherReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GopherReply) ProtoMessage() {}

func (x *GopherReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_gopher_gopher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GopherReply.ProtoReflect.Descriptor instead.
func (*GopherReply) Descriptor() ([]byte, []int) {
	return file_pkg_gopher_gopher_proto_rawDescGZIP(), []int{1}
}

func (x *GopherReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_gopher_gopher_proto protoreflect.FileDescriptor

var file_pkg_gopher_gopher_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x70,
	0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x6f, 0x70, 0x68, 0x65,
	0x72, 0x22, 0x23, 0x0a, 0x0d, 0x47, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x47, 0x6f, 0x70, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x43, 0x0a, 0x06, 0x47, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x47, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x2e,
	0x47, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4b, 0x72, 0x5f, 0x48, 0x61, 0x72, 0x73, 0x68, 0x69, 0x74, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x47, 0x52, 0x50, 0x43,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_gopher_gopher_proto_rawDescOnce sync.Once
	file_pkg_gopher_gopher_proto_rawDescData = file_pkg_gopher_gopher_proto_rawDesc
)

func file_pkg_gopher_gopher_proto_rawDescGZIP() []byte {
	file_pkg_gopher_gopher_proto_rawDescOnce.Do(func() {
		file_pkg_gopher_gopher_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_gopher_gopher_proto_rawDescData)
	})
	return file_pkg_gopher_gopher_proto_rawDescData
}

var file_pkg_gopher_gopher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_gopher_gopher_proto_goTypes = []interface{}{
	(*GopherRequest)(nil), // 0: gopher.GopherRequest
	(*GopherReply)(nil),   // 1: gopher.GopherReply
}
var file_pkg_gopher_gopher_proto_depIdxs = []int32{
	0, // 0: gopher.Gopher.GetGopher:input_type -> gopher.GopherRequest
	1, // 1: gopher.Gopher.GetGopher:output_type -> gopher.GopherReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_gopher_gopher_proto_init() }
func file_pkg_gopher_gopher_proto_init() {
	if File_pkg_gopher_gopher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_gopher_gopher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GopherRequest); i {
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
		file_pkg_gopher_gopher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GopherReply); i {
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
			RawDescriptor: file_pkg_gopher_gopher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_gopher_gopher_proto_goTypes,
		DependencyIndexes: file_pkg_gopher_gopher_proto_depIdxs,
		MessageInfos:      file_pkg_gopher_gopher_proto_msgTypes,
	}.Build()
	File_pkg_gopher_gopher_proto = out.File
	file_pkg_gopher_gopher_proto_rawDesc = nil
	file_pkg_gopher_gopher_proto_goTypes = nil
	file_pkg_gopher_gopher_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GopherClient is the client API for Gopher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GopherClient interface {
	// Get Gopher URL
	GetGopher(ctx context.Context, in *GopherRequest, opts ...grpc.CallOption) (*GopherReply, error)
}

type gopherClient struct {
	cc grpc.ClientConnInterface
}

func NewGopherClient(cc grpc.ClientConnInterface) GopherClient {
	return &gopherClient{cc}
}

func (c *gopherClient) GetGopher(ctx context.Context, in *GopherRequest, opts ...grpc.CallOption) (*GopherReply, error) {
	out := new(GopherReply)
	err := c.cc.Invoke(ctx, "/gopher.Gopher/GetGopher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GopherServer is the server API for Gopher service.
type GopherServer interface {
	// Get Gopher URL
	GetGopher(context.Context, *GopherRequest) (*GopherReply, error)
}

// UnimplementedGopherServer can be embedded to have forward compatible implementations.
type UnimplementedGopherServer struct {
}

func (*UnimplementedGopherServer) GetGopher(context.Context, *GopherRequest) (*GopherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGopher not implemented")
}

func RegisterGopherServer(s *grpc.Server, srv GopherServer) {
	s.RegisterService(&_Gopher_serviceDesc, srv)
}

func _Gopher_GetGopher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GopherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GopherServer).GetGopher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gopher.Gopher/GetGopher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GopherServer).GetGopher(ctx, req.(*GopherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gopher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gopher.Gopher",
	HandlerType: (*GopherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGopher",
			Handler:    _Gopher_GetGopher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/gopher/gopher.proto",
}
