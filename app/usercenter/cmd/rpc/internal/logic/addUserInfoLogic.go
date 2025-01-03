package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserInfoLogic {
	return &AddUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------userInfo-----------------------
func (l *AddUserInfoLogic) AddUserInfo(in *pb2.AddUserReq) (*pb2.AddUserResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.AddUserResp{}, nil
}
