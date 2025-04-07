// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2025-03-25 21:39:11
// 生成路径: internal/app/install/model/entity/sign_platform.go
// 生成人：smithy
// desc:平台
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SignPlatform is the golang structure for table sign_platform.
type SignPlatform struct {
	gmeta.Meta `orm:"table:sign_platform"`
	Id         uint        `orm:"id,primary" json:"id"`        // ID
	Name       string      `orm:"name" json:"name"`            // 平台名
	Code       string      `orm:"code" json:"code"`            // 平台标识
	BaseUrl    string      `orm:"base_url" json:"baseUrl"`     // 域名
	OpenSsl    int         `orm:"open_ssl" json:"openSsl"`     // 开启SSL
	Status     int         `orm:"status" json:"status"`        // 启用
	Token      string      `orm:"token" json:"token"`          // 对接Token
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt"` // 修改时间
}
