// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-02-18 19:57:00
// 生成路径: internal/app/appsource/logic/sys_config_group.go
// 生成人：smithy
// desc:配置组
// company:云南奇讯科技有限公司
// ==========================================================================

package sysConfigGroup

import (
	"context"
	"mangosmithy/internal/app/appsource/dao"
	"mangosmithy/internal/app/appsource/model"
	"mangosmithy/internal/app/appsource/model/do"
	"mangosmithy/internal/app/appsource/service"
	"mangosmithy/internal/app/system/consts"
	"mangosmithy/library/liberr"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	service.RegisterSysConfigGroup(New())
}

func New() service.ISysConfigGroup {
	return &sSysConfigGroup{}
}

type sSysConfigGroup struct{}

func (s *sSysConfigGroup) List(ctx context.Context, req *model.SysConfigGroupSearchReq) (listRes *model.SysConfigGroupSearchRes, err error) {
	listRes = new(model.SysConfigGroupSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysConfigGroup.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.SysConfigGroup.Columns().Id+" = ?", req.Id)
		}
		if req.GroupName != "" {
			m = m.Where(dao.SysConfigGroup.Columns().GroupName+" like ?", "%"+req.GroupName+"%")
		}
		if req.GroupKey != "" {
			m = m.Where(dao.SysConfigGroup.Columns().GroupKey+" = ?", req.GroupKey)
		}
		if req.CreateBy != "" {
			m = m.Where(dao.SysConfigGroup.Columns().CreateBy+" = ?", gconv.Uint(req.CreateBy))
		}
		if req.UpdateBy != "" {
			m = m.Where(dao.SysConfigGroup.Columns().UpdateBy+" = ?", gconv.Uint(req.UpdateBy))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.SysConfigGroup.Columns().CreatedAt+" >=? AND "+dao.SysConfigGroup.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
		}
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "id asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.SysConfigGroupListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.SysConfigGroupListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.SysConfigGroupListRes{
				Id:        v.Id,
				GroupName: v.GroupName,
				GroupKey:  v.GroupKey,
				CreatedAt: v.CreatedAt,
			}
		}
	})
	return
}

func (s *sSysConfigGroup) GetById(ctx context.Context, id uint) (res *model.SysConfigGroupInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysConfigGroup.Ctx(ctx).WithAll().Where(dao.SysConfigGroup.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sSysConfigGroup) Add(ctx context.Context, req *model.SysConfigGroupAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysConfigGroup.Ctx(ctx).Insert(do.SysConfigGroup{
			GroupName: req.GroupName,
			GroupKey:  req.GroupKey,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sSysConfigGroup) Edit(ctx context.Context, req *model.SysConfigGroupEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysConfigGroup.Ctx(ctx).WherePri(req.Id).Update(do.SysConfigGroup{
			GroupName: req.GroupName,
			GroupKey:  req.GroupKey,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sSysConfigGroup) Delete(ctx context.Context, ids []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysConfigGroup.Ctx(ctx).Delete(dao.SysConfigGroup.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}
