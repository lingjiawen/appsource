// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-02-18 20:01:49
// 生成路径: internal/app/appsource/controller/sys_config.go
// 生成人：smithy
// desc:配置项
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"
	"errors"
	"mangosmithy/api/v1/appsource"
	"mangosmithy/internal/app/appsource/service"
	systemController "mangosmithy/internal/app/system/controller"
	systemService "mangosmithy/internal/app/system/service"
)

type sysConfigController struct {
	systemController.BaseController
}

var SysConfig = new(sysConfigController)

// List 列表
func (c *sysConfigController) List(ctx context.Context, req *appsource.SysConfigSearchReq) (res *appsource.SysConfigSearchRes, err error) {
	res = new(appsource.SysConfigSearchRes)
	res.SysConfigSearchRes, err = service.SysConfig().List(ctx, &req.SysConfigSearchReq)
	return
}

// LinkedSysConfigDataSearch 相关连表查询数据
func (c *sysConfigController) LinkedSysConfigDataSearch(ctx context.Context, req *appsource.LinkedSysConfigDataSearchReq) (res *appsource.LinkedSysConfigDataSearchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/appsource/sysConfig/list") {
		err = errors.New("没有访问权限")
		return
	}
	res = new(appsource.LinkedSysConfigDataSearchRes)
	res.LinkedSysConfigDataSearchRes, err = service.SysConfig().LinkedSysConfigDataSearch(ctx)
	return
}

// Get 获取配置项
func (c *sysConfigController) Get(ctx context.Context, req *appsource.SysConfigGetReq) (res *appsource.SysConfigGetRes, err error) {
	res = new(appsource.SysConfigGetRes)
	res.SysConfigInfoRes, err = service.SysConfig().GetByConfigId(ctx, req.ConfigId)
	return
}

// Add 添加配置项
func (c *sysConfigController) Add(ctx context.Context, req *appsource.SysConfigAddReq) (res *appsource.SysConfigAddRes, err error) {
	err = service.SysConfig().Add(ctx, req.SysConfigAddReq)
	return
}

// Edit 修改配置项
func (c *sysConfigController) Edit(ctx context.Context, req *appsource.SysConfigEditReq) (res *appsource.SysConfigEditRes, err error) {
	err = service.SysConfig().Edit(ctx, req.SysConfigEditReq)
	return
}

// Delete 删除配置项
func (c *sysConfigController) Delete(ctx context.Context, req *appsource.SysConfigDeleteReq) (res *appsource.SysConfigDeleteRes, err error) {
	err = service.SysConfig().Delete(ctx, req.ConfigIds)
	return
}
