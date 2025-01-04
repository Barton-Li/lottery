package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"lottery/app/usercenter/cmd/rpc/usercenter"
	"lottery/common/xerr"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetIsSelectedLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetIsSelectedLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetIsSelectedLotteryLogic {
	return &SetIsSelectedLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SetIsSelectedLottery 是一个用于设置抽奖是否被选中的方法
// 它接受一个SetIsSelectedLotteryReq请求对象，并返回一个SetIsSelectedLotteryResp响应对象和一个错误对象
// 此方法主要用于更新数据库中特定抽奖的IsSelected字段，以标记抽奖是否被选中
func (l *SetIsSelectedLotteryLogic) SetIsSelectedLottery(in *pb.SetIsSelectedLotteryReq) (*pb.SetIsSelectedLotteryResp, error) {
	// 定义一个变量来存储抽奖是否被选中的状态
	var IsSelected int64

	// 使用事务来处理数据库操作，确保数据的一致性
	err := l.svcCtx.LotteryModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 通过RPC调用获取用户信息
		userInfo, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
			Id: in.UserId,
		})
		if err != nil {
			return err
		}

		// 检查用户是否是管理员
		if userInfo.User.IsAdmin != 0 {
			// 根据抽奖ID查询抽奖信息
			lottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.Id)
			if err != nil {
				// 如果查询失败，返回错误
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_LOTTERY_BYLOTTERYID_ERROR), "查询lottery信息失败err:%v,id:%d", err, in.Id)
			}

			// 根据抽奖的当前状态更新IsSelected字段
			if lottery.IsSelected == 0 {
				lottery.IsSelected = 1
				IsSelected = 1
			} else {
				lottery.IsSelected = 0
			}

			// 更新数据库中的抽奖信息
			err = l.svcCtx.LotteryModel.Update(l.ctx, lottery)
			if err != nil {
				// 如果更新失败，返回错误
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_LOTTERY_ERROR), "更新lottery信息失败err:%v, id:%d", err, in.Id)
			}
		} else {
			// 如果用户不是管理员，返回没有权限的错误
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_NO_SET_LOTTERY_ISSELECT_PERMISSION_ERROR), "没有设置中奖权限,id:%d,is_admin:%d,err:%v", userInfo.User.Id, userInfo.User.IsAdmin, err)
		}

		// 如果一切正常，返回nil
		return nil
	})

	// 如果事务中有任何错误，返回错误
	if err != nil {
		return nil, err
	}

	// 返回响应对象，包含是否被选中的状态
	return &pb.SetIsSelectedLotteryResp{
		IsSelected: IsSelected,
	}, nil
}
