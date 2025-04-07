// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-02-18 20:01:49
// 生成路径: internal/app/appsource/dao/sys_config.go
// 生成人：smithy
// desc:配置项
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"mangosmithy/internal/app/appsource/dao/internal"
)

// sysConfigDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type sysConfigDao struct {
	*internal.SysConfigDao
}

var (
	// SysConfig is globally public accessible object for table tools_gen_table operations.
	SysConfig = sysConfigDao{
		internal.NewSysConfigDao(),
	}
)

// Fill with you ideas below.
