// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-04-14 19:08:27
// 生成路径: internal/app/install/model/entity/sign_about.go
// 生成人：smithy
// desc:关于我们
// company:云南奇讯科技有限公司
// ==========================================================================

package do

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SignAbout is the golang structure for table sign_about.
type SignAbout struct {
	gmeta.Meta `orm:"table:sign_about, do:true"`
	Id         interface{} `orm:"id,primary" json:"id"`        // ID
	Icon       interface{} `orm:"icon" json:"icon"`            // 图标
	Title      interface{} `orm:"title" json:"title"`          // 标题
	Subtitle   interface{} `orm:"subtitle" json:"subtitle"`    // 内容
	IsLink     interface{} `orm:"is_link" json:"isLink"`       // 是否链接
	Url        interface{} `orm:"url" json:"url"`              // 链接
	Group      interface{} `orm:"group" json:"group"`          // 分组
	Weigh      interface{} `orm:"weigh" json:"weigh"`          // 权重
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 创建日期
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` // 修改日期
}
