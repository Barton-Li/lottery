package userSponsor

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lottery/app/usercenter/cmd/api/internal/logic/userSponsor"
	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"
)

// 添加 抽奖发起人（赞助商）
func AddSponsorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateSponsorReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userSponsor.NewAddSponsorLogic(r.Context(), svcCtx)
		resp, err := l.AddSponsor(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
