package userDynamic

import (
	"context"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/pb"
	"lottery/common/xerr"

	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDynamicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除用户动态
func NewDeleteDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDynamicLogic {
	return &DeleteDynamicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDynamicLogic) DeleteDynamic(req *types.DeleteDynamicReq) (resp *types.DeleteDynamicResp, err error) {
	_, err = l.svcCtx.UsercenterRpc.DelUserDynamic(l.ctx, &pb.DelUserDynamicReq{
		Id:     req.Id,
		UserId: req.UserId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to delete dynamic"), "%v,%+v", err, req)
	}
	return
}
