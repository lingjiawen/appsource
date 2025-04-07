// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-02-21 18:37:33
// 生成路径: internal/app/appsource/dao/as_application.go
// 生成人：smithy
// desc:应用管理
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"mangosmithy/internal/app/appsource/dao/internal"
)

// asApplicationDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type asApplicationDao struct {
	*internal.AsApplicationDao
}

var (
	// AsApplication is globally public accessible object for table tools_gen_table operations.
	AsApplication = asApplicationDao{
		internal.NewAsApplicationDao(),
	}
)

// Fill with you ideas below.
