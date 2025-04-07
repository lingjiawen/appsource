package signRedeemCode

import (
	"context"
	"errors"
	"fmt"
	commonService "mangosmithy/internal/app/common/service"
	"mangosmithy/internal/app/install/dao"
	"mangosmithy/internal/app/install/model"
	"mangosmithy/internal/app/install/model/do"
	"mangosmithy/internal/app/install/service"
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
	service.RegisterSignRedeemCode(New())
}

func New() service.ISignRedeemCode {
	return &sSignRedeemCode{}
}

type sSignRedeemCode struct{}

func (s *sSignRedeemCode) List(ctx context.Context, req *model.SignRedeemCodeSearchReq) (listRes *model.SignRedeemCodeSearchRes, err error) {
	listRes = new(model.SignRedeemCodeSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SignRedeemCode.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Id+" = ?", req.Id)
		}
		if req.Code != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Code+" = ?", req.Code)
		}
		if req.Udid != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Udid+" = ?", req.Udid)
		}
		if req.CertId != "" {
			m = m.Where(dao.SignRedeemCode.Columns().CertId+" = ?", req.CertId)
		}
		if req.AccountType != "" {
			m = m.Where(dao.SignRedeemCode.Columns().AccountType+" = ?", gconv.Int(req.AccountType))
		}
		if req.WarrantyType != "" {
			m = m.Where(dao.SignRedeemCode.Columns().WarrantyType+" = ?", gconv.Int(req.WarrantyType))
		}
		if req.DeviceType != "" {
			m = m.Where(dao.SignRedeemCode.Columns().DeviceType+" = ?", req.DeviceType)
		}
		if req.Pool != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Pool+" = ?", gconv.Int(req.Pool))
		}
		if req.ApiPlatform != "" {
			m = m.Where(dao.SignRedeemCode.Columns().ApiPlatform+" = ?", gconv.Int(req.ApiPlatform))
		}
		if req.ApiWarrantyType != "" {
			m = m.Where(dao.SignRedeemCode.Columns().ApiWarrantyType+" = ?", gconv.Int(req.ApiWarrantyType))
		}
		if req.Note != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Note+" = ?", req.Note)
		}
		if req.Banned != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Banned+" = ?", gconv.Int(req.Banned))
		}
		if req.Active != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Active+" = ?", gconv.Int(req.Active))
		}
		if req.ActiveAt != "" {
			m = m.Where(dao.SignRedeemCode.Columns().ActiveAt+" = ?", gconv.Time(req.ActiveAt))
		}
		if req.CreatedBy != "" {
			m = m.Where(dao.SignRedeemCode.Columns().CreatedBy+" = ?", gconv.Uint(req.CreatedBy))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.SignRedeemCode.Columns().CreatedAt+" >=? AND "+dao.SignRedeemCode.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		var res []*model.SignRedeemCodeListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.SignRedeemCodeListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.SignRedeemCodeListRes{
				Id:              v.Id,
				Code:            v.Code,
				Udid:            v.Udid,
				CertId:          v.CertId,
				AccountType:     v.AccountType,
				WarrantyType:    v.WarrantyType,
				DeviceType:      v.DeviceType,
				Pool:            v.Pool,
				ApiPlatform:     v.ApiPlatform,
				Note:            v.Note,
				ApiWarrantyType: v.ApiWarrantyType,
				Banned:          v.Banned,
				Active:          v.Active,
				ActiveAt:        v.ActiveAt,
				CreatedAt:       v.CreatedAt,
			}
		}
	})
	return
}

