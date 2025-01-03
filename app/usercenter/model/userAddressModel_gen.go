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
	userAddressFieldNames          = builder.RawFieldNames(&UserAddress{})
	userAddressRows                = strings.Join(userAddressFieldNames, ",")
	userAddressRowsExpectAutoSet   = strings.Join(stringx.Remove(userAddressFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userAddressRowsWithPlaceHolder = strings.Join(stringx.Remove(userAddressFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUsercenterUserAddressIdPrefix = "cache:usercenter:userAddress:id:"
)

type (
	userAddressModel interface {
		Insert(ctx context.Context, data *UserAddress) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *UserAddress) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserAddress, error)
		Update(ctx context.Context, data *UserAddress) error
		List(ctx context.Context, page, limit int64) ([]*UserAddress, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *UserAddress) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*UserAddress, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserAddress, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserAddress, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UserAddress, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UserAddress, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultUserAddressModel struct {
		sqlc.CachedConn
		table string
	}

	UserAddress struct {
		Id            int64     `db:"id"`             // 主键，地址记录唯一标识符，自动递增
		UserId        int64     `db:"user_id"`        // 用户ID，关联到user_info表中的id
		ContactName   string    `db:"contact_name"`   // 联系人姓名，默认为空字符串
		ContactMobile string    `db:"contact_mobile"` // 联系人手机号码，默认为空字符串
		District      string    `db:"district"`       // 地区信息，以JSON格式存储
		Detail        string    `db:"detail"`         // 详细地址，默认为空字符串
		Postcode      string    `db:"postcode"`       // 邮政编码，默认为空字符串
		IsDefault     int64     `db:"is_default"`     // 是否为默认地址，1表示是默认地址，0表示不是默认地址
		CreateTime    time.Time `db:"create_time"`    // 地址记录创建时间，默认为当前时间
		UpdateTime    time.Time `db:"update_time"`    // 地址记录最后更新时间，默认为当前时间，并且在每次更新时自动更新为当前时间
	}
)

func newUserAddressModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserAddressModel {
	return &defaultUserAddressModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_address`",
	}
}

func (m *defaultUserAddressModel) Delete(ctx context.Context, id int64) error {
	usercenterUserAddressIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserAddressIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, usercenterUserAddressIdKey)
	return err
}

func (m *defaultUserAddressModel) FindOne(ctx context.Context, id int64) (*UserAddress, error) {
	usercenterUserAddressIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserAddressIdPrefix, id)
	var resp UserAddress
	err := m.QueryRowCtx(ctx, &resp, usercenterUserAddressIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userAddressRows, m.table)
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

func (m *defaultUserAddressModel) Insert(ctx context.Context, data *UserAddress) (sql.Result, error) {
	usercenterUserAddressIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserAddressIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, userAddressRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ContactName, data.ContactMobile, data.District, data.Detail, data.Postcode, data.IsDefault)
	}, usercenterUserAddressIdKey)
	return ret, err
}

func (m *defaultUserAddressModel) TransInsert(ctx context.Context, session sqlx.Session, data *UserAddress) (sql.Result, error) {
	usercenterUserAddressIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserAddressIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, userAddressRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.UserId, data.ContactName, data.ContactMobile, data.District, data.Detail, data.Postcode, data.IsDefault)
	}, usercenterUserAddressIdKey)
	return ret, err
}
func (m *defaultUserAddressModel) Update(ctx context.Context, data *UserAddress) error {
	usercenterUserAddressIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserAddressIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userAddressRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.ContactName, data.ContactMobile, data.District, data.Detail, data.Postcode, data.IsDefault, data.Id)
	}, usercenterUserAddressIdKey)
	return err
}

func (m *defaultUserAddressModel) TransUpdate(ctx context.Context, session sqlx.Session, data *UserAddress) error {
	usercenterUserAddressIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserAddressIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userAddressRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.UserId, data.ContactName, data.ContactMobile, data.District, data.Detail, data.Postcode, data.IsDefault, data.Id)
	}, usercenterUserAddressIdKey)
	return err
}

func (m *defaultUserAddressModel) List(ctx context.Context, page, limit int64) ([]*UserAddress, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", userAddressRows, m.table)
	var resp []*UserAddress
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultUserAddressModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultUserAddressModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultUserAddressModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultUserAddressModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*UserAddress, error) {

	builder = builder.Columns(userAddressRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserAddress
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserAddressModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserAddress, error) {

	builder = builder.Columns(userAddressRows)

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

	var resp []*UserAddress
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserAddressModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserAddress, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(userAddressRows)

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

	var resp []*UserAddress
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultUserAddressModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UserAddress, error) {

	builder = builder.Columns(userAddressRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserAddress
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserAddressModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UserAddress, error) {

	builder = builder.Columns(userAddressRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserAddress
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserAddressModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultUserAddressModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUsercenterUserAddressIdPrefix, primary)
}

func (m *defaultUserAddressModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userAddressRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserAddressModel) tableName() string {
	return m.table
}
