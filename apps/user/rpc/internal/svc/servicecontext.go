package svc

import (
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/internal/config"
	"gitee.com/lfzizr/easy-chat/apps/user/models"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	UserModel  models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel:  models.NewUsersModel(sqlConn, c.Cache),
	}
}
