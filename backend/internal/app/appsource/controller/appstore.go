package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"mangosmithy/api/v1/appsource"
	"mangosmithy/internal/app/appsource/service"
	systemController "mangosmithy/internal/app/system/controller"
)

type appStoreController struct {
	systemController.BaseController
}

var AppStore = new(appStoreController)

func (c *appStoreController) AppStore(ctx context.Context, req *appsource.AppStoreReq) (res *appsource.AppStoreRes, err error) {
	res = new(appsource.AppStoreRes)
	res.AppStoreRes, err = service.AppStore().AppStore(ctx, &req.AppStoreReq)
	r := ghttp.RequestFromCtx(ctx)

	if err != nil {
		return nil, err
	} else {
		if res.Identifier == "activeSuccess" {
			r.Response.WriteJsonExit(g.Map{
				"code":    0,
				"msg":     "OK, 解锁成功",
				"message": "OK, 解锁成功",
			})
		}
		if res.Identifier == "encrypted" {
			r.Response.Header().Set("x-iv", res.Message)
			r.Response.WriteJsonExit(g.Map{
				"mango-source": res.Name,
			})
		}
		r.Response.WriteJsonExit(res.AppStoreRes)
	}
	return
}
