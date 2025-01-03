package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserInfoLogic {
	return &DelUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserInfoLogic) DelUserInfo(in *pb2.DelUserReq) (*pb2.DelUserResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.DelUserResp{}, nil
}
