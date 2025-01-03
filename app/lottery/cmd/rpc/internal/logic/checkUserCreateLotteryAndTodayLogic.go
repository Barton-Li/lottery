package logic

import (
	"context"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserCreateLotteryAndTodayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserCreateLotteryAndTodayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserCreateLotteryAndTodayLogic {
	return &CheckUserCreateLotteryAndTodayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserCreateLotteryAndTodayLogic) CheckUserCreateLotteryAndToday(in *pb.CheckUserCreateLotteryAndTodayReq) (*pb.CheckUserCreateLotteryAndTodayResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckUserCreateLotteryAndTodayResp{}, nil
}
