// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-02-21 18:37:33
// 生成路径: api/v1/appsource/as_application.go
// 生成人：smithy
// desc:应用管理相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package appsource

import (
	commonApi "mangosmithy/api/v1/common"
	"mangosmithy/internal/app/appsource/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// AsApplicationSearchReq 分页请求参数
type AsApplicationSearchReq struct {
	g.Meta `path:"/list" tags:"应用管理" method:"get" summary:"应用管理列表"`
	commonApi.Author
	model.AsApplicationSearchReq
}

// AsApplicationSearchRes 列表返回结果
type AsApplicationSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.AsApplicationSearchRes
}

// AsApplicationExportReq 导出请求
type AsApplicationExportReq struct {
	g.Meta `path:"/export" tags:"应用管理" method:"get" summary:"应用管理导出"`
	commonApi.Author
	model.AsApplicationSearchReq
}

// AsApplicationExportRes 导出响应
type AsApplicationExportRes struct {
	commonApi.EmptyRes
}
type AsApplicationExcelTemplateReq struct {
	g.Meta `path:"/excelTemplate" tags:"应用管理" method:"get" summary:"导出模板文件"`
	commonApi.Author
}
type AsApplicationExcelTemplateRes struct {
	commonApi.EmptyRes
}
type AsApplicationImportReq struct {
	g.Meta `path:"/import" tags:"应用管理" method:"post" summary:"应用管理导入"`
	commonApi.Author
	File *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件"  v:"required#上传文件必须"`
}
type AsApplicationImportRes struct {
	commonApi.EmptyRes
}

// AsApplicationAddReq 添加操作请求参数
type AsApplicationAddReq struct {
	g.Meta `path:"/add" tags:"应用管理" method:"post" summary:"应用管理添加"`
	commonApi.Author
	*model.AsApplicationAddReq
}

// AsApplicationAddRes 添加操作返回结果
type AsApplicationAddRes struct {
	commonApi.EmptyRes
}

// AsApplicationEditReq 修改操作请求参数
type AsApplicationEditReq struct {
	g.Meta `path:"/edit" tags:"应用管理" method:"put" summary:"应用管理修改"`
	commonApi.Author
	*model.AsApplicationEditReq
}

// AsApplicationEditRes 修改操作返回结果
type AsApplicationEditRes struct {
	commonApi.EmptyRes
}

// AsApplicationGetReq 获取一条数据请求
type AsApplicationGetReq struct {
	g.Meta `path:"/get" tags:"应用管理" method:"get" summary:"获取应用管理信息"`
	commonApi.Author
	Id uint `p:"id" v:"required#主键必须"` //通过主键获取
}

// AsApplicationGetRes 获取一条数据结果
type AsApplicationGetRes struct {
	g.Meta `mime:"application/json"`
	*model.AsApplicationInfoRes
}

// AsApplicationDeleteReq 删除数据请求
type AsApplicationDeleteReq struct {
	g.Meta `path:"/delete" tags:"应用管理" method:"delete" summary:"删除应用管理"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}

// AsApplicationDeleteRes 删除数据返回
type AsApplicationDeleteRes struct {
	commonApi.EmptyRes
}

// AsApplicationLockSwitchReq 应用管理付费修改（状态）
type AsApplicationLockSwitchReq struct {
	g.Meta `path:"/changeLock" tags:"应用管理" method:"put" summary:"修改付费"`
	commonApi.Author
	Id   uint `p:"id" v:"required#主键必须"`   //通过主键修改
	Lock int  `p:"lock" v:"required#状态必须"` //通过主键获取
}
type AsApplicationLockSwitchRes struct {
	commonApi.EmptyRes
}

// AsApplicationActiveSwitchReq 应用管理是否启用修改（状态）
type AsApplicationActiveSwitchReq struct {
	g.Meta `path:"/changeActive" tags:"应用管理" method:"put" summary:"修改是否启用"`
	commonApi.Author
	Id     uint `p:"id" v:"required#主键必须"`     //通过主键修改
	Active int  `p:"active" v:"required#状态必须"` //通过主键获取
}
type AsApplicationActiveSwitchRes struct {
	commonApi.EmptyRes
}

// AsApplicationExtractReq 解析软件请求参数
type AsApplicationExtractReq struct {
	g.Meta `path:"/extract" tags:"应用管理" method:"post" summary:"解析APP"`
	commonApi.Author
	*model.AsApplicationExtractReq
}

// AsApplicationExtractRes 解析软件返回结果
type AsApplicationExtractRes struct {
	g.Meta `mime:"application/json"`
	*model.AsApplicationExtractRes
}

// AsApplicationIconSearchReq 软件图标搜索请求参数
type AsApplicationIconSearchReq struct {
	g.Meta `path:"/search" tags:"应用管理" method:"post" summary:"APP图标搜索"`
	commonApi.Author
	*model.AsApplicationIconSearchReq
}

// AsApplicationIconSearchRes 软件图标搜索返回结果
type AsApplicationIconSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.EmptyRes
}

// AsApplicationSourceImportReq 软件源导入请求参数
type AsApplicationSourceImportReq struct {
	g.Meta `path:"/sourceImport" tags:"应用管理" method:"post" summary:"软件源导入"`
	commonApi.Author
	*model.AsApplicationSourceImportReq
}

// AsApplicationSourceImportRes 软件源导入返回结果
type AsApplicationSourceImportRes struct {
	g.Meta `mime:"application/json"`
	commonApi.EmptyRes
}
