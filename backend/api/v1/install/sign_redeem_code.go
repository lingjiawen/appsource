// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-03-25 13:14:15
// 生成路径: api/v1/install/sign_redeem_code.go
// 生成人：smithy
// desc:签名卡密相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package install

import (
	commonApi "mangosmithy/api/v1/common"
	"mangosmithy/internal/app/install/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// SignRedeemCodeSearchReq 分页请求参数
type SignRedeemCodeSearchReq struct {
	g.Meta `path:"/list" tags:"签名卡密" method:"get" summary:"签名卡密列表"`
	commonApi.Author
	model.SignRedeemCodeSearchReq
}

// SignRedeemCodeSearchRes 列表返回结果
type SignRedeemCodeSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.SignRedeemCodeSearchRes
}

// SignRedeemCodeExportReq 导出请求
type SignRedeemCodeExportReq struct {
	g.Meta `path:"/export" tags:"签名卡密" method:"get" summary:"签名卡密导出"`
	commonApi.Author
	model.SignRedeemCodeSearchReq
}

// SignRedeemCodeExportRes 导出响应
type SignRedeemCodeExportRes struct {
	commonApi.EmptyRes
}
type SignRedeemCodeExcelTemplateReq struct {
	g.Meta `path:"/excelTemplate" tags:"签名卡密" method:"get" summary:"导出模板文件"`
	commonApi.Author
}
type SignRedeemCodeExcelTemplateRes struct {
	commonApi.EmptyRes
}
type SignRedeemCodeImportReq struct {
	g.Meta `path:"/import" tags:"签名卡密" method:"post" summary:"签名卡密导入"`
	commonApi.Author
	File *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件"  v:"required#上传文件必须"`
}
type SignRedeemCodeImportRes struct {
	commonApi.EmptyRes
}

// SignRedeemCodeAddReq 添加操作请求参数
type SignRedeemCodeAddReq struct {
	g.Meta `path:"/add" tags:"签名卡密" method:"post" summary:"签名卡密添加"`
	commonApi.Author
	*model.SignRedeemCodeAddReq
}

// SignRedeemCodeAddRes 添加操作返回结果
type SignRedeemCodeAddRes struct {
	commonApi.EmptyRes
}

// SignRedeemCodeEditReq 修改操作请求参数
type SignRedeemCodeEditReq struct {
	g.Meta `path:"/edit" tags:"签名卡密" method:"put" summary:"签名卡密修改"`
	commonApi.Author
	*model.SignRedeemCodeEditReq
}

// SignRedeemCodeEditRes 修改操作返回结果
type SignRedeemCodeEditRes struct {
	commonApi.EmptyRes
}

// SignRedeemCodeGetReq 获取一条数据请求
type SignRedeemCodeGetReq struct {
	g.Meta `path:"/get" tags:"签名卡密" method:"get" summary:"获取签名卡密信息"`
	commonApi.Author
	Id uint `p:"id" v:"required#主键必须"` //通过主键获取
}

// SignRedeemCodeGetRes 获取一条数据结果
type SignRedeemCodeGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SignRedeemCodeInfoRes
}

// SignRedeemCodeDeleteReq 删除数据请求
type SignRedeemCodeDeleteReq struct {
	g.Meta `path:"/delete" tags:"签名卡密" method:"delete" summary:"删除签名卡密"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}

// SignRedeemCodeDeleteRes 删除数据返回
type SignRedeemCodeDeleteRes struct {
	commonApi.EmptyRes
}

// 签名卡密禁用修改（状态）
type SignRedeemCodeBannedSwitchReq struct {
	g.Meta `path:"/changeBanned" tags:"签名卡密" method:"put" summary:"修改禁用"`
	commonApi.Author
	Id     uint `p:"id" v:"required#主键必须"`     //通过主键修改
	Banned int  `p:"banned" v:"required#状态必须"` //通过主键获取
}
type SignRedeemCodeBannedSwitchRes struct {
	commonApi.EmptyRes
}

// 签名卡密激活修改（状态）
type SignRedeemCodeActiveSwitchReq struct {
	g.Meta `path:"/changeActive" tags:"签名卡密" method:"put" summary:"修改激活"`
	commonApi.Author
	Id     uint `p:"id" v:"required#主键必须"`     //通过主键修改
	Active int  `p:"active" v:"required#状态必须"` //通过主键获取
}
type SignRedeemCodeActiveSwitchRes struct {
	commonApi.EmptyRes
}
