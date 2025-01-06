package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/lottery/cmd/rpc/lottery"
	"lottery/app/lottery/cmd/rpc/pb"
	"lottery/common/ctxdata"
	"lottery/common/xerr"

	"lottery/app/lottery/cmd/api/internal/svc"
	"lottery/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLotteryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLotteryLogic {
	return &CreateLotteryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateLottery 创建lottery逻辑
// 该方法接收一个创建lottery的请求，并返回创建结果或错误
func (l *CreateLotteryLogic) CreateLottery(req *types.CreateLotteryReq) (resp *types.CreateLotteryResp, err error) {
	// 从上下文中获取用户ID
	userId := ctxdata.GetUidFromCtx(l.ctx)

	// 初始化一个空的奖品列表
	var pbPrizes []*pb.Prize

	// 转换奖品
	for _, prize := range req.Prizes {
		pbPrize := new(pb.Prize)
		// 将请求中的奖品信息复制到新的Prize对象中
		err := copier.Copy(&pbPrize, prize)
		if err != nil {
			// 如果复制过程中出现错误，则返回错误
			return nil, errors.Wrapf(xerr.NewErrMsg("Copy pbPrize Error"), "Copy pbPrize Error req: %+v , err : %v ", pbPrize, err)
		}
		// 将转换后的奖品添加到奖品列表中
		pbPrizes = append(pbPrizes, pbPrize)
	}

	// 初始化一个空的定时任务对象
	pbClockTask := new(pb.ClockTask)

	// 如果请求中包含定时任务信息，则进行转换
	if req.IsClocked == 1 && req.ClockTask != nil {
		// 将请求中的定时任务信息复制到新的ClockTask对象中
		err = copier.Copy(&pbClockTask, req.ClockTask)
		if err != nil {
			// 如果复制过程中出现错误，则返回错误
			return nil, errors.Wrapf(xerr.NewErrMsg("Copy pbClockTask Error"), "Copy pbClockTask Error req: %+v, err : %v ", pbClockTask, err)
		}
	} else {
		// 如果没有定时任务信息，则将定时任务对象设置为nil
		pbClockTask = nil
	}

	// 调用RPC服务添加lottery信息
	addLottery, err := l.svcCtx.LotteryRpc.AddLottery(l.ctx, &lottery.AddLotteryReq{
		UserId:        userId,
		Name:          req.Name,
		Thumb:         req.Thumb,
		AnnounceType:  req.AnnounceType,
		AnnounceTime:  req.AnnounceTime,
		JoinNumber:    req.JoinNumber,
		Introduce:     req.Introduce,
		AwardDeadline: req.AwardDeadline,
		Prizes:        pbPrizes,
		SponsorId:     req.SponsorId,
		IsClocked:     req.IsClocked,
		ClockTask:     pbClockTask,
		PublishType:   req.PublishType,
	})
	if err != nil {
		// 如果添加lottery过程中出现错误，则返回错误
		return nil, errors.Wrapf(xerr.NewErrMsg("AddLottery Error"), "AddLottery Error req: %+v, err : %v ", req, err)
	}

	// 返回创建lottery的响应
	return &types.CreateLotteryResp{
		Id: addLottery.Id,
	}, nil
}
