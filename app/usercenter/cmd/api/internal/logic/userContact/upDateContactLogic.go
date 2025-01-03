package userContact

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/pb"
	"lottery/common/xerr"

	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改抽奖发起人的联系方式
func NewUpDateContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateContactLogic {
	return &UpDateContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpDateContactLogic) UpDateContact(req *types.UpdateContactReq) (resp *types.UpdateContactResp, err error) {
	pbContactReq := new(pb.UpdateUserContactReq)
	err = copier.Copy(pbContactReq, req)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to copy contact"), "req:%+v,err:%v", req, err)
	}
	content, err := json.Marshal(req.Content)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to marshal contact"), "req:%+v, err:%v", req, err)
	}
	pbContactReq.Content = string(content)
	contact, err := l.svcCtx.UsercenterRpc.UpdateUserContact(l.ctx, pbContactReq)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to update contact"), "req:%+v, err:%v", req, err)
	}

	return &types.UpdateContactResp{
		Id: contact.Id,
	}, nil
}
