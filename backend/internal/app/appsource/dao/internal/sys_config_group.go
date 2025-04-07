// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2025-02-18 19:57:00
// 生成路径: internal/app/appsource/dao/internal/sys_config_group.go
// 生成人：smithy
// desc:配置组
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysConfigGroupDao is the manager for logic model data accessing and custom defined data operations functions management.
type SysConfigGroupDao struct {
	table   string                // Table is the underlying table name of the DAO.
	group   string                // Group is the database configuration group name of current DAO.
	columns SysConfigGroupColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// SysConfigGroupColumns defines and stores column names for table sys_config_group.
type SysConfigGroupColumns struct {
	Id        string // 主键
	GroupName string // 配置组名称
	GroupKey  string // 配置组键名
	CreateBy  string // 创建者
	UpdateBy  string // 更新者
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

var sysConfigGroupColumns = SysConfigGroupColumns{
	Id:        "id",
	GroupName: "group_name",
	GroupKey:  "group_key",
	CreateBy:  "create_by",
	UpdateBy:  "update_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSysConfigGroupDao creates and returns a new DAO object for table data access.
func NewSysConfigGroupDao() *SysConfigGroupDao {
	return &SysConfigGroupDao{
		group:   "default",
		table:   "sys_config_group",
		columns: sysConfigGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysConfigGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysConfigGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysConfigGroupDao) Columns() SysConfigGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysConfigGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysConfigGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysConfigGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
