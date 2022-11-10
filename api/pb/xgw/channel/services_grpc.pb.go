// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: xgw/channel/services.proto

package channel

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

// ChannelApiClient is the client API for ChannelApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelApiClient interface {
	CreateChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*CreateChannelResponse, error)
	UpdateChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*UpdateChannelResponse, error)
	DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error)
	GetAllChannels(ctx context.Context, in *GetAllChannelsRequest, opts ...grpc.CallOption) (*GetAllChannelsResponse, error)
	GetChannelsByTreasury(ctx context.Context, in *GetChannelsByTreasuryRequest, opts ...grpc.CallOption) (*GetChannelsByTreasuryResponse, error)
	GetProviderCoinInfo(ctx context.Context, in *GetProviderCoinInfoRequest, opts ...grpc.CallOption) (*GetProviderCoinInfoResponse, error)
}

type channelApiClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelApiClient(cc grpc.ClientConnInterface) ChannelApiClient {
	return &channelApiClient{cc}
}

func (c *channelApiClient) CreateChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*CreateChannelResponse, error) {
	out := new(CreateChannelResponse)
	err := c.cc.Invoke(ctx, "/channel.ChannelApi/CreateChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelApiClient) UpdateChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*UpdateChannelResponse, error) {
	out := new(UpdateChannelResponse)
	err := c.cc.Invoke(ctx, "/channel.ChannelApi/UpdateChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelApiClient) DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error) {
	out := new(DeleteChannelResponse)
	err := c.cc.Invoke(ctx, "/channel.ChannelApi/DeleteChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelApiClient) GetAllChannels(ctx context.Context, in *GetAllChannelsRequest, opts ...grpc.CallOption) (*GetAllChannelsResponse, error) {
	out := new(GetAllChannelsResponse)
	err := c.cc.Invoke(ctx, "/channel.ChannelApi/GetAllChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelApiClient) GetChannelsByTreasury(ctx context.Context, in *GetChannelsByTreasuryRequest, opts ...grpc.CallOption) (*GetChannelsByTreasuryResponse, error) {
	out := new(GetChannelsByTreasuryResponse)
	err := c.cc.Invoke(ctx, "/channel.ChannelApi/GetChannelsByTreasury", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelApiClient) GetProviderCoinInfo(ctx context.Context, in *GetProviderCoinInfoRequest, opts ...grpc.CallOption) (*GetProviderCoinInfoResponse, error) {
	out := new(GetProviderCoinInfoResponse)
	err := c.cc.Invoke(ctx, "/channel.ChannelApi/GetProviderCoinInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelApiServer is the server API for ChannelApi service.
// All implementations must embed UnimplementedChannelApiServer
// for forward compatibility
type ChannelApiServer interface {
	CreateChannel(context.Context, *Channel) (*CreateChannelResponse, error)
	UpdateChannel(context.Context, *Channel) (*UpdateChannelResponse, error)
	DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error)
	GetAllChannels(context.Context, *GetAllChannelsRequest) (*GetAllChannelsResponse, error)
	GetChannelsByTreasury(context.Context, *GetChannelsByTreasuryRequest) (*GetChannelsByTreasuryResponse, error)
	GetProviderCoinInfo(context.Context, *GetProviderCoinInfoRequest) (*GetProviderCoinInfoResponse, error)
	mustEmbedUnimplementedChannelApiServer()
}

// UnimplementedChannelApiServer must be embedded to have forward compatible implementations.
type UnimplementedChannelApiServer struct {
}

func (UnimplementedChannelApiServer) CreateChannel(context.Context, *Channel) (*CreateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedChannelApiServer) UpdateChannel(context.Context, *Channel) (*UpdateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChannel not implemented")
}
func (UnimplementedChannelApiServer) DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChannel not implemented")
}
func (UnimplementedChannelApiServer) GetAllChannels(context.Context, *GetAllChannelsRequest) (*GetAllChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllChannels not implemented")
}
func (UnimplementedChannelApiServer) GetChannelsByTreasury(context.Context, *GetChannelsByTreasuryRequest) (*GetChannelsByTreasuryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannelsByTreasury not implemented")
}
func (UnimplementedChannelApiServer) GetProviderCoinInfo(context.Context, *GetProviderCoinInfoRequest) (*GetProviderCoinInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProviderCoinInfo not implemented")
}
func (UnimplementedChannelApiServer) mustEmbedUnimplementedChannelApiServer() {}

// UnsafeChannelApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelApiServer will
// result in compilation errors.
type UnsafeChannelApiServer interface {
	mustEmbedUnimplementedChannelApiServer()
}

func RegisterChannelApiServer(s grpc.ServiceRegistrar, srv ChannelApiServer) {
	s.RegisterService(&ChannelApi_ServiceDesc, srv)
}

func _ChannelApi_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Channel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelApiServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.ChannelApi/CreateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelApiServer).CreateChannel(ctx, req.(*Channel))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelApi_UpdateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Channel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelApiServer).UpdateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.ChannelApi/UpdateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelApiServer).UpdateChannel(ctx, req.(*Channel))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelApi_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelApiServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.ChannelApi/DeleteChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelApiServer).DeleteChannel(ctx, req.(*DeleteChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelApi_GetAllChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelApiServer).GetAllChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.ChannelApi/GetAllChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelApiServer).GetAllChannels(ctx, req.(*GetAllChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelApi_GetChannelsByTreasury_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelsByTreasuryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelApiServer).GetChannelsByTreasury(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.ChannelApi/GetChannelsByTreasury",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelApiServer).GetChannelsByTreasury(ctx, req.(*GetChannelsByTreasuryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelApi_GetProviderCoinInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProviderCoinInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelApiServer).GetProviderCoinInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/channel.ChannelApi/GetProviderCoinInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelApiServer).GetProviderCoinInfo(ctx, req.(*GetProviderCoinInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelApi_ServiceDesc is the grpc.ServiceDesc for ChannelApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "channel.ChannelApi",
	HandlerType: (*ChannelApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChannel",
			Handler:    _ChannelApi_CreateChannel_Handler,
		},
		{
			MethodName: "UpdateChannel",
			Handler:    _ChannelApi_UpdateChannel_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _ChannelApi_DeleteChannel_Handler,
		},
		{
			MethodName: "GetAllChannels",
			Handler:    _ChannelApi_GetAllChannels_Handler,
		},
		{
			MethodName: "GetChannelsByTreasury",
			Handler:    _ChannelApi_GetChannelsByTreasury_Handler,
		},
		{
			MethodName: "GetProviderCoinInfo",
			Handler:    _ChannelApi_GetProviderCoinInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xgw/channel/services.proto",
}