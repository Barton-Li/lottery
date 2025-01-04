package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"lottery/common/xerr"
	"time"
)

var _ LotteryModel = (*customLotteryModel)(nil)

type (
	// LotteryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryModel.
	LotteryModel interface {
		lotteryModel
		TransUpdateClockTaskId(ctx context.Context, session sqlx.Session, data *Lottery) (sql.Result, error)
		UpdatePublishTime(ctx context.Context, id int64) error
		LotteryList(ctx context.Context, limit, selected, lastId int64) ([]*Lottery, error)
		GetLastId(ctx context.Context) (int64, error)
	}

	customLotteryModel struct {
		*defaultLotteryModel
	}
)

// NewLotteryModel returns a model for the database table.
func NewLotteryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryModel {
	return &customLotteryModel{
		defaultLotteryModel: newLotteryModel(conn, c, opts...),
	}
}

// TransUpdateClockTaskId 更新抽奖的时钟任务ID。
// 该函数使用事务执行更新操作，以确保数据一致性。
// 参数:
//   ctx - 上下文，用于传递请求范围的上下文信息。
//   session - SQL会话，用于执行数据库操作。
//   data - 包含要更新的抽奖信息的结构体指针。
// 返回值:
//   sql.Result - 数据库操作的结果。
//   error - 错误信息，如果操作成功，则为nil。
func (m *defaultLotteryModel) TransUpdateClockTaskId(ctx context.Context, session sqlx.Session, data *Lottery) (sql.Result, error) {
	// 生成缓存键，用于存储或获取抽奖的打卡任务ID。
	lotteryClockTaskIdKey := fmt.Sprintf("%s%v", cacheLotteryClockTaskIdPrefix, data.Id)

	// 执行更新操作，并返回结果和错误信息。
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		// 构建更新语句，更新数据库中的时钟任务ID。
		query := fmt.Sprintf("update %s set  clock_task_id=? where id=?", m.table)
		// 使用传入的session执行更新操作。
		return session.ExecCtx(ctx, query, data.ClockTaskId, data.Id)
	}, lotteryClockTaskIdKey)

	// 返回更新操作的结果和错误信息。
	return ret, err
}

func (m *defaultLotteryModel) UpdatePublishTime(ctx context.Context, id int64) error {
	updateSql := fmt.Sprintf("update %s set  publish_time=? where id=?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, updateSql, time.Now(), id)
	return err
}

// LotteryList 获取抽奖列表
// 该方法根据是否中选的标志(selected)来筛选抽奖列表，主要用于获取待公布中奖结果的抽奖。
// 参数:
//   ctx - 上下文，用于处理请求和响应的生命周期管理。
//   limit - 限制返回的抽奖数量。
//   selected - 筛选标志，非零表示筛选已选中的抽奖，零表示不筛选。
//   lastId - 上一次查询的最后一个抽奖ID，用于分页查询。
// 返回值:
//   []*Lottery - 抽奖列表，每个元素代表一个抽奖的详细信息。
//   error - 错误信息，如果查询过程中发生错误，则返回相应的错误。
func (c *customLotteryModel) LotteryList(ctx context.Context, limit, selected, lastId int64) ([]*Lottery, error) {
	var query string
	// 根据selected参数决定是否在查询中加入已选中(is_selected=1)的条件。
	if selected != 0 {
		query = fmt.Sprintf("select %s from %s where is_selected=1 and is_announced=0 and publish_time is not null and id<? order by id  desc limit ?", lotteryRows, c.table)
	} else {
		query = fmt.Sprintf("select %s from %s where is_announced=0 and publish_time is not null and id<? order by id  desc limit ?", lotteryRows, c.table)
	}
	var resp []*Lottery
	// 执行查询并解析结果为Lottery对象列表。
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lastId, limit)
	if err != nil {
		// 如果查询出错，包装错误并返回。
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_USERID_BYLOTTERYID_ERROR), "get lottery list error: %s", err)
	}
	return resp, nil
}

func (c *customLotteryModel) GetLastId(ctx context.Context) (int64, error) {
	var resp int64
	query := fmt.Sprintf("select id from %s order by id desc limit 1", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_GETLASTID_ERROR), "get last id error: %s", err)
	}
	return resp, nil
}
