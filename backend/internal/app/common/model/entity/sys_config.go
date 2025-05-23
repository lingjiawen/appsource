// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-04-18 21:09:17
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	ConfigId    uint        `json:"configId"    description:"参数主键"`
	ConfigName  string      `json:"configName"  description:"参数名称"`
	ConfigKey   string      `json:"configKey"   description:"参数键名"`
	ConfigValue string      `json:"configValue" description:"参数键值"`
	ConfigType  int         `json:"configType"  description:"系统内置（Y是 N否）"`
	Group       string      `json:"group"       description:"分组"`
	CreateBy    uint        `json:"createBy"    description:"创建者"`
	UpdateBy    uint        `json:"updateBy"    description:"更新者"`
	Remark      string      `json:"remark"      description:"备注"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"修改时间"`
}
