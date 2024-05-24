package role

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TeachersModel = (*customTeachersModel)(nil)

type (
	// TeachersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTeachersModel.
	TeachersModel interface {
		teachersModel
	}

	customTeachersModel struct {
		*defaultTeachersModel
	}
)

// NewTeachersModel returns a model for the database table.
func NewTeachersModel(conn sqlx.SqlConn) TeachersModel {
	return &customTeachersModel{
		defaultTeachersModel: newTeachersModel(conn),
	}
}
