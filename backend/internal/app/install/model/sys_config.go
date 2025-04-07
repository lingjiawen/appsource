package model

import (
	"github.com/gogf/gf/v2/frame/g"
	comModel "mangosmithy/internal/app/common/model"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysConfigInfoRes is the golang structure for table sys_config.
type SysConfigInfoRes struct {
	gmeta.Meta  `orm:"table:sys_config"`
	ConfigId    uint                           `orm:"config_id,primary" json:"configId" dc:"参数主键"`    // 参数主键
	ConfigName  string                         `orm:"config_name" json:"configName" dc:"参数名称"`        // 参数名称
	ConfigKey   string                         `orm:"config_key" json:"configKey" dc:"参数键名"`          // 参数键名
	ConfigValue string                         `orm:"config_value" json:"configValue" dc:"参数键值"`      // 参数键值
	ConfigType  int                            `orm:"config_type" json:"configType" dc:"系统内置（Y是 N否）"` // 系统内置（Y是 N否）
	AllowDel    int                            `orm:"allow_del" json:"allowDel" dc:"允许删除"`
	ElementType string                         `orm:"element_type" json:"elementType" dc:"参数类型"`
	Group       string                         `orm:"group" json:"group" dc:"参数分组"` // 参数分组
	LinkedGroup *LinkedSysConfigSysConfigGroup `orm:"with:group_key=group" json:"linkedGroup"`
	CreateBy    uint                           `orm:"create_by" json:"createBy" dc:"创建者"`    // 创建者
	UpdateBy    uint                           `orm:"update_by" json:"updateBy" dc:"更新者"`    // 更新者
	Remark      string                         `orm:"remark" json:"remark" dc:"备注"`          // 备注
	CreatedAt   *gtime.Time                    `orm:"created_at" json:"createdAt" dc:"创建时间"` // 创建时间
	UpdatedAt   *gtime.Time                    `orm:"updated_at" json:"updatedAt" dc:"修改时间"` // 修改时间
}
type LinkedSysConfigSysConfigGroup struct {
	gmeta.Meta `orm:"table:sys_config_group"`
	GroupKey   string `orm:"group_key" json:"groupKey" dc:"配置组键名"`   // 配置组键名
	GroupName  string `orm:"group_name" json:"groupName" dc:"配置组名称"` // 配置组名称
}

type SysConfigListRes struct {
	ConfigId    uint   `json:"configId" dc:"参数主键"`
	ConfigName  string `json:"configName" dc:"参数名称"`
	ConfigKey   string `json:"configKey" dc:"参数键名"`
	ConfigValue string `json:"configValue" dc:"参数键值"`
	ConfigType  int    `json:"configType" dc:"参数类型"`
	AllowDel    uint   `json:"allowDel" dc:"允许删除"`
	ElementType string `json:"elementType" dc:"参数类型"`
	Group       string `json:"group" dc:"参数分组"`
	Remark      string `json:"remark" dc:"备注"`
}

// SysConfigSearchReq 分页请求参数
type SysConfigSearchReq struct {
	comModel.PageReq
	ConfigId    string `p:"configId" dc:"参数主键"`                                                 //参数主键
	ConfigName  string `p:"configName" dc:"参数名称"`                                               //参数名称
	ConfigKey   string `p:"configKey" dc:"参数键名"`                                                //参数键名
	ConfigValue string `p:"configValue" dc:"参数键值"`                                              //参数键值
	ConfigType  string `p:"configType" v:"configType@integer#系统内置（Y是 N否）需为整数" dc:"系统内置（Y是 N否）"` //系统内置（Y是 N否）
	AllowDel    string `p:"allowDel" dc:"允许删除"`
	ElementType string `p:"elementType" dc:"参数类型"`
	Group       string `p:"group" dc:"参数分组"`                                                        //参数分组
	CreatedAt   string `p:"createdAt" v:"createdAt@datetime#创建时间需为YYYY-MM-DD hh:mm:ss格式" dc:"创建时间"` //创建时间
}

// SysConfigSearchRes 列表返回结果
type SysConfigSearchRes struct {
	List g.Map `json:"list"`
}

// LinkedSysConfigDataSearchRes 相关连表查询数据
type LinkedSysConfigDataSearchRes struct {
	LinkedSysConfigSysConfigGroup []*LinkedSysConfigSysConfigGroup `json:"linkedSysConfigSysConfigGroup"`
}

// SysConfigAddReq 添加操作请求参数
type SysConfigAddReq struct {
	ConfigName  string `p:"configName" v:"required#参数名称不能为空" dc:"参数名称"`
	ConfigKey   string `p:"configKey"  dc:"参数键名"`
	ConfigValue string `p:"configValue"  dc:"参数键值"`
	ConfigType  int    `p:"configType"  dc:"系统内置（Y是 N否）"`
	AllowDel    int    `p:"allowDel" dc:"允许删除"`
	ElementType string `p:"elementType" dc:"参数类型"`
	Group       string `p:"group"  dc:"参数分组"`
	Remark      string `p:"remark"  dc:"备注"`
}

// SysConfigEditReq 修改操作请求参数
type SysConfigEditReq struct {
	Configs []*SysConfigReqData
}

type SysConfigReqData struct {
	ConfigId    uint   `p:"configId" dc:"参数主键"`
	ConfigName  string `p:"configName" dc:"参数名称"`
	ConfigKey   string `p:"configKey"  dc:"参数键名"`
	ConfigValue string `p:"configValue"  dc:"参数键值"`
	ConfigType  int    `p:"configType"  dc:"系统内置（Y是 N否）"`
	AllowDel    int    `p:"allowDel" dc:"允许删除"`
	ElementType string `p:"elementType" dc:"参数类型"`
	Group       string `p:"group"  dc:"参数分组"`
	Remark      string `p:"remark"  dc:"备注"`
}
