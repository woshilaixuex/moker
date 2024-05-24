// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package usercenter

import (
	"context"
	pb2 "github.com/moker/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GenerateTokenReq  = pb2.GenerateTokenReq
	GenerateTokenResp = pb2.GenerateTokenResp
	GetUserInfoReq    = pb2.GetUserInfoReq
	GetUserInfoResp   = pb2.GetUserInfoResp
	GetVerCodeReq     = pb2.GetVerCodeReq
	GetVerCodeResp    = pb2.GetVerCodeResp
	LoginReq          = pb2.LoginReq
	LoginResp         = pb2.LoginResp
	RegisterReq       = pb2.RegisterReq
	RegisterResp      = pb2.RegisterResp
	UserAuth          = pb2.UserAuth
	UserInform        = pb2.UserInform

	Usercenter interface {
		GetVerCode(ctx context.Context, in *GetVerCodeReq, opts ...grpc.CallOption) (*GetVerCodeResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
	}

	defaultUsercenter struct {
		cli zrpc.Client
	}
)

func NewUsercenter(cli zrpc.Client) Usercenter {
	return &defaultUsercenter{
		cli: cli,
	}
}

func (m *defaultUsercenter) GetVerCode(ctx context.Context, in *GetVerCodeReq, opts ...grpc.CallOption) (*GetVerCodeResp, error) {
	client := pb2.NewUsercenterClient(m.cli.Conn())
	return client.GetVerCode(ctx, in, opts...)
}

func (m *defaultUsercenter) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb2.NewUsercenterClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUsercenter) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb2.NewUsercenterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb2.NewUsercenterClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUsercenter) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb2.NewUsercenterClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}
