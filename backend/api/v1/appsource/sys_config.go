// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2025-02-18 20:01:49
// 生成路径: api/v1/appsource/sys_config.go
// 生成人：smithy
// desc:配置项相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package appsource

import (
	commonApi "mangosmithy/api/v1/common"
	"mangosmithy/internal/app/appsource/model"

	"github.com/gogf/gf/v2/frame/g"
)

// SysConfigSearchReq 分页请求参数
type SysConfigSearchReq struct {
	g.Meta `path:"/list" tags:"配置项" method:"get" summary:"配置列表"`
	commonApi.Author
	model.SysConfigSearchReq
}

// SysConfigSearchRes 列表返回结果
type SysConfigSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.SysConfigSearchRes
}

// 相关连表查询数据
type LinkedSysConfigDataSearchReq struct {
	g.Meta `path:"/linkedData" tags:"配置项" method:"get" summary:"配置关联表数据"`
	commonApi.Author
}

// 相关连表查询数据
type LinkedSysConfigDataSearchRes struct {
	g.Meta `mime:"application/json"`
	*model.LinkedSysConfigDataSearchRes
}

// SysConfigAddReq 添加操作请求参数
type SysConfigAddReq struct {
	g.Meta `path:"/add" tags:"配置项" method:"post" summary:"配置添加"`
	commonApi.Author
	*model.SysConfigAddReq
}

// SysConfigAddRes 添加操作返回结果
type SysConfigAddRes struct {
	commonApi.EmptyRes
}

// SysConfigEditReq 修改操作请求参数
type SysConfigEditReq struct {
	g.Meta `path:"/edit" tags:"配置项" method:"put" summary:"配置修改"`
	commonApi.Author
	*model.SysConfigEditReq
}

// SysConfigEditRes 修改操作返回结果
type SysConfigEditRes struct {
	commonApi.EmptyRes
}

// SysConfigGetReq 获取一条数据请求
type SysConfigGetReq struct {
	g.Meta `path:"/get" tags:"配置项" method:"get" summary:"获取配置信息"`
	commonApi.Author
	ConfigId uint `p:"configId" v:"required#主键必须"` //通过主键获取
}

// SysConfigGetRes 获取一条数据结果
type SysConfigGetRes struct {
	g.Meta `mime:"application/json"`
	*model.SysConfigInfoRes
}

// SysConfigDeleteReq 删除数据请求
type SysConfigDeleteReq struct {
	g.Meta `path:"/delete" tags:"配置项" method:"delete" summary:"删除配置"`
	commonApi.Author
	ConfigIds []uint `p:"configIds" v:"required#主键必须"` //通过主键删除
}

// SysConfigDeleteRes 删除数据返回
type SysConfigDeleteRes struct {
	commonApi.EmptyRes
}

// ParseAppInfoReq 解析app信息请求
type ParseAppInfoReq struct {
	g.Meta `path:"/parseapp" tags:"配置项" method:"post" summary:"解析APP信息"`
	commonApi.Author
	model.ParseAppInfoReq
}

// ParseAppInfoRes 删除数据返回
type ParseAppInfoRes struct {
	g.Meta `mime:"application/json"`
	*model.ParseAppInfoRes
}
