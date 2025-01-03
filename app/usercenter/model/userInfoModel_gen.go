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
	userInfoFieldNames          = builder.RawFieldNames(&UserInfo{})
	userInfoRows                = strings.Join(userInfoFieldNames, ",")
	userInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(userInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(userInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUsercenterUserInfoIdPrefix     = "cache:usercenter:userInfo:id:"
	cacheUsercenterUserInfoMobilePrefix = "cache:usercenter:userInfo:mobile:"
)

type (
	userInfoModel interface {
		Insert(ctx context.Context, data *UserInfo) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *UserInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserInfo, error)
		FindOneByMobile(ctx context.Context, mobile string) (*UserInfo, error)
		Update(ctx context.Context, data *UserInfo) error
		List(ctx context.Context, page, limit int64) ([]*UserInfo, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *UserInfo) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*UserInfo, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserInfo, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserInfo, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UserInfo, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UserInfo, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultUserInfoModel struct {
		sqlc.CachedConn
		table string
	}

	UserInfo struct {
		Id               int64        `db:"id"`                // 主键，用户唯一标识符，自动递增
		CreateTime       time.Time    `db:"create_time"`       // 用户信息创建时间，默认为当前时间
		UpdateTime       time.Time    `db:"update_time"`       // 用户信息最后更新时间，默认为当前时间，并且在每次更新时自动更新为当前时间
		DeleteTime       sql.NullTime `db:"delete_time"`       // 用户信息删除时间，默认为空
		DelState         int64        `db:"del_state"`         // 删除状态，0表示未删除，1表示已删除
		Version          int64        `db:"version"`           // 版本号，用于数据版本控制
		Mobile           string       `db:"mobile"`            // 用户手机号码，默认为空字符串
		Password         string       `db:"password"`          // 用户密码，默认为空字符串
		Nickname         string       `db:"nickname"`          // 用户昵称，默认为空字符串
		Sex              int64        `db:"sex"`               // 性别，0表示男，1表示女
		Avatar           string       `db:"avatar"`            // 用户头像，默认为"头像"字符串
		Info             string       `db:"info"`              // 用户其他信息，默认为空字符串
		IsAdmin          int64        `db:"is_admin"`          // 是否是管理员，1表示是管理员，0表示不是管理员
		Signature        string       `db:"signature"`         // 个性签名，默认为空字符串
		LocationName     string       `db:"location_name"`     // 地址名称，默认为空字符串
		Longitude        float64      `db:"longitude"`         // 经度，默认为0
		Latitude         float64      `db:"latitude"`          // 纬度，默认为0
		TotalPrize       int64        `db:"total_prize"`       // 累计奖品数量，默认为0
		Fans             int64        `db:"fans"`              // 粉丝数量，默认为0
		AllLottery       int64        `db:"all_lottery"`       // 参与或发起的全部抽奖活动数量，默认为0
		InitiationRecord int64        `db:"initiation_record"` // 发起的抽奖活动数量，默认为0
		WinningRecord    int64        `db:"winning_record"`    // 中奖的记录数量，默认为0
	}
)

func newUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserInfoModel {
	return &defaultUserInfoModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_info`",
	}
}

func (m *defaultUserInfoModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	usercenterUserInfoIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserInfoIdPrefix, id)
	usercenterUserInfoMobileKey := fmt.Sprintf("%s%v", cacheUsercenterUserInfoMobilePrefix, data.Mobile)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, usercenterUserInfoIdKey, usercenterUserInfoMobileKey)
	return err
}

func (m *defaultUserInfoModel) FindOne(ctx context.Context, id int64) (*UserInfo, error) {
	usercenterUserInfoIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserInfoIdPrefix, id)
	var resp UserInfo
	err := m.QueryRowCtx(ctx, &resp, usercenterUserInfoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userInfoRows, m.table)
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

func (m *defaultUserInfoModel) FindOneByMobile(ctx context.Context, mobile string) (*UserInfo, error) {
	usercenterUserInfoMobileKey := fmt.Sprintf("%s%v", cacheUsercenterUserInfoMobilePrefix, mobile)
	var resp UserInfo
	err := m.QueryRowIndexCtx(ctx, &resp, usercenterUserInfoMobileKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", userInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, mobile); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserInfoModel) Insert(ctx context.Context, data *UserInfo) (sql.Result, error) {
	usercenterUserInfoIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserInfoIdPrefix, data.Id)
	usercenterUserInfoMobileKey := fmt.Sprintf("%s%v", cacheUsercenterUserInfoMobilePrefix, data.Mobile)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info, data.IsAdmin, data.Signature, data.LocationName, data.Longitude, data.Latitude, data.TotalPrize, data.Fans, data.AllLottery, data.InitiationRecord, data.WinningRecord)
	}, usercenterUserInfoIdKey, usercenterUserInfoMobileKey)
	return ret, err
}

func (m *defaultUserInfoModel) TransInsert(ctx context.Context, session sqlx.Session, data *UserInfo) (sql.Result, error) {
	usercenterUserInfoIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserInfoIdPrefix, data.Id)
	usercenterUserInfoMobileKey := fmt.Sprintf("%s%v", cacheUsercenterUserInfoMobilePrefix, data.Mobile)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userInfoRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info, data.IsAdmin, data.Signature, data.LocationName, data.Longitude, data.Latitude, data.TotalPrize, data.Fans, data.AllLottery, data.InitiationRecord, data.WinningRecord)
	}, usercenterUserInfoIdKey, usercenterUserInfoMobileKey)
	return ret, err
}
func (m *defaultUserInfoModel) Update(ctx context.Context, newData *UserInfo) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	usercenterUserInfoIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserInfoIdPrefix, data.Id)
	usercenterUserInfoMobileKey := fmt.Sprintf("%s%v", cacheUsercenterUserInfoMobilePrefix, data.Mobile)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.Mobile, newData.Password, newData.Nickname, newData.Sex, newData.Avatar, newData.Info, newData.IsAdmin, newData.Signature, newData.LocationName, newData.Longitude, newData.Latitude, newData.TotalPrize, newData.Fans, newData.AllLottery, newData.InitiationRecord, newData.WinningRecord, newData.Id)
	}, usercenterUserInfoIdKey, usercenterUserInfoMobileKey)
	return err
}

func (m *defaultUserInfoModel) TransUpdate(ctx context.Context, session sqlx.Session, newData *UserInfo) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	usercenterUserInfoIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserInfoIdPrefix, data.Id)
	usercenterUserInfoMobileKey := fmt.Sprintf("%s%v", cacheUsercenterUserInfoMobilePrefix, data.Mobile)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userInfoRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.Mobile, newData.Password, newData.Nickname, newData.Sex, newData.Avatar, newData.Info, newData.IsAdmin, newData.Signature, newData.LocationName, newData.Longitude, newData.Latitude, newData.TotalPrize, newData.Fans, newData.AllLottery, newData.InitiationRecord, newData.WinningRecord, newData.Id)
	}, usercenterUserInfoIdKey, usercenterUserInfoMobileKey)
	return err
}

func (m *defaultUserInfoModel) List(ctx context.Context, page, limit int64) ([]*UserInfo, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", userInfoRows, m.table)
	var resp []*UserInfo
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultUserInfoModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultUserInfoModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultUserInfoModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultUserInfoModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*UserInfo, error) {

	builder = builder.Columns(userInfoRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserInfoModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserInfo, error) {

	builder = builder.Columns(userInfoRows)

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

	var resp []*UserInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserInfoModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserInfo, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(userInfoRows)

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

	var resp []*UserInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultUserInfoModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UserInfo, error) {

	builder = builder.Columns(userInfoRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserInfoModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UserInfo, error) {

	builder = builder.Columns(userInfoRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserInfoModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultUserInfoModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUsercenterUserInfoIdPrefix, primary)
}

func (m *defaultUserInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserInfoModel) tableName() string {
	return m.table
}