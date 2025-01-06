package logic

import (
	"context"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserCreateLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserCreateLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserCreateLotteryLogic {
	return &CheckUserCreateLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserCreateLotteryLogic) CheckUserCreateLottery(in *pb.CheckUserCreateLotteryReq) (*pb.CheckUserCreateLotteryResp, error) {
	id, err := l.svcCtx.LotteryModel.GetLotteryIdByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	if id == nil {
		return &pb.CheckUserCreateLotteryResp{IsCreate: 0}, nil
	}
	return &pb.CheckUserCreateLotteryResp{IsCreate: 1}, nil
}
