package logic

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"lottery/common/constants"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"
	"github.com/jinzhu/now"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSelectedLotteryStatisticLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSelectedLotteryStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSelectedLotteryStatisticLogic {
	return &GetSelectedLotteryStatisticLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetSelectedLotteryStatistic 获取用户在当前日和周内选定的抽奖统计信息。
// 该方法首先计算用户在今天选定的抽奖数量，然后计算用户在本周选定的抽奖数量。
// in: 请求参数，包含需要统计的用户ID。
// 返回值: 包含日和周选定抽奖数量的响应对象，以及可能出现的错误。
func (l *GetSelectedLotteryStatisticLogic) GetSelectedLotteryStatistic(in *pb.GetSelectedLotteryStatisticReq) (*pb.GetSelectedLotteryStatisticResp, error) {
	// 获取当天的开始和结束时间
	start := now.BeginningOfDay()
	end := now.EndOfDay()

	// 构建查询用户今天参与的抽奖记录的查询构建器
	builder := l.svcCtx.LotteryParticipationModel.SelectBuilder().
		Where("user_id = ? AND create_time >= ? AND create_time <= ?", in.UserId, start, end)

	// 查询用户今天参与的抽奖记录
	participations, err := l.svcCtx.LotteryParticipationModel.FindAll(l.ctx, builder, "")
	if err != nil {
		return nil, err
	}

	// 提取参与的抽奖ID
	lotteryIds := make([]int64, len(participations))
	for i := range participations {
		lotteryIds[i] = participations[i].LotteryId
	}

	// 构建查询选定的抽奖记录的查询构建器
	builder = l.svcCtx.LotteryModel.SelectBuilder().
		Where(sq.Eq{"id": lotteryIds}).
		Where("is_selected = ?", constants.IsSelectedLottery)

	// 查询今天选定的抽奖数量
	dayCount, err := l.svcCtx.LotteryParticipationModel.FindCount(l.ctx, builder, "id")
	if err != nil {
		return nil, err
	}

	// 获取当周的开始和结束时间
	start = now.BeginningOfWeek()
	end = now.EndOfWeek()

	// 构建查询用户本周参与的抽奖记录的查询构建器
	builder = l.svcCtx.LotteryParticipationModel.SelectBuilder().
		Where("user_id = ? AND create_time >= ? AND create_time <= ?", in.UserId, start, end)

	// 查询用户本周参与的抽奖记录
	participations, err = l.svcCtx.LotteryParticipationModel.FindAll(l.ctx, builder, "")
	if err != nil {
		return nil, err
	}

	// 提取参与的抽奖ID
	lotteryIds = make([]int64, len(participations))
	for i := range participations {
		lotteryIds[i] = participations[i].LotteryId
	}

	// 构建查询选定的抽奖记录的查询构建器
	builder = l.svcCtx.LotteryModel.SelectBuilder().
		Where(sq.Eq{"id": lotteryIds}).
		Where("is_selected = ?", constants.IsSelectedLottery)

	// 查询本周选定的抽奖数量
	weekCount, err := l.svcCtx.LotteryParticipationModel.FindCount(l.ctx, builder, "id")
	if err != nil {
		return nil, err
	}

	// 构建响应对象
	resp := &pb.GetSelectedLotteryStatisticResp{
		DayCount:  dayCount,
		WeekCount: weekCount,
	}

	// 返回响应对象和nil错误，表示操作成功
	return resp, nil
}
