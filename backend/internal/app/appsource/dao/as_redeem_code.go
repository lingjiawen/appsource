// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-04-09 21:34:25
// 生成路径: internal/app/appsource/dao/as_redeem_code.go
// 生成人：smithy
// desc:卡密管理
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"mangosmithy/internal/app/appsource/dao/internal"
)

// asRedeemCodeDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type asRedeemCodeDao struct {
	*internal.AsRedeemCodeDao
}

var (
	// AsRedeemCode is globally public accessible object for table tools_gen_table operations.
	AsRedeemCode = asRedeemCodeDao{
		internal.NewAsRedeemCodeDao(),
	}
)

// Fill with you ideas below.
