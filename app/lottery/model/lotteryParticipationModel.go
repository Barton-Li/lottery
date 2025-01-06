package model

import (
	"context"
	"database/sql"
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
		GetParticipationUserIdsByLotteryId(ctx context.Context, LotteryId int64) ([]int64, error)
		UpdateWinners(ctx context.Context, LotteryId, UserId, PrizeId int64) error
		GetParticipatorsCountByLotteryId(ctx context.Context, LotteryId int64) (int64, error)
		GetWonListCountByUserId(ctx context.Context, UserId int64) (int64, error)
		GetWonListByUserId(ctx context.Context, UserId, Size, LastId int64) ([]*LotteryParticipation, error)
		CheckIsWonByUserIdAndLotteryId(ctx context.Context, LotteryId, UserId int64) (int64, error)
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
func (m *defaultLotteryParticipationModel) GetWonListCountByUserId(ctx context.Context, UserId int64) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where user_id=? and is_won=1", m.table)
	var resp int64
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, UserId)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.GET_WONLISTCOUNT_BYUSERID_ERROR), "GetWonListCountByUserId error, userId: %d, err: %v", UserId, err)
	}
	return resp, nil
}
func (m *defaultLotteryParticipationModel) GetWonListByUserId(ctx context.Context, UserId, Size, LastId int64) ([]*LotteryParticipation, error) {
	query := fmt.Sprintf("select * from %s where user_id=? and is_won=1 and id>? limit ?", m.table)
	var resp []*LotteryParticipation
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, UserId, LastId, Size)
	if err == sqlx.ErrNotFound {
		fmt.Println("未找到数据", err)
		return resp, nil
	}
	if err != nil {
		fmt.Println("查询数据失败", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GET_WONLIST_BYUSERID_ERROR), "GetWonListByUserId, UserId:%v, Size:%v, error: %v", UserId, Size, err)
	}
	return resp, nil
}
func (m *defaultLotteryParticipationModel) CheckIsWonByUserIdAndLotteryId(ctx context.Context, LotteryId, UserId int64) (int64, error) {
	query := fmt.Sprintf("select is_won from %s where lottery_id =? and user_id=?", m.table)
	var resp int64
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, LotteryId, UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.CHECK_ISWON_BYUSERID_ANDLOTTERYID_ERROR), "CheckIsWonByUserIdAndLotteryId, LotteryId:%v, UserId:%v, error: %v", LotteryId, UserId, err)
	}
	return resp, nil
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

// GetParticipationUserIdsByLotteryId 根据抽奖活动ID获取参与用户的ID列表。
// 该方法主要用于查询特定抽奖活动中所有参与用户的ID。
// 参数:
//   ctx - 上下文，用于传递请求范围的上下文信息。
//   LotteryId - 抽奖活动的唯一标识符。
// 返回值:
//   成功时返回参与用户的ID列表，失败时返回错误。
func (m *customLotteryParticipationModel) GetParticipationUserIdsByLotteryId(ctx context.Context, LotteryId int64) ([]int64, error) {
	// 构造SQL查询语句，用于从数据库中获取指定抽奖活动的参与用户ID列表。
	query := fmt.Sprintf("select user_id from %s where lottery_id=?", m.table)

	// 初始化一个整型切片用于存储查询结果。
	var resp []int64

	// 调用QueryRowsNoCacheCtx方法执行SQL查询，该方法不会使用缓存的数据。
	// 如果查询出错，将包装错误信息并返回。
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, LotteryId)
	if err != nil {
		// 当发生错误时，使用xerr.NewErrCode创建一个自定义错误，
		// 并使用errors.Wrapf包装错误，添加更多的错误上下文信息。
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GET_PARTICIPATION_USERIDS_BYLOTTERYID_ERROR), "GetParticipationUserIdsByLotteryId error, lotteryId: %d, err: %v", LotteryId, err)
	}

	// 如果查询成功，返回参与用户的ID列表。
	return resp, nil
}

// UpdateWinners 更新抽奖活动的中奖者信息。
// 该函数通过设置用户的奖品ID来标记用户是否中奖。
// 参数:
//   ctx - 上下文，用于传递请求范围的信息。
//   LotteryId - 抽奖活动的ID。
//   UserId - 参与抽奖的用户的ID。
//   PrizeId - 中奖的奖品ID。
// 返回值:
//   如果更新操作失败，返回错误。
func (m *customLotteryParticipationModel) UpdateWinners(ctx context.Context, LotteryId, UserId, PrizeId int64) error {
	// 查询特定抽奖活动和用户的相关数据。
	data, err := m.FindOneByLotteryIdUserId(ctx, LotteryId, UserId)
	if err != nil {
		return err
	}

	// 生成缓存键，用于后续更新缓存。
	goZeroLotteryParticipationIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryParticipationIdPrefix, data.Id)
	gozeroLotteryParticipationLotteryIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheLotteryLotteryParticipationLotteryIdUserIdPrefix, data.LotteryId, data.UserId)

	// 构建更新语句，标记用户为中奖并指定奖品ID。
	query := fmt.Sprintf("update %s set is_won=1,prize_id=? where lottery_id=?and user_id=?", m.table)

	// 执行更新操作，并传入必要的参数。
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		res, err := conn.ExecCtx(ctx, query, PrizeId, LotteryId, UserId)
		if err != nil {
			return nil, err
		}
		return res, nil
	}, goZeroLotteryParticipationIdKey, gozeroLotteryParticipationLotteryIdUserIdKey)
	if err != nil {
		// 如果更新操作失败，包装错误并返回。
		return errors.Wrapf(xerr.NewErrCode(xerr.UPDATE_WINNER_ERROR), "UpdateWinners error, lotteryId: %d, userId: %d, prizeId: %d, err: %v", LotteryId, UserId, PrizeId, err)
	}
	return nil
}

func (m *customLotteryParticipationModel) GetParticipatorsCountByLotteryId(ctx context.Context, LotteryId int64) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where lottery_id=?", m.table)
	var resp int64
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, LotteryId)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.GET_PARTICIPATORS_COUNT_BYLOTTERYID_ERROR), "GetParticipatorsCountByLotteryId error, lotteryId: %d, err: %v", LotteryId, err)
	}
	return resp, nil
}
