// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: cgw/ethereum/services.proto

package ethereum

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

// EthereumApiClient is the client API for EthereumApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EthereumApiClient interface {
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error)
	Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	TransactionDetail(ctx context.Context, in *TransactionDetailRequest, opts ...grpc.CallOption) (*TransactionDetailResponse, error)
	GetWithdrawFee(ctx context.Context, in *GetWithdrawFeeRequest, opts ...grpc.CallOption) (*GetWithdrawFeeResponse, error)
}

type ethereumApiClient struct {
	cc grpc.ClientConnInterface
}

func NewEthereumApiClient(cc grpc.ClientConnInterface) EthereumApiClient {
	return &ethereumApiClient{cc}
}

func (c *ethereumApiClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error) {
	out := new(WithdrawResponse)
	err := c.cc.Invoke(ctx, "/ethereum.EthereumApi/Withdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ethereumApiClient) Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/ethereum.EthereumApi/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ethereumApiClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/ethereum.EthereumApi/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ethereumApiClient) TransactionDetail(ctx context.Context, in *TransactionDetailRequest, opts ...grpc.CallOption) (*TransactionDetailResponse, error) {
	out := new(TransactionDetailResponse)
	err := c.cc.Invoke(ctx, "/ethereum.EthereumApi/TransactionDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ethereumApiClient) GetWithdrawFee(ctx context.Context, in *GetWithdrawFeeRequest, opts ...grpc.CallOption) (*GetWithdrawFeeResponse, error) {
	out := new(GetWithdrawFeeResponse)
	err := c.cc.Invoke(ctx, "/ethereum.EthereumApi/GetWithdrawFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EthereumApiServer is the server API for EthereumApi service.
// All implementations must embed UnimplementedEthereumApiServer
// for forward compatibility
type EthereumApiServer interface {
	Withdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error)
	Balance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	TransactionDetail(context.Context, *TransactionDetailRequest) (*TransactionDetailResponse, error)
	GetWithdrawFee(context.Context, *GetWithdrawFeeRequest) (*GetWithdrawFeeResponse, error)
	mustEmbedUnimplementedEthereumApiServer()
}

// UnimplementedEthereumApiServer must be embedded to have forward compatible implementations.
type UnimplementedEthereumApiServer struct {
}

func (UnimplementedEthereumApiServer) Withdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedEthereumApiServer) Balance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedEthereumApiServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedEthereumApiServer) TransactionDetail(context.Context, *TransactionDetailRequest) (*TransactionDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactionDetail not implemented")
}
func (UnimplementedEthereumApiServer) GetWithdrawFee(context.Context, *GetWithdrawFeeRequest) (*GetWithdrawFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdrawFee not implemented")
}
func (UnimplementedEthereumApiServer) mustEmbedUnimplementedEthereumApiServer() {}

// UnsafeEthereumApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EthereumApiServer will
// result in compilation errors.
type UnsafeEthereumApiServer interface {
	mustEmbedUnimplementedEthereumApiServer()
}

func RegisterEthereumApiServer(s grpc.ServiceRegistrar, srv EthereumApiServer) {
	s.RegisterService(&EthereumApi_ServiceDesc, srv)
}

func _EthereumApi_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthereumApiServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.EthereumApi/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthereumApiServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EthereumApi_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthereumApiServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.EthereumApi/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthereumApiServer).Balance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EthereumApi_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthereumApiServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.EthereumApi/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthereumApiServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EthereumApi_TransactionDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthereumApiServer).TransactionDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.EthereumApi/TransactionDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthereumApiServer).TransactionDetail(ctx, req.(*TransactionDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EthereumApi_GetWithdrawFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthereumApiServer).GetWithdrawFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.EthereumApi/GetWithdrawFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthereumApiServer).GetWithdrawFee(ctx, req.(*GetWithdrawFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EthereumApi_ServiceDesc is the grpc.ServiceDesc for EthereumApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EthereumApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.EthereumApi",
	HandlerType: (*EthereumApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Withdraw",
			Handler:    _EthereumApi_Withdraw_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _EthereumApi_Balance_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _EthereumApi_CreateAccount_Handler,
		},
		{
			MethodName: "TransactionDetail",
			Handler:    _EthereumApi_TransactionDetail_Handler,
		},
		{
			MethodName: "GetWithdrawFee",
			Handler:    _EthereumApi_GetWithdrawFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cgw/ethereum/services.proto",
}