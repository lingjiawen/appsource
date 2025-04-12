package model

import (
	comModel "mangosmithy/internal/app/common/model"
	systemModel "mangosmithy/internal/app/system/model"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// AsRedeemCodeInfoRes is the golang structure for table as_redeem_code.
type AsRedeemCodeInfoRes struct {
	gmeta.Meta  `orm:"table:as_redeem_code"`
	Id          uint                     `orm:"id,primary" json:"id" dc:"ID"`         // ID
	Code        string                   `orm:"code" json:"code" dc:"兑换码"`            // 兑换码
	Udid        string                   `orm:"udid" json:"udid" dc:"设备码"`            // 设备码
	Type        int                      `orm:"type" json:"type" dc:"类型"`             // 类型
	Active      int                      `orm:"active" json:"active" dc:"是否激活"`       // 是否激活
	ActiveAt    *gtime.Time              `orm:"active_at" json:"activeAt" dc:"激活时间"`  // 激活时间
	CreatedBy   uint                     `orm:"created_by" json:"createdBy" dc:"创建人"` // 创建人
	CreatedUser *systemModel.LinkUserRes `orm:"with:id=created_by" json:"createdUser"`
	UpdatedBy   uint                     `orm:"updated_by" json:"updatedBy" dc:"修改人"` // 修改人
	UpdatedUser *systemModel.LinkUserRes `orm:"with:id=updated_by" json:"updatedUser"`
	CreatedAt   *gtime.Time              `orm:"created_at" json:"createdAt" dc:"生成时间"` // 生成时间
	UpdatedAt   *gtime.Time              `orm:"updated_at" json:"updatedAt" dc:"修改时间"` // 修改时间
	DeletedAt   *gtime.Time              `orm:"deleted_at" json:"deletedAt" dc:"删除时间"` // 删除时间
	Note        string                   `orm:"note" json:"note" dc:"备注"`              // 备注
}

type AsRedeemCodeListRes struct {
	Id        uint        `json:"id" dc:"ID"`
	Code      string      `json:"code" dc:"兑换码"`
	Udid      string      `json:"udid" dc:"设备码"`
	Type      int         `json:"type" dc:"类型"`
	Note      string      `json:"note" dc:"备注"`
	Active    int         `json:"active" dc:"是否激活"`
	ActiveAt  *gtime.Time `json:"activeAt" dc:"激活时间"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"生成时间"`
}

// AsRedeemCodeSearchReq 分页请求参数
type AsRedeemCodeSearchReq struct {
	comModel.PageReq
	Id        string `p:"id" dc:"ID"`                                                             //ID
	Code      string `p:"code" dc:"兑换码"`                                                          //兑换码
	Udid      string `p:"udid" dc:"设备码"`                                                          //设备码
	Type      string `p:"type" v:"type@integer#类型需为整数" dc:"类型"`                                   //类型
	Active    string `p:"active" v:"active@integer#是否激活需为整数" dc:"是否激活"`                           //是否激活
	ActiveAt  string `p:"activeAt" v:"activeAt@datetime#激活时间需为YYYY-MM-DD hh:mm:ss格式" dc:"激活时间"`   //激活时间
	CreatedBy string `p:"createdBy" v:"createdBy@integer#创建人需为整数" dc:"创建人"`                       //创建人
	CreatedAt string `p:"createdAt" v:"createdAt@datetime#生成时间需为YYYY-MM-DD hh:mm:ss格式" dc:"生成时间"` //生成时间
}

// AsRedeemCodeSearchRes 列表返回结果
type AsRedeemCodeSearchRes struct {
	comModel.ListRes
	List []*AsRedeemCodeListRes `json:"list"`
}

// AsRedeemCodeAddReq 添加操作请求参数
type AsRedeemCodeAddReq struct {
	Prefix   string `p:"prefix"  dc:"前缀"`
	Quantity int    `p:"quantity"  dc:"数量"`
	Type     int    `p:"type"  dc:"类型"`
	Note     string `p:"note"  dc:"备注"`
}

// AsRedeemCodeEditReq 修改操作请求参数
type AsRedeemCodeEditReq struct {
	Id        uint        `p:"id" v:"required#主键ID不能为空" dc:"ID"`
	Code      string      `p:"code"  dc:"兑换码"`
	Udid      string      `p:"udid"  dc:"设备码"`
	Type      int         `p:"type"  dc:"类型"`
	Active    int         `p:"active"  dc:"是否激活"`
	ActiveAt  *gtime.Time `p:"activeAt"  dc:"激活时间"`
	Note      string      `p:"note"  dc:"备注"`
	UpdatedBy uint64
}
