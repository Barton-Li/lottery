package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserDynamicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserDynamicLogic {
	return &UpdateUserDynamicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserDynamicLogic) UpdateUserDynamic(in *pb2.UpdateUserDynamicReq) (*pb2.UpdateUserDynamicResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.UpdateUserDynamicResp{}, nil
}
