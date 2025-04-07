// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-02-21 18:37:33
// 生成路径: internal/app/appsource/service/as_application.go
// 生成人：smithy
// desc:应用管理
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"mangosmithy/internal/app/appsource/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type IAsApplication interface {
	List(ctx context.Context, req *model.AsApplicationSearchReq) (res *model.AsApplicationSearchRes, err error)
	GetExportData(ctx context.Context, req *model.AsApplicationSearchReq) (listRes []*model.AsApplicationInfoRes, err error)
	Import(ctx context.Context, file *ghttp.UploadFile) (err error)
	GetById(ctx context.Context, Id uint) (res *model.AsApplicationInfoRes, err error)
	Add(ctx context.Context, req *model.AsApplicationAddReq) (err error)
	Edit(ctx context.Context, req *model.AsApplicationEditReq) (err error)
	Delete(ctx context.Context, Id []uint) (err error)
	Extract(ctx context.Context, req *model.AsApplicationExtractReq) (res *model.AsApplicationExtractRes, err error)
	IconSearch(ctx context.Context, req *model.AsApplicationIconSearchReq) (res g.Map, err error)
	SourceImport(ctx context.Context, req *model.AsApplicationSourceImportReq) (err error)
	// ChangeLock 应用管理付费修改（状态）
	ChangeLock(ctx context.Context, id uint, lock int) (err error)
	// ChangeActive 应用管理是否启用修改（状态）
	ChangeActive(ctx context.Context, id uint, active int) (err error)
}

var localAsApplication IAsApplication

func AsApplication() IAsApplication {
	if localAsApplication == nil {
		panic("implement not found for interface IAsApplication, forgot register?")
	}
	return localAsApplication
}

func RegisterAsApplication(i IAsApplication) {
	localAsApplication = i
}
