package logic

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"lottery/common/constants"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckSelectedLotteryParticipatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckSelectedLotteryParticipatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckSelectedLotteryParticipatedLogic {
	return &CheckSelectedLotteryParticipatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CheckSelectedLotteryParticipatedLogic 检查用户是否参与了选定的抽奖逻辑。
// 该方法通过查询用户参与的抽奖记录，并检查这些抽奖是否被选定，来判断用户是否参与了选定的抽奖。
func (l *CheckSelectedLotteryParticipatedLogic) CheckSelectedLotteryParticipated(in *pb.CheckSelectedLotteryParticipatedReq) (*pb.CheckSelectedLotteryParticipatedResp, error) {
	// 构建查询条件以获取用户参与的抽奖记录。
	builder := l.svcCtx.LotteryParticipationModel.SelectBuilder().
		Where("user_id = ?", in.UserId)

	// 执行查询，获取用户参与的抽奖记录。
	participations, err := l.svcCtx.LotteryParticipationModel.FindAll(l.ctx, builder, "")
	if err != nil {
		return nil, err
	}

	// 提取用户参与的抽奖ID。
	lotteryIds := make([]int64, len(participations))
	for i := range participations {
		lotteryIds[i] = participations[i].LotteryId
	}

	// 构建查询条件以获取被选定的抽奖记录。
	builder = l.svcCtx.LotteryModel.SelectBuilder().
		Where(sq.Eq{"id": lotteryIds}).
		Where("is_selected = ?", constants.IsSelectedLottery)

	// 计算符合条件的抽奖记录数量。
	count, err := l.svcCtx.LotteryParticipationModel.FindCount(l.ctx, builder, "id")
	if err != nil {
		return nil, err
	}

	// 如果存在符合条件的抽奖记录，将计数设置为1，否则为0。
	if count > 0 {
		count = 1
	}

	// 返回用户是否参与了选定的抽奖的结果。
	return &pb.CheckSelectedLotteryParticipatedResp{
		Participated: count,
	}, nil

}
