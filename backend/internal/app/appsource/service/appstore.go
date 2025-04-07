package service

import (
	"context"
	"mangosmithy/internal/app/appsource/model"
)

type IAppStore interface {
	AppStore(ctx context.Context, req *model.AppStoreReq) (res *model.AppStoreRes, err error)
}

var localAppStore IAppStore

func AppStore() IAppStore {
	if localAppStore == nil {
		panic("implement not found for interface IAppStore, forgot register?")
	}
	return localAppStore
}

func RegisterAppStore(i IAppStore) {
	localAppStore = i
}
