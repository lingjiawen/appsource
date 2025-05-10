// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2025-04-14 19:12:45
// 生成路径: internal/app/install/dao/internal/sign_help.go
// 生成人：smithy
// desc:帮助中心
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SignHelpDao is the manager for logic model data accessing and custom defined data operations functions management.
type SignHelpDao struct {
	table   string          // Table is the underlying table name of the DAO.
	group   string          // Group is the database configuration group name of current DAO.
	columns SignHelpColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// SignHelpColumns defines and stores column names for table sign_help.
type SignHelpColumns struct {
	Id        string // ID
	Title     string // 标题
	Content   string // 内容
	IsExpand  string // 默认展开
	Weigh     string // 权重
	CreatedAt string // 创建日期
	UpdatedAt string // 修改日期
}

var signHelpColumns = SignHelpColumns{
	Id:        "id",
	Title:     "title",
	Content:   "content",
	IsExpand:  "is_expand",
	Weigh:     "weigh",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSignHelpDao creates and returns a new DAO object for table data access.
func NewSignHelpDao() *SignHelpDao {
	return &SignHelpDao{
		group:   "default",
		table:   "sign_help",
		columns: signHelpColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SignHelpDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SignHelpDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SignHelpDao) Columns() SignHelpColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SignHelpDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SignHelpDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SignHelpDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
