package logic

import (
	"context"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"
	"lottery/common/constants"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryStatisticLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryStatisticLogic {
	return &GetLotteryStatisticLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetLotteryStatistic 获取用户彩票参与统计信息
// 该方法根据用户ID查询用户参与的彩票数量、创建的彩票数量和中奖的彩票数量
// 参数:
//   in *pb.GetLotteryStatisticReq: 包含用户ID的请求对象
// 返回值:
//   *pb.GetLotteryStatisticResp: 包含参与、创建和中奖彩票数量的响应对象
//   error: 错误信息，如果执行过程中遇到错误则返回
func (l *GetLotteryStatisticLogic) GetLotteryStatistic(in *pb.GetLotteryStatisticReq) (*pb.GetLotteryStatisticResp, error) {
	// 查询用户参与的彩票数量
	builder := l.svcCtx.LotteryParticipationModel.SelectBuilder().Where("user_id = ?", in.UserId)
	ParticipationCount, err := l.svcCtx.LotteryParticipationModel.FindCount(l.ctx, builder, "id")
	if err != nil {
		return nil, err
	}

	// 查询用户创建的彩票数量
	builder = l.svcCtx.LotteryModel.SelectBuilder().Where("user_id = ?", in.UserId)
	CreatedCount, err := l.svcCtx.LotteryModel.FindCount(l.ctx, builder, "id")
	if err != nil {
		return nil, err
	}

	// 查询用户中奖的彩票数量
	builder = l.svcCtx.LotteryParticipationModel.SelectBuilder().Where("user_id = ? and is_won = ?", in.UserId, constants.IsWon)
	WonCount, err := l.svcCtx.LotteryParticipationModel.FindCount(l.ctx, builder, "id")
	if err != nil {
		return nil, err
	}

	// 返回用户彩票参与统计信息
	return &pb.GetLotteryStatisticResp{
		ParticipationCount: ParticipationCount,
		CreatedCount:       CreatedCount,
		WonCount:           WonCount,
	}, nil
}
