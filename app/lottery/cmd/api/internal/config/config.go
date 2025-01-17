package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	jwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf

	UserCenterRpcConf zrpc.RpcClientConf
	LotteryRpcConf    zrpc.RpcClientConf
}
