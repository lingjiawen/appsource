// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2025-02-18 19:57:00
// 生成路径: internal/app/appsource/model/sys_config_group.go
// 生成人：smithy
// desc:配置组
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	comModel "mangosmithy/internal/app/common/model"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysConfigGroupInfoRes is the golang structure for table sys_config_group.
type SysConfigGroupInfoRes struct {
	gmeta.Meta `orm:"table:sys_config_group"`
	Id         uint        `orm:"id,primary" json:"id" dc:"主键"`           // 主键
	GroupName  string      `orm:"group_name" json:"groupName" dc:"配置组名称"` // 配置组名称
	GroupKey   string      `orm:"group_key" json:"groupKey" dc:"配置组键名"`   // 配置组键名
	CreateBy   uint        `orm:"create_by" json:"createBy" dc:"创建者"`     // 创建者
	UpdateBy   uint        `orm:"update_by" json:"updateBy" dc:"更新者"`     // 更新者
	CreatedAt  *gtime.Time `orm:"created_at" json:"createdAt" dc:"创建时间"`  // 创建时间
	UpdatedAt  *gtime.Time `orm:"updated_at" json:"updatedAt" dc:"修改时间"`  // 修改时间
}

type SysConfigGroupListRes struct {
	Id        uint        `json:"id" dc:"主键"`
	GroupName string      `json:"groupName" dc:"配置组名称"`
	GroupKey  string      `json:"groupKey" dc:"配置组键名"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// SysConfigGroupSearchReq 分页请求参数
type SysConfigGroupSearchReq struct {
	comModel.PageReq
	Id        string `p:"id" dc:"主键"`                                                             //主键
	GroupName string `p:"groupName" dc:"配置组名称"`                                                   //配置组名称
	GroupKey  string `p:"groupKey" dc:"配置组键名"`                                                    //配置组键名
	CreateBy  string `p:"createBy" v:"createBy@integer#创建者需为整数" dc:"创建者"`                         //创建者
	UpdateBy  string `p:"updateBy" v:"updateBy@integer#更新者需为整数" dc:"更新者"`                         //更新者
	CreatedAt string `p:"createdAt" v:"createdAt@datetime#创建时间需为YYYY-MM-DD hh:mm:ss格式" dc:"创建时间"` //创建时间
}

// SysConfigGroupSearchRes 列表返回结果
type SysConfigGroupSearchRes struct {
	comModel.ListRes
	List []*SysConfigGroupListRes `json:"list"`
}

// SysConfigGroupAddReq 添加操作请求参数
type SysConfigGroupAddReq struct {
	GroupName string `p:"groupName" v:"required#配置组名称不能为空" dc:"配置组名称"`
	GroupKey  string `p:"groupKey"  dc:"配置组键名"`
}

// SysConfigGroupEditReq 修改操作请求参数
type SysConfigGroupEditReq struct {
	Id        uint   `p:"id" v:"required#主键ID不能为空" dc:"主键"`
	GroupName string `p:"groupName" v:"required#配置组名称不能为空" dc:"配置组名称"`
	GroupKey  string `p:"groupKey"  dc:"配置组键名"`
}
