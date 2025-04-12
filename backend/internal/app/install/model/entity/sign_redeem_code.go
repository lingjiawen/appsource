// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-03-25 13:14:15
// 生成路径: internal/app/install/model/entity/sign_redeem_code.go
// 生成人：smithy
// desc:签名卡密
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SignRedeemCode is the golang structure for table sign_redeem_code.
type SignRedeemCode struct {
	gmeta.Meta      `orm:"table:sign_redeem_code"`
	Id              uint        `orm:"id,primary" json:"id"`                     // ID
	Code            string      `orm:"code" json:"code"`                         // 兑换码
	Udid            string      `orm:"udid" json:"udid"`                         // 设备码
	CertId          string      `orm:"cert_id" json:"certId"`                    // 证书标识
	AccountType     int         `orm:"account_type" json:"accountType"`          // 时效类型
	WarrantyType    int         `orm:"warranty_type" json:"warrantyType"`        // 售后类型
	DeviceType      string      `orm:"device_type" json:"deviceType"`            // 设备类型
	Pool            int         `orm:"pool" json:"pool"`                         // 出书方式
	ApiPlatform     int         `orm:"api_platform" json:"apiPlatform"`          // 对接平台
	Note            string      `orm:"note" json:"note"`                         // 备注
	ApiWarrantyType int         `orm:"api_warranty_type" json:"apiWarrantyType"` // 对接售后类型
	Force           int         `orm:"force" json:"force"`                       // 强制添加
	Banned          int         `orm:"banned" json:"banned"`                     // 禁用
	Active          int         `orm:"active" json:"active"`                     // 激活
	ActiveAt        *gtime.Time `orm:"active_at" json:"activeAt"`                // 激活时间
	CreatedBy       uint        `orm:"created_by" json:"createdBy"`              // 创建人
	UpdatedBy       uint        `orm:"updated_by" json:"updatedBy"`              // 修改人
	CreatedAt       *gtime.Time `orm:"created_at" json:"createdAt"`              // 创建时间
	UpdatedAt       *gtime.Time `orm:"updated_at" json:"updatedAt"`              // 修改时间
	DeletedAt       *gtime.Time `orm:"deleted_at" json:"deletedAt"`              // 删除时间
}
