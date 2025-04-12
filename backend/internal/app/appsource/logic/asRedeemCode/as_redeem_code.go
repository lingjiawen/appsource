package asRedeemCode

import (
	"context"
	"errors"
	"mangosmithy/internal/app/appsource/dao"
	"mangosmithy/internal/app/appsource/model"
	"mangosmithy/internal/app/appsource/model/do"
	"mangosmithy/internal/app/appsource/service"
	commonService "mangosmithy/internal/app/common/service"
	"mangosmithy/internal/app/system/consts"
	systemService "mangosmithy/internal/app/system/service"
	"mangosmithy/library/liberr"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
)

func init() {
	service.RegisterAsRedeemCode(New())
}

func New() service.IAsRedeemCode {
	return &sAsRedeemCode{}
}

type sAsRedeemCode struct{}

func (s *sAsRedeemCode) List(ctx context.Context, req *model.AsRedeemCodeSearchReq) (listRes *model.AsRedeemCodeSearchRes, err error) {
	listRes = new(model.AsRedeemCodeSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.AsRedeemCode.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.AsRedeemCode.Columns().Id+" = ?", req.Id)
		}
		if req.Code != "" {
			m = m.Where(dao.AsRedeemCode.Columns().Code+" = ?", req.Code)
		}
		if req.Udid != "" {
			m = m.Where(dao.AsRedeemCode.Columns().Udid+" = ?", req.Udid)
		}
		if req.Type != "" {
			m = m.Where(dao.AsRedeemCode.Columns().Type+" = ?", gconv.Int(req.Type))
		}
		if req.Active != "" {
			m = m.Where(dao.AsRedeemCode.Columns().Active+" = ?", gconv.Int(req.Active))
		}
		if req.ActiveAt != "" {
			m = m.Where(dao.AsRedeemCode.Columns().ActiveAt+" = ?", gconv.Time(req.ActiveAt))
		}
		if req.CreatedBy != "" {
			m = m.Where(dao.AsRedeemCode.Columns().CreatedBy+" = ?", gconv.Uint(req.CreatedBy))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.AsRedeemCode.Columns().CreatedAt+" >=? AND "+dao.AsRedeemCode.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		order := "id desc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.AsRedeemCodeListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.AsRedeemCodeListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.AsRedeemCodeListRes{
				Id:        v.Id,
				Code:      v.Code,
				Udid:      v.Udid,
				Type:      v.Type,
				Note:      v.Note,
				Active:    v.Active,
				ActiveAt:  v.ActiveAt,
				CreatedAt: v.CreatedAt,
			}
		}
	})
	return
}

func (s *sAsRedeemCode) GetExportData(ctx context.Context, req *model.AsRedeemCodeSearchReq) (listRes []*model.AsRedeemCodeInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.AsRedeemCode.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.AsRedeemCode.Columns().Id+" = ?", req.Id)
		}
		if req.Code != "" {
			m = m.Where(dao.AsRedeemCode.Columns().Code+" = ?", req.Code)
		}
		if req.Udid != "" {
			m = m.Where(dao.AsRedeemCode.Columns().Udid+" = ?", req.Udid)
		}
		if req.Type != "" {
			m = m.Where(dao.AsRedeemCode.Columns().Type+" = ?", gconv.Int(req.Type))
		}
		if req.Active != "" {
			m = m.Where(dao.AsRedeemCode.Columns().Active+" = ?", gconv.Int(req.Active))
		}
		if req.ActiveAt != "" {
			m = m.Where(dao.AsRedeemCode.Columns().ActiveAt+" = ?", gconv.Time(req.ActiveAt))
		}
		if req.CreatedBy != "" {
			m = m.Where(dao.AsRedeemCode.Columns().CreatedBy+" = ?", gconv.Uint(req.CreatedBy))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.AsRedeemCode.Columns().CreatedAt+" >=? AND "+dao.AsRedeemCode.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "id desc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&listRes)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	return
}

func (s *sAsRedeemCode) Import(ctx context.Context, file *ghttp.UploadFile) (err error) {
	if file == nil {
		err = errors.New("请上传数据文件")
		return
	}
	var data []do.AsRedeemCode
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
		data = make([]do.AsRedeemCode, len(rows)-1)
		for k, v := range rows {
			if k == 0 {
				continue
			}
			for kv, vv := range v {
				d[kv] = vv
			}
			data[k-1] = do.AsRedeemCode{
				Code:      d[0],
				Udid:      d[1],
				Type:      gconv.Int64(d[2]),
				Active:    gconv.Int64(d[3]),
				ActiveAt:  gconv.GTime(d[4]),
				CreatedBy: gconv.Int64(d[5]),
				UpdatedBy: gconv.Int64(d[6]),
				CreatedAt: gconv.GTime(d[7]),
				UpdatedAt: gconv.GTime(d[8]),
				DeletedAt: gconv.GTime(d[9]),
				Note:      d[10],
			}
		}
		if len(data) > 0 {
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = dao.AsRedeemCode.Ctx(ctx).Batch(500).Insert(data)
				return err
			})
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (s *sAsRedeemCode) GetById(ctx context.Context, id uint) (res *model.AsRedeemCodeInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.AsRedeemCode.Ctx(ctx).WithAll().Where(dao.AsRedeemCode.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sAsRedeemCode) Add(ctx context.Context, req *model.AsRedeemCodeAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var codes []do.AsRedeemCode
		userID := systemService.Context().GetUserId(ctx)

		for i := 0; i < req.Quantity; i++ {
			code, err := commonService.SnowID().GenBase36ID()
			if err != nil {
				return
			}
			if req.Prefix != "" {
				code = req.Prefix + code
			}

			codes = append(codes, do.AsRedeemCode{
				Code:      code,
				Type:      req.Type,
				Note:      req.Note,
				CreatedBy: userID,
			})
		}

		// 批量插入
		_, err = dao.AsRedeemCode.Ctx(ctx).Insert(codes)
		liberr.ErrIsNil(ctx, err, "生成失败")
	})
	return
}

func (s *sAsRedeemCode) Edit(ctx context.Context, req *model.AsRedeemCodeEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AsRedeemCode.Ctx(ctx).WherePri(req.Id).Update(do.AsRedeemCode{
			Code:      req.Code,
			Udid:      req.Udid,
			Type:      req.Type,
			Active:    req.Active,
			ActiveAt:  req.ActiveAt,
			Note:      req.Note,
			UpdatedBy: systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sAsRedeemCode) Delete(ctx context.Context, ids []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AsRedeemCode.Ctx(ctx).Delete(dao.AsRedeemCode.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// ChangeActive 卡密管理是否激活修改（状态）
func (s *sAsRedeemCode) ChangeActive(ctx context.Context, id uint, active int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AsRedeemCode.Ctx(ctx).WherePri(id).
			Update(do.AsRedeemCode{
				Active: active,
			})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}
