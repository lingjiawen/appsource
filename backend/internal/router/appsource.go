package router

import (
	"context"
	appsourceRouter "mangosmithy/internal/app/appsource/router"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindAppsourceModuleController(ctx context.Context, group *ghttp.RouterGroup) {
	appsourceRouter.R.BindController(ctx, group)
}
