package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserShopLogic {
	return &DelUserShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserShopLogic) DelUserShop(in *pb2.DelUserShopReq) (*pb2.DelUserShopResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.DelUserShopResp{}, nil
}
