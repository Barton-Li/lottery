package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/model"
	"lottery/common/xerr"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryByIdLogic {
	return &GetLotteryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryByIdLogic) GetLotteryById(in *pb.GetLotteryByIdReq) (*pb.GetLotteryByIdResp, error) {
	lottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.Id)
	if err != nil || err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询抽奖活动失败:%d,信息失败, err:%v", in.Id, err)
	}
	pbLottery := new(pb.Lottery)
	_ = copier.Copy(pbLottery, lottery)
	return &pb.GetLotteryByIdResp{Lottery: pbLottery}, nil
}
