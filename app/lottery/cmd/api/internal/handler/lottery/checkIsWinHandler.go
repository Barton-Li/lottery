package lottery

import (
	"net/http"

	"lottery/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lottery/app/lottery/cmd/api/internal/logic/lottery"
	"lottery/app/lottery/cmd/api/internal/svc"
	"lottery/app/lottery/cmd/api/internal/types"
)

func CheckIsWinHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckIsWinReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		validateErr := translator.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(r, w, validateErr)
			return
		}

		l := lottery.NewCheckIsWinLogic(r.Context(), svcCtx)
		resp, err := l.CheckIsWin(&req)

		result.HttpResult(r, w, resp, err)
	}
}
