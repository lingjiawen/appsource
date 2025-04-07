package service

import (
	"context"
	"mangosmithy/internal/app/install/model"
)

type IPublic interface {
	Install(ctx context.Context, req *model.ApplicationInstallReq) (res *model.ApplicationInstallRes, err error)
	PrivateInstall(ctx context.Context, req *model.ApplicationPrivateInstallReq) (res *model.ApplicationPrivateInstallRes, err error)
	GetPlist(ctx context.Context, req *model.GetPlistReq) (res *model.GetPlistRes, err error)
	CheckDevices(ctx context.Context, req *model.CheckDevicesReq) (res *model.CheckDevicesRes, err error)
}

var localPublic IPublic

func Public() IPublic {
	if localPublic == nil {
		panic("implement not found for interface IPublic, forgot register?")
	}
	return localPublic
}

func RegisterPublic(i IPublic) {
	localPublic = i
}
