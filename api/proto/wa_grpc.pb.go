// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: api/proto/wa.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WhatsAppService_SendTextMessage_FullMethodName         = "/proto.WhatsAppService/SendTextMessage"
	WhatsAppService_SendExtendedTextMessage_FullMethodName = "/proto.WhatsAppService/SendExtendedTextMessage"
	WhatsAppService_UploadMedia_FullMethodName             = "/proto.WhatsAppService/UploadMedia"
	WhatsAppService_StoreFileMetadata_FullMethodName       = "/proto.WhatsAppService/StoreFileMetadata"
)

// WhatsAppServiceClient is the client API for WhatsAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WhatsAppServiceClient interface {
	// ✅ Send a text message
	SendTextMessage(ctx context.Context, in *TextMessageRequest, opts ...grpc.CallOption) (*CommonMessageResponse, error)
	// ✅ Send an extended text message
	SendExtendedTextMessage(ctx context.Context, in *ExtendedTextMessageRequest, opts ...grpc.CallOption) (*CommonMessageResponse, error)
	// ✅ Upload media (image, video, audio, document, sticker)
	UploadMedia(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[MediaUploadRequest, MediaUploadResponse], error)
	// ✅ Store file metadata
	StoreFileMetadata(ctx context.Context, in *FileMetadataRequest, opts ...grpc.CallOption) (*FileMetadataResponse, error)
}

type whatsAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWhatsAppServiceClient(cc grpc.ClientConnInterface) WhatsAppServiceClient {
	return &whatsAppServiceClient{cc}
}

func (c *whatsAppServiceClient) SendTextMessage(ctx context.Context, in *TextMessageRequest, opts ...grpc.CallOption) (*CommonMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonMessageResponse)
	err := c.cc.Invoke(ctx, WhatsAppService_SendTextMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whatsAppServiceClient) SendExtendedTextMessage(ctx context.Context, in *ExtendedTextMessageRequest, opts ...grpc.CallOption) (*CommonMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonMessageResponse)
	err := c.cc.Invoke(ctx, WhatsAppService_SendExtendedTextMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whatsAppServiceClient) UploadMedia(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[MediaUploadRequest, MediaUploadResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WhatsAppService_ServiceDesc.Streams[0], WhatsAppService_UploadMedia_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[MediaUploadRequest, MediaUploadResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WhatsAppService_UploadMediaClient = grpc.ClientStreamingClient[MediaUploadRequest, MediaUploadResponse]

func (c *whatsAppServiceClient) StoreFileMetadata(ctx context.Context, in *FileMetadataRequest, opts ...grpc.CallOption) (*FileMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileMetadataResponse)
	err := c.cc.Invoke(ctx, WhatsAppService_StoreFileMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WhatsAppServiceServer is the server API for WhatsAppService service.
// All implementations must embed UnimplementedWhatsAppServiceServer
// for forward compatibility.
type WhatsAppServiceServer interface {
	// ✅ Send a text message
	SendTextMessage(context.Context, *TextMessageRequest) (*CommonMessageResponse, error)
	// ✅ Send an extended text message
	SendExtendedTextMessage(context.Context, *ExtendedTextMessageRequest) (*CommonMessageResponse, error)
	// ✅ Upload media (image, video, audio, document, sticker)
	UploadMedia(grpc.ClientStreamingServer[MediaUploadRequest, MediaUploadResponse]) error
	// ✅ Store file metadata
	StoreFileMetadata(context.Context, *FileMetadataRequest) (*FileMetadataResponse, error)
	mustEmbedUnimplementedWhatsAppServiceServer()
}

// UnimplementedWhatsAppServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWhatsAppServiceServer struct{}

func (UnimplementedWhatsAppServiceServer) SendTextMessage(context.Context, *TextMessageRequest) (*CommonMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTextMessage not implemented")
}
func (UnimplementedWhatsAppServiceServer) SendExtendedTextMessage(context.Context, *ExtendedTextMessageRequest) (*CommonMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendExtendedTextMessage not implemented")
}
func (UnimplementedWhatsAppServiceServer) UploadMedia(grpc.ClientStreamingServer[MediaUploadRequest, MediaUploadResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UploadMedia not implemented")
}
func (UnimplementedWhatsAppServiceServer) StoreFileMetadata(context.Context, *FileMetadataRequest) (*FileMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreFileMetadata not implemented")
}
func (UnimplementedWhatsAppServiceServer) mustEmbedUnimplementedWhatsAppServiceServer() {}
func (UnimplementedWhatsAppServiceServer) testEmbeddedByValue()                         {}

// UnsafeWhatsAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WhatsAppServiceServer will
// result in compilation errors.
type UnsafeWhatsAppServiceServer interface {
	mustEmbedUnimplementedWhatsAppServiceServer()
}

func RegisterWhatsAppServiceServer(s grpc.ServiceRegistrar, srv WhatsAppServiceServer) {
	// If the following call pancis, it indicates UnimplementedWhatsAppServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WhatsAppService_ServiceDesc, srv)
}

func _WhatsAppService_SendTextMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhatsAppServiceServer).SendTextMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WhatsAppService_SendTextMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhatsAppServiceServer).SendTextMessage(ctx, req.(*TextMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WhatsAppService_SendExtendedTextMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtendedTextMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhatsAppServiceServer).SendExtendedTextMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WhatsAppService_SendExtendedTextMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhatsAppServiceServer).SendExtendedTextMessage(ctx, req.(*ExtendedTextMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WhatsAppService_UploadMedia_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WhatsAppServiceServer).UploadMedia(&grpc.GenericServerStream[MediaUploadRequest, MediaUploadResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WhatsAppService_UploadMediaServer = grpc.ClientStreamingServer[MediaUploadRequest, MediaUploadResponse]

func _WhatsAppService_StoreFileMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhatsAppServiceServer).StoreFileMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WhatsAppService_StoreFileMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhatsAppServiceServer).StoreFileMetadata(ctx, req.(*FileMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WhatsAppService_ServiceDesc is the grpc.ServiceDesc for WhatsAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WhatsAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WhatsAppService",
	HandlerType: (*WhatsAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTextMessage",
			Handler:    _WhatsAppService_SendTextMessage_Handler,
		},
		{
			MethodName: "SendExtendedTextMessage",
			Handler:    _WhatsAppService_SendExtendedTextMessage_Handler,
		},
		{
			MethodName: "StoreFileMetadata",
			Handler:    _WhatsAppService_StoreFileMetadata_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadMedia",
			Handler:       _WhatsAppService_UploadMedia_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/proto/wa.proto",
}
