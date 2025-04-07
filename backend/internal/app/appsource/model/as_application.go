// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-02-21 18:37:33
// 生成路径: internal/app/appsource/model/as_application.go
// 生成人：smithy
// desc:应用管理
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	comModel "mangosmithy/internal/app/common/model"
	systemModel "mangosmithy/internal/app/system/model"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// AsApplicationInfoRes is the golang structure for table as_application.
type AsApplicationInfoRes struct {
	gmeta.Meta  `orm:"table:as_application"`
	Id          uint                     `orm:"id,primary" json:"id" dc:"ID"`           // ID
	FileUrl     string                   `orm:"file_url" json:"fileUrl" dc:"文件"`        // 文件
	Name        string                   `orm:"name" json:"name" dc:"应用名称"`             // 应用名称
	BundleId    string                   `orm:"bundle_id" json:"bundleId" dc:"包名"`      // 包名
	Version     string                   `orm:"version" json:"version" dc:"版本号"`        // 版本号
	Size        int                      `orm:"size" json:"size" dc:"文件大小"`             // 文件大小
	Type        int                      `orm:"type" json:"type" dc:"类型"`               // 类型
	Description string                   `orm:"description" json:"description" dc:"描述"` // 描述
	IconUrl     string                   `orm:"icon_url" json:"iconUrl" dc:"图标"`        // 图标
	Lock        int                      `orm:"lock" json:"lock" dc:"付费"`               // 付费
	Lanzou      int                      `orm:"lanzou" json:"lanzou" dc:"是否蓝奏云链接"`      // 是否蓝奏云链接
	Weigh       int                      `orm:"weigh" json:"weigh" dc:"权重"`             // 权重
	Active      int                      `orm:"active" json:"active" dc:"是否启用"`         // 是否启用
	Note        string                   `orm:"note" json:"note" dc:"备注"`               // 备注
	CreatedBy   uint                     `orm:"created_by" json:"createdBy" dc:"创建人"`   // 创建人
	CreatedUser *systemModel.LinkUserRes `orm:"with:id=created_by" json:"createdUser"`
	UpdatedBy   uint                     `orm:"updated_by" json:"updatedBy" dc:"修改人"` // 修改人
	UpdatedUser *systemModel.LinkUserRes `orm:"with:id=updated_by" json:"updatedUser"`
	CreatedAt   *gtime.Time              `orm:"created_at" json:"createdAt" dc:"创建时间"` // 创建时间
	UpdatedAt   *gtime.Time              `orm:"updated_at" json:"updatedAt" dc:"修改时间"` // 修改时间
}

type AsApplicationListRes struct {
	Id          uint        `json:"id" dc:"ID"`
	FileUrl     string      `json:"fileUrl" dc:"文件"`
	Name        string      `json:"name" dc:"应用名称"`
	BundleId    string      `json:"bundleId" dc:"包名"`
	Version     string      `json:"version" dc:"版本号"`
	Size        int         `json:"size" dc:"文件大小"`
	Type        int         `json:"type" dc:"类型"`
	Description string      `json:"description" dc:"描述"`
	IconUrl     string      `json:"iconUrl" dc:"图标"`
	Lock        int         `json:"lock" dc:"付费"`
	Lanzou      int         `json:"lanzou" dc:"是否蓝奏云链接"`
	Weigh       int         `json:"weigh" dc:"权重"`
	Active      int         `json:"active" dc:"是否启用"`
	Note        string      `json:"note" dc:"备注"`
	CreatedAt   *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt" dc:"修改时间"` // 修改时间
}

