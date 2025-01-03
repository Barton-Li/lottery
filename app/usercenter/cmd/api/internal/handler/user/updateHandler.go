package user

import (
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lottery/app/usercenter/cmd/api/internal/logic/user"
	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		validateErr := translator.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(r, w, validateErr)
			return
		}

		l := user.NewUpdateLogic(r.Context(), svcCtx)
		resp, err := l.Update(&req)

		result.HttpResult(r, w, resp, err)
	}
}
