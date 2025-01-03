package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"lottery/app/usercenter/cmd/rpc/internal/config"
	"lottery/app/usercenter/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config           config.Config
	RedisClient      *redis.Redis
	UserModel        model.UserInfoModel
	UserAuthModel    model.UserAuthModel
	UserAddressModel model.UserAddressModel
	UserSponsorModel model.UserSponsorModel
	UserContactModel model.UserContactModel
	UserDynamicModel model.UserDynamicModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:           c,
		UserModel:        model.NewUserInfoModel(sqlConn, c.Cache),
		UserAuthModel:    model.NewUserAuthModel(sqlConn, c.Cache),
		UserAddressModel: model.NewUserAddressModel(sqlConn, c.Cache),
		UserSponsorModel: model.NewUserSponsorModel(sqlConn, c.Cache),
		UserContactModel: model.NewUserContactModel(sqlConn, c.Cache),
		UserDynamicModel: model.NewUserDynamicModel(sqlConn, c.Cache),
		RedisClient: redis.MustNewRedis(c.Redis.RedisConf, func(r *redis.Redis) {
			r.Type = c.Redis.RedisConf.Type
			r.Pass = c.Redis.RedisConf.Pass
		}),
	}
}
