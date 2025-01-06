package model

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ClockTaskRecordModel = (*customClockTaskRecordModel)(nil)

type (
	// ClockTaskRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClockTaskRecordModel.
	ClockTaskRecordModel interface {
		clockTaskRecordModel
		GetClockTaskRecordByLotteryIdAndUserIds(lotteryId int64, userIds []int64) ([]*ClockTaskRecord, error)
	}

	customClockTaskRecordModel struct {
		*defaultClockTaskRecordModel
	}
)

// NewClockTaskRecordModel returns a model for the database table.
func NewClockTaskRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ClockTaskRecordModel {
	return &customClockTaskRecordModel{
		defaultClockTaskRecordModel: newClockTaskRecordModel(conn, c, opts...),
	}
}

// GetClockTaskRecordByLotteryIdAndUserIds 根据抽奖ID和用户ID列表获取打卡任务记录。
// 参数:
//   lotteryId - 抽奖活动的ID。
//   userIds - 用户ID的切片。
// 返回值:
//   []*ClockTaskRecord - 打卡任务记录的切片。
//   error - 如果查询过程中发生错误，返回该错误。
func (m *customClockTaskRecordModel) GetClockTaskRecordByLotteryIdAndUserIds(lotteryId int64, userIds []int64) ([]*ClockTaskRecord, error) {
    // 如果用户ID列表为空，则直接返回空切片和nil错误。
    if len(userIds) == 0 {
        return nil, nil
    }

    // 构建用户ID的字符串表示，用于SQL查询。
    userIdsStr := ""
    for i, UserId := range userIds {
        if i == 0 {
            userIdsStr = fmt.Sprintf("%d", UserId)
        } else {
            userIdsStr = fmt.Sprintf("%s,%d", userIdsStr, UserId)
        }
    }

    // 构建查询语句，获取指定抽奖活动和用户列表的打卡任务记录。
    query := fmt.Sprintf("select * from %s where lottery_id=? and user_id in (%s)", clockTaskRecordRows, m.table, userIdsStr)

    // 定义变量records来存储查询结果。
    var records []*ClockTaskRecord

    // 使用QueryRowNoCache方法执行查询，并将结果存储到records变量中。
    err := m.QueryRowNoCache(&records, query, lotteryId)
    if err != nil {
        // 如果查询过程中发生错误，返回nil和错误信息。
        return nil, err
    }

    // 返回查询结果和nil错误。
    return records, nil
}

