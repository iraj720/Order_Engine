// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: accounting/toman/services.proto

package toman

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

// TomanAPIClient is the client API for TomanAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TomanAPIClient interface {
	Deposit(ctx context.Context, in *TomanDepositRequest, opts ...grpc.CallOption) (*TomanDepositResponse, error)
	Withdraw(ctx context.Context, in *TomanWithdrawRequest, opts ...grpc.CallOption) (*TomanWithdrawResponse, error)
	CorrectiveDeposit(ctx context.Context, in *TomanCorrectiveDepositRequest, opts ...grpc.CallOption) (*TomanCorrectiveDepositResponse, error)
}

type tomanAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTomanAPIClient(cc grpc.ClientConnInterface) TomanAPIClient {
	return &tomanAPIClient{cc}
}

func (c *tomanAPIClient) Deposit(ctx context.Context, in *TomanDepositRequest, opts ...grpc.CallOption) (*TomanDepositResponse, error) {
	out := new(TomanDepositResponse)
	err := c.cc.Invoke(ctx, "/toman.TomanAPI/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tomanAPIClient) Withdraw(ctx context.Context, in *TomanWithdrawRequest, opts ...grpc.CallOption) (*TomanWithdrawResponse, error) {
	out := new(TomanWithdrawResponse)
	err := c.cc.Invoke(ctx, "/toman.TomanAPI/Withdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tomanAPIClient) CorrectiveDeposit(ctx context.Context, in *TomanCorrectiveDepositRequest, opts ...grpc.CallOption) (*TomanCorrectiveDepositResponse, error) {
	out := new(TomanCorrectiveDepositResponse)
	err := c.cc.Invoke(ctx, "/toman.TomanAPI/CorrectiveDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TomanAPIServer is the server API for TomanAPI service.
// All implementations must embed UnimplementedTomanAPIServer
// for forward compatibility
type TomanAPIServer interface {
	Deposit(context.Context, *TomanDepositRequest) (*TomanDepositResponse, error)
	Withdraw(context.Context, *TomanWithdrawRequest) (*TomanWithdrawResponse, error)
	CorrectiveDeposit(context.Context, *TomanCorrectiveDepositRequest) (*TomanCorrectiveDepositResponse, error)
	mustEmbedUnimplementedTomanAPIServer()
}

// UnimplementedTomanAPIServer must be embedded to have forward compatible implementations.
type UnimplementedTomanAPIServer struct {
}

func (UnimplementedTomanAPIServer) Deposit(context.Context, *TomanDepositRequest) (*TomanDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedTomanAPIServer) Withdraw(context.Context, *TomanWithdrawRequest) (*TomanWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedTomanAPIServer) CorrectiveDeposit(context.Context, *TomanCorrectiveDepositRequest) (*TomanCorrectiveDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CorrectiveDeposit not implemented")
}
func (UnimplementedTomanAPIServer) mustEmbedUnimplementedTomanAPIServer() {}

// UnsafeTomanAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TomanAPIServer will
// result in compilation errors.
type UnsafeTomanAPIServer interface {
	mustEmbedUnimplementedTomanAPIServer()
}

func RegisterTomanAPIServer(s grpc.ServiceRegistrar, srv TomanAPIServer) {
	s.RegisterService(&TomanAPI_ServiceDesc, srv)
}

func _TomanAPI_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TomanDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TomanAPIServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toman.TomanAPI/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TomanAPIServer).Deposit(ctx, req.(*TomanDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TomanAPI_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TomanWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TomanAPIServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toman.TomanAPI/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TomanAPIServer).Withdraw(ctx, req.(*TomanWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TomanAPI_CorrectiveDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TomanCorrectiveDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TomanAPIServer).CorrectiveDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toman.TomanAPI/CorrectiveDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TomanAPIServer).CorrectiveDeposit(ctx, req.(*TomanCorrectiveDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TomanAPI_ServiceDesc is the grpc.ServiceDesc for TomanAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TomanAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "toman.TomanAPI",
	HandlerType: (*TomanAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deposit",
			Handler:    _TomanAPI_Deposit_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _TomanAPI_Withdraw_Handler,
		},
		{
			MethodName: "CorrectiveDeposit",
			Handler:    _TomanAPI_CorrectiveDeposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounting/toman/services.proto",
}
