// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-03-25 13:14:15
// 生成路径: internal/app/install/model/sign_redeem_code.go
// 生成人：smithy
// desc:签名卡密
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	comModel "mangosmithy/internal/app/common/model"
	systemModel "mangosmithy/internal/app/system/model"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SignRedeemCodeInfoRes is the golang structure for table sign_redeem_code.
type SignRedeemCodeInfoRes struct {
	gmeta.Meta      `orm:"table:sign_redeem_code"`
	Id              uint                     `orm:"id,primary" json:"id" dc:"ID"`                         // ID
	Code            string                   `orm:"code" json:"code" dc:"兑换码"`                            // 兑换码
	Udid            string                   `orm:"udid" json:"udid" dc:"设备码"`                            // 设备码
	CertId          string                   `orm:"cert_id" json:"certId" dc:"证书标识"`                      // 证书标识
	AccountType     int                      `orm:"account_type" json:"accountType" dc:"时效类型"`            // 时效类型
	WarrantyType    int                      `orm:"warranty_type" json:"warrantyType" dc:"售后类型"`          // 售后类型
	DeviceType      string                   `orm:"device_type" json:"deviceType" dc:"设备类型"`              // 设备类型
	Pool            int                      `orm:"pool" json:"pool" dc:"出书方式"`                           // 出书方式
	ApiPlatform     int                      `orm:"api_platform" json:"apiPlatform" dc:"对接平台"`            // 对接平台
	Note            string                   `orm:"note" json:"note" dc:"备注"`                             // 备注
	ApiWarrantyType int                      `orm:"api_warranty_type" json:"apiWarrantyType" dc:"对接售后类型"` // 对接售后类型
	Force           int                      `orm:"force" json:"force" dc:"强制添加"`                         // 强制
	Banned          int                      `orm:"banned" json:"banned" dc:"禁用"`                         // 禁用
	Active          int                      `orm:"active" json:"active" dc:"激活"`                         // 激活
	ActiveAt        *gtime.Time              `orm:"active_at" json:"activeAt" dc:"激活时间"`                  // 激活时间
	CreatedBy       uint                     `orm:"created_by" json:"createdBy" dc:"创建人"`                 // 创建人
	CreatedUser     *systemModel.LinkUserRes `orm:"with:id=created_by" json:"createdUser"`
	UpdatedBy       uint                     `orm:"updated_by" json:"updatedBy" dc:"修改人"` // 修改人
	UpdatedUser     *systemModel.LinkUserRes `orm:"with:id=updated_by" json:"updatedUser"`
	CreatedAt       *gtime.Time              `orm:"created_at" json:"createdAt" dc:"创建时间"` // 创建时间
	UpdatedAt       *gtime.Time              `orm:"updated_at" json:"updatedAt" dc:"修改时间"` // 修改时间
	DeletedAt       *gtime.Time              `orm:"deleted_at" json:"deletedAt" dc:"删除时间"` // 删除时间
}

