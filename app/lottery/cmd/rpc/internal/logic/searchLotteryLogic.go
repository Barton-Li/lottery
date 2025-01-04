package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"lottery/app/lottery/model"
	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLotteryLogic {
	return &SearchLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLotteryLogic) SearchLottery(in *pb.SearchLotteryReq) (*pb.SearchLotteryResp, error) {
	if in.LastId == 0 {
		id, err := l.svcCtx.LotteryModel.GetLastId(l.ctx)
		if err != nil {
			return nil, err
		}
		in.LastId = id + 1
	}
	list, err := l.svcCtx.LotteryModel.LotteryList(l.ctx, in.Limit, in.IsSelected, in.LastId)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	var resp []*pb.Lottery
	if len(list) > 0 {
		for _, lottery := range list {
			var pblottery pb.Lottery
			_ = copier.Copy(&pblottery, lottery)
			pblottery.PubilshTime = lottery.PublishTime.Time.Unix()
			pblottery.AwardDeadline = lottery.AwardDeadline.Unix()
			pblottery.Announcetype = lottery.AnnounceType
			pblottery.AnnounceTime = lottery.AnnounceTime.Unix()
			pblottery.IsAnnounced = lottery.IsAnnounced
			resp = append(resp, &pblottery)
		}
		return &pb.SearchLotteryResp{
			Lottery: resp,
		}, nil
	}
	return &pb.SearchLotteryResp{}, nil
}
