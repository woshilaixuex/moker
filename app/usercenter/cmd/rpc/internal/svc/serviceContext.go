package svc

import (
	"github.com/moker/app/usercenter/cmd/rpc/internal/config"
	"github.com/moker/app/usercenter/model/user"
	"github.com/moker/common/security"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	RedisClient   *redis.Redis
	UserModel     user.UserModel
	UserSaltModel user.UserSaltModel
	Etools        *security.EncryptTools
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		RedisClient:   redis.MustNewRedis(c.RedisConfig),
		UserModel:     user.NewUserModel(sqlConn, c.Cache),
		UserSaltModel: user.NewUserSaltModel(sqlConn, c.Cache),
		Etools:        security.NewEncryptTools(),
	}
}
