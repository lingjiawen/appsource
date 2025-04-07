// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-03-25 21:39:11
// 生成路径: internal/app/install/model/sign_platform.go
// 生成人：smithy
// desc:平台
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	comModel "mangosmithy/internal/app/common/model"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SignPlatformInfoRes is the golang structure for table sign_platform.
type SignPlatformInfoRes struct {
	gmeta.Meta `orm:"table:sign_platform"`
	Id         uint        `orm:"id,primary" json:"id" dc:"ID"`          // ID
	Name       string      `orm:"name" json:"name" dc:"平台名"`             // 平台名
	Code       string      `orm:"code" json:"code" dc:"平台标识"`            // 平台标识
	BaseUrl    string      `orm:"base_url" json:"baseUrl" dc:"域名"`       // 域名
	OpenSsl    int         `orm:"open_ssl" json:"openSsl" dc:"开启SSL"`    // 开启SSL
	Status     int         `orm:"status" json:"status" dc:"启用"`          // 启用
	Token      string      `orm:"token" json:"token" dc:"对接Token"`       // 对接Token
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt" dc:"创建时间"` // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt" dc:"修改时间"` // 修改时间
}

type SignPlatformListRes struct {
	Id        uint        `json:"id" dc:"ID"`
	Name      string      `json:"name" dc:"平台名"`
	Code      string      `json:"code" dc:"平台标识"`
	BaseUrl   string      `json:"baseUrl" dc:"域名"`
	OpenSsl   int         `json:"openSsl" dc:"开启SSL"`
	Status    int         `json:"status" dc:"启用"`
	Token     string      `json:"token" dc:"对接Token"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// SignPlatformSearchReq 分页请求参数
type SignPlatformSearchReq struct {
	comModel.PageReq
	Id        string `p:"id" dc:"ID"`                                                             //ID
	Name      string `p:"name" dc:"平台名"`                                                          //平台名
	Code      string `p:"code" dc:"平台标识"`                                                         //平台标识
	BaseUrl   string `p:"baseUrl" dc:"域名"`                                                        //域名
	OpenSsl   string `p:"openSsl" v:"openSsl@integer#开启SSL需为整数" dc:"开启SSL"`                       //开启SSL
	Status    string `p:"status" v:"status@integer#启用需为整数" dc:"启用"`                               //启用
	Token     string `p:"token" dc:"对接Token"`                                                     //对接Token
	CreatedAt string `p:"createdAt" v:"createdAt@datetime#创建时间需为YYYY-MM-DD hh:mm:ss格式" dc:"创建时间"` //创建时间
}

// SignPlatformSearchRes 列表返回结果
type SignPlatformSearchRes struct {
	comModel.ListRes
	List []*SignPlatformListRes `json:"list"`
}

// SignPlatformAddReq 添加操作请求参数
type SignPlatformAddReq struct {
	Name    string `p:"name" v:"required#平台名不能为空" dc:"平台名"`
	Code    string `p:"code"  dc:"平台标识"`
	BaseUrl string `p:"baseUrl" v:"required#域名不能为空" dc:"域名"`
	OpenSsl int    `p:"openSsl"  dc:"开启SSL"`
	Status  int    `p:"status" v:"required#启用不能为空" dc:"启用"`
	Token   string `p:"token" v:"required#对接Token不能为空" dc:"对接Token"`
}

// SignPlatformEditReq 修改操作请求参数
type SignPlatformEditReq struct {
	Id      uint   `p:"id" v:"required#主键ID不能为空" dc:"ID"`
	Name    string `p:"name" v:"required#平台名不能为空" dc:"平台名"`
	Code    string `p:"code"  dc:"平台标识"`
	BaseUrl string `p:"baseUrl" v:"required#域名不能为空" dc:"域名"`
	OpenSsl int    `p:"openSsl"  dc:"开启SSL"`
	Status  int    `p:"status" v:"required#启用不能为空" dc:"启用"`
	Token   string `p:"token" v:"required#对接Token不能为空" dc:"对接Token"`
}
