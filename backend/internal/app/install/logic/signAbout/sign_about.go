// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-04-14 19:08:27
// 生成路径: internal/app/install/logic/sign_about.go
// 生成人：smithy
// desc:关于我们
// company:云南奇讯科技有限公司
// ==========================================================================

package signAbout

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
	service.RegisterSignAbout(New())
}

func New() service.ISignAbout {
	return &sSignAbout{}
}

type sSignAbout struct{}

func (s *sSignAbout) List(ctx context.Context, req *model.SignAboutSearchReq) (listRes *model.SignAboutSearchRes, err error) {
	listRes = new(model.SignAboutSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SignAbout.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.SignAbout.Columns().Id+" = ?", req.Id)
		}
		if req.Icon != "" {
			m = m.Where(dao.SignAbout.Columns().Icon+" = ?", req.Icon)
		}
		if req.Title != "" {
			m = m.Where(dao.SignAbout.Columns().Title+" = ?", req.Title)
		}
		if req.Subtitle != "" {
			m = m.Where(dao.SignAbout.Columns().Subtitle+" = ?", req.Subtitle)
		}
		if req.IsLink != "" {
			m = m.Where(dao.SignAbout.Columns().IsLink+" = ?", gconv.Int(req.IsLink))
		}
		if req.Url != "" {
			m = m.Where(dao.SignAbout.Columns().Url+" = ?", req.Url)
		}
		if req.Group != "" {
			m = m.Where(dao.SignAbout.Columns().Group+" = ?", gconv.Int(req.Group))
		}
		if req.Weigh != "" {
			m = m.Where(dao.SignAbout.Columns().Weigh+" = ?", gconv.Int(req.Weigh))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.SignAbout.Columns().CreatedAt+" >=? AND "+dao.SignAbout.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		order := "weigh desc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.SignAboutListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.SignAboutListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.SignAboutListRes{
				Id:        v.Id,
				Icon:      v.Icon,
				Title:     v.Title,
				Subtitle:  v.Subtitle,
				IsLink:    v.IsLink,
				Url:       v.Url,
				Group:     v.Group,
				Weigh:     v.Weigh,
				CreatedAt: v.CreatedAt,
			}
		}
	})
	return
}

func (s *sSignAbout) GetExportData(ctx context.Context, req *model.SignAboutSearchReq) (listRes []*model.SignAboutInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SignAbout.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.SignAbout.Columns().Id+" = ?", req.Id)
		}
		if req.Icon != "" {
			m = m.Where(dao.SignAbout.Columns().Icon+" = ?", req.Icon)
		}
		if req.Title != "" {
			m = m.Where(dao.SignAbout.Columns().Title+" = ?", req.Title)
		}
		if req.Subtitle != "" {
			m = m.Where(dao.SignAbout.Columns().Subtitle+" = ?", req.Subtitle)
		}
		if req.IsLink != "" {
			m = m.Where(dao.SignAbout.Columns().IsLink+" = ?", gconv.Int(req.IsLink))
		}
		if req.Url != "" {
			m = m.Where(dao.SignAbout.Columns().Url+" = ?", req.Url)
		}
		if req.Group != "" {
			m = m.Where(dao.SignAbout.Columns().Group+" = ?", gconv.Int(req.Group))
		}
		if req.Weigh != "" {
			m = m.Where(dao.SignAbout.Columns().Weigh+" = ?", gconv.Int(req.Weigh))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.SignAbout.Columns().CreatedAt+" >=? AND "+dao.SignAbout.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "weigh desc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&listRes)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

func (s *sSignAbout) Import(ctx context.Context, file *ghttp.UploadFile) (err error) {
	if file == nil {
		err = errors.New("请上传数据文件")
		return
	}
	var data []do.SignAbout
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
		data = make([]do.SignAbout, len(rows)-1)
		for k, v := range rows {
			if k == 0 {
				continue
			}
			for kv, vv := range v {
				d[kv] = vv
			}
			data[k-1] = do.SignAbout{
				Icon:      d[0],
				Title:     d[1],
				Subtitle:  d[2],
				IsLink:    gconv.Int64(d[3]),
				Url:       d[4],
				Group:     gconv.Int64(d[5]),
				Weigh:     gconv.Int64(d[6]),
				CreatedAt: gconv.GTime(d[7]),
				UpdatedAt: gconv.GTime(d[8]),
			}
		}
		if len(data) > 0 {
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = dao.SignAbout.Ctx(ctx).Batch(500).Insert(data)
				return err
			})
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (s *sSignAbout) GetById(ctx context.Context, id uint) (res *model.SignAboutInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SignAbout.Ctx(ctx).WithAll().Where(dao.SignAbout.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sSignAbout) Add(ctx context.Context, req *model.SignAboutAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignAbout.Ctx(ctx).Insert(do.SignAbout{
			Icon:     req.Icon,
			Title:    req.Title,
			Subtitle: req.Subtitle,
			IsLink:   req.IsLink,
			Url:      req.Url,
			Group:    req.Group,
			Weigh:    req.Weigh,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sSignAbout) Edit(ctx context.Context, req *model.SignAboutEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignAbout.Ctx(ctx).WherePri(req.Id).Update(do.SignAbout{
			Icon:     req.Icon,
			Title:    req.Title,
			Subtitle: req.Subtitle,
			IsLink:   req.IsLink,
			Url:      req.Url,
			Group:    req.Group,
			Weigh:    req.Weigh,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sSignAbout) Delete(ctx context.Context, ids []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignAbout.Ctx(ctx).Delete(dao.SignAbout.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// 关于我们是否链接修改（状态）
func (s *sSignAbout) ChangeIsLink(ctx context.Context, id uint, isLink int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignAbout.Ctx(ctx).WherePri(id).
			Update(do.SignAbout{
				IsLink: isLink,
			})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}
