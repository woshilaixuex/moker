package role

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StudentsModel = (*customStudentsModel)(nil)

type (
	// StudentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStudentsModel.
	StudentsModel interface {
		studentsModel
	}

	customStudentsModel struct {
		*defaultStudentsModel
	}
)

// NewStudentsModel returns a model for the database table.
func NewStudentsModel(conn sqlx.SqlConn, c cache.CacheConf) StudentsModel {
	return &customStudentsModel{
		defaultStudentsModel: newStudentsModel(conn, c),
	}
}
