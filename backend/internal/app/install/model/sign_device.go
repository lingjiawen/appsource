// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-04-01 19:38:11
// 生成路径: internal/app/install/model/sign_device.go
// 生成人：smithy
// desc:设备管理
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	comModel "mangosmithy/internal/app/common/model"
	systemModel "mangosmithy/internal/app/system/model"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SignDeviceInfoRes is the golang structure for table sign_device.
type SignDeviceInfoRes struct {
	gmeta.Meta      `orm:"table:sign_device"`
	Id              uint                     `orm:"id,primary" json:"id" dc:"ID"`                         // ID
	Udid            string                   `orm:"udid" json:"udid" dc:"设备码"`                            // 设备码
	Name            string                   `orm:"name" json:"name" dc:"证书名"`                            // 证书名
	CertId          string                   `orm:"cert_id" json:"certId" dc:"证书标识"`                      // 证书标识
	DeviceId        string                   `orm:"device_id" json:"deviceId" dc:"设备标识"`                  // 设备标识
	P12             string                   `orm:"p12" json:"p12" dc:"证书文件"`                             // 证书文件
	Mobileprovision string                   `orm:"mobileprovision" json:"mobileprovision" dc:"描述文件"`     // 描述文件
	AddTime         uint64                   `orm:"add_time" json:"addTime" dc:"添加时间"`                    // 添加时间
	P12Password     string                   `orm:"p12_password" json:"p12Password" dc:"证书密码"`            // 证书密码
	Model           string                   `orm:"model" json:"model" dc:"设备型号"`                         // 设备型号
	ExpireTime      uint64                   `orm:"expire_time" json:"expireTime" dc:"过期时间"`              // 过期时间
	SerialNumber    string                   `orm:"serial_number" json:"serialNumber" dc:"序列号"`           // 序列号
	RedeemCode      string                   `orm:"redeem_code" json:"redeemCode" dc:"兑换卡密"`              // 兑换卡密
	AccountType     int                      `orm:"account_type" json:"accountType" dc:"时效类型"`            // 时效类型
	WarrantyType    int                      `orm:"warranty_type" json:"warrantyType" dc:"售后类型"`          // 售后类型
	DeviceType      string                   `orm:"device_type" json:"deviceType" dc:"设备类型"`              // 设备类型
	Status          string                   `orm:"status" json:"status" dc:"状态"`                         // 状态
	Pool            int                      `orm:"pool" json:"pool" dc:"出书方式"`                           // 出书方式
	ApiPlatform     int                      `orm:"api_platform" json:"apiPlatform" dc:"对接平台"`            // 对接平台
	ApiWarrantyType int                      `orm:"api_warranty_type" json:"apiWarrantyType" dc:"对接售后类型"` // 对接售后类型
	Active          int                      `orm:"active" json:"active" dc:"禁用"`                         // 禁用
	CreatedBy       uint                     `orm:"created_by" json:"createdBy" dc:"创建人"`                 // 创建人
	CreatedUser     *systemModel.LinkUserRes `orm:"with:id=created_by" json:"createdUser"`
	UpdatedBy       uint                     `orm:"updated_by" json:"updatedBy" dc:"修改人"` // 修改人
	UpdatedUser     *systemModel.LinkUserRes `orm:"with:id=updated_by" json:"updatedUser"`
	CreatedAt       *gtime.Time              `orm:"created_at" json:"createdAt" dc:"创建时间"` // 创建时间
	UpdatedAt       *gtime.Time              `orm:"updated_at" json:"updatedAt" dc:"修改时间"` // 修改时间
	DeletedAt       *gtime.Time              `orm:"deleted_at" json:"deletedAt" dc:"删除时间"` // 删除时间
}

type SignDeviceListRes struct {
	Id              uint                     `json:"id" dc:"ID"`
	Udid            string                   `json:"udid" dc:"设备码"`
	Name            string                   `json:"name" dc:"证书名"`
	CertId          string                   `json:"certId" dc:"证书标识"`
	AddTime         string                   `json:"addTime" dc:"添加时间"`
	Model           string                   `json:"model" dc:"设备型号"`
	ExpireTime      string                   `json:"expireTime" dc:"过期时间"`
	RedeemCode      string                   `json:"redeemCode" dc:"兑换卡密"`
	AccountType     int                      `json:"accountType" dc:"时效类型"`
	WarrantyType    int                      `json:"warrantyType" dc:"售后类型"`
	DeviceType      string                   `json:"deviceType" dc:"设备类型"`
	Status          string                   `json:"status" dc:"状态"`
	Pool            int                      `json:"pool" dc:"出书方式"`
	ApiPlatform     int                      `json:"apiPlatform" dc:"对接平台"`
	ApiWarrantyType int                      `json:"apiWarrantyType" dc:"对接售后类型"`
	Active          int                      `json:"active" dc:"禁用"`
	CreatedUser     *systemModel.LinkUserRes `orm:"with:id=created_by" json:"createdUser"`
	CreatedBy       uint                     `json:"createdBy" dc:"创建人"`
	CreatedAt       *gtime.Time              `json:"createdAt" dc:"创建时间"`
}

