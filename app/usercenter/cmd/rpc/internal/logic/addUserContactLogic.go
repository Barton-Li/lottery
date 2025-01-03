package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"
	"lottery/app/usercenter/model"
	"lottery/common/xerr"


	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserContactLogic {
	return &AddUserContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------userContact-----------------------
func (l *AddUserContactLogic) AddUserContact(in *pb2.AddUserContactReq) (*pb2.AddUserContactResp, error) {
	userContact := new(model.UserContact)
	err := copier.Copy(userContact, in)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "copy error:%v,in%+v", err, in)
	}
	insert, err := l.svcCtx.UserContactModel.Insert(l.ctx, userContact)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert error:%v, in%+v", err, in)
	}
	insertId, err := insert.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert error:%v, in%+v", err, in)
	}
	return &pb2.AddUserContactResp{
		Id: insertId,
	}, nil
}
