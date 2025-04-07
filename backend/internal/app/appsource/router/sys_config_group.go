// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2025-02-18 19:57:00
// 生成路径: internal/app/appsource/router/sys_config_group.go
// 生成人：smithy
// desc:配置组
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"
	"mangosmithy/internal/app/appsource/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindSysConfigGroupController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/sysConfigGroup", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SysConfigGroup,
		)
	})
}
