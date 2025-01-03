package logic

import (
	"context"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"
	"lottery/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserSponsorLogic {
	return &UpdateUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserSponsorLogic) UpdateUserSponsor(in *pb2.UpdateUserSponsorReq) (*pb2.UpdateUserSponsorResp, error) {
	sponsor, err := l.svcCtx.UserSponsorModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR_NOT_FOUND), "查询赞助商失败:%v,in:%+v", err, in)
	}
	sponsor.Id = in.Id
	sponsor.UserId = in.UserId
	sponsor.Type = in.Type
	sponsor.AppletType = in.AppletType
	sponsor.IsShow = in.IsShow
	sponsor.Name = in.Name
	sponsor.Desc = in.Desc
	sponsor.QrCode = in.QrCode
	sponsor.InputA = in.InputA
	sponsor.InputB = in.InputB

	err = l.svcCtx.UserSponsorModel.Update(l.ctx, sponsor)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新赞助商失败:%v, in:%+v", err, in)
	}
	return &pb2.UpdateUserSponsorResp{}, nil
}
