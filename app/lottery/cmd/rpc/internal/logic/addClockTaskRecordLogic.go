package logic

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"lottery/app/lottery/model"
	"github.com/pkg/errors"
	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"
	"lottery/common/constants"
	"lottery/common/xerr"
	"math/rand"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClockTaskRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddClockTaskRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClockTaskRecordLogic {
	return &AddClockTaskRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------完成打卡任务-----------------------
// AddClockTaskRecord 添加用户完成打卡任务的记录
// 该方法首先验证抽奖ID和打卡任务ID的存在性和一致性，然后检查用户是否已经完成过该打卡任务。
// 如果用户未完成过该任务，将创建一个新的打卡任务记录并保存到数据库中。
func (l *AddClockTaskRecordLogic) AddClockTaskRecord(in *pb.AddClockTaskRecordReq) (*pb.AddClockTaskRecordResp, error) {
	// 根据抽奖ID查询抽奖信息
	lottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.LotteryId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "lottery_id:%d,err:%v", in.LotteryId, err)

	}
	if errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrMsg("抽奖ID不存在"), "抽奖ID不存在, id: %v", in.LotteryId)
	}
	// 验证抽奖关联的ID是否一致
	if lottery.ClockTaskId == 0 || lottery.ClockTaskId != in.ClockTaskId {
		return nil, errors.Wrapf(xerr.NewErrMsg("抽奖关联任务ID不一致"), "抽奖关联任务ID不一致,lotteryInfo.ClockTaskId:%d, in.ClockTaskId:%d", lottery.ClockTaskId, in.ClockTaskId)
	}
	// 根据打卡任务ID查询打卡任务信息
	clockTask, err := l.svcCtx.ClockTaskModel.FindOne(l.ctx, in.ClockTaskId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "clock_task_id:%d,err:%v", in.ClockTaskId, err)
	}
	if errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrMsg("打卡任务ID不存在"), "打卡任务ID不存在, id:%v", in.ClockTaskId)
	}
	// 检查用户是否已经完成过该打卡任务
	builder := l.svcCtx.ClockTaskRecordModel.SelectBuilder().
		Where(sq.Eq{"user_id": in.UserId}).
		Where(sq.Eq{"lottery_id": lottery.Id}).
		Where(sq.Eq{"clock_task_id": clockTask.Id})
	count, err := l.svcCtx.ClockTaskRecordModel.FindCount(l.ctx, builder, "id")
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户已完成过打卡任务"), "用户已完成过打卡任务, user_id:%v, lottery_id:%v, clock_task_id:%v", in.UserId, lottery.Id, clockTask.Id)
	}
	// 新增完成打卡任务记录
	clockTaskRecord := new(model.ClockTaskRecord)
	clockTaskRecord.LotteryId = lottery.Id
	clockTaskRecord.UserId = in.UserId
	clockTaskRecord.ClockTaskId = clockTask.Id
	// 获取中奖倍率
	if clockTask.ChanceType == constants.Appoint {
		// 指定中奖倍率
		clockTaskRecord.IncreaseMultiple = clockTask.IncreaseMultiple
	} else {
		// 随机1～10倍
		clockTaskRecord.IncreaseMultiple = rand.Int63n(10) + 1
	}
	// 插入打卡任务记录到数据库
	insert, err := l.svcCtx.ClockTaskRecordModel.Insert(l.ctx, clockTaskRecord)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_INSERTCLOCKTASKRECORD_ERROR), "Lottery Database Exception clockTaskRecord : %+v , err: %v", clockTaskRecord, err)
	}
	// 获取插入记录的ID
	clockTaskRecordId, _ := insert.LastInsertId()
	return &pb.AddClockTaskRecordResp{
		Id: clockTaskRecordId,
	}, nil
}
