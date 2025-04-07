// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-02-18 20:01:49
// 生成路径: internal/app/appsource/service/sys_config.go
// 生成人：smithy
// desc:配置项
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"
	"mangosmithy/internal/app/appsource/model"
)

type ISysConfig interface {
	List(ctx context.Context, req *model.SysConfigSearchReq) (res *model.SysConfigSearchRes, err error)
	GetByConfigId(ctx context.Context, ConfigId uint) (res *model.SysConfigInfoRes, err error)
	Add(ctx context.Context, req *model.SysConfigAddReq) (err error)
	Edit(ctx context.Context, req *model.SysConfigEditReq) (err error)
	Delete(ctx context.Context, ConfigId []uint) (err error)
	ParseAppInfo(ctx context.Context, req *model.ParseAppInfoReq) (res *model.ParseAppInfoRes, err error)
	LinkedSysConfigDataSearch(ctx context.Context) (res *model.LinkedSysConfigDataSearchRes, err error)
}

var localSysConfig ISysConfig

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}
