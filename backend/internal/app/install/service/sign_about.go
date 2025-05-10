// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-04-14 19:08:27
// 生成路径: internal/app/install/service/sign_about.go
// 生成人：smithy
// desc:关于我们
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"
	"mangosmithy/internal/app/install/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type ISignAbout interface {
	List(ctx context.Context, req *model.SignAboutSearchReq) (res *model.SignAboutSearchRes, err error)
	GetExportData(ctx context.Context, req *model.SignAboutSearchReq) (listRes []*model.SignAboutInfoRes, err error)
	Import(ctx context.Context, file *ghttp.UploadFile) (err error)
	GetById(ctx context.Context, Id uint) (res *model.SignAboutInfoRes, err error)
	Add(ctx context.Context, req *model.SignAboutAddReq) (err error)
	Edit(ctx context.Context, req *model.SignAboutEditReq) (err error)
	Delete(ctx context.Context, Id []uint) (err error)
	// 关于我们是否链接修改（状态）
	ChangeIsLink(ctx context.Context, id uint, isLink int) (err error)
}

var localSignAbout ISignAbout

func SignAbout() ISignAbout {
	if localSignAbout == nil {
		panic("implement not found for interface ISignAbout, forgot register?")
	}
	return localSignAbout
}

func RegisterSignAbout(i ISignAbout) {
	localSignAbout = i
}
