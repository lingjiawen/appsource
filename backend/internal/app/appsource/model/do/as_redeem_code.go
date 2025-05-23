// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-04-09 21:34:25
// 生成路径: internal/app/appsource/model/entity/as_redeem_code.go
// 生成人：smithy
// desc:卡密管理
// company:云南奇讯科技有限公司
// ==========================================================================

package do

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// AsRedeemCode is the golang structure for table as_redeem_code.
type AsRedeemCode struct {
	gmeta.Meta `orm:"table:as_redeem_code, do:true"`
	Id         interface{} `orm:"id,primary" json:"id"`        // ID
	Code       interface{} `orm:"code" json:"code"`            // 兑换码
	Udid       interface{} `orm:"udid" json:"udid"`            // 设备码
	Type       interface{} `orm:"type" json:"type"`            // 类型
	Active     interface{} `orm:"active" json:"active"`        // 是否激活
	ActiveAt   *gtime.Time `orm:"active_at" json:"activeAt"`   // 激活时间
	CreatedBy  interface{} `orm:"created_by" json:"createdBy"` // 创建人
	UpdatedBy  interface{} `orm:"updated_by" json:"updatedBy"` // 修改人
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 生成时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` // 修改时间
	DeletedAt  *gtime.Time `orm:"deleted_at" json:"deletedAt"` // 删除时间
	Note       interface{} `orm:"note" json:"note"`            // 备注
}
