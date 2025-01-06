package lottery

import (
	"context"
	"lottery/app/lottery/cmd/rpc/lottery"
	"lottery/common/ctxdata"

	"lottery/app/lottery/cmd/api/internal/svc"
	"lottery/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetLotteryIsSelectedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetLotteryIsSelectedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetLotteryIsSelectedLogic {
	return &SetLotteryIsSelectedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetLotteryIsSelectedLogic) SetLotteryIsSelected(req *types.SetLotteryIsSelectedReq) (resp *types.SetLotteryIsSelectedResp, err error) {

	userId := ctxdata.GetUidFromCtx(l.ctx)
	selectedLottery, err := l.svcCtx.LotteryRpc.SetIsSelectedLottery(l.ctx, &lottery.SetIsSelectedLotteryReq{
		Id:     req.Id,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.SetLotteryIsSelectedResp{

		IsSelected: selectedLottery.IsSelected,
	}, nil
}
