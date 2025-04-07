package router

import (
	"context"
	installRouter "mangosmithy/internal/app/install/router"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (router *Router) BindInstallModuleController(ctx context.Context, group *ghttp.RouterGroup) {
	installRouter.R.BindController(ctx, group)
}
