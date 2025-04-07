package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	websocketRouter "mangosmithy/internal/app/websocket/router"
)

func (router *Router) BindWebsocketModuleController(ctx context.Context, group *ghttp.RouterGroup) {
	websocketRouter.R.BindController(ctx, group)
}
