// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2025-04-09 21:34:25
// 生成路径: internal/app/appsource/router/as_redeem_code.go
// 生成人：smithy
// desc:卡密管理
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"
	"mangosmithy/internal/app/appsource/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindAsRedeemCodeController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/asRedeemCode", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.AsRedeemCode,
		)
	})
}
