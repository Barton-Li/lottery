// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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
	goodsCategoryFieldNames          = builder.RawFieldNames(&GoodsCategory{})
	goodsCategoryRows                = strings.Join(goodsCategoryFieldNames, ",")
	goodsCategoryRowsExpectAutoSet   = strings.Join(stringx.Remove(goodsCategoryFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	goodsCategoryRowsWithPlaceHolder = strings.Join(stringx.Remove(goodsCategoryFieldNames, "`category_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheShopGoodsCategoryCategoryIdPrefix = "cache:shop:goodsCategory:categoryId:"
)

type (
	goodsCategoryModel interface {
		Insert(ctx context.Context, data *GoodsCategory) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *GoodsCategory) (sql.Result, error)
		FindOne(ctx context.Context, categoryId int64) (*GoodsCategory, error)
		Update(ctx context.Context, data *GoodsCategory) error
		List(ctx context.Context, page, limit int64) ([]*GoodsCategory, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *GoodsCategory) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*GoodsCategory, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*GoodsCategory, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*GoodsCategory, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*GoodsCategory, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*GoodsCategory, error)
		Delete(ctx context.Context, categoryId int64) error
	}

	defaultGoodsCategoryModel struct {
		sqlc.CachedConn
		table string
	}

	GoodsCategory struct {
		CategoryId   int64  `db:"category_id"`
		CategoryName string `db:"category_name"`
	}
)

func newGoodsCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultGoodsCategoryModel {
	return &defaultGoodsCategoryModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`goods_category`",
	}
}

func (m *defaultGoodsCategoryModel) Delete(ctx context.Context, categoryId int64) error {
	shopGoodsCategoryCategoryIdKey := fmt.Sprintf("%s%v", cacheShopGoodsCategoryCategoryIdPrefix, categoryId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `category_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, categoryId)
	}, shopGoodsCategoryCategoryIdKey)
	return err
}

func (m *defaultGoodsCategoryModel) FindOne(ctx context.Context, categoryId int64) (*GoodsCategory, error) {
	shopGoodsCategoryCategoryIdKey := fmt.Sprintf("%s%v", cacheShopGoodsCategoryCategoryIdPrefix, categoryId)
	var resp GoodsCategory
	err := m.QueryRowCtx(ctx, &resp, shopGoodsCategoryCategoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `category_id` = ? limit 1", goodsCategoryRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, categoryId)
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

func (m *defaultGoodsCategoryModel) Insert(ctx context.Context, data *GoodsCategory) (sql.Result, error) {
	shopGoodsCategoryCategoryIdKey := fmt.Sprintf("%s%v", cacheShopGoodsCategoryCategoryIdPrefix, data.CategoryId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, goodsCategoryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CategoryId, data.CategoryName)
	}, shopGoodsCategoryCategoryIdKey)
	return ret, err
}

func (m *defaultGoodsCategoryModel) TransInsert(ctx context.Context, session sqlx.Session, data *GoodsCategory) (sql.Result, error) {
	shopGoodsCategoryCategoryIdKey := fmt.Sprintf("%s%v", cacheShopGoodsCategoryCategoryIdPrefix, data.CategoryId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, goodsCategoryRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.CategoryId, data.CategoryName)
	}, shopGoodsCategoryCategoryIdKey)
	return ret, err
}
func (m *defaultGoodsCategoryModel) Update(ctx context.Context, data *GoodsCategory) error {
	shopGoodsCategoryCategoryIdKey := fmt.Sprintf("%s%v", cacheShopGoodsCategoryCategoryIdPrefix, data.CategoryId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `category_id` = ?", m.table, goodsCategoryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CategoryName, data.CategoryId)
	}, shopGoodsCategoryCategoryIdKey)
	return err
}

func (m *defaultGoodsCategoryModel) TransUpdate(ctx context.Context, session sqlx.Session, data *GoodsCategory) error {
	shopGoodsCategoryCategoryIdKey := fmt.Sprintf("%s%v", cacheShopGoodsCategoryCategoryIdPrefix, data.CategoryId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `category_id` = ?", m.table, goodsCategoryRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.CategoryName, data.CategoryId)
	}, shopGoodsCategoryCategoryIdKey)
	return err
}

func (m *defaultGoodsCategoryModel) List(ctx context.Context, page, limit int64) ([]*GoodsCategory, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", goodsCategoryRows, m.table)
	var resp []*GoodsCategory
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultGoodsCategoryModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultGoodsCategoryModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultGoodsCategoryModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultGoodsCategoryModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*GoodsCategory, error) {

	builder = builder.Columns(goodsCategoryRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*GoodsCategory
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultGoodsCategoryModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*GoodsCategory, error) {

	builder = builder.Columns(goodsCategoryRows)

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

	var resp []*GoodsCategory
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultGoodsCategoryModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*GoodsCategory, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(goodsCategoryRows)

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

	var resp []*GoodsCategory
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultGoodsCategoryModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*GoodsCategory, error) {

	builder = builder.Columns(goodsCategoryRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*GoodsCategory
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultGoodsCategoryModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*GoodsCategory, error) {

	builder = builder.Columns(goodsCategoryRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*GoodsCategory
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultGoodsCategoryModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultGoodsCategoryModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheShopGoodsCategoryCategoryIdPrefix, primary)
}

func (m *defaultGoodsCategoryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `category_id` = ? limit 1", goodsCategoryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultGoodsCategoryModel) tableName() string {
	return m.table
}
