package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSponsorByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserSponsorByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSponsorByIdLogic {
	return &GetUserSponsorByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserSponsorByIdLogic) GetUserSponsorById(in *pb2.GetUserSponsorByIdReq) (*pb2.GetUserSponsorByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.GetUserSponsorByIdResp{}, nil
}
