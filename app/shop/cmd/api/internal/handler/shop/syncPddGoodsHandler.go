package shop

import (
	"net/http"

	"lottery/common/result"

	"lottery/app/shop/cmd/api/internal/logic/shop"
	"lottery/app/shop/cmd/api/internal/svc"
)

func SyncPddGoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := shop.NewSyncPddGoodsLogic(r.Context(), svcCtx)
		err := l.SyncPddGoods()

		result.HttpResult(r, w, nil, err)
	}
}