// AsApplicationSearchReq 分页请求参数
type AsApplicationSearchReq struct {
	comModel.PageReq
	Id          string `p:"id" dc:"ID"`                                                             //ID
	FileUrl     string `p:"fileUrl" dc:"文件"`                                                        //文件
	Name        string `p:"name" dc:"应用名称"`                                                         //应用名称
	BundleId    string `p:"bundleId" dc:"包名"`                                                       //包名
	Version     string `p:"version" dc:"版本号"`                                                       //版本号
	Size        string `p:"size" v:"size@integer#文件大小需为整数" dc:"文件大小"`                               //文件大小
	Type        string `p:"type" v:"type@integer#类型需为整数" dc:"类型"`                                   //类型
	Description string `p:"description" dc:"描述"`                                                    //描述
	IconUrl     string `p:"iconUrl" dc:"图标"`                                                        //图标
	Lock        string `p:"lock" v:"lock@integer#付费需为整数" dc:"付费"`                                   //付费
	Lanzou      string `p:"lanzou" v:"lanzou@integer#是否蓝奏云链接需为整数" dc:"是否蓝奏云链接"`                     //是否蓝奏云链接
	Weigh       string `p:"weigh" v:"weigh@integer#权重需为整数" dc:"权重"`                                 //权重
	Active      string `p:"active" v:"active@integer#是否启用需为整数" dc:"是否启用"`                           //是否启用
	Note        string `p:"note" dc:"备注"`                                                           //备注
	CreatedBy   string `p:"createdBy" v:"createdBy@integer#创建人需为整数" dc:"创建人"`                       //创建人
	CreatedAt   string `p:"createdAt" v:"createdAt@datetime#创建时间需为YYYY-MM-DD hh:mm:ss格式" dc:"创建时间"` //创建时间
}

// AsApplicationSearchRes 列表返回结果
type AsApplicationSearchRes struct {
	comModel.ListRes
	List []*AsApplicationListRes `json:"list"`
}

// AsApplicationAddReq 添加操作请求参数
type AsApplicationAddReq struct {
	FileUrl     string `p:"fileUrl"  dc:"文件"`
	Name        string `p:"name" v:"required#应用名称不能为空" dc:"应用名称"`
	BundleId    string `p:"bundleId"  dc:"包名"`
	Version     string `p:"version"  dc:"版本号"`
	Size        int    `p:"size"  dc:"文件大小"`
	Type        int    `p:"type"  dc:"类型"`
	Description string `p:"description"  dc:"描述"`
	IconUrl     string `p:"iconUrl"  dc:"图标"`
	Lock        int    `p:"lock"  dc:"付费"`
	Lanzou      int    `p:"lanzou"  dc:"是否蓝奏云链接"`
	Weigh       int    `p:"weigh"  dc:"权重"`
	Active      int    `p:"active"  dc:"是否启用"`
	Note        string `p:"note"  dc:"备注"`
	CreatedBy   uint64
}

// AsApplicationEditReq 修改操作请求参数
type AsApplicationEditReq struct {
	Id          uint   `p:"id" v:"required#主键ID不能为空" dc:"ID"`
	FileUrl     string `p:"fileUrl"  dc:"文件"`
	Name        string `p:"name" v:"required#应用名称不能为空" dc:"应用名称"`
	BundleId    string `p:"bundleId"  dc:"包名"`
	Version     string `p:"version"  dc:"版本号"`
	Size        int    `p:"size"  dc:"文件大小"`
	Type        int    `p:"type"  dc:"类型"`
	Description string `p:"description"  dc:"描述"`
	IconUrl     string `p:"iconUrl"  dc:"图标"`
	Lock        int    `p:"lock"  dc:"付费"`
	Lanzou      int    `p:"lanzou"  dc:"是否蓝奏云链接"`
	Weigh       int    `p:"weigh"  dc:"权重"`
	Active      int    `p:"active"  dc:"是否启用"`
	Note        string `p:"note"  dc:"备注"`
	UpdatedBy   uint64
}

// AsApplicationExtractReq 解析软件请求参数
type AsApplicationExtractReq struct {
	FileUrl string `p:"fileUrl"  dc:"文件"`
}

// AsApplicationExtractRes 解析软件请求响应参数
type AsApplicationExtractRes struct {
	Name     string `json:"name" dc:"软件名"`
	BundleId string `json:"bundleId" dc:"包名"`
	Version  string `json:"version" dc:"版本号"`
	IconUrl  string `json:"iconUrl" dc:"图标"`
	AppSize  int64  `json:"appSize" dc:"大小"`
}

// AsApplicationIconSearchReq 图标搜索请求参数
type AsApplicationIconSearchReq struct {
	Name    string `p:"name"  dc:"名称"`
	Country string `p:"country"  dc:"国家"`
}

// AsApplicationSourceImportReq 软件源导入请求参数
type AsApplicationSourceImportReq struct {
	SourceUrl     string `p:"sourceUrl"  dc:"软件源地址"`
	SettingImport bool   `p:"settingImport"  dc:"导入配置"`
}
