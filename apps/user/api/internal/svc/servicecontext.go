package svc

import (
	"gitee.com/lfzizr/easy-chat/apps/user/api/internal/config"

	"gitee.com/lfzizr/easy-chat/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	UserClient userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserClient: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
