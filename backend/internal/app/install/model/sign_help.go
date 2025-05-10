// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-04-14 19:12:45
// 生成路径: internal/app/install/model/sign_help.go
// 生成人：smithy
// desc:帮助中心
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	comModel "mangosmithy/internal/app/common/model"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SignHelpInfoRes is the golang structure for table sign_help.
type SignHelpInfoRes struct {
	gmeta.Meta `orm:"table:sign_help"`
	Id         uint        `orm:"id,primary" json:"id" dc:"ID"`          // ID
	Title      string      `orm:"title" json:"title" dc:"标题"`            // 标题
	Content    string      `orm:"content" json:"content" dc:"内容"`        // 内容
	IsExpand   int         `orm:"is_expand" json:"isExpand" dc:"默认展开"`   // 默认展开
	Weigh      int         `orm:"weigh" json:"weigh" dc:"权重"`            // 权重
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt" dc:"创建日期"` // 创建日期
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt" dc:"修改日期"` // 修改日期
}

type SignHelpListRes struct {
	Id        uint        `json:"id" dc:"ID"`
	Title     string      `json:"title" dc:"标题"`
	Content   string      `json:"content" dc:"内容"`
	IsExpand  int         `json:"isExpand" dc:"默认展开"`
	Weigh     int         `json:"weigh" dc:"权重"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建日期"`
}

// SignHelpSearchReq 分页请求参数
type SignHelpSearchReq struct {
	comModel.PageReq
	Id        string `p:"id" dc:"ID"`                                                             //ID
	Title     string `p:"title" dc:"标题"`                                                          //标题
	Content   string `p:"content" dc:"内容"`                                                        //内容
	IsExpand  string `p:"isExpand" v:"isExpand@integer#默认展开需为整数" dc:"默认展开"`                       //默认展开
	Weigh     string `p:"weigh" v:"weigh@integer#权重需为整数" dc:"权重"`                                 //权重
	CreatedAt string `p:"createdAt" v:"createdAt@datetime#创建日期需为YYYY-MM-DD hh:mm:ss格式" dc:"创建日期"` //创建日期
}

// SignHelpSearchRes 列表返回结果
type SignHelpSearchRes struct {
	comModel.ListRes
	List []*SignHelpListRes `json:"list"`
}

// SignHelpAddReq 添加操作请求参数
type SignHelpAddReq struct {
	Title    string `p:"title" v:"required#标题不能为空" dc:"标题"`
	Content  string `p:"content" v:"required#内容不能为空" dc:"内容"`
	IsExpand int    `p:"isExpand"  dc:"默认展开"`
	Weigh    int    `p:"weigh"  dc:"权重"`
}

// SignHelpEditReq 修改操作请求参数
type SignHelpEditReq struct {
	Id       uint   `p:"id" v:"required#主键ID不能为空" dc:"ID"`
	Title    string `p:"title" v:"required#标题不能为空" dc:"标题"`
	Content  string `p:"content" v:"required#内容不能为空" dc:"内容"`
	IsExpand int    `p:"isExpand"  dc:"默认展开"`
	Weigh    int    `p:"weigh"  dc:"权重"`
}
