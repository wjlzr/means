package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"means/app/user/model"
	"means/app/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,

		UserModel: model.NewUserModel(sqlConn, c.Cache),
	}
}
