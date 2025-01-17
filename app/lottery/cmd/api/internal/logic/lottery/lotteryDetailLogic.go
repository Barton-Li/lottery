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

type LotteryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLotteryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryDetailLogic {
	return &LotteryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LotteryDetailLogic) LotteryDetail(req *types.LotteryDetailReq) (resp *types.LotteryDetailResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	res, err := l.svcCtx.LotteryRpc.LotteryDetail(l.ctx, &lottery.LotteryDetailReq{
		Id:     req.Id,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.LotteryDetailResp)
	// todo 返回成功，但是json反序列化提示error
	_ = copier.Copy(resp, res)
	_ = copier.Copy(resp, res.Lottery)
	resp.IsParticipated = res.IsParticipated

	// 根据获取到的lottery信息中的sponsorId获取赞助商信息
	res2, err := l.svcCtx.LotteryRpc.LotterySponsor(l.ctx, &lottery.LotterySponsorReq{
		SponsorId: res.Lottery.SponsorId,
	})

	//

	if err != nil {
		return nil, err
	}
	resp.Sponsor = new(types.LotterySponsor)
	_ = copier.Copy(resp.Sponsor, res2)

	return
}
