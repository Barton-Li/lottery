package userDynamicComment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lottery/app/usercenter/cmd/api/internal/logic/userDynamicComment"
	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"
)

// 累计奖品发布动态用户晒单列表
func UserWonDynamicCommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserWonDynamicCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userDynamicComment.NewUserWonDynamicCommentListLogic(r.Context(), svcCtx)
		resp, err := l.UserWonDynamicCommentList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
