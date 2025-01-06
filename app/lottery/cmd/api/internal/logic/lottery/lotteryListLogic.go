package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"lottery/app/lottery/cmd/rpc/lottery"

	"lottery/app/lottery/cmd/api/internal/svc"
	"lottery/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotteryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLotteryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryListLogic {
	return &LotteryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LotteryListLogic) LotteryList(req *types.LotteryListReq) (*types.LotteryListResp, error) {
	resp, err := l.svcCtx.LotteryRpc.SearchLottery(l.ctx, &lottery.SearchLotteryReq{
		LastId:     req.LastId,
		Limit:      req.PageSize,
		IsSelected: req.IsSelected,
	})
	if err != nil {
		return nil, err
	}
	var LotteryList []types.Lottery
	if len(resp.Lottery) > 0 {
		for _, v := range resp.Lottery {
			var t types.Lottery
			_ = copier.Copy(&t, v)
			LotteryList = append(LotteryList, t)
		}
	}
	return &types.LotteryListResp{
		List: LotteryList,
	}, nil

}
