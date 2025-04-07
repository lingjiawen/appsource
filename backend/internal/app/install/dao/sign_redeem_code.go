// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2025-03-25 13:14:15
// 生成路径: internal/app/install/dao/sign_redeem_code.go
// 生成人：smithy
// desc:签名卡密
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"mangosmithy/internal/app/install/dao/internal"
)

// signRedeemCodeDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type signRedeemCodeDao struct {
	*internal.SignRedeemCodeDao
}

var (
	// SignRedeemCode is globally public accessible object for table tools_gen_table operations.
	SignRedeemCode = signRedeemCodeDao{
		internal.NewSignRedeemCodeDao(),
	}
)

// Fill with you ideas below.
