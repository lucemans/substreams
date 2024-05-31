// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sf/remotebuild/v1/remotebuild.proto

package pbbuildv1

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

const (
	BuildService_Build_FullMethodName = "/sf.remotebuild.v1.BuildService/Build"
)

// BuildServiceClient is the client API for BuildService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuildServiceClient interface {
	Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error)
}

type buildServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBuildServiceClient(cc grpc.ClientConnInterface) BuildServiceClient {
	return &buildServiceClient{cc}
}

func (c *buildServiceClient) Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := c.cc.Invoke(ctx, BuildService_Build_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildServiceServer is the server API for BuildService service.
// All implementations should embed UnimplementedBuildServiceServer
// for forward compatibility
type BuildServiceServer interface {
	Build(context.Context, *BuildRequest) (*BuildResponse, error)
}

// UnimplementedBuildServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBuildServiceServer struct {
}

func (UnimplementedBuildServiceServer) Build(context.Context, *BuildRequest) (*BuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Build not implemented")
}

// UnsafeBuildServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuildServiceServer will
// result in compilation errors.
type UnsafeBuildServiceServer interface {
	mustEmbedUnimplementedBuildServiceServer()
}

func RegisterBuildServiceServer(s grpc.ServiceRegistrar, srv BuildServiceServer) {
	s.RegisterService(&BuildService_ServiceDesc, srv)
}

func _BuildService_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServiceServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BuildService_Build_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServiceServer).Build(ctx, req.(*BuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BuildService_ServiceDesc is the grpc.ServiceDesc for BuildService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BuildService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sf.remotebuild.v1.BuildService",
	HandlerType: (*BuildServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Build",
			Handler:    _BuildService_Build_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sf/remotebuild/v1/remotebuild.proto",
}
