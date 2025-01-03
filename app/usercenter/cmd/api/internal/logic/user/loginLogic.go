package user

import (
	"context"
	"github.com/jinzhu/copier"
	"lottery/app/usercenter/cmd/rpc/usercenter"
	"lottery/app/usercenter/model"

	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginResp,err:=l.svcCtx.UsercenterRpc.Login(l.ctx,&usercenter.LoginReq{
		AuthKey: req.Mobile,
		Password: req.Password,
		AuthType: model.UserAuthTypeSystem,
	})
	if err != nil {
		return nil, err
	}
	_=copier.Copy(resp,loginResp)

	return resp,nil
}
