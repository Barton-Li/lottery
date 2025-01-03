package userDynamicComment

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/usercenter"
	"lottery/common/xerr"

	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserWonDynamicCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 累计奖品发布动态用户晒单列表
func NewUserWonDynamicCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserWonDynamicCommentListLogic {
	return &UserWonDynamicCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserWonDynamicCommentListLogic) UserWonDynamicCommentList(req *types.UserWonDynamicCommentReq) (resp *types.UserWonDynamicCommentResp, err error) {
	dynamicList, err := l.svcCtx.UsercenterRpc.SearchUserDynamic(l.ctx, &usercenter.SearchUserDynamicReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to search user dynamic"), "Failed to search user dynamic, err: %v,req%+v", err, req)
	}
	var userDynamicList []types.DynamicInfo
	if len(dynamicList.UserDynamic) > 0 {
		for _, dynamic := range dynamicList.UserDynamic {
			var t types.DynamicInfo
			_ = copier.Copy(&t, dynamic)
			userDynamicList = append(userDynamicList, t)
		}
	}

	return
}
