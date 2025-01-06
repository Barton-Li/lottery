package logic

import (
	"context"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWonListCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWonListCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWonListCountLogic {
	return &GetWonListCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWonListCountLogic) GetWonListCount(in *pb.GetWonListCountReq) (*pb.GetWonListCountResp, error) {
	count, err := l.svcCtx.LotteryParticipationModel.GetWonListCountByUserId(l.ctx, in.UserId)
	if err != nil {
		// 如果获取用户中奖列表数量出错，返回错误。
		return nil, err
	}

	return &pb.GetWonListCountResp{
		Count: count,
	}, nil
}
