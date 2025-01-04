package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"lottery/app/lottery/model"
	"lottery/common/xerr"
	"time"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLotteryLogic {
	return &AddLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -------------------------抽奖表-------------------------
// AddLotteryLogic 的 AddLottery 方法用于添加新的抽奖活动。
// 该方法接收一个 AddLotteryReq 请求对象，包含抽奖活动的相关信息，
// 并返回一个 AddLotteryResp 响应对象，其中包含新添加抽奖活动的ID。
func (l *AddLotteryLogic) AddLottery(in *pb.AddLotteryReq) (*pb.AddLotteryResp, error) {
    var lotteryId int64
    // 使用事务处理抽奖活动的添加过程，确保数据一致性。
    err := l.svcCtx.LotteryModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
        // 抽奖基本信息
        lottery := new(model.Lottery)
        lottery.UserId = in.UserId
        lottery.Name = in.Name
        lottery.AwardDeadline = time.Unix(in.AwardDeadline, 0)
        lottery.Introduce = in.Introduce
        lottery.JoinNumber = in.JoinNumber
        lottery.AnnounceType = in.AnnounceType
        lottery.AnnounceTime = time.Unix(in.AnnounceTime, 0)
        lottery.Thumb = in.Thumb
        lottery.IsSelected = 0
        lottery.IsAnnounced = 0
        lottery.SponsorId = in.SponsorId
        lottery.IsClocked = in.IsClocked
        if in.PublishType == 1 {
            lottery.PublishTime.Time = time.Now()
            lottery.PublishTime.Valid = true
        }
        // 插入抽奖活动到数据库
        insert, err := l.svcCtx.LotteryModel.TransInsert(l.ctx, session, lottery)
        if err != nil {
            return errors.Wrapf(xerr.NewErrCode(xerr.DB_INSERTLOTTERY_ERROR), "insert lottery error: %v,lottery:%+v", err, lottery)
        }
        // 获取插入记录的ID
        lotteryId, err := insert.LastInsertId()
        if err != nil {
            return errors.Wrapf(err, "insert lottery error: %v, lottery:%+v", err, lottery)
        }
        // 处理奖品信息
        for _, pbPrize := range in.Prizes {
            prize := new(model.Prize)
            err := copier.Copy(prize, pbPrize)
            if err != nil {
                return errors.Wrapf(err, "copy prize error: %v, in:%+v", err, pbPrize)
            }
            prize.LotteryId = lotteryId
            _, err = l.svcCtx.PrizeModel.TransInsert(l.ctx, session, prize)
            if err != nil {
                return errors.Wrapf(xerr.NewErrCode(xerr.DB_INSERTPRIZE_ERROR), "insert prize error: %v, prize:%+v", err, prize)
            }
        }
        // 创建打卡任务
        if in.ClockTask != nil {
            clockTask := new(model.ClockTask)
            clockTask.LotteryId = lotteryId
            clockTask.Seconds = in.ClockTask.Seconds
            clockTask.Type = in.ClockTask.Type
            clockTask.AppletType = in.ClockTask.AppletType
            clockTask.PageLink = in.ClockTask.PageLink
            clockTask.AppId = in.ClockTask.AppId
            clockTask.PagePath = in.ClockTask.PagePath
            clockTask.Image = in.ClockTask.Image
            clockTask.VideoAccountId = in.ClockTask.VideoAccountId
            clockTask.VideoId = in.ClockTask.VideoId
            clockTask.ArticleLink = in.ClockTask.ArticleLink
            clockTask.Copywriting = in.ClockTask.Copywriting
            clockTask.ChanceType = in.ClockTask.ChanceType
            clockTask.IncreaseMultiple = in.ClockTask.IncreaseMultiple
            // 插入打卡任务到数据库
            insert, err = l.svcCtx.ClockTaskModel.TransInsert(l.ctx, session, clockTask)
            if err != nil {
                return errors.Wrapf(xerr.NewErrCode(xerr.DB_INSERTLOTTERY_ERROR), "Lottery Database Exception clockTask : %+v , err: %v", clockTask, err)
            }
            clockTaskId, _ := insert.LastInsertId()
            // 更新lottery表的打卡任务ID
            lottery := new(model.Lottery)
            lottery.Id = lotteryId
            lottery.ClockTaskId = clockTaskId
            _, err = l.svcCtx.LotteryModel.TransUpdateClockTaskId(l.ctx, session, lottery)
            if err != nil {
                logx.Error("更新打卡任务ID失败:%v", err)
                return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Lottery Database Exception lottery : %+v , err: %v", lottery, err)
            }
        }
        return nil
    })
    if err != nil {
        return nil, err
    }
    // 返回新添加抽奖活动的ID
    return &pb.AddLotteryResp{
        Id: lotteryId,
    }, nil
}

