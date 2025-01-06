package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"lottery/app/lottery/model"
	"lottery/common/constants"
	"lottery/common/xerr"
	"math/rand"
	"sort"
	"time"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnnounceLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}
type LotteryStrategy interface {
	Run() error
}
type TimeLotteryStrategy struct {
	*AnnounceLotteryLogic
	CurrentTime time.Time
}

type PeopleLotteryStrategy struct {
	*AnnounceLotteryLogic
	CurrentTime time.Time
}

type Winner struct {
	LotteryId int64
	UserId    int64
	PrizeId   int64
}

func NewAnnounceLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnnounceLotteryLogic {
	return &AnnounceLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AnnounceLotteryLogic) AnnounceLottery(in *pb.AnnounceLotteryReq) (*pb.AnnounceLotteryResp, error) {
	var strategy LotteryStrategy
	switch in.AnnounceType {
	case constants.AnnounceTypeTimeLottery:
		strategy = &TimeLotteryStrategy{
			AnnounceLotteryLogic: l,
			CurrentTime:          time.Now(),
		}
	case constants.AnnounceTypePeopleLottery:
		strategy = &PeopleLotteryStrategy{
			AnnounceLotteryLogic: l,
			CurrentTime:          time.Now(),
		}

	}
	err := strategy.Run()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "AnnounceStrategy run error: %v", err)

	}
	return &pb.AnnounceLotteryResp{}, nil
}
func (s *TimeLotteryStrategy) Run() error {
	litteries, err := s.svcCtx.LotteryModel.GetLotterysByLessThanCurrentTime(s.ctx, s.CurrentTime, constants.AnnounceTypeTimeLottery)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "GetLotterysByLessThanCurrentTime error: %v", err)
	}
	for _, l := range litteries {
		var participators []int64
		err = s.svcCtx.LotteryModel.Trans(s.ctx, func(ctx context.Context, session sqlx.Session) error {
			prizes, err := s.svcCtx.PrizeModel.FindByLotteryId(s.ctx, l)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "FindByLotteryId error: %v", err)
			}
			participators, err = s.svcCtx.LotteryParticipationModel.GetParticipationUserIdsByLotteryId(s.ctx, l)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "GetParticipationUserIdsByLotteryId error: %v", err)
			}
			winners, err := s.DrawLottery(s.ctx, l, prizes, participators)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "DrawLottery error: %v,lotteryId:%d,prizeIds:%v,participators:%v", err, l, prizes, participators)
			}
			err = s.svcCtx.LotteryModel.UpdateLotteryStatus(s.ctx, l)
			if err != nil {
				return err
			}
			err = s.WriteWinnersToLotteryParticipation(winners)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "Trans error: %v", err)
		}
	}
	return nil
}
func (l *AnnounceLotteryLogic) WriteWinnersToLotteryParticipation(winners []Winner) error {
	for _, winner := range winners {
		err := l.svcCtx.LotteryParticipationModel.UpdateWinners(l.ctx, winner.LotteryId, winner.UserId, winner.PrizeId)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "UpdateLotteryParticipation error: %v", err)
		}
	}
	return nil
}
func (s *PeopleLotteryStrategy) CheckLottery(lotteries []*model.Lottery) (CheckedLotterys []*model.Lottery, err error) {
	for _, l := range lotteries {
		if l.AnnounceTime.Before(s.CurrentTime) || l.AnnounceTime.Equal(s.CurrentTime) {
			CheckedLotterys = append(CheckedLotterys, l)
		} else {
			ParticipatorCount, err := s.svcCtx.LotteryParticipationModel.GetParticipatorsCountByLotteryId(s.ctx, l.Id)
			if err != nil {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "GetParticipatedLotteryIdsByUserId error: %v", err)
			}
			if ParticipatorCount >= l.JoinNumber {
				CheckedLotterys = append(CheckedLotterys, l)

			}
		}
	}
	return
}
func (s *PeopleLotteryStrategy) Run() error {
	lotteries, err := s.svcCtx.LotteryModel.GetTypeIs2AndIsNotAnnounceLotterys(s.ctx, constants.AnnounceTypePeopleLottery)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "GetTypeIs2AndIsNotAnnounceLotterys error: %v", err)
	}
	CheckLottery, err := s.CheckLottery(lotteries)
	if err != nil {
		return err
	}
	for _, lottery := range CheckLottery {
		var participators []int64
		err = s.svcCtx.PrizeModel.Trans(s.ctx, func(context context.Context, session sqlx.Session) error {
			prizes, err := s.svcCtx.PrizeModel.FindByLotteryId(s.ctx, lottery.Id)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "FindByLotteryId error: %v", err)
			}
			participators, err = s.svcCtx.LotteryParticipationModel.GetParticipationUserIdsByLotteryId(s.ctx, lottery.Id)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "GetParticipationUserIdsByLotteryId error: %v", err)
			}
			winners, err := s.DrawLottery(s.ctx, lottery.Id, prizes, participators)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "DrawLottery error: %v, lotteryId:%d, prizeIds:%v, participators:%v", err, lottery.Id, prizes, participators)
			}
			err = s.svcCtx.LotteryModel.UpdateLotteryStatus(s.ctx, lottery.Id)
			if err != nil {
				return err
			}
			err = s.WriteWinnersToLotteryParticipation(winners)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "Trans error: %v", err)
		}

	}
	return nil

}

