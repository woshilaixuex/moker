package user

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserSaltModel = (*customUserSaltModel)(nil)

type (
	// UserSaltModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserSaltModel.
	UserSaltModel interface {
		userSaltModel
	}

	customUserSaltModel struct {
		*defaultUserSaltModel
	}
)

// NewUserSaltModel returns a model for the database table.
func NewUserSaltModel(conn sqlx.SqlConn, c cache.CacheConf) UserSaltModel {
	return &customUserSaltModel{
		defaultUserSaltModel: newUserSaltModel(conn, c),
	}
}
