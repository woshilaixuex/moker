// Code generated by goctl. DO NOT EDIT.

package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/moker/common/globalkey"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheMokerUsercenterUserIdPrefix = "cache:mokerUsercenter:user:id:"
	cacheMokerUsercenterUserAccountPrefix = "cache:mokerUsercenter:user:account:"
)

type (
	userModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *User)  (sql.Result, error)
		FindOne(ctx context.Context, account string) (*User, error)
		FindOneById(ctx context.Context, id int64) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id         int64     `db:"id"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		DeleteTime time.Time `db:"delete_time"`
		DelState   int64     `db:"del_state"`
		Version    int64     `db:"version"` // 版本号
		Email      string    `db:"email"`
		Password   string    `db:"password"`
		Account    string    `db:"account"`
		Username   string    `db:"username"`
		Sex        int64     `db:"sex"` // 性别  0:男  1:女  3:未知
		Avatar     string    `db:"avatar"`
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	mokerUsercenterUserIdKey := fmt.Sprintf("%s%v", cacheMokerUsercenterUserIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, mokerUsercenterUserIdKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, account string) (*User, error) {
	mokerUsercenterUserAccountKey := fmt.Sprintf("%s%v", cacheMokerUsercenterUserAccountPrefix, account)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, mokerUsercenterUserAccountKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `account` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, account)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func(m *defaultUserModel)FindOneById(ctx context.Context, id int64) (*User, error){
	mokerUsercenterUserIdKey := fmt.Sprintf("%s%v", cacheMokerUsercenterUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, mokerUsercenterUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context,session sqlx.Session, data *User) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo
	mokerUsercenterUserAccountKey := fmt.Sprintf("%s%v", cacheMokerUsercenterUserAccountPrefix, data.Account)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		if session != nil{
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Email, data.Password, data.Account, data.Username, data.Sex, data.Avatar)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Email, data.Password, data.Account, data.Username, data.Sex, data.Avatar)
	}, mokerUsercenterUserAccountKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	mokerUsercenterUserIdKey := fmt.Sprintf("%s%v", cacheMokerUsercenterUserIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Email, data.Password, data.Account, data.Username, data.Sex, data.Avatar, data.Id)
	}, mokerUsercenterUserIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMokerUsercenterUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
func (m *defaultUserModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}