package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"lottery/app/lottery/cmd/rpc/lottery"

	"lottery/app/lottery/cmd/api/internal/svc"
	"lottery/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryWinlist2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLotteryWinlist2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryWinlist2Logic {
	return &GetLotteryWinlist2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLotteryWinlist2Logic) GetLotteryWinlist2(req *types.GetLotteryWinList2Req) (resp *types.GetLotteryWinList2Resp, err error) {
	list, err := l.svcCtx.LotteryRpc.GetWonListByLotteryId(l.ctx, &lottery.GetWonListByLotteryIdReq{

		LotteryId: req.LotteryId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.GetLotteryWinList2Resp{}
	for _, v := range list.List {
		var t types.WonList2
		err = copier.Copy(&t, v)
		if err != nil {
			return nil, err
		}
		resp.List = append(resp.List, &t)
	}

	return
}
