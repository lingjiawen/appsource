package install

import (
	commonApi "mangosmithy/api/v1/common"
	"mangosmithy/internal/app/install/model"

	"github.com/gogf/gf/v2/frame/g"
)

// SysConfigSearchReq 分页请求参数
type SysConfigSearchReq struct {
	g.Meta `path:"/getconfig" tags:"签名工具" method:"get" summary:"配置项"`
	model.SysConfigSearchReq
}

// SysConfigSearchRes 列表返回结果
type SysConfigSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.SysConfigSearchRes
}

// ApplicationInstallReq 安装请求参数
type ApplicationInstallReq struct {
	g.Meta `path:"/install" tags:"签名工具" method:"post" summary:"兑换安装"`
	model.ApplicationInstallReq
}

// ApplicationInstallRes 安装返回结果
type ApplicationInstallRes struct {
	g.Meta `mime:"application/json"`
	*model.ApplicationInstallRes
}

// ApplicationPrivateInstallReq 安装请求参数
type ApplicationPrivateInstallReq struct {
	g.Meta `path:"/installPrivate" tags:"签名工具" method:"post" summary:"上传证书安装"`
	model.ApplicationPrivateInstallReq
}

// ApplicationPrivateInstallRes 安装返回结果
type ApplicationPrivateInstallRes struct {
	g.Meta `mime:"application/json"`
	*model.ApplicationPrivateInstallRes
}

// GetInstallPlistReq 获取安装链接请求参数
type GetInstallPlistReq struct {
	g.Meta `path:"/getInstallPlist/:uuid" tags:"签名工具" method:"get" summary:"获取安装Plist"`
	model.GetPlistReq
}

// GetInstallPlistRes 获取安装链接返回结果
type GetInstallPlistRes struct {
	commonApi.EmptyRes
}

// CheckDevicesReq 获取设备请求参数
type CheckDevicesReq struct {
	g.Meta `path:"/deviceCheck" tags:"签名工具" method:"post" summary:"获取设备"`
	model.CheckDevicesReq
}

// CheckDevicesRes 获取设备返回结果
type CheckDevicesRes struct {
	g.Meta `mime:"application/json"`
	*model.CheckDevicesRes
}
