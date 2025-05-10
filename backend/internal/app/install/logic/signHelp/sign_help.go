// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-04-14 19:12:45
// 生成路径: internal/app/install/logic/sign_help.go
// 生成人：smithy
// desc:帮助中心
// company:云南奇讯科技有限公司
// ==========================================================================

package signHelp

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
	service.RegisterSignHelp(New())
}

func New() service.ISignHelp {
	return &sSignHelp{}
}

type sSignHelp struct{}

func (s *sSignHelp) List(ctx context.Context, req *model.SignHelpSearchReq) (listRes *model.SignHelpSearchRes, err error) {
	listRes = new(model.SignHelpSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SignHelp.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.SignHelp.Columns().Id+" = ?", req.Id)
		}
		if req.Title != "" {
			m = m.Where(dao.SignHelp.Columns().Title+" = ?", req.Title)
		}
		if req.Content != "" {
			m = m.Where(dao.SignHelp.Columns().Content+" = ?", req.Content)
		}
		if req.IsExpand != "" {
			m = m.Where(dao.SignHelp.Columns().IsExpand+" = ?", gconv.Int(req.IsExpand))
		}
		if req.Weigh != "" {
			m = m.Where(dao.SignHelp.Columns().Weigh+" = ?", gconv.Int(req.Weigh))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.SignHelp.Columns().CreatedAt+" >=? AND "+dao.SignHelp.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		var res []*model.SignHelpListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.SignHelpListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.SignHelpListRes{
				Id:        v.Id,
				Title:     v.Title,
				Content:   v.Content,
				IsExpand:  v.IsExpand,
				Weigh:     v.Weigh,
				CreatedAt: v.CreatedAt,
			}
		}
	})
	return
}

func (s *sSignHelp) GetExportData(ctx context.Context, req *model.SignHelpSearchReq) (listRes []*model.SignHelpInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SignHelp.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.SignHelp.Columns().Id+" = ?", req.Id)
		}
		if req.Title != "" {
			m = m.Where(dao.SignHelp.Columns().Title+" = ?", req.Title)
		}
		if req.Content != "" {
			m = m.Where(dao.SignHelp.Columns().Content+" = ?", req.Content)
		}
		if req.IsExpand != "" {
			m = m.Where(dao.SignHelp.Columns().IsExpand+" = ?", gconv.Int(req.IsExpand))
		}
		if req.Weigh != "" {
			m = m.Where(dao.SignHelp.Columns().Weigh+" = ?", gconv.Int(req.Weigh))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.SignHelp.Columns().CreatedAt+" >=? AND "+dao.SignHelp.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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

func (s *sSignHelp) Import(ctx context.Context, file *ghttp.UploadFile) (err error) {
	if file == nil {
		err = errors.New("请上传数据文件")
		return
	}
	var data []do.SignHelp
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
		data = make([]do.SignHelp, len(rows)-1)
		for k, v := range rows {
			if k == 0 {
				continue
			}
			for kv, vv := range v {
				d[kv] = vv
			}
			data[k-1] = do.SignHelp{
				Title:     d[0],
				Content:   d[1],
				IsExpand:  gconv.Int64(d[2]),
				Weigh:     gconv.Int64(d[3]),
				CreatedAt: gconv.GTime(d[4]),
				UpdatedAt: gconv.GTime(d[5]),
			}
		}
		if len(data) > 0 {
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = dao.SignHelp.Ctx(ctx).Batch(500).Insert(data)
				return err
			})
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (s *sSignHelp) GetById(ctx context.Context, id uint) (res *model.SignHelpInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SignHelp.Ctx(ctx).WithAll().Where(dao.SignHelp.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sSignHelp) Add(ctx context.Context, req *model.SignHelpAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignHelp.Ctx(ctx).Insert(do.SignHelp{
			Title:    req.Title,
			Content:  req.Content,
			IsExpand: req.IsExpand,
			Weigh:    req.Weigh,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sSignHelp) Edit(ctx context.Context, req *model.SignHelpEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignHelp.Ctx(ctx).WherePri(req.Id).Update(do.SignHelp{
			Title:    req.Title,
			Content:  req.Content,
			IsExpand: req.IsExpand,
			Weigh:    req.Weigh,
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sSignHelp) Delete(ctx context.Context, ids []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignHelp.Ctx(ctx).Delete(dao.SignHelp.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// 帮助中心默认展开修改（状态）
func (s *sSignHelp) ChangeIsExpand(ctx context.Context, id uint, isExpand int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignHelp.Ctx(ctx).WherePri(id).
			Update(do.SignHelp{
				IsExpand: isExpand,
			})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}
