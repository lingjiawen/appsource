// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-02-18 19:57:00
// 生成路径: api/v1/appsource/sys_config_group.go
// 生成人：smithy
// desc:配置组相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package appsource

import (
	commonApi "mangosmithy/api/v1/common"
	"mangosmithy/internal/app/appsource/model"

	"github.com/gogf/gf/v2/frame/g"
)

// SysConfigGroupSearchReq 分页请求参数
type SysConfigGroupSearchReq struct {
	g.Meta `path:"/list" tags:"配置组" method:"get" summary:"配置分组列表"`
	commonApi.Author
	model.SysConfigGroupSearchReq
}

// SysConfigGroupSearchRes 列表返回结果
type SysConfigGroupSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.SysConfigGroupSearchRes
}

// SysConfigGroupAddReq 添加操作请求参数
type SysConfigGroupAddReq struct {
	g.Meta `path:"/add" tags:"配置组" method:"post" summary:"配置分组添加"`
	commonApi.Author
	*model.SysConfigGroupAddReq
}

// SysConfigGroupAddRes 添加操作返回结果
type SysConfigGroupAddRes struct {
	commonApi.EmptyRes
}

// SysConfigGroupEditReq 修改操作请求参数
type SysConfigGroupEditReq struct {
	g.Meta `path:"/edit" tags:"配置组" method:"put" summary:"配置分组修改"`
	commonApi.Author
	*model.SysConfigGroupEditReq
}

// SysConfigGroupEditRes 修改操作返回结果
type SysConfigGroupEditRes struct {
	commonApi.EmptyRes
}

// SysConfigGroupGetReq 获取一条数据请求
type SysConfigGroupGetReq struct {
	g.Meta `path:"/get" tags:"配置组" method:"get" summary:"获取配置分组信息"`
	commonApi.Author
	Id uint `p:"id" v:"required#主键必须"` //通过主键获取
}

// SysConfigGroupGetRes 获取一条数据结果
type SysConfigGroupGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SysConfigGroupInfoRes
}

// SysConfigGroupDeleteReq 删除数据请求
type SysConfigGroupDeleteReq struct {
	g.Meta `path:"/delete" tags:"配置组" method:"delete" summary:"删除配置分组"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}

// SysConfigGroupDeleteRes 删除数据返回
type SysConfigGroupDeleteRes struct {
	commonApi.EmptyRes
}
