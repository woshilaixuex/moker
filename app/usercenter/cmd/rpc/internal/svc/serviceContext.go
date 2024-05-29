package svc

import (
	"github.com/moker/app/usercenter/cmd/rpc/internal/config"
	"github.com/moker/app/usercenter/model/role"
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
	StudentsModel role.StudentsModel
	TeahcersModel role.TeachersModel
	Etools        *security.EncryptTools
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     user.NewUserModel(sqlConn, c.Cache),
		UserSaltModel: user.NewUserSaltModel(sqlConn, c.Cache),
		StudentsModel: role.NewStudentsModel(sqlConn, c.Cache),
		TeahcersModel: role.NewTeachersModel(sqlConn, c.Cache),
		Etools:        security.NewEncryptTools(),
	}
}
