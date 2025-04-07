package do

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	gmeta.Meta  `orm:"table:sys_config, do:true"`
	ConfigId    interface{} `orm:"config_id,primary" json:"configId"` // 参数主键
	ConfigName  interface{} `orm:"config_name" json:"configName"`     // 参数名称
	ConfigKey   interface{} `orm:"config_key" json:"configKey"`       // 参数键名
	ConfigValue interface{} `orm:"config_value" json:"configValue"`   // 参数键值
	ConfigType  interface{} `orm:"config_type" json:"configType"`     // 系统内置（Y是 N否）
	Group       interface{} `orm:"group" json:"group"`                // 参数分组
	CreateBy    interface{} `orm:"create_by" json:"createBy"`         // 创建者
	UpdateBy    interface{} `orm:"update_by" json:"updateBy"`         // 更新者
	Remark      interface{} `orm:"remark" json:"remark"`              // 备注
	CreatedAt   *gtime.Time `orm:"created_at" json:"createdAt"`       // 创建时间
	UpdatedAt   *gtime.Time `orm:"updated_at" json:"updatedAt"`       // 修改时间
}
