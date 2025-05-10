// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-04-14 19:12:45
// 生成路径: internal/app/install/model/entity/sign_help.go
// 生成人：smithy
// desc:帮助中心
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SignHelp is the golang structure for table sign_help.
type SignHelp struct {
	gmeta.Meta `orm:"table:sign_help"`
	Id         uint        `orm:"id,primary" json:"id"`        // ID
	Title      string      `orm:"title" json:"title"`          // 标题
	Content    string      `orm:"content" json:"content"`      // 内容
	IsExpand   int         `orm:"is_expand" json:"isExpand"`   // 默认展开
	Weigh      int         `orm:"weigh" json:"weigh"`          // 权重
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 创建日期
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` // 修改日期
}
