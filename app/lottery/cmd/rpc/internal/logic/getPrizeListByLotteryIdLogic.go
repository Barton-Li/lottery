package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/common/xerr"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeListByLotteryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizeListByLotteryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeListByLotteryIdLogic {
	return &GetPrizeListByLotteryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizeListByLotteryIdLogic) GetPrizeListByLotteryId(in *pb.GetPrizeListByLotteryIdReq) (*pb.GetPrizeListByLotteryIdResp, error) {
	lotteryId, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, in.LotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindByLotteryId failed, lotteryId: %d, err: %v", in.LotteryId, err)
	}
	prizes := make([]*pb.Prize, 0)
	for _, v := range lotteryId {
		pbPrize := new(pb.Prize)
		err := copier.Copy(pbPrize, v)
		if err != nil {
			return nil, err
		}
		prizes = append(prizes, pbPrize)
	}

	return &pb.GetPrizeListByLotteryIdResp{
		Prizes: prizes,
	}, nil
}
