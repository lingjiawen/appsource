// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2025-03-25 13:14:15
// 生成路径: internal/app/install/router/sign_redeem_code.go
// 生成人：smithy
// desc:签名卡密
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"
	"mangosmithy/internal/app/install/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindSignRedeemCodeController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/signRedeemCode", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SignRedeemCode,
		)
	})
}
