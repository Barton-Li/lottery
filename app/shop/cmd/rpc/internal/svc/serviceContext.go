package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"lottery/app/shop/cmd/rpc/internal/config"
	"lottery/app/shop/model"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	GoodsModel model.GoodsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		GoodsModel: model.NewGoodsModel(sqlConn, c.Cache),
	}
}
