// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2025-02-21 18:37:33
// 生成路径: internal/app/appsource/router/as_application.go
// 生成人：smithy
// desc:应用管理
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"
	"mangosmithy/internal/app/appsource/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindAsApplicationController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/asApplication", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.AsApplication,
		)
	})
}
