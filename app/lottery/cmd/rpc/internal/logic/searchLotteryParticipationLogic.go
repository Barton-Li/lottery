package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLotteryParticipationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLotteryParticipationLogic {
	return &SearchLotteryParticipationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SearchLotteryParticipation 用于搜索抽奖参与信息。
// 它接受一个 SearchLotteryParticipationReq 请求对象作为输入，
// 并返回一个 SearchLotteryParticipationResp 响应对象，其中包含参与信息的列表和总数。
// 主要功能包括根据抽奖ID筛选参与信息、分页查询、以及获取参与该抽奖的总人数。
func (l *SearchLotteryParticipationLogic) SearchLotteryParticipation(in *pb.SearchLotteryParticipationReq) (*pb.SearchLotteryParticipationResp, error) {
    // 计算查询的起始位置（offset）和查询数量（limit），以支持分页功能。
    offset := (in.PageIndex - 1) * in.PageSize
    limit := in.PageSize

    // 构建查询语句，选择符合条件的记录。
    // 这里指定了抽奖ID作为筛选条件，并设置了查询的数量和起始位置。
    builder := l.svcCtx.LotteryParticipationModel.SelectBuilder().Where("lottery_id=?", in.LotteryId).Limit(uint64(limit)).Offset(uint64(offset))

    // 执行查询，获取参与信息列表。
    list, err := l.svcCtx.LotteryParticipationModel.FindAll(l.ctx, builder, "")
    if err != nil {
        return nil, err
    }

    // 获取参与指定抽奖的总人数。
    count, err := l.svcCtx.LotteryParticipationModel.GetParticipatorsCountByLotteryId(l.ctx, in.LotteryId)
    if err != nil {
        return nil, err
    }

    // 初始化响应对象，包含总人数和参与信息列表。
    resp := &pb.SearchLotteryParticipationResp{
        Count: count,
        List:  []*pb.LotteryParticipation{},
    }

    // 将查询到的参与信息列表复制到响应对象中。
    if err = copier.Copy(&resp.List, list); err != nil {
        return nil, err
    }

    // 返回响应对象。
    return resp, nil
}

