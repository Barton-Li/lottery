package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserDynamicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserDynamicLogic {
	return &DelUserDynamicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserDynamicLogic) DelUserDynamic(in *pb2.DelUserDynamicReq) (*pb2.DelUserDynamicResp, error) {

	err := l.svcCtx.UserDynamicModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb2.DelUserDynamicResp{}, nil
}
