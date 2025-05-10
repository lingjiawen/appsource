// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-04-14 19:12:45
// 生成路径: internal/app/install/service/sign_help.go
// 生成人：smithy
// desc:帮助中心
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"
	"mangosmithy/internal/app/install/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type ISignHelp interface {
	List(ctx context.Context, req *model.SignHelpSearchReq) (res *model.SignHelpSearchRes, err error)
	GetExportData(ctx context.Context, req *model.SignHelpSearchReq) (listRes []*model.SignHelpInfoRes, err error)
	Import(ctx context.Context, file *ghttp.UploadFile) (err error)
	GetById(ctx context.Context, Id uint) (res *model.SignHelpInfoRes, err error)
	Add(ctx context.Context, req *model.SignHelpAddReq) (err error)
	Edit(ctx context.Context, req *model.SignHelpEditReq) (err error)
	Delete(ctx context.Context, Id []uint) (err error)
	// 帮助中心默认展开修改（状态）
	ChangeIsExpand(ctx context.Context, id uint, isExpand int) (err error)
}

var localSignHelp ISignHelp

func SignHelp() ISignHelp {
	if localSignHelp == nil {
		panic("implement not found for interface ISignHelp, forgot register?")
	}
	return localSignHelp
}

func RegisterSignHelp(i ISignHelp) {
	localSignHelp = i
}
