// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-04-01 19:38:11
// 生成路径: internal/app/install/dao/sign_device.go
// 生成人：smithy
// desc:设备管理
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"mangosmithy/internal/app/install/dao/internal"
)

// signDeviceDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type signDeviceDao struct {
	*internal.SignDeviceDao
}

var (
	// SignDevice is globally public accessible object for table tools_gen_table operations.
	SignDevice = signDeviceDao{
		internal.NewSignDeviceDao(),
	}
)

// Fill with you ideas below.
