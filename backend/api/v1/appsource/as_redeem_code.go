// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-02-16 21:29:59
// 生成路径: api/v1/appsource/as_redeem_code.go
// 生成人：smithy
// desc:卡密管理相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package appsource

import (
	commonApi "mangosmithy/api/v1/common"
	"mangosmithy/internal/app/appsource/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// AsRedeemCodeSearchReq 分页请求参数
type AsRedeemCodeSearchReq struct {
	g.Meta `path:"/list" tags:"卡密管理" method:"get" summary:"卡密管理列表"`
	commonApi.Author
	model.AsRedeemCodeSearchReq
}

// AsRedeemCodeSearchRes 列表返回结果
type AsRedeemCodeSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.AsRedeemCodeSearchRes
}

// AsRedeemCodeExportReq 导出请求
type AsRedeemCodeExportReq struct {
	g.Meta `path:"/export" tags:"卡密管理" method:"get" summary:"卡密管理导出"`
	commonApi.Author
	model.AsRedeemCodeSearchReq
}

// AsRedeemCodeExportRes 导出响应
type AsRedeemCodeExportRes struct {
	commonApi.EmptyRes
}
type AsRedeemCodeExcelTemplateReq struct {
	g.Meta `path:"/excelTemplate" tags:"卡密管理" method:"get" summary:"导出模板文件"`
	commonApi.Author
}
type AsRedeemCodeExcelTemplateRes struct {
	commonApi.EmptyRes
}
type AsRedeemCodeImportReq struct {
	g.Meta `path:"/import" tags:"卡密管理" method:"post" summary:"卡密管理导入"`
	commonApi.Author
	File *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件"  v:"required#上传文件必须"`
}
type AsRedeemCodeImportRes struct {
	commonApi.EmptyRes
}

// AsRedeemCodeAddReq 添加操作请求参数
type AsRedeemCodeAddReq struct {
	g.Meta `path:"/add" tags:"卡密管理" method:"post" summary:"卡密管理添加"`
	commonApi.Author
	*model.AsRedeemCodeAddReq
}

// AsRedeemCodeAddRes 添加操作返回结果
type AsRedeemCodeAddRes struct {
	commonApi.EmptyRes
}

// AsRedeemCodeEditReq 修改操作请求参数
type AsRedeemCodeEditReq struct {
	g.Meta `path:"/edit" tags:"卡密管理" method:"put" summary:"卡密管理修改"`
	commonApi.Author
	*model.AsRedeemCodeEditReq
}

// AsRedeemCodeEditRes 修改操作返回结果
type AsRedeemCodeEditRes struct {
	commonApi.EmptyRes
}

// AsRedeemCodeGetReq 获取一条数据请求
type AsRedeemCodeGetReq struct {
	g.Meta `path:"/get" tags:"卡密管理" method:"get" summary:"获取卡密管理信息"`
	commonApi.Author
	Id uint `p:"id" v:"required#主键必须"` //通过主键获取
}

// AsRedeemCodeGetRes 获取一条数据结果
type AsRedeemCodeGetRes struct {
	g.Meta `mime:"application/json"`
	*model.AsRedeemCodeInfoRes
}

// AsRedeemCodeDeleteReq 删除数据请求
type AsRedeemCodeDeleteReq struct {
	g.Meta `path:"/delete" tags:"卡密管理" method:"delete" summary:"删除卡密管理"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}

// AsRedeemCodeDeleteRes 删除数据返回
type AsRedeemCodeDeleteRes struct {
	commonApi.EmptyRes
}

// 卡密管理是否激活修改（状态）
type AsRedeemCodeActiveSwitchReq struct {
	g.Meta `path:"/changeActive" tags:"卡密管理" method:"put" summary:"修改是否激活"`
	commonApi.Author
	Id     uint `p:"id" v:"required#主键必须"`     //通过主键修改
	Active int  `p:"active" v:"required#状态必须"` //通过主键获取
}
type AsRedeemCodeActiveSwitchRes struct {
	commonApi.EmptyRes
}
