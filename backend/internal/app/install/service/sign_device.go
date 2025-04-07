// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2025-04-01 19:38:11
// 生成路径: internal/app/install/service/sign_device.go
// 生成人：smithy
// desc:设备管理
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"
	"mangosmithy/internal/app/install/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type ISignDevice interface {
	List(ctx context.Context, req *model.SignDeviceSearchReq) (res *model.SignDeviceSearchRes, err error)
	GetExportData(ctx context.Context, req *model.SignDeviceSearchReq) (listRes []*model.SignDeviceInfoRes, err error)
	Import(ctx context.Context, file *ghttp.UploadFile) (err error)
	GetById(ctx context.Context, Id uint) (res *model.SignDeviceInfoRes, err error)
	Add(ctx context.Context, req *model.SignDeviceAddReq) (err error)
	Edit(ctx context.Context, req *model.SignDeviceEditReq) (err error)
	Delete(ctx context.Context, Id []uint) (err error)
	// 设备管理禁用修改（状态）
	ChangeActive(ctx context.Context, id uint, active int) (err error)
}

var localSignDevice ISignDevice

func SignDevice() ISignDevice {
	if localSignDevice == nil {
		panic("implement not found for interface ISignDevice, forgot register?")
	}
	return localSignDevice
}

func RegisterSignDevice(i ISignDevice) {
	localSignDevice = i
}
