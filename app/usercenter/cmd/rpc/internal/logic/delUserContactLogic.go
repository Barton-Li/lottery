package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"
)

type DelUserContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserContactLogic {
	return &DelUserContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserContactLogic) DelUserContact(in *pb2.DelUserContactReq) (*pb2.DelUserContactResp, error) {
	//只能删除自己的Contact，所以这里只需要删除自己的Contact即可
	//查询自己的Contact

	for _, id := range in.Id {
		if err := l.svcCtx.UserContactModel.Delete(l.ctx, id); err != nil {
			return nil, err
		}
	}

	return &pb2.DelUserContactResp{}, nil
}
