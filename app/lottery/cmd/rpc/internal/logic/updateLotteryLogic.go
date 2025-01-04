package logic

import (
	"context"
	"github.com/pkg/errors"
	"lottery/common/xerr"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLotteryLogic {
	return &UpdateLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLotteryLogic) UpdateLottery(in *pb.UpdateLotteryReq) (*pb.UpdateLotteryResp, error) {
	lottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.Id)
	if err!=nil{
		return nil,errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindOne error: id(%d), err: %v", in.Id, err)
	}
	if lottery.UserId!=in.UserId{
		return nil,errors.Wrapf(xerr.NewErrMsg("不是本人操作"), "不是本人操作")
	}
	err = l.svcCtx.LotteryModel.UpdatePublishTime(l.ctx, in.Id)
	if err!=nil{
		return nil,errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "UpdatePublishTime error: id(%d), err: %v", in.Id, err)
	}
	return &pb.UpdateLotteryResp{}, nil
}
