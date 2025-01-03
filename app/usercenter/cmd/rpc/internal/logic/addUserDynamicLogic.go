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

type AddUserDynamicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserDynamicLogic {
	return &AddUserDynamicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------userDynamic-----------------------
func (l *AddUserDynamicLogic) AddUserDynamic(in *pb2.AddUserDynamicReq) (*pb2.AddUserDynamicResp, error) {
	userDynamic := new(model.UserDynamic)
	err := copier.Copy(userDynamic, in)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "copy error:%v,in:%+v", err, in)
	}
	insert, err := l.svcCtx.UserDynamicModel.Insert(l.ctx, userDynamic)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert error:%v, in:%+v", err, in)
	}
	lastInsertId, err := insert.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "LastInsertId error:%v, in:%+v", err, in)
	}
	return &pb2.AddUserDynamicResp{
		Id: lastInsertId,
	}, nil
}
