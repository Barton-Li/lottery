package lottery

import (
	"net/http"

	"lottery/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lottery/app/lottery/cmd/api/internal/logic/lottery"
	"lottery/app/lottery/cmd/api/internal/svc"
	"lottery/app/lottery/cmd/api/internal/types"
)

func GetLotteryListByUserIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetLotteryListByUserIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		validateErr := translator.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(r, w, validateErr)
			return
		}

		l := lottery.NewGetLotteryListByUserIdLogic(r.Context(), svcCtx)
		resp, err := l.GetLotteryListByUserId(&req)

		result.HttpResult(r, w, resp, err)
	}
}
