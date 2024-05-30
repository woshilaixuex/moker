package svc

import (
	"github.com/moker/app/course/cmd/api/internal/config"
	"github.com/moker/app/course/cmd/rpc/getgreeter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	CourseCenterRPC getgreeter.GetGreeter
}

func NewServiceContext(c config.Config) *ServiceContext {
	CourseCenterClient := zrpc.MustNewClient(c.CourseRpc)
	return &ServiceContext{
		Config:          c,
		CourseCenterRPC: getgreeter.NewGetGreeter(CourseCenterClient),
	}
}
