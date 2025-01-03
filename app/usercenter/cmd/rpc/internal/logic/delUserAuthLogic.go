package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserAuthLogic {
	return &DelUserAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserAuthLogic) DelUserAuth(in *pb2.DelUserAuthReq) (*pb2.DelUserAuthResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.DelUserAuthResp{}, nil
}