// DrawLottery 执行抽奖逻辑，根据奖品和参与者列表确定中奖者。
// ctx: 上下文，用于处理请求、超时等。
// lotteryId: 抽奖活动的ID。
// prizes: 奖品列表，每个奖品包含其ID和数量。
// participators: 参与者列表，每个参与者有一个用户ID。
// 返回值: 中奖者列表和可能的错误。
func (l *AnnounceLotteryLogic) DrawLottery(ctx context.Context, lotteryId int64, prizes []*model.Prize, participators []int64) ([]Winner, error) {
	// 初始化随机数生成器。
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// 计算所有奖品的中奖者总数。
	var winnersNum int64
	for _, prize := range prizes {
		winnersNum += prize.Count
	}

	// 初始化中奖者列表。
	winners := make([]Winner, 0)

	// 获取参与者的打卡任务记录，以便计算中奖概率。
	records, err := l.svcCtx.ClockTaskRecordModel.GetClockTaskRecordByLotteryIdAndUserIds(lotteryId, participators)
	if err != nil {
		return nil, err
	}

	// 初始化参与者中奖概率映射。
	RationsMap := make(map[int64]int64)
	for _, participator := range participators {
		RationsMap[participator] = 1
	}

	// 根据打卡记录增加参与者的中奖概率。
	for _, record := range records {
		RationsMap[record.UserId] += record.IncreaseMultiple
	}

	// 将中奖概率映射转换为概率数组。
	Ratios := make([]int64, len(participators))
	for i, participator := range participators {
		Ratios[i] = RationsMap[participator]
	}

	// 计算总中奖概率。
	toatlRatio := int64(0)
	for _, ratio := range Ratios {
		toatlRatio += ratio
	}

	// 计算每个参与者的最终中奖概率。
	FinalRatios := make([]float64, len(participators))
	for idx := range Ratios {
		FinalRatios[idx] = float64(Ratios[idx]) / float64(toatlRatio)
	}

	// 开始抽奖过程。
	for i := 0; i < int(winnersNum); i++ {
		var rangdomWinnerIndex int
		var winnerUserId int64

		// 如果没有参与者，则跳出循环。
		if len(participators) == 0 {
			break
		}

		// 生成随机概率并确定中奖者。
		// 生成一个随机概率值，用于后续选择获胜者
		rangdomPrrobability := rand.Float64()
		// 初始化概率累加和，用于计算累计概率
		probabilitySum := 0.0
		// 遍历所有参与者，根据FinalRatios中的概率选择获胜者
		for ids := range participators {
			// 累加当前参与者的概率到总概率中
			probabilitySum += FinalRatios[ids]
			// 如果当前累计概率大于等于随机概率值，则选择当前参与者为获胜者
			if rangdomPrrobability <= probabilitySum {
				// 记录获胜者的索引
				rangdomWinnerIndex = ids
				// 根据获胜者索引获取并记录获胜者的用户ID
				winnerUserId = participators[rangdomWinnerIndex]
				// 选择出获胜者后，结束循环
				break
			}
		}

		// 如果未确定中奖者，则默认为第一个参与者。
		if winnerUserId == 0 {
			winnerUserId = participators[0]
		}

		// 对剩余奖品进行排序，确保按类型顺序分配。
		sort.Slice(prizes, func(i, j int) bool {
			return prizes[i].Type < prizes[j].Type
		})

		// 更新奖品列表，移除已分配完毕的奖品。
		if prizes[0].Count == 0 {
			prizes = prizes[1:]
		}
		prizes[0].Count--
		prizesId := prizes[0].Id

		// 创建并添加中奖者到中奖者列表。
		winner := Winner{
			LotteryId: lotteryId,
			UserId:    winnerUserId,
			PrizeId:   prizesId,
		}
		winners = append(winners, winner)

		// 移除已中奖的参与者和对应的概率。
		participators = append(participators[:rangdomWinnerIndex], participators[rangdomWinnerIndex+1:]...)
		FinalRatios = append(FinalRatios[:rangdomWinnerIndex], FinalRatios[rangdomWinnerIndex+1:]...)
	}

	// 返回中奖者列表和nil错误。
	return winners, nil
}

//func (l *AnnounceLotteryLogic) NotifyParticipators(participators []int64, lotteryId int64) error {
//	fmt.Println("NotifyParticipators", participators, lotteryId)
//	_, err := l.svcCtx.NoticeRpc.NoticeLotteryDraw(l.ctx, &notice.NoticeLotteryDrawReq{
//		LotteryId: lotteryId,
//		UserIds:   participators,
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//}