// SignDeviceSearchReq 分页请求参数
type SignDeviceSearchReq struct {
	comModel.PageReq
	Id              string `p:"id" dc:"ID"`                                                             //ID
	Udid            string `p:"udid" dc:"设备码"`                                                          //设备码
	Name            string `p:"name" dc:"证书名"`                                                          //证书名
	CertId          string `p:"certId" dc:"证书标识"`                                                       //证书标识
	DeviceId        string `p:"deviceId" dc:"设备标识"`                                                     //设备标识
	AddTime         string `p:"addTime" v:"addTime@integer#添加时间需为整数" dc:"添加时间"`                         //添加时间
	Model           string `p:"model" dc:"设备型号"`                                                        //设备型号
	ExpireTime      string `p:"expireTime" v:"expireTime@integer#过期时间需为整数" dc:"过期时间"`                   //过期时间
	RedeemCode      string `p:"redeemCode" dc:"兑换卡密"`                                                   //兑换卡密
	AccountType     string `p:"accountType" v:"accountType@integer#时效类型需为整数" dc:"时效类型"`                 //时效类型
	WarrantyType    string `p:"warrantyType" v:"warrantyType@integer#售后类型需为整数" dc:"售后类型"`               //售后类型
	DeviceType      string `p:"deviceType" dc:"设备类型"`                                                   //设备类型
	Status          string `p:"status" dc:"状态"`                                                         //状态
	Pool            string `p:"pool" v:"pool@integer#出书方式需为整数" dc:"出书方式"`                               //出书方式
	ApiPlatform     string `p:"apiPlatform" v:"apiPlatform@integer#对接平台需为整数" dc:"对接平台"`                 //对接平台
	ApiWarrantyType string `p:"apiWarrantyType" v:"apiWarrantyType@integer#对接售后类型需为整数" dc:"对接售后类型"`     //对接售后类型
	Active          string `p:"active" v:"active@integer#禁用需为整数" dc:"禁用"`                               //禁用
	CreatedBy       string `p:"createdBy" v:"createdBy@integer#创建人需为整数" dc:"创建人"`                       //创建人
	CreatedAt       string `p:"createdAt" v:"createdAt@datetime#创建时间需为YYYY-MM-DD hh:mm:ss格式" dc:"创建时间"` //创建时间
}

// SignDeviceSearchRes 列表返回结果
type SignDeviceSearchRes struct {
	comModel.ListRes
	List []*SignDeviceListRes `json:"list"`
}

// SignDeviceAddReq 添加操作请求参数
type SignDeviceAddReq struct {
	Udid            string `p:"udid"  dc:"设备码"`
	Name            string `p:"name" v:"required#证书名不能为空" dc:"证书名"`
	CertId          string `p:"certId"  dc:"证书标识"`
	DeviceId        string `p:"deviceId"  dc:"设备标识"`
	P12             string `p:"p12"  dc:"证书文件"`
	Mobileprovision string `p:"mobileprovision"  dc:"描述文件"`
	AddTime         uint64 `p:"addTime"  dc:"添加时间"`
	Model           string `p:"model"  dc:"设备型号"`
	ExpireTime      uint64 `p:"expireTime"  dc:"过期时间"`
	RedeemCode      string `p:"redeemCode"  dc:"兑换卡密"`
	AccountType     int    `p:"accountType"  dc:"时效类型"`
	WarrantyType    int    `p:"warrantyType"  dc:"售后类型"`
	DeviceType      string `p:"deviceType"  dc:"设备类型"`
	Status          string `p:"status" v:"required#状态不能为空" dc:"状态"`
	Pool            int    `p:"pool"  dc:"出书方式"`
	ApiPlatform     int    `p:"apiPlatform"  dc:"对接平台"`
	ApiWarrantyType int    `p:"apiWarrantyType"  dc:"对接售后类型"`
	Active          int    `p:"active"  dc:"禁用"`
	CreatedBy       uint64
}

// SignDeviceEditReq 修改操作请求参数
type SignDeviceEditReq struct {
	Id              uint   `p:"id" v:"required#主键ID不能为空" dc:"ID"`
	Udid            string `p:"udid"  dc:"设备码"`
	Name            string `p:"name" v:"required#证书名不能为空" dc:"证书名"`
	CertId          string `p:"certId"  dc:"证书标识"`
	DeviceId        string `p:"deviceId"  dc:"设备标识"`
	P12             string `p:"p12"  dc:"证书文件"`
	Mobileprovision string `p:"mobileprovision"  dc:"描述文件"`
	AddTime         uint64 `p:"addTime"  dc:"添加时间"`
	Model           string `p:"model"  dc:"设备型号"`
	ExpireTime      uint64 `p:"expireTime"  dc:"过期时间"`
	RedeemCode      string `p:"redeemCode"  dc:"兑换卡密"`
	AccountType     int    `p:"accountType"  dc:"时效类型"`
	WarrantyType    int    `p:"warrantyType"  dc:"售后类型"`
	DeviceType      string `p:"deviceType"  dc:"设备类型"`
	Status          string `p:"status" v:"required#状态不能为空" dc:"状态"`
	Pool            int    `p:"pool"  dc:"出书方式"`
	ApiPlatform     int    `p:"apiPlatform"  dc:"对接平台"`
	ApiWarrantyType int    `p:"apiWarrantyType"  dc:"对接售后类型"`
	Active          int    `p:"active"  dc:"禁用"`
	UpdatedBy       uint64
}
