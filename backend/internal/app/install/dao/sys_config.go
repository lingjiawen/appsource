package dao

import (
	"mangosmithy/internal/app/install/dao/internal"
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
