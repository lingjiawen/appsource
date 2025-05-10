// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2025-04-14 19:08:27
// 生成路径: internal/app/install/router/sign_about.go
// 生成人：smithy
// desc:关于我们
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"
	"mangosmithy/internal/app/install/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindSignAboutController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/signAbout", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SignAbout,
		)
	})
}
