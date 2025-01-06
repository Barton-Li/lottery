package logic

import (
	"context"
	"lottery/common/constants"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserCreateLotteryAndTodayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserCreateLotteryAndTodayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserCreateLotteryAndTodayLogic {
	return &CheckUserCreateLotteryAndTodayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserCreateLotteryAndTodayLogic) CheckUserCreateLotteryAndToday(in *pb.CheckUserCreateLotteryAndTodayReq) (*pb.CheckUserCreateLotteryAndTodayResp, error) {
	userId := in.UserId
	// 根据uid获取当前用户并且今天发布的的所有抽奖id
	LotteryIds, err := l.svcCtx.LotteryModel.GetTodayLotteryIdsByUserId(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	//fmt.Println("lotterys:", LotteryIds)
	// 判断是否有一个抽奖符合，有一个符合就跳出循环，返回yes = 1
	for _, lotteryId := range LotteryIds {
		yes, err := l.CheckLotteryToday(lotteryId)
		if err != nil {
			return nil, err
		}
		if yes {
			return &pb.CheckUserCreateLotteryAndTodayResp{
				Yes: 1,
			}, nil
		}
	}
	return &pb.CheckUserCreateLotteryAndTodayResp{
		Yes: 0,
	}, nil
}
func (l *CheckUserCreateLotteryAndTodayLogic) CheckLotteryToday(lotteryID int64) (bool, error) {
	participantsCount, err := l.svcCtx.LotteryParticipationModel.GetParticipatorsCountByLotteryId(l.ctx, lotteryID)
	if err != nil {
		return false, err
	}
	//fmt.Println("participantsCount:", participantsCount)
	// 判断抽奖是否在今天之内发起并有超过五个人参加
	return participantsCount > constants.LotteryTodayParticipantsCount, nil
}
