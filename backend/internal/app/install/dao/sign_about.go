// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-04-14 19:08:27
// 生成路径: internal/app/install/dao/sign_about.go
// 生成人：smithy
// desc:关于我们
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"mangosmithy/internal/app/install/dao/internal"
)

// signAboutDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type signAboutDao struct {
	*internal.SignAboutDao
}

var (
	// SignAbout is globally public accessible object for table tools_gen_table operations.
	SignAbout = signAboutDao{
		internal.NewSignAboutDao(),
	}
)

// Fill with you ideas below.
