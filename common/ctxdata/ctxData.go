package ctxdata

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	// 从上下文中获取 JwtUserId 值，并尝试将其断言为 json.Number 类型
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		// 将 json.Number 类型的值转换为 int64 类型
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			// 如果转换失败，记录错误日志
			logx.WithContext(ctx).Errorf("从上下文中获取 uid 失败, 错误:%v", err)
		}
	}
	return uid
}
