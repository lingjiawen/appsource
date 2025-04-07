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
	"mangosmithy/internal/app/system/service"
	"mangosmithy/library/libRouter"

	"github.com/gogf/gf/v2/net/ghttp"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/install", func(group *ghttp.RouterGroup) {
		//context拦截器
		group.Middleware(service.Middleware().Ctx)
		//自动绑定定义的控制器
		if err := libRouter.RouterAutoBindBefore(ctx, router, group); err != nil {
			panic(err)
		}
		//登录验证拦截
		service.GfToken().Middleware(group)
		group.Middleware(service.Middleware().Auth)
		//后台操作日志记录
		group.Hook("/*", ghttp.HookAfterOutput, service.OperateLog().OperationLog)
		//自动绑定定义的控制器
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
