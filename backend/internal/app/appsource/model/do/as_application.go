// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-02-21 18:37:33
// 生成路径: internal/app/appsource/model/entity/as_application.go
// 生成人：smithy
// desc:应用管理
// company:云南奇讯科技有限公司
// ==========================================================================

package do

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// AsApplication is the golang structure for table as_application.
type AsApplication struct {
	gmeta.Meta  `orm:"table:as_application, do:true"`
	Id          interface{} `orm:"id,primary" json:"id"`           // ID
	FileUrl     interface{} `orm:"file_url" json:"fileUrl"`        // 文件
	Name        interface{} `orm:"name" json:"name"`               // 应用名称
	BundleId    interface{} `orm:"bundle_id" json:"bundleId"`      // 包名
	Version     interface{} `orm:"version" json:"version"`         // 版本号
	Size        interface{} `orm:"size" json:"size"`               // 文件大小
	Type        interface{} `orm:"type" json:"type"`               // 类型
	Description interface{} `orm:"description" json:"description"` // 描述
	IconUrl     interface{} `orm:"icon_url" json:"iconUrl"`        // 图标
	Lock        interface{} `orm:"lock" json:"lock"`               // 付费
	Lanzou      interface{} `orm:"lanzou" json:"lanzou"`           // 是否蓝奏云链接
	Weigh       interface{} `orm:"weigh" json:"weigh"`             // 权重
	Active      interface{} `orm:"active" json:"active"`           // 是否启用
	Note        interface{} `orm:"note" json:"note"`               // 备注
	CreatedBy   interface{} `orm:"created_by" json:"createdBy"`    // 创建人
	UpdatedBy   interface{} `orm:"updated_by" json:"updatedBy"`    // 修改人
	CreatedAt   *gtime.Time `orm:"created_at" json:"createdAt"`    // 创建时间
	UpdatedAt   *gtime.Time `orm:"updated_at" json:"updatedAt"`    // 修改时间
}
