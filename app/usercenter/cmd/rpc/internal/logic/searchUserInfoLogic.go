package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"



	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserInfoLogic {
	return &SearchUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserInfoLogic) SearchUserInfo(in *pb2.SearchUserReq) (*pb2.SearchUserResp, error) {
	// todo: add your logic here and delete this line

	return &pb2.SearchUserResp{}, nil
}
