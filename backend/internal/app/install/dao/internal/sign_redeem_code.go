// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2025-03-25 13:14:15
// 生成路径: internal/app/install/dao/internal/sign_redeem_code.go
// 生成人：smithy
// desc:签名卡密
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SignRedeemCodeDao is the manager for logic model data accessing and custom defined data operations functions management.
type SignRedeemCodeDao struct {
	table   string                // Table is the underlying table name of the DAO.
	group   string                // Group is the database configuration group name of current DAO.
	columns SignRedeemCodeColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// SignRedeemCodeColumns defines and stores column names for table sign_redeem_code.
type SignRedeemCodeColumns struct {
	Id              string // ID
	Code            string // 兑换码
	Udid            string // 设备码
	CertId          string // 证书标识
	AccountType     string // 时效类型
	WarrantyType    string // 售后类型
	DeviceType      string // 设备类型
	Pool            string // 出书方式
	ApiPlatform     string // 对接平台
	Note            string // 备注
	ApiWarrantyType string // 对接售后类型
	Force           string // 强制
	Banned          string // 禁用
	Active          string // 激活
	ActiveAt        string // 激活时间
	CreatedBy       string // 创建人
	UpdatedBy       string // 修改人
	CreatedAt       string // 创建时间
	UpdatedAt       string // 修改时间
	DeletedAt       string // 删除时间
}

var signRedeemCodeColumns = SignRedeemCodeColumns{
	Id:              "id",
	Code:            "code",
	Udid:            "udid",
	CertId:          "cert_id",
	AccountType:     "account_type",
	WarrantyType:    "warranty_type",
	DeviceType:      "device_type",
	Pool:            "pool",
	ApiPlatform:     "api_platform",
	Note:            "note",
	ApiWarrantyType: "api_warranty_type",
	Force:           "force",
	Banned:          "banned",
	Active:          "active",
	ActiveAt:        "active_at",
	CreatedBy:       "created_by",
	UpdatedBy:       "updated_by",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

// NewSignRedeemCodeDao creates and returns a new DAO object for table data access.
func NewSignRedeemCodeDao() *SignRedeemCodeDao {
	return &SignRedeemCodeDao{
		group:   "default",
		table:   "sign_redeem_code",
		columns: signRedeemCodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SignRedeemCodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SignRedeemCodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SignRedeemCodeDao) Columns() SignRedeemCodeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SignRedeemCodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SignRedeemCodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SignRedeemCodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
