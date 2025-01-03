package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/usercenter"
	"lottery/app/usercenter/model"

	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	registerResp,err:=	l.svcCtx.UsercenterRpc.Register(l.ctx,&usercenter.RegisterReq{
		Mobile: req.Mobile,
		Password: req.Password,
		AuthType: model.UserAuthTypeSystem,
		AuthKey: req.Mobile,

	})
	if err!=nil{
		return nil,errors.Wrapf(err,"req:%+v",req)
	}

	_=copier.Copy(resp,registerResp)
	l.Logger.Debugf("Error 创建用户成功")
	return resp,nil

}
