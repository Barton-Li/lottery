package logic

import (
	"context"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPrizeLogic {
	return &AddPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -------------------------奖品表-----------------------------------------------
func (l *AddPrizeLogic) AddPrize(in *pb.AddPrizeReq) (*pb.AddPrizeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddPrizeResp{}, nil
}
