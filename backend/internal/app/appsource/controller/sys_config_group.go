// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-02-18 19:57:00
// 生成路径: internal/app/appsource/controller/sys_config_group.go
// 生成人：smithy
// desc:配置组
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"
	"mangosmithy/api/v1/appsource"
	"mangosmithy/internal/app/appsource/service"
	systemController "mangosmithy/internal/app/system/controller"
)

type sysConfigGroupController struct {
	systemController.BaseController
}

var SysConfigGroup = new(sysConfigGroupController)

// List 列表
func (c *sysConfigGroupController) List(ctx context.Context, req *appsource.SysConfigGroupSearchReq) (res *appsource.SysConfigGroupSearchRes, err error) {
	res = new(appsource.SysConfigGroupSearchRes)
	res.SysConfigGroupSearchRes, err = service.SysConfigGroup().List(ctx, &req.SysConfigGroupSearchReq)
	return
}

// Get 获取配置组
func (c *sysConfigGroupController) Get(ctx context.Context, req *appsource.SysConfigGroupGetReq) (res *appsource.SysConfigGroupGetRes, err error) {
	res = new(appsource.SysConfigGroupGetRes)
	res.SysConfigGroupInfoRes, err = service.SysConfigGroup().GetById(ctx, req.Id)
	return
}

// Add 添加配置组
func (c *sysConfigGroupController) Add(ctx context.Context, req *appsource.SysConfigGroupAddReq) (res *appsource.SysConfigGroupAddRes, err error) {
	err = service.SysConfigGroup().Add(ctx, req.SysConfigGroupAddReq)
	return
}

// Edit 修改配置组
func (c *sysConfigGroupController) Edit(ctx context.Context, req *appsource.SysConfigGroupEditReq) (res *appsource.SysConfigGroupEditRes, err error) {
	err = service.SysConfigGroup().Edit(ctx, req.SysConfigGroupEditReq)
	return
}

// Delete 删除配置组
func (c *sysConfigGroupController) Delete(ctx context.Context, req *appsource.SysConfigGroupDeleteReq) (res *appsource.SysConfigGroupDeleteRes, err error) {
	err = service.SysConfigGroup().Delete(ctx, req.Ids)
	return
}
