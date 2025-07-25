package config

import "github.com/zeromicro/go-zero/rest"
import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
