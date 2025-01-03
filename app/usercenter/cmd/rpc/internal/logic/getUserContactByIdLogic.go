package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"
	"lottery/app/usercenter/cmd/rpc/usercenter"
	"lottery/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserContactByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserContactByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserContactByIdLogic {
	return &GetUserContactByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserContactByIdLogic) GetUserContactById(in *pb2.GetUserContactByIdReq) (*pb2.GetUserContactByIdResp, error) {
	userContact, err := l.svcCtx.UserContactModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserContactById FindOne error:%v,in:%+v", err, in)
	}
	if userContact == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户联系人不存在"), "用户联系人不存在 id:%d", in.Id)
	}
	var respUserContact usercenter.UserContact
	_ = copier.Copy(&respUserContact, userContact)

	return &pb2.GetUserContactByIdResp{
		UserContact: &respUserContact,
	}, nil
}
