// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-03-25 13:14:15
// 生成路径: internal/app/install/service/sign_redeem_code.go
// 生成人：smithy
// desc:签名卡密
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"
	"mangosmithy/internal/app/install/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type ISignRedeemCode interface {
	List(ctx context.Context, req *model.SignRedeemCodeSearchReq) (res *model.SignRedeemCodeSearchRes, err error)
	GetExportData(ctx context.Context, req *model.SignRedeemCodeSearchReq) (listRes []*model.SignRedeemCodeInfoRes, err error)
	Import(ctx context.Context, file *ghttp.UploadFile) (err error)
	GetById(ctx context.Context, Id uint) (res *model.SignRedeemCodeInfoRes, err error)
	Add(ctx context.Context, req *model.SignRedeemCodeAddReq) (err error)
	Edit(ctx context.Context, req *model.SignRedeemCodeEditReq) (err error)
	Delete(ctx context.Context, Id []uint) (err error)
	// 签名卡密禁用修改（状态）
	ChangeBanned(ctx context.Context, id uint, banned int) (err error)
	// 签名卡密激活修改（状态）
	ChangeActive(ctx context.Context, id uint, active int) (err error)
}

var localSignRedeemCode ISignRedeemCode

func SignRedeemCode() ISignRedeemCode {
	if localSignRedeemCode == nil {
		panic("implement not found for interface ISignRedeemCode, forgot register?")
	}
	return localSignRedeemCode
}

func RegisterSignRedeemCode(i ISignRedeemCode) {
	localSignRedeemCode = i
}
