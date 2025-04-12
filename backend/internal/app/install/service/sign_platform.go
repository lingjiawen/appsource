// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-04-12 18:50:08
// 生成路径: internal/app/install/service/sign_platform.go
// 生成人：smithy
// desc:平台
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"
	"mangosmithy/internal/app/install/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type ISignPlatform interface {
	List(ctx context.Context, req *model.SignPlatformSearchReq) (res *model.SignPlatformSearchRes, err error)
	GetExportData(ctx context.Context, req *model.SignPlatformSearchReq) (listRes []*model.SignPlatformInfoRes, err error)
	Import(ctx context.Context, file *ghttp.UploadFile) (err error)
	GetById(ctx context.Context, Id uint) (res *model.SignPlatformInfoRes, err error)
	Add(ctx context.Context, req *model.SignPlatformAddReq) (err error)
	Edit(ctx context.Context, req *model.SignPlatformEditReq) (err error)
	Delete(ctx context.Context, Id []uint) (err error)
	// 平台开启SSL修改（状态）
	ChangeOpenSsl(ctx context.Context, id uint, openSsl int) (err error)
	// 平台启用修改（状态）
	ChangeStatus(ctx context.Context, id uint, status int) (err error)
}

var localSignPlatform ISignPlatform

func SignPlatform() ISignPlatform {
	if localSignPlatform == nil {
		panic("implement not found for interface ISignPlatform, forgot register?")
	}
	return localSignPlatform
}

func RegisterSignPlatform(i ISignPlatform) {
	localSignPlatform = i
}
