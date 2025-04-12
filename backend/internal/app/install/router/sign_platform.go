// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2025-04-12 18:50:08
// 生成路径: internal/app/install/router/sign_platform.go
// 生成人：smithy
// desc:平台
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"
	"mangosmithy/internal/app/install/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindSignPlatformController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/signPlatform", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SignPlatform,
		)
	})
}
