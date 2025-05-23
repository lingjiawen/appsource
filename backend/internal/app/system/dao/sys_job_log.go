// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mangosmithy/internal/app/system/dao/internal"
)

// internalSysJobLogDao is internal type for wrapping internal DAO implements.
type internalSysJobLogDao = *internal.SysJobLogDao

// sysJobLogDao is the data access object for table sys_job_log.
// You can define custom methods on it to extend its functionality as you wish.
type sysJobLogDao struct {
	internalSysJobLogDao
}

var (
	// SysJobLog is globally public accessible object for table sys_job_log operations.
	SysJobLog = sysJobLogDao{
		internal.NewSysJobLogDao(),
	}
)

// Fill with you ideas below.
