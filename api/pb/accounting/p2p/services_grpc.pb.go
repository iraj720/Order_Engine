// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: accounting/p2p/services.proto

package p2p

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

// P2PAPIClient is the client API for P2PAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type P2PAPIClient interface {
	P2PTransaction(ctx context.Context, in *P2PTransactionRequest, opts ...grpc.CallOption) (*P2PTransactionResponse, error)
	P2PTrade(ctx context.Context, in *P2PTradeRequest, opts ...grpc.CallOption) (*P2PTradeResponse, error)
}

type p2PAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewP2PAPIClient(cc grpc.ClientConnInterface) P2PAPIClient {
	return &p2PAPIClient{cc}
}

func (c *p2PAPIClient) P2PTransaction(ctx context.Context, in *P2PTransactionRequest, opts ...grpc.CallOption) (*P2PTransactionResponse, error) {
	out := new(P2PTransactionResponse)
	err := c.cc.Invoke(ctx, "/p2p.P2PAPI/P2PTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PAPIClient) P2PTrade(ctx context.Context, in *P2PTradeRequest, opts ...grpc.CallOption) (*P2PTradeResponse, error) {
	out := new(P2PTradeResponse)
	err := c.cc.Invoke(ctx, "/p2p.P2PAPI/P2PTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// P2PAPIServer is the server API for P2PAPI service.
// All implementations must embed UnimplementedP2PAPIServer
// for forward compatibility
type P2PAPIServer interface {
	P2PTransaction(context.Context, *P2PTransactionRequest) (*P2PTransactionResponse, error)
	P2PTrade(context.Context, *P2PTradeRequest) (*P2PTradeResponse, error)
	mustEmbedUnimplementedP2PAPIServer()
}

// UnimplementedP2PAPIServer must be embedded to have forward compatible implementations.
type UnimplementedP2PAPIServer struct {
}

func (UnimplementedP2PAPIServer) P2PTransaction(context.Context, *P2PTransactionRequest) (*P2PTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method P2PTransaction not implemented")
}
func (UnimplementedP2PAPIServer) P2PTrade(context.Context, *P2PTradeRequest) (*P2PTradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method P2PTrade not implemented")
}
func (UnimplementedP2PAPIServer) mustEmbedUnimplementedP2PAPIServer() {}

// UnsafeP2PAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to P2PAPIServer will
// result in compilation errors.
type UnsafeP2PAPIServer interface {
	mustEmbedUnimplementedP2PAPIServer()
}

func RegisterP2PAPIServer(s grpc.ServiceRegistrar, srv P2PAPIServer) {
	s.RegisterService(&P2PAPI_ServiceDesc, srv)
}

func _P2PAPI_P2PTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(P2PTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PAPIServer).P2PTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2p.P2PAPI/P2PTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PAPIServer).P2PTransaction(ctx, req.(*P2PTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PAPI_P2PTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(P2PTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PAPIServer).P2PTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2p.P2PAPI/P2PTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PAPIServer).P2PTrade(ctx, req.(*P2PTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// P2PAPI_ServiceDesc is the grpc.ServiceDesc for P2PAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var P2PAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "p2p.P2PAPI",
	HandlerType: (*P2PAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "P2PTransaction",
			Handler:    _P2PAPI_P2PTransaction_Handler,
		},
		{
			MethodName: "P2PTrade",
			Handler:    _P2PAPI_P2PTrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounting/p2p/services.proto",
}
