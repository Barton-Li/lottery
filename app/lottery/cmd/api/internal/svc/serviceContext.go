package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lottery/app/lottery/cmd/api/internal/config"
	"lottery/app/lottery/cmd/rpc/lottery"
	"lottery/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config config.Config
	UsercenterRpc usercenter.Usercenter
	LotteryRpc    lottery.LotteryZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		LotteryRpc:    lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
	}
}
