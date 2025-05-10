package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"mangosmithy/api/v1/install"
	"mangosmithy/internal/app/install/service"
	systemController "mangosmithy/internal/app/system/controller"
)

type publicController struct {
	systemController.BaseController
}

var PublicAPI = new(publicController)

// GetConfig 获取配置
func (c *publicController) GetConfig(ctx context.Context, req *install.SysConfigSearchReq) (res *install.SysConfigSearchRes, err error) {
	res = new(install.SysConfigSearchRes)
	res.SysConfigSearchRes, err = service.SysConfig().List(ctx, &req.SysConfigSearchReq)
	return
}

// GetHelp 获取配置
func (c *publicController) GetHelp(ctx context.Context, req *install.GetHelpReq) (res *install.GetHelpRes, err error) {
	res = new(install.GetHelpRes)
	res.GetHelpRes, err = service.Public().GetHelp(ctx)
	return
}

// GetAbout 获取配置
func (c *publicController) GetAbout(ctx context.Context, req *install.GetAboutReq) (res *install.GetAboutRes, err error) {
	res = new(install.GetAboutRes)
	res.GetAboutRes, err = service.Public().GetAbout(ctx)
	return
}

// ApplicationInstall 安装
func (c *publicController) ApplicationInstall(ctx context.Context, req *install.ApplicationInstallReq) (res *install.ApplicationInstallRes, err error) {
	res = new(install.ApplicationInstallRes)
	res.ApplicationInstallRes, err = service.Public().Install(ctx, &req.ApplicationInstallReq)
	return
}

// ApplicationPrivateInstall 上传安装
func (c *publicController) ApplicationPrivateInstall(ctx context.Context, req *install.ApplicationPrivateInstallReq) (res *install.ApplicationPrivateInstallRes, err error) {
	res = new(install.ApplicationPrivateInstallRes)
	res.ApplicationPrivateInstallRes, err = service.Public().PrivateInstall(ctx, &req.ApplicationPrivateInstallReq)
	return
}

// GetInstallPlist 获取安装链接
func (c *publicController) GetInstallPlist(ctx context.Context, req *install.GetInstallPlistReq) (res *install.GetInstallPlistRes, err error) {
	res = new(install.GetInstallPlistRes)

	// 获取 plist 内容（根据 uuid 获取）
	plistContent, err := service.Public().GetPlist(ctx, &req.GetPlistReq)
	if err != nil {
		return
	}

	r := g.RequestFromCtx(ctx)
	r.Response.Header().Set("Content-Type", "application/xml")
	//r.Response.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.plist\"", req.UUID))
	r.Response.Write([]byte(plistContent.PlistContent))
	return
}

// CheckDevices 获取设备
func (c *publicController) CheckDevices(ctx context.Context, req *install.CheckDevicesReq) (res *install.CheckDevicesRes, err error) {
	res = new(install.CheckDevicesRes)
	res.CheckDevicesRes, err = service.Public().CheckDevices(ctx, &req.CheckDevicesReq)
	return
}
