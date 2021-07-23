// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package geometry

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

// GeometryClient is the client API for Geometry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeometryClient interface {
	ComputeArea(ctx context.Context, in *AreaRequest, opts ...grpc.CallOption) (*AreaResponse, error)
	ComputePerimeter(ctx context.Context, in *PerimeterRequest, opts ...grpc.CallOption) (*PerimeterResponse, error)
}

type geometryClient struct {
	cc grpc.ClientConnInterface
}

func NewGeometryClient(cc grpc.ClientConnInterface) GeometryClient {
	return &geometryClient{cc}
}

func (c *geometryClient) ComputeArea(ctx context.Context, in *AreaRequest, opts ...grpc.CallOption) (*AreaResponse, error) {
	out := new(AreaResponse)
	err := c.cc.Invoke(ctx, "/Geometry/ComputeArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geometryClient) ComputePerimeter(ctx context.Context, in *PerimeterRequest, opts ...grpc.CallOption) (*PerimeterResponse, error) {
	out := new(PerimeterResponse)
	err := c.cc.Invoke(ctx, "/Geometry/ComputePerimeter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeometryServer is the server API for Geometry service.
// All implementations must embed UnimplementedGeometryServer
// for forward compatibility
type GeometryServer interface {
	ComputeArea(context.Context, *AreaRequest) (*AreaResponse, error)
	ComputePerimeter(context.Context, *PerimeterRequest) (*PerimeterResponse, error)
	mustEmbedUnimplementedGeometryServer()
}

// UnimplementedGeometryServer must be embedded to have forward compatible implementations.
type UnimplementedGeometryServer struct {
}

func (UnimplementedGeometryServer) ComputeArea(context.Context, *AreaRequest) (*AreaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeArea not implemented")
}
func (UnimplementedGeometryServer) ComputePerimeter(context.Context, *PerimeterRequest) (*PerimeterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputePerimeter not implemented")
}
func (UnimplementedGeometryServer) mustEmbedUnimplementedGeometryServer() {}

// UnsafeGeometryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeometryServer will
// result in compilation errors.
type UnsafeGeometryServer interface {
	mustEmbedUnimplementedGeometryServer()
}

func RegisterGeometryServer(s grpc.ServiceRegistrar, srv GeometryServer) {
	s.RegisterService(&Geometry_ServiceDesc, srv)
}

func _Geometry_ComputeArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeometryServer).ComputeArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Geometry/ComputeArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeometryServer).ComputeArea(ctx, req.(*AreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Geometry_ComputePerimeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerimeterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeometryServer).ComputePerimeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Geometry/ComputePerimeter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeometryServer).ComputePerimeter(ctx, req.(*PerimeterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Geometry_ServiceDesc is the grpc.ServiceDesc for Geometry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Geometry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Geometry",
	HandlerType: (*GeometryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeArea",
			Handler:    _Geometry_ComputeArea_Handler,
		},
		{
			MethodName: "ComputePerimeter",
			Handler:    _Geometry_ComputePerimeter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geometry.proto",
}