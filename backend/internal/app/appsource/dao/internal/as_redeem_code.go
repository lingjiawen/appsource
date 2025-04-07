// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2025-02-16 21:29:59
// 生成路径: internal/app/appsource/dao/internal/as_redeem_code.go
// 生成人：smithy
// desc:卡密管理
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AsRedeemCodeDao is the manager for logic model data accessing and custom defined data operations functions management.
type AsRedeemCodeDao struct {
	table   string              // Table is the underlying table name of the DAO.
	group   string              // Group is the database configuration group name of current DAO.
	columns AsRedeemCodeColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// AsRedeemCodeColumns defines and stores column names for table as_redeem_code.
type AsRedeemCodeColumns struct {
	Id        string // ID
	Code      string // 兑换码
	Udid      string // 设备码
	Type      string // 类型
	Active    string // 是否激活
	ActiveAt  string // 激活时间
	CreatedBy string // 创建人
	UpdatedBy string // 修改人
	CreatedAt string // 生成时间
	UpdatedAt string // 修改时间
	DeletedAt string // 删除时间
}

var asRedeemCodeColumns = AsRedeemCodeColumns{
	Id:        "id",
	Code:      "code",
	Udid:      "udid",
	Type:      "type",
	Active:    "active",
	ActiveAt:  "active_at",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewAsRedeemCodeDao creates and returns a new DAO object for table data access.
func NewAsRedeemCodeDao() *AsRedeemCodeDao {
	return &AsRedeemCodeDao{
		group:   "default",
		table:   "as_redeem_code",
		columns: asRedeemCodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AsRedeemCodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AsRedeemCodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AsRedeemCodeDao) Columns() AsRedeemCodeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AsRedeemCodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AsRedeemCodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AsRedeemCodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
