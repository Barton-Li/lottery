package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"lottery/common/xerr"
)

var _ LotteryParticipationModel = (*customLotteryParticipationModel)(nil)

type (
	// LotteryParticipationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryParticipationModel.
	LotteryParticipationModel interface {
		lotteryParticipationModel
		CheckIsParticipatedByUserIdAndLotteryId(ctx context.Context, UserId, LotteryId int64) (int64, error)
		GetParticipatedLotteryIdsByUserId(ctx context.Context, UserId int64) ([]int64, error)
	}

	customLotteryParticipationModel struct {
		*defaultLotteryParticipationModel
	}
)

// NewLotteryParticipationModel returns a model for the database table.
func NewLotteryParticipationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryParticipationModel {
	return &customLotteryParticipationModel{
		defaultLotteryParticipationModel: newLotteryParticipationModel(conn, c, opts...),
	}
}

// CheckIsParticipatedByUserIdAndLotteryId 检查用户是否已参与特定的抽奖活动。
// 该方法通过用户ID和抽奖活动ID来确定用户是否已经参与了某个抽奖活动。
// 参数:
//   ctx - 上下文，用于处理请求和传递请求范围的值。
//   UserId - 用户的唯一标识符。
//   LotteryId - 抽奖活动的唯一标识符。
// 返回值:
//   int64 - 如果用户已参与抽奖，返回参与记录的数量；未参与则返回0。
//   error - 如果查询过程中发生错误，返回错误信息。
func (m *customLotteryParticipationModel) CheckIsParticipatedByUserIdAndLotteryId(ctx context.Context, UserId, LotteryId int64) (int64, error) {
	// 构造SQL查询语句，用于检查用户是否已参与特定的抽奖活动。
	query := fmt.Sprintf("select count(*) from %s where user_id=? and lottery_id=?", m.table)
	// 初始化变量resp以存储查询结果。
	var resp int64

	// 执行查询，检查用户是否已参与特定的抽奖活动。
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, UserId, LotteryId)
	// 如果查询过程中发生错误，处理错误并返回。
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.CHECK_ISPARTICIPATED_BYUSERID_ANDLOTTERYID_ERROR), "CheckIsParticipatedByUserIdAndLotteryId error, userId: %d, lotteryId: %d, err: %v", UserId, LotteryId, err)
	}

	// 返回查询结果，即用户是否已参与抽奖的记录数。
	return resp, nil
}

func (m *customLotteryParticipationModel) GetParticipatedLotteryIdsByUserId(ctx context.Context, UserId int64) ([]int64, error) {
	query := fmt.Sprintf("select lottery_id from %s where user_id=?", m.table)
	var resp []int64
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GET_PARTICIPATED_LOTTERYIDS_BYUSERID_ERROR), "GetParticipatedLotteryIdsByUserId error, userId: %d, err: %v", UserId, err)
	}
	return resp, nil
}
