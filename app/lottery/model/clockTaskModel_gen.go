// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"lottery/common/globalkey"
)

var (
	clockTaskFieldNames          = builder.RawFieldNames(&ClockTask{})
	clockTaskRows                = strings.Join(clockTaskFieldNames, ",")
	clockTaskRowsExpectAutoSet   = strings.Join(stringx.Remove(clockTaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	clockTaskRowsWithPlaceHolder = strings.Join(stringx.Remove(clockTaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLotteryClockTaskIdPrefix = "cache:lottery:clockTask:id:"
)

type (
	clockTaskModel interface {
		Insert(ctx context.Context, data *ClockTask) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *ClockTask) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ClockTask, error)
		Update(ctx context.Context, data *ClockTask) error
		List(ctx context.Context, page, limit int64) ([]*ClockTask, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *ClockTask) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*ClockTask, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*ClockTask, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*ClockTask, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*ClockTask, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*ClockTask, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultClockTaskModel struct {
		sqlc.CachedConn
		table string
	}

	ClockTask struct {
		Id               int64     `db:"id"`
		LotteryId        int64     `db:"lottery_id"`        // 抽奖ID
		Type             int64     `db:"type"`              // 任务类型 1: 体验小程序 2： 浏览指定公众号文章 3: 浏览图片（微信图片二维码等） 4： 浏览视频号视频
		Seconds          int64     `db:"seconds"`           // 任务秒数
		AppletType       int64     `db:"applet_type"`       // type=1时该字段才有意义，1小程序链接 2小程序路径
		PageLink         string    `db:"page_link"`         // type=1 并且 applet_type=1时 该字段才有意义，配置要跳转小程序的页面链接 （如 #小程序://抽奖/oM....）
		AppId            string    `db:"app_id"`            // type=1 并且 applet_type=2时 该字段才有意义，配置要跳转的小程序AppID
		PagePath         string    `db:"page_path"`         // type=1 并且 applet_type=2时 该字段才有意义，配置要跳转的小程序路径（如：/pages/index）
		Image            string    `db:"image"`             // type=3时 该字段才有意义，添加要查看的图片
		VideoAccountId   string    `db:"video_account_id"`  // type=4时 该字段才有意义，视频号ID
		VideoId          string    `db:"video_id"`          // type=4时 该字段才有意义，视频ID
		ArticleLink      string    `db:"article_link"`      // type=2时 该字段才有意义，公众号文章链接
		Copywriting      string    `db:"copywriting"`       // 引导参与者完成打卡任务的文案
		ChanceType       int64     `db:"chance_type"`       // 概率类型 1: 随机 2: 指定
		IncreaseMultiple int64     `db:"increase_multiple"` // chance_type=2时 该字段才有意义，概率增加倍数
		CreateTime       time.Time `db:"create_time"`
		UpdateTime       time.Time `db:"update_time"`
	}
)

func newClockTaskModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultClockTaskModel {
	return &defaultClockTaskModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`clock_task`",
	}
}

func (m *defaultClockTaskModel) Delete(ctx context.Context, id int64) error {
	lotteryClockTaskIdKey := fmt.Sprintf("%s%v", cacheLotteryClockTaskIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, lotteryClockTaskIdKey)
	return err
}

func (m *defaultClockTaskModel) FindOne(ctx context.Context, id int64) (*ClockTask, error) {
	lotteryClockTaskIdKey := fmt.Sprintf("%s%v", cacheLotteryClockTaskIdPrefix, id)
	var resp ClockTask
	err := m.QueryRowCtx(ctx, &resp, lotteryClockTaskIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", clockTaskRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultClockTaskModel) Insert(ctx context.Context, data *ClockTask) (sql.Result, error) {
	lotteryClockTaskIdKey := fmt.Sprintf("%s%v", cacheLotteryClockTaskIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, clockTaskRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.LotteryId, data.Type, data.Seconds, data.AppletType, data.PageLink, data.AppId, data.PagePath, data.Image, data.VideoAccountId, data.VideoId, data.ArticleLink, data.Copywriting, data.ChanceType, data.IncreaseMultiple)
	}, lotteryClockTaskIdKey)
	return ret, err
}

func (m *defaultClockTaskModel) TransInsert(ctx context.Context, session sqlx.Session, data *ClockTask) (sql.Result, error) {
	lotteryClockTaskIdKey := fmt.Sprintf("%s%v", cacheLotteryClockTaskIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, clockTaskRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.LotteryId, data.Type, data.Seconds, data.AppletType, data.PageLink, data.AppId, data.PagePath, data.Image, data.VideoAccountId, data.VideoId, data.ArticleLink, data.Copywriting, data.ChanceType, data.IncreaseMultiple)
	}, lotteryClockTaskIdKey)
	return ret, err
}
func (m *defaultClockTaskModel) Update(ctx context.Context, data *ClockTask) error {
	lotteryClockTaskIdKey := fmt.Sprintf("%s%v", cacheLotteryClockTaskIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, clockTaskRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.LotteryId, data.Type, data.Seconds, data.AppletType, data.PageLink, data.AppId, data.PagePath, data.Image, data.VideoAccountId, data.VideoId, data.ArticleLink, data.Copywriting, data.ChanceType, data.IncreaseMultiple, data.Id)
	}, lotteryClockTaskIdKey)
	return err
}

func (m *defaultClockTaskModel) TransUpdate(ctx context.Context, session sqlx.Session, data *ClockTask) error {
	lotteryClockTaskIdKey := fmt.Sprintf("%s%v", cacheLotteryClockTaskIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, clockTaskRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.LotteryId, data.Type, data.Seconds, data.AppletType, data.PageLink, data.AppId, data.PagePath, data.Image, data.VideoAccountId, data.VideoId, data.ArticleLink, data.Copywriting, data.ChanceType, data.IncreaseMultiple, data.Id)
	}, lotteryClockTaskIdKey)
	return err
}

func (m *defaultClockTaskModel) List(ctx context.Context, page, limit int64) ([]*ClockTask, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", clockTaskRows, m.table)
	var resp []*ClockTask
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultClockTaskModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultClockTaskModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultClockTaskModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultClockTaskModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*ClockTask, error) {

	builder = builder.Columns(clockTaskRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ClockTask
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultClockTaskModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*ClockTask, error) {

	builder = builder.Columns(clockTaskRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ClockTask
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultClockTaskModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*ClockTask, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(clockTaskRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*ClockTask
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultClockTaskModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*ClockTask, error) {

	builder = builder.Columns(clockTaskRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ClockTask
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultClockTaskModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*ClockTask, error) {

	builder = builder.Columns(clockTaskRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ClockTask
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultClockTaskModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultClockTaskModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLotteryClockTaskIdPrefix, primary)
}

func (m *defaultClockTaskModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", clockTaskRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultClockTaskModel) tableName() string {
	return m.table
}
