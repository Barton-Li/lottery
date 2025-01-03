package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"
	"lottery/app/usercenter/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserDynamicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserDynamicLogic {
	return &SearchUserDynamicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserDynamicLogic) SearchUserDynamic(in *pb2.SearchUserDynamicReq) (*pb2.SearchUserDynamicResp, error) {

	list, err := l.svcCtx.UserDynamicModel.FindListByUserId(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	var resp []*pb2.UserDynamic
	if len(list) > 0 {
		for _, dynamic := range list {
			var pbDynamic pb2.UserDynamic
			logx.Error("dynamic:", dynamic)
			_ = copier.Copy(&pbDynamic, dynamic)
			pbDynamic.UpdateTime = dynamic.UpdateTime.Unix()
			resp = append(resp, &pbDynamic)
		}
	}
	return &pb2.SearchUserDynamicResp{
		UserDynamic: resp,
	}, nil
}
