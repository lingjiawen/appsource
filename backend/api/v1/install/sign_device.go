// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-04-01 19:38:11
// 生成路径: api/v1/install/sign_device.go
// 生成人：smithy
// desc:设备管理相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package install

import (
	commonApi "mangosmithy/api/v1/common"
	"mangosmithy/internal/app/install/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// SignDeviceSearchReq 分页请求参数
type SignDeviceSearchReq struct {
	g.Meta `path:"/list" tags:"设备管理" method:"get" summary:"设备管理列表"`
	commonApi.Author
	model.SignDeviceSearchReq
}

// SignDeviceSearchRes 列表返回结果
type SignDeviceSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.SignDeviceSearchRes
}

// SignDeviceExportReq 导出请求
type SignDeviceExportReq struct {
	g.Meta `path:"/export" tags:"设备管理" method:"get" summary:"设备管理导出"`
	commonApi.Author
	model.SignDeviceSearchReq
}

// SignDeviceExportRes 导出响应
type SignDeviceExportRes struct {
	commonApi.EmptyRes
}
type SignDeviceExcelTemplateReq struct {
	g.Meta `path:"/excelTemplate" tags:"设备管理" method:"get" summary:"导出模板文件"`
	commonApi.Author
}
type SignDeviceExcelTemplateRes struct {
	commonApi.EmptyRes
}
type SignDeviceImportReq struct {
	g.Meta `path:"/import" tags:"设备管理" method:"post" summary:"设备管理导入"`
	commonApi.Author
	File *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件"  v:"required#上传文件必须"`
}
type SignDeviceImportRes struct {
	commonApi.EmptyRes
}

// SignDeviceAddReq 添加操作请求参数
type SignDeviceAddReq struct {
	g.Meta `path:"/add" tags:"设备管理" method:"post" summary:"设备管理添加"`
	commonApi.Author
	*model.SignDeviceAddReq
}

// SignDeviceAddRes 添加操作返回结果
type SignDeviceAddRes struct {
	commonApi.EmptyRes
}

// SignDeviceEditReq 修改操作请求参数
type SignDeviceEditReq struct {
	g.Meta `path:"/edit" tags:"设备管理" method:"put" summary:"设备管理修改"`
	commonApi.Author
	*model.SignDeviceEditReq
}

// SignDeviceEditRes 修改操作返回结果
type SignDeviceEditRes struct {
	commonApi.EmptyRes
}

// SignDeviceGetReq 获取一条数据请求
type SignDeviceGetReq struct {
	g.Meta `path:"/get" tags:"设备管理" method:"get" summary:"获取设备管理信息"`
	commonApi.Author
	Id uint `p:"id" v:"required#主键必须"` //通过主键获取
}

// SignDeviceGetRes 获取一条数据结果
type SignDeviceGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SignDeviceInfoRes
}

// SignDeviceDeleteReq 删除数据请求
type SignDeviceDeleteReq struct {
	g.Meta `path:"/delete" tags:"设备管理" method:"delete" summary:"删除设备管理"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}

// SignDeviceDeleteRes 删除数据返回
type SignDeviceDeleteRes struct {
	commonApi.EmptyRes
}

// 设备管理禁用修改（状态）
type SignDeviceActiveSwitchReq struct {
	g.Meta `path:"/changeActive" tags:"设备管理" method:"put" summary:"修改禁用"`
	commonApi.Author
	Id     uint `p:"id" v:"required#主键必须"`     //通过主键修改
	Active int  `p:"active" v:"required#状态必须"` //通过主键获取
}
type SignDeviceActiveSwitchRes struct {
	commonApi.EmptyRes
}
