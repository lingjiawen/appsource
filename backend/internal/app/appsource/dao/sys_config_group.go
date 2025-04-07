// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-02-18 19:57:00
// 生成路径: internal/app/appsource/dao/sys_config_group.go
// 生成人：smithy
// desc:配置组
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"mangosmithy/internal/app/appsource/dao/internal"
)

// sysConfigGroupDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type sysConfigGroupDao struct {
	*internal.SysConfigGroupDao
}

var (
	// SysConfigGroup is globally public accessible object for table tools_gen_table operations.
	SysConfigGroup = sysConfigGroupDao{
		internal.NewSysConfigGroupDao(),
	}
)

// Fill with you ideas below.