type SignRedeemCodeListRes struct {
	Id              uint        `json:"id" dc:"ID"`
	Code            string      `json:"code" dc:"兑换码"`
	Udid            string      `json:"udid" dc:"设备码"`
	CertId          string      `json:"certId" dc:"证书标识"`
	AccountType     int         `json:"accountType" dc:"时效类型"`
	WarrantyType    int         `json:"warrantyType" dc:"售后类型"`
	DeviceType      string      `json:"deviceType" dc:"设备类型"`
	Pool            int         `json:"pool" dc:"出书方式"`
	ApiPlatform     int         `json:"apiPlatform" dc:"对接平台"`
	Note            string      `json:"note" dc:"备注"`
	ApiWarrantyType int         `json:"apiWarrantyType" dc:"对接售后类型"`
	Force           int         `json:"force" dc:"强制添加"`
	Banned          int         `json:"banned" dc:"禁用"`
	Active          int         `json:"active" dc:"激活"`
	ActiveAt        *gtime.Time `json:"activeAt" dc:"激活时间"`
	CreatedAt       *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// SignRedeemCodeSearchReq 分页请求参数
type SignRedeemCodeSearchReq struct {
	comModel.PageReq
	Id              string `p:"id" dc:"ID"`                                                             //ID
	Code            string `p:"code" dc:"兑换码"`                                                          //兑换码
	Udid            string `p:"udid" dc:"设备码"`                                                          //设备码
	CertId          string `p:"certId" dc:"证书标识"`                                                       //证书标识
	AccountType     string `p:"accountType" v:"accountType@integer#时效类型需为整数" dc:"时效类型"`                 //时效类型
	WarrantyType    string `p:"warrantyType" v:"warrantyType@integer#售后类型需为整数" dc:"售后类型"`               //售后类型
	DeviceType      string `p:"deviceType" dc:"设备类型"`                                                   //设备类型
	Pool            string `p:"pool" v:"pool@integer#出书方式需为整数" dc:"出书方式"`                               //出书方式
	ApiPlatform     string `p:"apiPlatform" v:"apiPlatform@integer#对接平台需为整数" dc:"对接平台"`                 //对接平台
	ApiWarrantyType string `p:"apiWarrantyType" v:"apiWarrantyType@integer#对接售后类型需为整数" dc:"对接售后类型"`     //对接售后类型
	Force           string `p:"force" v:"force@integer#强制添加需为整数" dc:"强制添加"`                             //强制添加
	Note            string `p:"note" dc:"备注"`                                                           //备注
	Banned          string `p:"banned" v:"banned@integer#禁用需为整数" dc:"禁用"`                               //禁用
	Active          string `p:"active" v:"active@integer#激活需为整数" dc:"激活"`                               //激活
	ActiveAt        string `p:"activeAt" v:"activeAt@datetime#激活时间需为YYYY-MM-DD hh:mm:ss格式" dc:"激活时间"`   //激活时间
	CreatedBy       string `p:"createdBy" v:"createdBy@integer#创建人需为整数" dc:"创建人"`                       //创建人
	CreatedAt       string `p:"createdAt" v:"createdAt@datetime#创建时间需为YYYY-MM-DD hh:mm:ss格式" dc:"创建时间"` //创建时间
}

// SignRedeemCodeSearchRes 列表返回结果
type SignRedeemCodeSearchRes struct {
	comModel.ListRes
	List []*SignRedeemCodeListRes `json:"list"`
}

// SignRedeemCodeAddReq 添加操作请求参数
type SignRedeemCodeAddReq struct {
	Prefix          string `p:"prefix"  dc:"前缀"`
	Quantity        int    `p:"quantity"  dc:"数量"`
	AccountType     int    `p:"accountType"  dc:"时效类型"`
	WarrantyType    int    `p:"warrantyType"  dc:"售后类型"`
	DeviceType      string `p:"deviceType"  dc:"设备类型"`
	Pool            int    `p:"pool"  dc:"出书方式"`
	ApiPlatform     int    `p:"apiPlatform"  dc:"对接平台"`
	ApiWarrantyType int    `p:"apiWarrantyType"  dc:"对接售后类型"`
	Force           int    `p:"force"  dc:"强制添加"`
	Note            string `p:"note"  dc:"备注"`
}

// SignRedeemCodeEditReq 修改操作请求参数
type SignRedeemCodeEditReq struct {
	Id              uint        `p:"id" v:"required#主键ID不能为空" dc:"ID"`
	Code            string      `p:"code"  dc:"兑换码"`
	Udid            string      `p:"udid"  dc:"设备码"`
	CertId          string      `p:"certId"  dc:"证书标识"`
	AccountType     int         `p:"accountType"  dc:"时效类型"`
	WarrantyType    int         `p:"warrantyType"  dc:"售后类型"`
	DeviceType      string      `p:"deviceType"  dc:"设备类型"`
	Pool            int         `p:"pool"  dc:"出书方式"`
	ApiPlatform     int         `p:"apiPlatform"  dc:"对接平台"`
	Note            string      `p:"note"  dc:"备注"`
	ApiWarrantyType int         `p:"apiWarrantyType"  dc:"对接售后类型"`
	Force           int         `p:"force"  dc:"强制添加"`
	Banned          int         `p:"banned"  dc:"禁用"`
	Active          int         `p:"active"  dc:"激活"`
	ActiveAt        *gtime.Time `p:"activeAt"  dc:"激活时间"`
	UpdatedBy       uint64
}
