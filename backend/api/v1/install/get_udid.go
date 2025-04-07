package install

import (
	"github.com/gogf/gf/v2/frame/g"
	"mangosmithy/internal/app/install/model"
)

// GetUDIDReq 请求参数
type GetUDIDReq struct {
	g.Meta `path:"/getudid" tags:"签名工具" method:"get" summary:"获取UDID"`
}

// GetUDIDRes 返回结果
type GetUDIDRes struct {
	g.Meta `mime:"application/json"`
}

// GetUDIDResultReq 请求参数
type GetUDIDResultReq struct {
	g.Meta `path:"/udidResult" tags:"签名工具" method:"post" summary:"获取UDID结果"`
	model.GetUDIDResultReq
}

// GetUDIDResultRes 返回结果
type GetUDIDResultRes struct {
	g.Meta `mime:"application/json"`
}
