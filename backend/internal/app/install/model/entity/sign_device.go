// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-04-01 19:38:11
// 生成路径: internal/app/install/model/entity/sign_device.go
// 生成人：smithy
// desc:设备管理
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SignDevice is the golang structure for table sign_device.
type SignDevice struct {
	gmeta.Meta      `orm:"table:sign_device"`
	Id              uint        `orm:"id,primary" json:"id"`                     // ID
	Udid            string      `orm:"udid" json:"udid"`                         // 设备码
	Name            string      `orm:"name" json:"name"`                         // 证书名
	CertId          string      `orm:"cert_id" json:"certId"`                    // 证书标识
	DeviceId        string      `orm:"device_id" json:"deviceId"`                // 设备标识
	P12             string      `orm:"p12" json:"p12"`                           // 证书文件
	Mobileprovision string      `orm:"mobileprovision" json:"mobileprovision"`   // 描述文件
	AddTime         uint64      `orm:"add_time" json:"addTime"`                  // 添加时间
	P12Password     string      `orm:"p12_password" json:"p12Password"`          // 证书密码
	Model           string      `orm:"model" json:"model"`                       // 设备型号
	ExpireTime      uint64      `orm:"expire_time" json:"expireTime"`            // 过期时间
	SerialNumber    string      `orm:"serial_number" json:"serialNumber"`        // 序列号
	RedeemCode      string      `orm:"redeem_code" json:"redeemCode"`            // 兑换卡密
	AccountType     int         `orm:"account_type" json:"accountType"`          // 时效类型
	WarrantyType    int         `orm:"warranty_type" json:"warrantyType"`        // 售后类型
	DeviceType      string      `orm:"device_type" json:"deviceType"`            // 设备类型
	Status          string      `orm:"status" json:"status"`                     // 状态
	Pool            int         `orm:"pool" json:"pool"`                         // 出书方式
	ApiPlatform     int         `orm:"api_platform" json:"apiPlatform"`          // 对接平台
	ApiWarrantyType int         `orm:"api_warranty_type" json:"apiWarrantyType"` // 对接售后类型
	Active          int         `orm:"active" json:"active"`                     // 禁用
	CreatedBy       uint        `orm:"created_by" json:"createdBy"`              // 创建人
	UpdatedBy       uint        `orm:"updated_by" json:"updatedBy"`              // 修改人
	CreatedAt       *gtime.Time `orm:"created_at" json:"createdAt"`              // 创建时间
	UpdatedAt       *gtime.Time `orm:"updated_at" json:"updatedAt"`              // 修改时间
	DeletedAt       *gtime.Time `orm:"deleted_at" json:"deletedAt"`              // 删除时间
}
