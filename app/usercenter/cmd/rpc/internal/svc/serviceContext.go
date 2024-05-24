package svc

import (
	"github.com/moker/app/usercenter/cmd/rpc/internal/config"
	"github.com/moker/common/security"
)

type ServiceContext struct {
	Config config.Config
	Etools *security.EncryptTools
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Etools: security.NewEncryptTools(),
	}
}
