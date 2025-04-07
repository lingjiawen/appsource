// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2025-04-01 19:38:11
// 生成路径: internal/app/install/router/sign_device.go
// 生成人：smithy
// desc:设备管理
// company:云南奇讯科技有限公司
// ==========================================================================

package router

import (
	"context"
	"mangosmithy/internal/app/install/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindSignDeviceController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/signDevice", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SignDevice,
		)
	})
}
