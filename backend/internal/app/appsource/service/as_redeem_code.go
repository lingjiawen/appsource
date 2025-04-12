// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-04-09 21:34:25
// 生成路径: internal/app/appsource/service/as_redeem_code.go
// 生成人：smithy
// desc:卡密管理
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"
	"mangosmithy/internal/app/appsource/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type IAsRedeemCode interface {
	List(ctx context.Context, req *model.AsRedeemCodeSearchReq) (res *model.AsRedeemCodeSearchRes, err error)
	GetExportData(ctx context.Context, req *model.AsRedeemCodeSearchReq) (listRes []*model.AsRedeemCodeInfoRes, err error)
	Import(ctx context.Context, file *ghttp.UploadFile) (err error)
	GetById(ctx context.Context, Id uint) (res *model.AsRedeemCodeInfoRes, err error)
	Add(ctx context.Context, req *model.AsRedeemCodeAddReq) (err error)
	Edit(ctx context.Context, req *model.AsRedeemCodeEditReq) (err error)
	Delete(ctx context.Context, Id []uint) (err error)
	// 卡密管理是否激活修改（状态）
	ChangeActive(ctx context.Context, id uint, active int) (err error)
}

var localAsRedeemCode IAsRedeemCode

func AsRedeemCode() IAsRedeemCode {
	if localAsRedeemCode == nil {
		panic("implement not found for interface IAsRedeemCode, forgot register?")
	}
	return localAsRedeemCode
}

func RegisterAsRedeemCode(i IAsRedeemCode) {
	localAsRedeemCode = i
}
