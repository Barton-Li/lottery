package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserShopLogic {
	return &AddUserShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------userShop-----------------------
func (l *AddUserShopLogic) AddUserShop(in *pb2.AddUserShopReq) (*pb2.AddUserShopResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.AddUserShopResp{}, nil
}
