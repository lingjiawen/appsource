// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2025-04-01 19:38:11
// 生成路径: internal/app/install/dao/internal/sign_device.go
// 生成人：smithy
// desc:设备管理
// company:云南奇讯科技有限公司
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SignDeviceDao is the manager for logic model data accessing and custom defined data operations functions management.
type SignDeviceDao struct {
	table   string            // Table is the underlying table name of the DAO.
	group   string            // Group is the database configuration group name of current DAO.
	columns SignDeviceColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// SignDeviceColumns defines and stores column names for table sign_device.
type SignDeviceColumns struct {
	Id              string // ID
	Udid            string // 设备码
	Name            string // 证书名
	CertId          string // 证书标识
	DeviceId        string // 设备标识
	P12             string // 证书文件
	Mobileprovision string // 描述文件
	AddTime         string // 添加时间
	P12Password     string // 证书密码
	Model           string // 设备型号
	ExpireTime      string // 过期时间
	SerialNumber    string // 序列号
	RedeemCode      string // 兑换卡密
	AccountType     string // 时效类型
	WarrantyType    string // 售后类型
	DeviceType      string // 设备类型
	Status          string // 状态
	Pool            string // 出书方式
	ApiPlatform     string // 对接平台
	ApiWarrantyType string // 对接售后类型
	Active          string // 禁用
	CreatedBy       string // 创建人
	UpdatedBy       string // 修改人
	CreatedAt       string // 创建时间
	UpdatedAt       string // 修改时间
	DeletedAt       string // 删除时间
}

var signDeviceColumns = SignDeviceColumns{
	Id:              "id",
	Udid:            "udid",
	Name:            "name",
	CertId:          "cert_id",
	DeviceId:        "device_id",
	P12:             "p12",
	Mobileprovision: "mobileprovision",
	AddTime:         "add_time",
	P12Password:     "p12_password",
	Model:           "model",
	ExpireTime:      "expire_time",
	SerialNumber:    "serial_number",
	RedeemCode:      "redeem_code",
	AccountType:     "account_type",
	WarrantyType:    "warranty_type",
	DeviceType:      "device_type",
	Status:          "status",
	Pool:            "pool",
	ApiPlatform:     "api_platform",
	ApiWarrantyType: "api_warranty_type",
	Active:          "active",
	CreatedBy:       "created_by",
	UpdatedBy:       "updated_by",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

// NewSignDeviceDao creates and returns a new DAO object for table data access.
func NewSignDeviceDao() *SignDeviceDao {
	return &SignDeviceDao{
		group:   "default",
		table:   "sign_device",
		columns: signDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SignDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SignDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SignDeviceDao) Columns() SignDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SignDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SignDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SignDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
