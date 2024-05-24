// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package server

import (
	"context"
	logic2 "github.com/moker/app/usercenter/cmd/rpc/internal/logic"
	"github.com/moker/app/usercenter/cmd/rpc/internal/svc"
	pb2 "github.com/moker/app/usercenter/cmd/rpc/pb"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
	pb2.UnimplementedUsercenterServer
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UsercenterServer) GetVerCode(ctx context.Context, in *pb2.GetVerCodeReq) (*pb2.GetVerCodeResp, error) {
	l := logic2.NewGetVerCodeLogic(ctx, s.svcCtx)
	return l.GetVerCode(in)
}

func (s *UsercenterServer) Login(ctx context.Context, in *pb2.LoginReq) (*pb2.LoginResp, error) {
	l := logic2.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UsercenterServer) Register(ctx context.Context, in *pb2.RegisterReq) (*pb2.RegisterResp, error) {
	l := logic2.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UsercenterServer) GetUserInfo(ctx context.Context, in *pb2.GetUserInfoReq) (*pb2.GetUserInfoResp, error) {
	l := logic2.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UsercenterServer) GenerateToken(ctx context.Context, in *pb2.GenerateTokenReq) (*pb2.GenerateTokenResp, error) {
	l := logic2.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}
