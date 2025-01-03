package userContact

import (
	"context"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/pb"

	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContactDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量删除抽奖发起人的联系方式
func NewContactDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContactDelLogic {
	return &ContactDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContactDelLogic) ContactDel(req *types.ContactDelReq) (resp *types.ContactDelResp, err error) {
	_,err=l.svcCtx.UsercenterRpc.DelUserContact(l.ctx,&pb.DelUserContactReq{
		Id: req.Id,
	})
	if err!= nil {
		return nil,errors.Wrapf(err, "req:%+v", req)
	}
	return
}
