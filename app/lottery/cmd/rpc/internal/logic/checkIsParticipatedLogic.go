package logic

import (
	"context"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsParticipatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIsParticipatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsParticipatedLogic {
	return &CheckIsParticipatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIsParticipatedLogic) CheckIsParticipated(in *pb.CheckIsParticipatedReq) (*pb.CheckIsParticipatedResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckIsParticipatedResp{}, nil
}
