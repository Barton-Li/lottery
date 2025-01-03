package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserAuthLogic {
	return &AddUserAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户授权表----------------------
func (l *AddUserAuthLogic) AddUserAuth(in *pb2.AddUserAuthReq) (*pb2.AddUserAuthResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.AddUserAuthResp{}, nil
}
