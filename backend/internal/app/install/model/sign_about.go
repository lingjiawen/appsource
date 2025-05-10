// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-04-14 19:08:27
// 生成路径: internal/app/install/model/sign_about.go
// 生成人：smithy
// desc:关于我们
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	comModel "mangosmithy/internal/app/common/model"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SignAboutInfoRes is the golang structure for table sign_about.
type SignAboutInfoRes struct {
	gmeta.Meta `orm:"table:sign_about"`
	Id         uint        `orm:"id,primary" json:"id" dc:"ID"`          // ID
	Icon       string      `orm:"icon" json:"icon" dc:"图标"`              // 图标
	Title      string      `orm:"title" json:"title" dc:"标题"`            // 标题
	Subtitle   string      `orm:"subtitle" json:"subtitle" dc:"内容"`      // 内容
	IsLink     int         `orm:"is_link" json:"isLink" dc:"是否链接"`       // 是否链接
	Url        string      `orm:"url" json:"url" dc:"链接"`                // 链接
	Group      int         `orm:"group" json:"group" dc:"分组"`            // 分组
	Weigh      int         `orm:"weigh" json:"weigh" dc:"权重"`            // 权重
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt" dc:"创建日期"` // 创建日期
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt" dc:"修改日期"` // 修改日期
}

type SignAboutListRes struct {
	Id        uint        `json:"id" dc:"ID"`
	Icon      string      `json:"icon" dc:"图标"`
	Title     string      `json:"title" dc:"标题"`
	Subtitle  string      `json:"subtitle" dc:"内容"`
	IsLink    int         `json:"isLink" dc:"是否链接"`
	Url       string      `json:"url" dc:"链接"`
	Group     int         `json:"group" dc:"分组"`
	Weigh     int         `json:"weigh" dc:"权重"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建日期"`
}

// SignAboutSearchReq 分页请求参数
type SignAboutSearchReq struct {
	comModel.PageReq
	Id        string `p:"id" dc:"ID"`                                                             //ID
	Icon      string `p:"icon" dc:"图标"`                                                           //图标
	Title     string `p:"title" dc:"标题"`                                                          //标题
	Subtitle  string `p:"subtitle" dc:"内容"`                                                       //内容
	IsLink    string `p:"isLink" v:"isLink@integer#是否链接需为整数" dc:"是否链接"`                           //是否链接
	Url       string `p:"url" dc:"链接"`                                                            //链接
	Group     string `p:"group" v:"group@integer#分组需为整数" dc:"分组"`                                 //分组
	Weigh     string `p:"weigh" v:"weigh@integer#权重需为整数" dc:"权重"`                                 //权重
	CreatedAt string `p:"createdAt" v:"createdAt@datetime#创建日期需为YYYY-MM-DD hh:mm:ss格式" dc:"创建日期"` //创建日期
}

// SignAboutSearchRes 列表返回结果
type SignAboutSearchRes struct {
	comModel.ListRes
	List []*SignAboutListRes `json:"list"`
}

// SignAboutAddReq 添加操作请求参数
type SignAboutAddReq struct {
	Icon     string `p:"icon"  dc:"图标"`
	Title    string `p:"title" v:"required#标题不能为空" dc:"标题"`
	Subtitle string `p:"subtitle" v:"required#内容不能为空" dc:"内容"`
	IsLink   int    `p:"isLink"  dc:"是否链接"`
	Url      string `p:"url"  dc:"链接"`
	Group    int    `p:"group"  dc:"分组"`
	Weigh    int    `p:"weigh"  dc:"权重"`
}

// SignAboutEditReq 修改操作请求参数
type SignAboutEditReq struct {
	Id       uint   `p:"id" v:"required#主键ID不能为空" dc:"ID"`
	Icon     string `p:"icon"  dc:"图标"`
	Title    string `p:"title" v:"required#标题不能为空" dc:"标题"`
	Subtitle string `p:"subtitle" v:"required#内容不能为空" dc:"内容"`
	IsLink   int    `p:"isLink"  dc:"是否链接"`
	Url      string `p:"url"  dc:"链接"`
	Group    int    `p:"group"  dc:"分组"`
	Weigh    int    `p:"weigh"  dc:"权重"`
}
