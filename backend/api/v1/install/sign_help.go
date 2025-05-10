// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-04-14 19:12:45
// 生成路径: api/v1/install/sign_help.go
// 生成人：smithy
// desc:帮助中心相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package install

import (
	commonApi "mangosmithy/api/v1/common"
	"mangosmithy/internal/app/install/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// SignHelpSearchReq 分页请求参数
type SignHelpSearchReq struct {
	g.Meta `path:"/list" tags:"帮助中心" method:"get" summary:"帮助中心列表"`
	commonApi.Author
	model.SignHelpSearchReq
}

// SignHelpSearchRes 列表返回结果
type SignHelpSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.SignHelpSearchRes
}

// SignHelpExportReq 导出请求
type SignHelpExportReq struct {
	g.Meta `path:"/export" tags:"帮助中心" method:"get" summary:"帮助中心导出"`
	commonApi.Author
	model.SignHelpSearchReq
}

// SignHelpExportRes 导出响应
type SignHelpExportRes struct {
	commonApi.EmptyRes
}
type SignHelpExcelTemplateReq struct {
	g.Meta `path:"/excelTemplate" tags:"帮助中心" method:"get" summary:"导出模板文件"`
	commonApi.Author
}
type SignHelpExcelTemplateRes struct {
	commonApi.EmptyRes
}
type SignHelpImportReq struct {
	g.Meta `path:"/import" tags:"帮助中心" method:"post" summary:"帮助中心导入"`
	commonApi.Author
	File *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件"  v:"required#上传文件必须"`
}
type SignHelpImportRes struct {
	commonApi.EmptyRes
}

// SignHelpAddReq 添加操作请求参数
type SignHelpAddReq struct {
	g.Meta `path:"/add" tags:"帮助中心" method:"post" summary:"帮助中心添加"`
	commonApi.Author
	*model.SignHelpAddReq
}

// SignHelpAddRes 添加操作返回结果
type SignHelpAddRes struct {
	commonApi.EmptyRes
}

// SignHelpEditReq 修改操作请求参数
type SignHelpEditReq struct {
	g.Meta `path:"/edit" tags:"帮助中心" method:"put" summary:"帮助中心修改"`
	commonApi.Author
	*model.SignHelpEditReq
}

// SignHelpEditRes 修改操作返回结果
type SignHelpEditRes struct {
	commonApi.EmptyRes
}

// SignHelpGetReq 获取一条数据请求
type SignHelpGetReq struct {
	g.Meta `path:"/get" tags:"帮助中心" method:"get" summary:"获取帮助中心信息"`
	commonApi.Author
	Id uint `p:"id" v:"required#主键必须"` //通过主键获取
}

// SignHelpGetRes 获取一条数据结果
type SignHelpGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SignHelpInfoRes
}

// SignHelpDeleteReq 删除数据请求
type SignHelpDeleteReq struct {
	g.Meta `path:"/delete" tags:"帮助中心" method:"delete" summary:"删除帮助中心"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}

// SignHelpDeleteRes 删除数据返回
type SignHelpDeleteRes struct {
	commonApi.EmptyRes
}

// 帮助中心默认展开修改（状态）
type SignHelpIsExpandSwitchReq struct {
	g.Meta `path:"/changeIsExpand" tags:"帮助中心" method:"put" summary:"修改默认展开"`
	commonApi.Author
	Id       uint `p:"id" v:"required#主键必须"`       //通过主键修改
	IsExpand int  `p:"isExpand" v:"required#状态必须"` //通过主键获取
}
type SignHelpIsExpandSwitchRes struct {
	commonApi.EmptyRes
}
