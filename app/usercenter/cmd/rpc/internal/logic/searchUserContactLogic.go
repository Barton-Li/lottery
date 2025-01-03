package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"
	"lottery/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserContactLogic {
	return &SearchUserContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserContactLogic) SearchUserContact(in *pb2.SearchUserContactReq) (*pb2.SearchUserContactResp, error) {
	list, err := l.svcCtx.UserContactModel.FindPageByUserId(l.ctx, in.UserId, in.Page, in.Limit)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "search user contact error:%v, in:%+v", err, in)
	}
	var resp []*pb2.UserContact
	if len(list) > 0 {
		for _, contact := range list {
			var pbcontact pb2.UserContact
			_ = copier.Copy(&pbcontact, contact)
			resp = append(resp, &pbcontact)
		}
	}
	return &pb2.SearchUserContactResp{
		UserContact: resp,
	}, nil
}
