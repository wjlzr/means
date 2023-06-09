// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	User struct {
		Id        int64     `db:"id"`         // ID
		Nickname  string    `db:"nickname"`   // 昵称
		Phone     string    `db:"phone"`      // 手机号
		Sex       int64     `db:"sex"`        // 性别 1 男 2 女
		Avatar    string    `db:"avatar"`     // 头像
		Password  string    `db:"password"`   // 密码
		DeletedAt int       `db:"deleted_at"` // 1 已删除
		CreatedAt time.Time `db:"created_at"` // 创建时间
		UpdatedAt time.Time `db:"updated_at"` // 更新时间
	}
)

func newUserModel(conn sqlx.SqlConn) *defaultUserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "`user`",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {

	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
	if session != nil {
		return session.ExecCtx(ctx, query, data.Nickname, data.Phone, data.Sex, data.Avatar, data.Password, data.DeletedAt, data.CreatedAt, data.UpdatedAt)
	}
	return m.conn.ExecCtx(ctx, query, data.Nickname, data.Phone, data.Sex, data.Avatar, data.Password, data.DeletedAt, data.CreatedAt, data.UpdatedAt)
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(ctx context.Context, session sqlx.Session, data *User) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, data.Nickname, data.Phone, data.Sex, data.Avatar, data.Password, data.DeletedAt, data.CreatedAt, data.UpdatedAt, data.Id)
		return err
	}
	_, err := m.conn.ExecCtx(ctx, query, data.Nickname, data.Phone, data.Sex, data.Avatar, data.Password, data.DeletedAt, data.CreatedAt, data.UpdatedAt, data.Id)
	return err
}

func (m *defaultUserModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	if session != nil {
		_, err := session.ExecCtx(ctx, query, id)
		return err
	}
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserModel) tableName() string {
	return m.table
}

func (m *defaultUserModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}
