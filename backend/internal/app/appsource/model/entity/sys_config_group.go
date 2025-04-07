// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-02-18 19:57:00
// 生成路径: internal/app/appsource/model/entity/sys_config_group.go
// 生成人：smithy
// desc:配置组
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysConfigGroup is the golang structure for table sys_config_group.
type SysConfigGroup struct {
	gmeta.Meta `orm:"table:sys_config_group"`
	Id         uint        `orm:"id,primary" json:"id"`        // 主键
	GroupName  string      `orm:"group_name" json:"groupName"` // 配置组名称
	GroupKey   string      `orm:"group_key" json:"groupKey"`   // 配置组键名
	CreateBy   uint        `orm:"create_by" json:"createBy"`   // 创建者
	UpdateBy   uint        `orm:"update_by" json:"updateBy"`   // 更新者
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` // 修改时间
}
