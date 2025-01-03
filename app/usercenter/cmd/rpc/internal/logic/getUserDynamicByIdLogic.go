package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDynamicByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserDynamicByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDynamicByIdLogic {
	return &GetUserDynamicByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserDynamicByIdLogic) GetUserDynamicById(in *pb2.GetUserDynamicByIdReq) (*pb2.GetUserDynamicByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.GetUserDynamicByIdResp{}, nil
}
