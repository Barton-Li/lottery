package shop

import (
	"lottery/app/shop/cmd/api/internal/handler/translator"
	"net/http"

	"lottery/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lottery/app/shop/cmd/api/internal/logic/shop"
	"lottery/app/shop/cmd/api/internal/svc"
	"lottery/app/shop/cmd/api/internal/types"
)

func GetGoodsByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		validateErr := translator.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(r, w, validateErr)
			return
		}

		l := shop.NewGetGoodsByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetGoodsById(&req)

		result.HttpResult(r, w, resp, err)
	}
}
