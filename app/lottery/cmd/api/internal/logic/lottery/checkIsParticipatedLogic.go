package lottery

import (
	"context"
	"lottery/app/lottery/cmd/rpc/lottery"
	"lottery/common/ctxdata"

	"lottery/app/lottery/cmd/api/internal/svc"
	"lottery/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsParticipatedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckIsParticipatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsParticipatedLogic {
	return &CheckIsParticipatedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckIsParticipatedLogic) CheckIsParticipated(req *types.CheckIsParticipatedReq) (resp *types.CheckIsParticipatedResp, err error) {

	userId := ctxdata.GetUidFromCtx(l.ctx)
	participated, err := l.svcCtx.LotteryRpc.CheckIsParticipated(l.ctx, &lottery.CheckIsParticipatedReq{
		UserId:    userId,
		LotteryId: req.LotteryId,
	})
	if err != nil {
		return nil, err
	}
	return &types.CheckIsParticipatedResp{IsParticipated: participated.IsParticipated}, nil
}
