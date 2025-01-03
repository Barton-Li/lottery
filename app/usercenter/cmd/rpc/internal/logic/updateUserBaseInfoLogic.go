package logic

import (
	"context"

	"lottery/app/usercenter/cmd/rpc/internal/svc"
	"lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserBaseInfoLogic {
	return &UpdateUserBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserBaseInfoLogic) UpdateUserBaseInfo(in *pb.UpdateUserBaseInfoReq) (*pb.UpdateUserBaseInfoResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	user.Nickname = in.Nickname
	user.Sex = in.Sex
	user.Info = in.Info
	user.Avatar = in.Avatar
	user.Signature = in.Signature
	user.Longitude = in.Longitude
	user.Latitude = in.Latitude

	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserBaseInfoResp{}, nil
}
