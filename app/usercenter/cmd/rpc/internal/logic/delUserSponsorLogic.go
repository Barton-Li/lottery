package logic

import (
	"context"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"
	"lottery/common/xerr"



	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserSponsorLogic {
	return &DelUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserSponsorLogic) DelUserSponsor(in *pb2.DelUserSponsorReq) (*pb2.DelUserSponsorResp, error) {
	err := l.svcCtx.UserSponsorModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "删除赞助商失败:%v,rpc:%+v", err, in)
	}

	return &pb2.DelUserSponsorResp{}, nil
}
