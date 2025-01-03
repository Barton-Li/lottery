package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserAuthLogic {
	return &SearchUserAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserAuthLogic) SearchUserAuth(in *pb2.SearchUserAuthReq) (*pb2.SearchUserAuthResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.SearchUserAuthResp{}, nil
}
