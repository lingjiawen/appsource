// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2025-02-21 18:37:33
// 生成路径: internal/app/appsource/dao/internal/as_application.go
// 生成人：smithy
// desc:应用管理
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AsApplicationDao is the manager for logic model data accessing and custom defined data operations functions management.
type AsApplicationDao struct {
	table   string               // Table is the underlying table name of the DAO.
	group   string               // Group is the database configuration group name of current DAO.
	columns AsApplicationColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// AsApplicationColumns defines and stores column names for table as_application.
type AsApplicationColumns struct {
	Id          string // ID
	FileUrl     string // 文件
	Name        string // 应用名称
	BundleId    string // 包名
	Version     string // 版本号
	Size        string // 文件大小
	Type        string // 类型
	Description string // 描述
	IconUrl     string // 图标
	Lock        string // 付费
	Lanzou      string // 是否蓝奏云链接
	Weigh       string // 权重
	Active      string // 是否启用
	Note        string // 备注
	CreatedBy   string // 创建人
	UpdatedBy   string // 修改人
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

var asApplicationColumns = AsApplicationColumns{
	Id:          "id",
	FileUrl:     "file_url",
	Name:        "name",
	BundleId:    "bundle_id",
	Version:     "version",
	Size:        "size",
	Type:        "type",
	Description: "description",
	IconUrl:     "icon_url",
	Lock:        "lock",
	Lanzou:      "lanzou",
	Weigh:       "weigh",
	Active:      "active",
	Note:        "note",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewAsApplicationDao creates and returns a new DAO object for table data access.
func NewAsApplicationDao() *AsApplicationDao {
	return &AsApplicationDao{
		group:   "default",
		table:   "as_application",
		columns: asApplicationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AsApplicationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AsApplicationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AsApplicationDao) Columns() AsApplicationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AsApplicationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AsApplicationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AsApplicationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
