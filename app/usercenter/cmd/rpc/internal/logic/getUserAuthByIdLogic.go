package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByIdLogic {
	return &GetUserAuthByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByIdLogic) GetUserAuthById(in *pb2.GetUserAuthByIdReq) (*pb2.GetUserAuthByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.GetUserAuthByIdResp{}, nil
}
