package logic

import (
	"context"

	"lottery/app/usercenter/cmd/rpc/internal/svc"
	"lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIsAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsAdminLogic {
	return &CheckIsAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIsAdminLogic) CheckIsAdmin(in *pb.CheckIsAdminReq) (*pb.CheckIsAdminResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckIsAdminResp{}, nil
}
