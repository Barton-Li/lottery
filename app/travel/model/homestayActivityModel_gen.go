// Code generated by goctl. DO NOT EDIT!

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
	homestayActivityFieldNames          = builder.RawFieldNames(&HomestayActivity{})
	homestayActivityRows                = strings.Join(homestayActivityFieldNames, ",")
	homestayActivityRowsExpectAutoSet   = strings.Join(stringx.Remove(homestayActivityFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	homestayActivityRowsWithPlaceHolder = strings.Join(stringx.Remove(homestayActivityFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cachelotteryTravelHomestayActivityIdPrefix = "cache:lotteryTravel:homestayActivity:id:"
)

type (
	homestayActivityModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *HomestayActivity) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*HomestayActivity, error)
		Update(ctx context.Context, session sqlx.Session, data *HomestayActivity) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *HomestayActivity) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *HomestayActivity) error
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*HomestayActivity, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayActivity, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayActivity, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayActivity, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HomestayActivity, error)
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultHomestayActivityModel struct {
		sqlc.CachedConn
		table string
	}

	HomestayActivity struct {
		Id         int64     `db:"id"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		DeleteTime time.Time `db:"delete_time"`
		DelState   int64     `db:"del_state"`
		RowType    string    `db:"row_type"`   // 活动类型
		DataId     int64     `db:"data_id"`    // 业务表id（id跟随活动类型走）
		RowStatus  int64     `db:"row_status"` // 0:下架 1:上架
		Version    int64     `db:"version"`    // 版本号
	}
)

func newHomestayActivityModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultHomestayActivityModel {
	return &defaultHomestayActivityModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`homestay_activity`",
	}
}

func (m *defaultHomestayActivityModel) Insert(ctx context.Context, session sqlx.Session, data *HomestayActivity) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo
	lotteryTravelHomestayActivityIdKey := fmt.Sprintf("%s%v", cachelotteryTravelHomestayActivityIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, homestayActivityRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.RowType, data.DataId, data.RowStatus, data.Version)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.RowType, data.DataId, data.RowStatus, data.Version)
	}, lotteryTravelHomestayActivityIdKey)
}

func (m *defaultHomestayActivityModel) FindOne(ctx context.Context, id int64) (*HomestayActivity, error) {
	lotteryTravelHomestayActivityIdKey := fmt.Sprintf("%s%v", cachelotteryTravelHomestayActivityIdPrefix, id)
	var resp HomestayActivity
	err := m.QueryRowCtx(ctx, &resp, lotteryTravelHomestayActivityIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", homestayActivityRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
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

func (m *defaultHomestayActivityModel) Update(ctx context.Context, session sqlx.Session, data *HomestayActivity) (sql.Result, error) {
	lotteryTravelHomestayActivityIdKey := fmt.Sprintf("%s%v", cachelotteryTravelHomestayActivityIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, homestayActivityRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.RowType, data.DataId, data.RowStatus, data.Version, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.RowType, data.DataId, data.RowStatus, data.Version, data.Id)
	}, lotteryTravelHomestayActivityIdKey)
}

func (m *defaultHomestayActivityModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *HomestayActivity) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	lotteryTravelHomestayActivityIdKey := fmt.Sprintf("%s%v", cachelotteryTravelHomestayActivityIdPrefix, data.Id)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, homestayActivityRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.RowType, data.DataId, data.RowStatus, data.Version, data.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.RowType, data.DataId, data.RowStatus, data.Version, data.Id, oldVersion)
	}, lotteryTravelHomestayActivityIdKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return ErrNoRowsUpdate
	}

	return nil
}

func (m *defaultHomestayActivityModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *HomestayActivity) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "HomestayActivityModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultHomestayActivityModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultHomestayActivityModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultHomestayActivityModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*HomestayActivity, error) {

	builder = builder.Columns(homestayActivityRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayActivity
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayActivityModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayActivity, error) {

	builder = builder.Columns(homestayActivityRows)

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

	var resp []*HomestayActivity
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayActivityModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*HomestayActivity, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(homestayActivityRows)

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

	var resp []*HomestayActivity
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultHomestayActivityModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*HomestayActivity, error) {

	builder = builder.Columns(homestayActivityRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayActivity
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayActivityModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*HomestayActivity, error) {

	builder = builder.Columns(homestayActivityRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*HomestayActivity
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultHomestayActivityModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultHomestayActivityModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
func (m *defaultHomestayActivityModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	lotteryTravelHomestayActivityIdKey := fmt.Sprintf("%s%v", cachelotteryTravelHomestayActivityIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, lotteryTravelHomestayActivityIdKey)
	return err
}
func (m *defaultHomestayActivityModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachelotteryTravelHomestayActivityIdPrefix, primary)
}
func (m *defaultHomestayActivityModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", homestayActivityRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultHomestayActivityModel) tableName() string {
	return m.table
}
