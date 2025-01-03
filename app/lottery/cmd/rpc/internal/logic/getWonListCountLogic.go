package logic

import (
	"context"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWonListCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWonListCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWonListCountLogic {
	return &GetWonListCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWonListCountLogic) GetWonListCount(in *pb.GetWonListCountReq) (*pb.GetWonListCountResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetWonListCountResp{}, nil
}
