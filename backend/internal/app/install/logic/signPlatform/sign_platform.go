// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-03-25 21:39:11
// 生成路径: internal/app/install/logic/sign_platform.go
// 生成人：smithy
// desc:平台
// company:云南奇讯科技有限公司
// ==========================================================================

package signPlatform

import (
	"context"
	"errors"
	"mangosmithy/internal/app/install/dao"
	"mangosmithy/internal/app/install/model"
	"mangosmithy/internal/app/install/model/do"
	"mangosmithy/internal/app/install/service"
	"mangosmithy/internal/app/system/consts"
	"mangosmithy/library/liberr"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
)

func init() {
	service.RegisterSignPlatform(New())
}

func New() service.ISignPlatform {
	return &sSignPlatform{}
}

type sSignPlatform struct{}

func (s *sSignPlatform) List(ctx context.Context, req *model.SignPlatformSearchReq) (listRes *model.SignPlatformSearchRes, err error) {
	listRes = new(model.SignPlatformSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SignPlatform.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.SignPlatform.Columns().Id+" = ?", req.Id)
		}
		if req.Name != "" {
			m = m.Where(dao.SignPlatform.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.Code != "" {
			m = m.Where(dao.SignPlatform.Columns().Code+" = ?", req.Code)
		}
		if req.BaseUrl != "" {
			m = m.Where(dao.SignPlatform.Columns().BaseUrl+" = ?", req.BaseUrl)
		}
		if req.OpenSsl != "" {
			m = m.Where(dao.SignPlatform.Columns().OpenSsl+" = ?", gconv.Int(req.OpenSsl))
		}
		if req.Status != "" {
			m = m.Where(dao.SignPlatform.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		if req.Token != "" {
			m = m.Where(dao.SignPlatform.Columns().Token+" = ?", req.Token)
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.SignPlatform.Columns().CreatedAt+" >=? AND "+dao.SignPlatform.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		var res []*model.SignPlatformListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.SignPlatformListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.SignPlatformListRes{
				Id:        v.Id,
				Name:      v.Name,
				Code:      v.Code,
				BaseUrl:   v.BaseUrl,
				OpenSsl:   v.OpenSsl,
				Status:    v.Status,
				Token:     v.Token,
				CreatedAt: v.CreatedAt,
			}
		}
	})
	return
}

func (s *sSignPlatform) GetExportData(ctx context.Context, req *model.SignPlatformSearchReq) (listRes []*model.SignPlatformInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SignPlatform.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.SignPlatform.Columns().Id+" = ?", req.Id)
		}
		if req.Name != "" {
			m = m.Where(dao.SignPlatform.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.Code != "" {
			m = m.Where(dao.SignPlatform.Columns().Code+" = ?", req.Code)
		}
		if req.BaseUrl != "" {
			m = m.Where(dao.SignPlatform.Columns().BaseUrl+" = ?", req.BaseUrl)
		}
		if req.OpenSsl != "" {
			m = m.Where(dao.SignPlatform.Columns().OpenSsl+" = ?", gconv.Int(req.OpenSsl))
		}
		if req.Status != "" {
			m = m.Where(dao.SignPlatform.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		if req.Token != "" {
			m = m.Where(dao.SignPlatform.Columns().Token+" = ?", req.Token)
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.SignPlatform.Columns().CreatedAt+" >=? AND "+dao.SignPlatform.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "id asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&listRes)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

func (s *sSignPlatform) Import(ctx context.Context, file *ghttp.UploadFile) (err error) {
	if file == nil {
		err = errors.New("请上传数据文件")
		return
	}
	var data []do.SignPlatform
	err = g.Try(ctx, func(ctx context.Context) {
		f, err := file.Open()
		liberr.ErrIsNil(ctx, err)
		defer f.Close()
		exFile, err := excelize.OpenReader(f)
		liberr.ErrIsNil(ctx, err)
		defer exFile.Close()
		rows, err := exFile.GetRows("Sheet1")
		liberr.ErrIsNil(ctx, err)
		if len(rows) == 0 {
			liberr.ErrIsNil(ctx, errors.New("表格内容不能为空"))
		}
		d := make([]interface{}, len(rows[0]))
		data = make([]do.SignPlatform, len(rows)-1)
		for k, v := range rows {
			if k == 0 {
				continue
			}
			for kv, vv := range v {
				d[kv] = vv
			}
			data[k-1] = do.SignPlatform{
				Name:      d[0],
				Code:      d[1],
				BaseUrl:   d[2],
				OpenSsl:   gconv.Int64(d[3]),
				Status:    gconv.Int64(d[4]),
				Token:     d[5],
				CreatedAt: gconv.GTime(d[6]),
				UpdatedAt: gconv.GTime(d[7]),
			}
		}
		if len(data) > 0 {
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = dao.SignPlatform.Ctx(ctx).Batch(500).Insert(data)
				return err
			})
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (s *sSignPlatform) GetById(ctx context.Context, id uint) (res *model.SignPlatformInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SignPlatform.Ctx(ctx).WithAll().Where(dao.SignPlatform.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sSignPlatform) Add(ctx context.Context, req *model.SignPlatformAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignPlatform.Ctx(ctx).Insert(do.SignPlatform{
			Name:    req.Name,
			Code:    req.Code,
			BaseUrl: req.BaseUrl,
			OpenSsl: req.OpenSsl,
			Status:  req.Status,
			Token:   req.Token,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sSignPlatform) Edit(ctx context.Context, req *model.SignPlatformEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignPlatform.Ctx(ctx).WherePri(req.Id).Update(do.SignPlatform{
			Name:    req.Name,
			Code:    req.Code,
			BaseUrl: req.BaseUrl,
			OpenSsl: req.OpenSsl,
			Status:  req.Status,
			Token:   req.Token,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sSignPlatform) Delete(ctx context.Context, ids []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignPlatform.Ctx(ctx).Delete(dao.SignPlatform.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// 平台开启SSL修改（状态）
func (s *sSignPlatform) ChangeOpenSsl(ctx context.Context, id uint, openSsl int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignPlatform.Ctx(ctx).WherePri(id).
			Update(do.SignPlatform{
				OpenSsl: openSsl,
			})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

// 平台启用修改（状态）
func (s *sSignPlatform) ChangeStatus(ctx context.Context, id uint, status int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignPlatform.Ctx(ctx).WherePri(id).
			Update(do.SignPlatform{
				Status: status,
			})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}
