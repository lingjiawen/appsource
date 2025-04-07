// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2025-02-18 20:01:49
// 生成路径: internal/app/appsource/router/sys_config.go
// 生成人：smithy
// desc:配置项
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"
	"mangosmithy/internal/app/appsource/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindSysConfigController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/sysConfig", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SysConfig,
		)
	})
}
