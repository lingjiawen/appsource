// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-02-18 19:57:00
// 生成路径: internal/app/appsource/service/sys_config_group.go
// 生成人：smithy
// desc:配置组
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"
	"mangosmithy/internal/app/appsource/model"
)

type ISysConfigGroup interface {
	List(ctx context.Context, req *model.SysConfigGroupSearchReq) (res *model.SysConfigGroupSearchRes, err error)
	GetById(ctx context.Context, Id uint) (res *model.SysConfigGroupInfoRes, err error)
	Add(ctx context.Context, req *model.SysConfigGroupAddReq) (err error)
	Edit(ctx context.Context, req *model.SysConfigGroupEditReq) (err error)
	Delete(ctx context.Context, Id []uint) (err error)
}

var localSysConfigGroup ISysConfigGroup

func SysConfigGroup() ISysConfigGroup {
	if localSysConfigGroup == nil {
		panic("implement not found for interface ISysConfigGroup, forgot register?")
	}
	return localSysConfigGroup
}

func RegisterSysConfigGroup(i ISysConfigGroup) {
	localSysConfigGroup = i
}
