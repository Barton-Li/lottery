package user

import (
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lottery/app/usercenter/cmd/api/internal/logic/user"
	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"
)

func SetAdminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetAdminReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		validateErr := translator.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(r, w, validateErr)
			return
		}

		l := user.NewSetAdminLogic(r.Context(), svcCtx)
		resp, err := l.SetAdmin(&req)

		result.HttpResult(r, w, resp, err)
	}
}
