package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"lottery/app/lottery/cmd/rpc/lottery"
	"lottery/common/ctxdata"

	"lottery/app/lottery/cmd/api/internal/svc"
	"lottery/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotteryListAfterLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLotteryListAfterLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryListAfterLoginLogic {
	return &LotteryListAfterLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LotteryListAfterLoginLogic) LotteryListAfterLogin(req *types.LotteryListReq) (resp *types.LotteryListResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	res, err := l.svcCtx.LotteryRpc.GetLotteryListAfterLogin(l.ctx, &lottery.GetLotteryListAfterLoginReq{
		LastId:     req.LastId,
		Size:       req.PageSize,
		IsSelected: req.IsSelected,
		UserId:     userId,
	})
	if err != nil {
		return nil, err
	}
	var lotteryList []types.Lottery
	if len(res.List) > 0 {
		for _, v := range res.List {
			var t types.Lottery
			_ = copier.Copy(&t, v)
			lotteryList = append(lotteryList, t)
		}
	}

	return &types.LotteryListResp{
		List: lotteryList,
	}, nil
}