func (s *sSignRedeemCode) GetExportData(ctx context.Context, req *model.SignRedeemCodeSearchReq) (listRes []*model.SignRedeemCodeInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SignRedeemCode.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Id+" = ?", req.Id)
		}
		if req.Code != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Code+" = ?", req.Code)
		}
		if req.Udid != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Udid+" = ?", req.Udid)
		}
		if req.CertId != "" {
			m = m.Where(dao.SignRedeemCode.Columns().CertId+" = ?", req.CertId)
		}
		if req.AccountType != "" {
			m = m.Where(dao.SignRedeemCode.Columns().AccountType+" = ?", gconv.Int(req.AccountType))
		}
		if req.WarrantyType != "" {
			m = m.Where(dao.SignRedeemCode.Columns().WarrantyType+" = ?", gconv.Int(req.WarrantyType))
		}
		if req.DeviceType != "" {
			m = m.Where(dao.SignRedeemCode.Columns().DeviceType+" = ?", req.DeviceType)
		}
		if req.Pool != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Pool+" = ?", gconv.Int(req.Pool))
		}
		if req.ApiPlatform != "" {
			m = m.Where(dao.SignRedeemCode.Columns().ApiPlatform+" = ?", gconv.Int(req.ApiPlatform))
		}
		if req.ApiWarrantyType != "" {
			m = m.Where(dao.SignRedeemCode.Columns().ApiWarrantyType+" = ?", gconv.Int(req.ApiWarrantyType))
		}
		if req.Note != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Note+" = ?", req.Note)
		}
		if req.Banned != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Banned+" = ?", gconv.Int(req.Banned))
		}
		if req.Active != "" {
			m = m.Where(dao.SignRedeemCode.Columns().Active+" = ?", gconv.Int(req.Active))
		}
		if req.ActiveAt != "" {
			m = m.Where(dao.SignRedeemCode.Columns().ActiveAt+" = ?", gconv.Time(req.ActiveAt))
		}
		if req.CreatedBy != "" {
			m = m.Where(dao.SignRedeemCode.Columns().CreatedBy+" = ?", gconv.Uint(req.CreatedBy))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.SignRedeemCode.Columns().CreatedAt+" >=? AND "+dao.SignRedeemCode.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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

func (s *sSignRedeemCode) Import(ctx context.Context, file *ghttp.UploadFile) (err error) {
	if file == nil {
		err = errors.New("请上传数据文件")
		return
	}
	var data []do.SignRedeemCode
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
		data = make([]do.SignRedeemCode, len(rows)-1)
		for k, v := range rows {
			if k == 0 {
				continue
			}
			for kv, vv := range v {
				d[kv] = vv
			}
			data[k-1] = do.SignRedeemCode{
				Code:            d[0],
				Udid:            d[1],
				CertId:          d[2],
				AccountType:     gconv.Int64(d[3]),
				WarrantyType:    gconv.Int64(d[4]),
				DeviceType:      d[5],
				Pool:            gconv.Int64(d[6]),
				ApiPlatform:     gconv.Int64(d[7]),
				Note:            d[8],
				ApiWarrantyType: gconv.Int64(d[9]),
				Banned:          gconv.Int64(d[10]),
				Active:          gconv.Int64(d[11]),
				ActiveAt:        gconv.GTime(d[12]),
				CreatedBy:       gconv.Int64(d[13]),
				UpdatedBy:       gconv.Int64(d[14]),
				CreatedAt:       gconv.GTime(d[15]),
				UpdatedAt:       gconv.GTime(d[16]),
				DeletedAt:       gconv.GTime(d[17]),
			}
		}
		if len(data) > 0 {
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = dao.SignRedeemCode.Ctx(ctx).Batch(500).Insert(data)
				return err
			})
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (s *sSignRedeemCode) GetById(ctx context.Context, id uint) (res *model.SignRedeemCodeInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SignRedeemCode.Ctx(ctx).WithAll().Where(dao.SignRedeemCode.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sSignRedeemCode) Add(ctx context.Context, req *model.SignRedeemCodeAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var codes []do.SignRedeemCode
		userID := systemService.Context().GetUserId(ctx)
		for i := 0; i < req.Quantity; i++ {
			code, err := commonService.SnowID().GenBase36ID()
			if err != nil {
				return
			}
			if req.Prefix != "" {
				code = req.Prefix + code
			}

			codes = append(codes, do.SignRedeemCode{
				Code:            code,
				AccountType:     req.AccountType,
				WarrantyType:    req.WarrantyType,
				DeviceType:      req.DeviceType,
				Pool:            req.Pool,
				ApiPlatform:     req.ApiPlatform,
				Note:            req.Note,
				ApiWarrantyType: req.ApiWarrantyType,
				CreatedBy:       userID,
			})

		}

		// 批量插入
		_, err = dao.SignRedeemCode.Ctx(ctx).Insert(codes)
		if err != nil {
			liberr.ErrIsNil(ctx, err, fmt.Sprintf("添加失败, err: %v", err))
		}
	})
	return
}

func (s *sSignRedeemCode) Edit(ctx context.Context, req *model.SignRedeemCodeEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignRedeemCode.Ctx(ctx).WherePri(req.Id).Update(do.SignRedeemCode{
			Code:            req.Code,
			Udid:            req.Udid,
			CertId:          req.CertId,
			AccountType:     req.AccountType,
			WarrantyType:    req.WarrantyType,
			DeviceType:      req.DeviceType,
			Pool:            req.Pool,
			ApiPlatform:     req.ApiPlatform,
			Note:            req.Note,
			ApiWarrantyType: req.ApiWarrantyType,
			Banned:          req.Banned,
			Active:          req.Active,
			ActiveAt:        req.ActiveAt,
			UpdatedBy:       systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sSignRedeemCode) Delete(ctx context.Context, ids []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignRedeemCode.Ctx(ctx).Delete(dao.SignRedeemCode.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// ChangeBanned 签名卡密禁用修改（状态）
func (s *sSignRedeemCode) ChangeBanned(ctx context.Context, id uint, banned int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignRedeemCode.Ctx(ctx).WherePri(id).
			Update(do.SignRedeemCode{
				Banned: banned,
			})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

// ChangeActive 签名卡密激活修改（状态）
func (s *sSignRedeemCode) ChangeActive(ctx context.Context, id uint, active int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignRedeemCode.Ctx(ctx).WherePri(id).
			Update(do.SignRedeemCode{
				Active: active,
			})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}
