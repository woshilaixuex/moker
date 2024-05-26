package svc

import (
	"github.com/moker/app/usercenter/cmd/api/internal/config"
	"github.com/moker/app/usercenter/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserCenterRPC usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	UserCenterClient := zrpc.MustNewClient(c.UserRPC)
	return &ServiceContext{
		Config:        c,
		UserCenterRPC: usercenter.NewUsercenter(UserCenterClient),
	}
}
