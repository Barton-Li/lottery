package userSponsor

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/pb"
	"lottery/common/xerr"

	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateSponsorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改抽奖发起人（赞助商）
func NewUpDateSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateSponsorLogic {
	return &UpDateSponsorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpDateSponsorLogic) UpDateSponsor(req *types.UpdateSponsorReq) (resp *types.UpdateSponsorResp, err error) {
	pbSponsorReq := new(pb.UpdateUserSponsorReq)
	err = copier.Copy(pbSponsorReq, req)
	if err != nil {
		l.Logger.Errorf("Error 创建用户失败")
		return nil, errors.Wrapf(xerr.NewErrMsg("创建用户失败"), "copy to sponsor failed: req=%+v, err=%v", req, err)
	}
	sponsor, err := l.svcCtx.UsercenterRpc.UpdateUserSponsor(l.ctx, pbSponsorReq)
	if err != nil {
		// 如果添加赞助商失败，则返回错误
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to add sponsor"), "failed to add sponsor:%+v,err:%s", sponsor, err)
	}

	err = copier.Copy(sponsor, resp)
	if err != nil {
		// 如果添加赞助商失败，则返回错误
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to add sponsor"), "failed to add sponsor:%+v,err:%s", sponsor, err)
	}

	return resp, nil
}
