// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheMeansUserIdPrefix = "cache:means:user:id:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id        int64          `db:"id"`         // ID
		Nickname  string         `db:"nickname"`   // 昵称
		Phone     string         `db:"phone"`      // 手机号
		Sex       int64          `db:"sex"`        // 性别 1 男 2 女
		Avatar    sql.NullString `db:"avatar"`     // 头像
		Password  string         `db:"password"`   // 密码
		DeletedAt sql.NullInt64  `db:"deleted_at"` // 1 已删除
		CreatedAt sql.NullTime   `db:"created_at"` // 创建时间
		UpdatedAt sql.NullTime   `db:"updated_at"` // 更新时间
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	meansUserIdKey := fmt.Sprintf("%s%v", cacheMeansUserIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Nickname, data.Phone, data.Sex, data.Avatar, data.Password, data.DeletedAt, data.CreatedAt, data.UpdatedAt)
	}, meansUserIdKey)
	return ret, err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	meansUserIdKey := fmt.Sprintf("%s%v", cacheMeansUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, meansUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
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

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	meansUserIdKey := fmt.Sprintf("%s%v", cacheMeansUserIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Nickname, data.Phone, data.Sex, data.Avatar, data.Password, data.DeletedAt, data.CreatedAt, data.UpdatedAt, data.Id)
	}, meansUserIdKey)
	return err
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	meansUserIdKey := fmt.Sprintf("%s%v", cacheMeansUserIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, meansUserIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMeansUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
