// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-04-12 18:50:08
// 生成路径: internal/app/install/dao/sign_platform.go
// 生成人：smithy
// desc:平台
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"mangosmithy/internal/app/install/dao/internal"
)

// signPlatformDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type signPlatformDao struct {
	*internal.SignPlatformDao
}

var (
	// SignPlatform is globally public accessible object for table tools_gen_table operations.
	SignPlatform = signPlatformDao{
		internal.NewSignPlatformDao(),
	}
)

// Fill with you ideas below.
