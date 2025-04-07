package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	gmeta.Meta  `orm:"table:sys_config"`
	ConfigId    uint                           `orm:"config_id,primary" json:"configId"` // 参数主键
	ConfigName  string                         `orm:"config_name" json:"configName"`     // 参数名称
	ConfigKey   string                         `orm:"config_key" json:"configKey"`       // 参数键名
	ConfigValue string                         `orm:"config_value" json:"configValue"`   // 参数键值
	ConfigType  int                            `orm:"config_type" json:"configType"`     // 系统内置（Y是 N否）
	Group       string                         `orm:"group" json:"group"`                // 参数分组
	LinkedGroup *LinkedSysConfigSysConfigGroup `orm:"with:group_key=group" json:"linkedGroup"`
	CreateBy    uint                           `orm:"create_by" json:"createBy"`   // 创建者
	UpdateBy    uint                           `orm:"update_by" json:"updateBy"`   // 更新者
	Remark      string                         `orm:"remark" json:"remark"`        // 备注
	CreatedAt   *gtime.Time                    `orm:"created_at" json:"createdAt"` // 创建时间
	UpdatedAt   *gtime.Time                    `orm:"updated_at" json:"updatedAt"` // 修改时间
}

type LinkedSysConfigSysConfigGroup struct {
	gmeta.Meta `orm:"table:sys_config_group"`
	GroupKey   string `orm:"group_key" json:"groupKey"`   // 配置组键名
	GroupName  string `orm:"group_name" json:"groupName"` // 配置组名称
}
