package logic

import (
	"context"

	"lottery/app/usercenter/cmd/rpc/internal/svc"
	"lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetAdminLogic {
	return &SetAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetAdminLogic) SetAdmin(in *pb.SetAdminReq) (*pb.SetAdminResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SetAdminResp{}, nil
}
