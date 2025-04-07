package appsource

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "mangosmithy/api/v1/common"
	"mangosmithy/internal/app/appsource/model"
)

// AppStoreReq 请求参数
type AppStoreReq struct {
	g.Meta `path:"/appstore" tags:"应用管理" method:"get" summary:"软件列表"`
	commonApi.Author
	model.AppStoreReq
}

// AppStoreRes 返回结果
type AppStoreRes struct {
	g.Meta `mime:"application/json"`
	*model.AppStoreRes
}
