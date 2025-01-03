package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserShopLogic {
	return &UpdateUserShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserShopLogic) UpdateUserShop(in *pb2.UpdateUserShopReq) (*pb2.UpdateUserShopResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.UpdateUserShopResp{}, nil
}
