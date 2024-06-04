// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: example/speakers.proto

package example

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SpeakerServiceClient is the client API for SpeakerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpeakerServiceClient interface {
	// ListTalks is used to get all talks according to any provided query params within ListTalksRequest
	ListSpeakers(ctx context.Context, in *ListSpeakersRequest, opts ...grpc.CallOption) (*ListSpeakersResponse, error)
	// CreateSpeaker is used to get all talks according to any provided query params within ListTalksRequest
	CreateSpeaker(ctx context.Context, in *CreateSpeakerRequest, opts ...grpc.CallOption) (*Speaker, error)
}

type speakerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpeakerServiceClient(cc grpc.ClientConnInterface) SpeakerServiceClient {
	return &speakerServiceClient{cc}
}

func (c *speakerServiceClient) ListSpeakers(ctx context.Context, in *ListSpeakersRequest, opts ...grpc.CallOption) (*ListSpeakersResponse, error) {
	out := new(ListSpeakersResponse)
	err := c.cc.Invoke(ctx, "/example.SpeakerService/ListSpeakers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *speakerServiceClient) CreateSpeaker(ctx context.Context, in *CreateSpeakerRequest, opts ...grpc.CallOption) (*Speaker, error) {
	out := new(Speaker)
	err := c.cc.Invoke(ctx, "/example.SpeakerService/CreateSpeaker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpeakerServiceServer is the server API for SpeakerService service.
// All implementations must embed UnimplementedSpeakerServiceServer
// for forward compatibility
type SpeakerServiceServer interface {
	// ListTalks is used to get all talks according to any provided query params within ListTalksRequest
	ListSpeakers(context.Context, *ListSpeakersRequest) (*ListSpeakersResponse, error)
	// CreateSpeaker is used to get all talks according to any provided query params within ListTalksRequest
	CreateSpeaker(context.Context, *CreateSpeakerRequest) (*Speaker, error)
	mustEmbedUnimplementedSpeakerServiceServer()
}

// UnimplementedSpeakerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSpeakerServiceServer struct {
}

func (UnimplementedSpeakerServiceServer) ListSpeakers(context.Context, *ListSpeakersRequest) (*ListSpeakersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSpeakers not implemented")
}
func (UnimplementedSpeakerServiceServer) CreateSpeaker(context.Context, *CreateSpeakerRequest) (*Speaker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpeaker not implemented")
}
func (UnimplementedSpeakerServiceServer) mustEmbedUnimplementedSpeakerServiceServer() {}

// UnsafeSpeakerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpeakerServiceServer will
// result in compilation errors.
type UnsafeSpeakerServiceServer interface {
	mustEmbedUnimplementedSpeakerServiceServer()
}

func RegisterSpeakerServiceServer(s grpc.ServiceRegistrar, srv SpeakerServiceServer) {
	s.RegisterService(&SpeakerService_ServiceDesc, srv)
}

func _SpeakerService_ListSpeakers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSpeakersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpeakerServiceServer).ListSpeakers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.SpeakerService/ListSpeakers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpeakerServiceServer).ListSpeakers(ctx, req.(*ListSpeakersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpeakerService_CreateSpeaker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSpeakerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpeakerServiceServer).CreateSpeaker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.SpeakerService/CreateSpeaker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpeakerServiceServer).CreateSpeaker(ctx, req.(*CreateSpeakerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpeakerService_ServiceDesc is the grpc.ServiceDesc for SpeakerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpeakerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.SpeakerService",
	HandlerType: (*SpeakerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSpeakers",
			Handler:    _SpeakerService_ListSpeakers_Handler,
		},
		{
			MethodName: "CreateSpeaker",
			Handler:    _SpeakerService_CreateSpeaker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/speakers.proto",
}