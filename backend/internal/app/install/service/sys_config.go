package service

import (
	"context"
	"mangosmithy/internal/app/install/model"
)

type ISysConfig interface {
	List(ctx context.Context, req *model.SysConfigSearchReq) (res *model.SysConfigSearchRes, err error)
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
