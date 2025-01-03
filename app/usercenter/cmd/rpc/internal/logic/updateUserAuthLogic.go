package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAuthLogic {
	return &UpdateUserAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserAuthLogic) UpdateUserAuth(in *pb2.UpdateUserAuthReq) (*pb2.UpdateUserAuthResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.UpdateUserAuthResp{}, nil
}
