// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-04-12 18:50:08
// 生成路径: api/v1/install/sign_platform.go
// 生成人：smithy
// desc:平台相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package install

import (
	commonApi "mangosmithy/api/v1/common"
	"mangosmithy/internal/app/install/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// SignPlatformSearchReq 分页请求参数
type SignPlatformSearchReq struct {
	g.Meta `path:"/list" tags:"平台" method:"get" summary:"平台列表"`
	commonApi.Author
	model.SignPlatformSearchReq
}

// SignPlatformSearchRes 列表返回结果
type SignPlatformSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.SignPlatformSearchRes
}

// SignPlatformExportReq 导出请求
type SignPlatformExportReq struct {
	g.Meta `path:"/export" tags:"平台" method:"get" summary:"平台导出"`
	commonApi.Author
	model.SignPlatformSearchReq
}

// SignPlatformExportRes 导出响应
type SignPlatformExportRes struct {
	commonApi.EmptyRes
}
type SignPlatformExcelTemplateReq struct {
	g.Meta `path:"/excelTemplate" tags:"平台" method:"get" summary:"导出模板文件"`
	commonApi.Author
}
type SignPlatformExcelTemplateRes struct {
	commonApi.EmptyRes
}
type SignPlatformImportReq struct {
	g.Meta `path:"/import" tags:"平台" method:"post" summary:"平台导入"`
	commonApi.Author
	File *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件"  v:"required#上传文件必须"`
}
type SignPlatformImportRes struct {
	commonApi.EmptyRes
}

// SignPlatformAddReq 添加操作请求参数
type SignPlatformAddReq struct {
	g.Meta `path:"/add" tags:"平台" method:"post" summary:"平台添加"`
	commonApi.Author
	*model.SignPlatformAddReq
}

// SignPlatformAddRes 添加操作返回结果
type SignPlatformAddRes struct {
	commonApi.EmptyRes
}

// SignPlatformEditReq 修改操作请求参数
type SignPlatformEditReq struct {
	g.Meta `path:"/edit" tags:"平台" method:"put" summary:"平台修改"`
	commonApi.Author
	*model.SignPlatformEditReq
}

// SignPlatformEditRes 修改操作返回结果
type SignPlatformEditRes struct {
	commonApi.EmptyRes
}

// SignPlatformGetReq 获取一条数据请求
type SignPlatformGetReq struct {
	g.Meta `path:"/get" tags:"平台" method:"get" summary:"获取平台信息"`
	commonApi.Author
	Id uint `p:"id" v:"required#主键必须"` //通过主键获取
}

// SignPlatformGetRes 获取一条数据结果
type SignPlatformGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SignPlatformInfoRes
}

// SignPlatformDeleteReq 删除数据请求
type SignPlatformDeleteReq struct {
	g.Meta `path:"/delete" tags:"平台" method:"delete" summary:"删除平台"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}

// SignPlatformDeleteRes 删除数据返回
type SignPlatformDeleteRes struct {
	commonApi.EmptyRes
}

// 平台开启SSL修改（状态）
type SignPlatformOpenSslSwitchReq struct {
	g.Meta `path:"/changeOpenSsl" tags:"平台" method:"put" summary:"修改开启SSL"`
	commonApi.Author
	Id      uint `p:"id" v:"required#主键必须"`      //通过主键修改
	OpenSsl int  `p:"openSsl" v:"required#状态必须"` //通过主键获取
}
type SignPlatformOpenSslSwitchRes struct {
	commonApi.EmptyRes
}

// 平台启用修改（状态）
type SignPlatformStatusSwitchReq struct {
	g.Meta `path:"/changeStatus" tags:"平台" method:"put" summary:"修改启用"`
	commonApi.Author
	Id     uint `p:"id" v:"required#主键必须"`     //通过主键修改
	Status int  `p:"status" v:"required#状态必须"` //通过主键获取
}
type SignPlatformStatusSwitchRes struct {
	commonApi.EmptyRes
}
