// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-04-14 19:08:27
// 生成路径: api/v1/install/sign_about.go
// 生成人：smithy
// desc:关于我们相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package install

import (
	commonApi "mangosmithy/api/v1/common"
	"mangosmithy/internal/app/install/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// SignAboutSearchReq 分页请求参数
type SignAboutSearchReq struct {
	g.Meta `path:"/list" tags:"关于我们" method:"get" summary:"关于我们列表"`
	commonApi.Author
	model.SignAboutSearchReq
}

// SignAboutSearchRes 列表返回结果
type SignAboutSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.SignAboutSearchRes
}

// SignAboutExportReq 导出请求
type SignAboutExportReq struct {
	g.Meta `path:"/export" tags:"关于我们" method:"get" summary:"关于我们导出"`
	commonApi.Author
	model.SignAboutSearchReq
}

// SignAboutExportRes 导出响应
type SignAboutExportRes struct {
	commonApi.EmptyRes
}
type SignAboutExcelTemplateReq struct {
	g.Meta `path:"/excelTemplate" tags:"关于我们" method:"get" summary:"导出模板文件"`
	commonApi.Author
}
type SignAboutExcelTemplateRes struct {
	commonApi.EmptyRes
}
type SignAboutImportReq struct {
	g.Meta `path:"/import" tags:"关于我们" method:"post" summary:"关于我们导入"`
	commonApi.Author
	File *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件"  v:"required#上传文件必须"`
}
type SignAboutImportRes struct {
	commonApi.EmptyRes
}

// SignAboutAddReq 添加操作请求参数
type SignAboutAddReq struct {
	g.Meta `path:"/add" tags:"关于我们" method:"post" summary:"关于我们添加"`
	commonApi.Author
	*model.SignAboutAddReq
}

// SignAboutAddRes 添加操作返回结果
type SignAboutAddRes struct {
	commonApi.EmptyRes
}

// SignAboutEditReq 修改操作请求参数
type SignAboutEditReq struct {
	g.Meta `path:"/edit" tags:"关于我们" method:"put" summary:"关于我们修改"`
	commonApi.Author
	*model.SignAboutEditReq
}

// SignAboutEditRes 修改操作返回结果
type SignAboutEditRes struct {
	commonApi.EmptyRes
}

// SignAboutGetReq 获取一条数据请求
type SignAboutGetReq struct {
	g.Meta `path:"/get" tags:"关于我们" method:"get" summary:"获取关于我们信息"`
	commonApi.Author
	Id uint `p:"id" v:"required#主键必须"` //通过主键获取
}

// SignAboutGetRes 获取一条数据结果
type SignAboutGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SignAboutInfoRes
}

// SignAboutDeleteReq 删除数据请求
type SignAboutDeleteReq struct {
	g.Meta `path:"/delete" tags:"关于我们" method:"delete" summary:"删除关于我们"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}

// SignAboutDeleteRes 删除数据返回
type SignAboutDeleteRes struct {
	commonApi.EmptyRes
}

// 关于我们是否链接修改（状态）
type SignAboutIsLinkSwitchReq struct {
	g.Meta `path:"/changeIsLink" tags:"关于我们" method:"put" summary:"修改是否链接"`
	commonApi.Author
	Id     uint `p:"id" v:"required#主键必须"`     //通过主键修改
	IsLink int  `p:"isLink" v:"required#状态必须"` //通过主键获取
}
type SignAboutIsLinkSwitchRes struct {
	commonApi.EmptyRes
}
