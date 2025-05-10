// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-04-14 19:12:45
// 生成路径: internal/app/install/dao/sign_help.go
// 生成人：smithy
// desc:帮助中心
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"mangosmithy/internal/app/install/dao/internal"
)

// signHelpDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type signHelpDao struct {
	*internal.SignHelpDao
}

var (
	// SignHelp is globally public accessible object for table tools_gen_table operations.
	SignHelp = signHelpDao{
		internal.NewSignHelpDao(),
	}
)

// Fill with you ideas below.
