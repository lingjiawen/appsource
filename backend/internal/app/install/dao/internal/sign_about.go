// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2025-04-14 19:08:27
// 生成路径: internal/app/install/dao/internal/sign_about.go
// 生成人：smithy
// desc:关于我们
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SignAboutDao is the manager for logic model data accessing and custom defined data operations functions management.
type SignAboutDao struct {
	table   string           // Table is the underlying table name of the DAO.
	group   string           // Group is the database configuration group name of current DAO.
	columns SignAboutColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// SignAboutColumns defines and stores column names for table sign_about.
type SignAboutColumns struct {
	Id        string // ID
	Icon      string // 图标
	Title     string // 标题
	Subtitle  string // 内容
	IsLink    string // 是否链接
	Url       string // 链接
	Group     string // 分组
	Weigh     string // 权重
	CreatedAt string // 创建日期
	UpdatedAt string // 修改日期
}

var signAboutColumns = SignAboutColumns{
	Id:        "id",
	Icon:      "icon",
	Title:     "title",
	Subtitle:  "subtitle",
	IsLink:    "is_link",
	Url:       "url",
	Group:     "group",
	Weigh:     "weigh",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSignAboutDao creates and returns a new DAO object for table data access.
func NewSignAboutDao() *SignAboutDao {
	return &SignAboutDao{
		group:   "default",
		table:   "sign_about",
		columns: signAboutColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SignAboutDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SignAboutDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SignAboutDao) Columns() SignAboutColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SignAboutDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SignAboutDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SignAboutDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
