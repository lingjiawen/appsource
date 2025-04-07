// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-04-01 19:38:11
// 生成路径: internal/app/install/logic/sign_device.go
// 生成人：smithy
// desc:设备管理
// company:云南奇讯科技有限公司
// ==========================================================================

package signDevice

import (
	"context"
	"errors"
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
	service.RegisterSignDevice(New())
}

func New() service.ISignDevice {
	return &sSignDevice{}
}

type sSignDevice struct{}

func (s *sSignDevice) List(ctx context.Context, req *model.SignDeviceSearchReq) (listRes *model.SignDeviceSearchRes, err error) {
	listRes = new(model.SignDeviceSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SignDevice.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.SignDevice.Columns().Id+" = ?", req.Id)
		}
		if req.Udid != "" {
			m = m.Where(dao.SignDevice.Columns().Udid+" = ?", req.Udid)
		}
		if req.Name != "" {
			m = m.Where(dao.SignDevice.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.CertId != "" {
			m = m.Where(dao.SignDevice.Columns().CertId+" = ?", req.CertId)
		}
		if req.DeviceId != "" {
			m = m.Where(dao.SignDevice.Columns().DeviceId+" = ?", req.DeviceId)
		}
		if req.AddTime != "" {
			m = m.Where(dao.SignDevice.Columns().AddTime+" = ?", gconv.Uint64(req.AddTime))
		}
		if req.Model != "" {
			m = m.Where(dao.SignDevice.Columns().Model+" = ?", req.Model)
		}
		if req.ExpireTime != "" {
			m = m.Where(dao.SignDevice.Columns().ExpireTime+" = ?", gconv.Uint64(req.ExpireTime))
		}
		if req.RedeemCode != "" {
			m = m.Where(dao.SignDevice.Columns().RedeemCode+" = ?", req.RedeemCode)
		}
		if req.AccountType != "" {
			m = m.Where(dao.SignDevice.Columns().AccountType+" = ?", gconv.Int(req.AccountType))
		}
		if req.WarrantyType != "" {
			m = m.Where(dao.SignDevice.Columns().WarrantyType+" = ?", gconv.Int(req.WarrantyType))
		}
		if req.DeviceType != "" {
			m = m.Where(dao.SignDevice.Columns().DeviceType+" = ?", req.DeviceType)
		}
		if req.Status != "" {
			m = m.Where(dao.SignDevice.Columns().Status+" = ?", req.Status)
		}
		if req.Pool != "" {
			m = m.Where(dao.SignDevice.Columns().Pool+" = ?", gconv.Int(req.Pool))
		}
		if req.ApiPlatform != "" {
			m = m.Where(dao.SignDevice.Columns().ApiPlatform+" = ?", gconv.Int(req.ApiPlatform))
		}
		if req.ApiWarrantyType != "" {
			m = m.Where(dao.SignDevice.Columns().ApiWarrantyType+" = ?", gconv.Int(req.ApiWarrantyType))
		}
		if req.Active != "" {
			m = m.Where(dao.SignDevice.Columns().Active+" = ?", gconv.Int(req.Active))
		}
		if req.CreatedBy != "" {
			m = m.Where(dao.SignDevice.Columns().CreatedBy+" = ?", gconv.Uint(req.CreatedBy))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.SignDevice.Columns().CreatedAt+" >=? AND "+dao.SignDevice.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		var res []*model.SignDeviceListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.SignDeviceListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.SignDeviceListRes{
				Id:              v.Id,
				Udid:            v.Udid,
				Name:            v.Name,
				CertId:          v.CertId,
				AddTime:         v.AddTime,
				Model:           v.Model,
				ExpireTime:      v.ExpireTime,
				RedeemCode:      v.RedeemCode,
				AccountType:     v.AccountType,
				WarrantyType:    v.WarrantyType,
				DeviceType:      v.DeviceType,
				Status:          v.Status,
				Pool:            v.Pool,
				ApiPlatform:     v.ApiPlatform,
				ApiWarrantyType: v.ApiWarrantyType,
				Active:          v.Active,
				CreatedUser:     v.CreatedUser,
				CreatedBy:       v.CreatedBy,
				CreatedAt:       v.CreatedAt,
			}
		}
	})
	return
}

func (s *sSignDevice) GetExportData(ctx context.Context, req *model.SignDeviceSearchReq) (listRes []*model.SignDeviceInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SignDevice.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.SignDevice.Columns().Id+" = ?", req.Id)
		}
		if req.Udid != "" {
			m = m.Where(dao.SignDevice.Columns().Udid+" = ?", req.Udid)
		}
		if req.Name != "" {
			m = m.Where(dao.SignDevice.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.CertId != "" {
			m = m.Where(dao.SignDevice.Columns().CertId+" = ?", req.CertId)
		}
		if req.DeviceId != "" {
			m = m.Where(dao.SignDevice.Columns().DeviceId+" = ?", req.DeviceId)
		}
		if req.AddTime != "" {
			m = m.Where(dao.SignDevice.Columns().AddTime+" = ?", gconv.Uint64(req.AddTime))
		}
		if req.Model != "" {
			m = m.Where(dao.SignDevice.Columns().Model+" = ?", req.Model)
		}
		if req.ExpireTime != "" {
			m = m.Where(dao.SignDevice.Columns().ExpireTime+" = ?", gconv.Uint64(req.ExpireTime))
		}
		if req.RedeemCode != "" {
			m = m.Where(dao.SignDevice.Columns().RedeemCode+" = ?", req.RedeemCode)
		}
		if req.AccountType != "" {
			m = m.Where(dao.SignDevice.Columns().AccountType+" = ?", gconv.Int(req.AccountType))
		}
		if req.WarrantyType != "" {
			m = m.Where(dao.SignDevice.Columns().WarrantyType+" = ?", gconv.Int(req.WarrantyType))
		}
		if req.DeviceType != "" {
			m = m.Where(dao.SignDevice.Columns().DeviceType+" = ?", req.DeviceType)
		}
		if req.Status != "" {
			m = m.Where(dao.SignDevice.Columns().Status+" = ?", req.Status)
		}
		if req.Pool != "" {
			m = m.Where(dao.SignDevice.Columns().Pool+" = ?", gconv.Int(req.Pool))
		}
		if req.ApiPlatform != "" {
			m = m.Where(dao.SignDevice.Columns().ApiPlatform+" = ?", gconv.Int(req.ApiPlatform))
		}
		if req.ApiWarrantyType != "" {
			m = m.Where(dao.SignDevice.Columns().ApiWarrantyType+" = ?", gconv.Int(req.ApiWarrantyType))
		}
		if req.Active != "" {
			m = m.Where(dao.SignDevice.Columns().Active+" = ?", gconv.Int(req.Active))
		}
		if req.CreatedBy != "" {
			m = m.Where(dao.SignDevice.Columns().CreatedBy+" = ?", gconv.Uint(req.CreatedBy))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.SignDevice.Columns().CreatedAt+" >=? AND "+dao.SignDevice.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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

func (s *sSignDevice) Import(ctx context.Context, file *ghttp.UploadFile) (err error) {
	if file == nil {
		err = errors.New("请上传数据文件")
		return
	}
	var data []do.SignDevice
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
		data = make([]do.SignDevice, len(rows)-1)
		for k, v := range rows {
			if k == 0 {
				continue
			}
			for kv, vv := range v {
				d[kv] = vv
			}
			data[k-1] = do.SignDevice{
				Udid:            d[0],
				Name:            d[1],
				CertId:          d[2],
				DeviceId:        d[3],
				P12:             d[4],
				Mobileprovision: d[5],
				AddTime:         gconv.Int64(d[6]),
				P12Password:     d[7],
				Model:           d[8],
				ExpireTime:      gconv.Int64(d[9]),
				SerialNumber:    d[10],
				RedeemCode:      d[11],
				AccountType:     gconv.Int64(d[12]),
				WarrantyType:    gconv.Int64(d[13]),
				DeviceType:      d[14],
				Status:          d[15],
				Pool:            gconv.Int64(d[16]),
				ApiPlatform:     gconv.Int64(d[17]),
				ApiWarrantyType: gconv.Int64(d[18]),
				Active:          gconv.Int64(d[19]),
				CreatedBy:       gconv.Int64(d[20]),
				UpdatedBy:       gconv.Int64(d[21]),
				CreatedAt:       gconv.GTime(d[22]),
				UpdatedAt:       gconv.GTime(d[23]),
				DeletedAt:       gconv.GTime(d[24]),
			}
		}
		if len(data) > 0 {
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = dao.SignDevice.Ctx(ctx).Batch(500).Insert(data)
				return err
			})
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (s *sSignDevice) GetById(ctx context.Context, id uint) (res *model.SignDeviceInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SignDevice.Ctx(ctx).WithAll().Where(dao.SignDevice.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sSignDevice) Add(ctx context.Context, req *model.SignDeviceAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignDevice.Ctx(ctx).Insert(do.SignDevice{
			Udid:            req.Udid,
			Name:            req.Name,
			CertId:          req.CertId,
			DeviceId:        req.DeviceId,
			P12:             req.P12,
			Mobileprovision: req.Mobileprovision,
			AddTime:         req.AddTime,
			Model:           req.Model,
			ExpireTime:      req.ExpireTime,
			RedeemCode:      req.RedeemCode,
			AccountType:     req.AccountType,
			WarrantyType:    req.WarrantyType,
			DeviceType:      req.DeviceType,
			Status:          req.Status,
			Pool:            req.Pool,
			ApiPlatform:     req.ApiPlatform,
			ApiWarrantyType: req.ApiWarrantyType,
			Active:          req.Active,
			CreatedBy:       systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sSignDevice) Edit(ctx context.Context, req *model.SignDeviceEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignDevice.Ctx(ctx).WherePri(req.Id).Update(do.SignDevice{
			Udid:            req.Udid,
			Name:            req.Name,
			CertId:          req.CertId,
			DeviceId:        req.DeviceId,
			P12:             req.P12,
			Mobileprovision: req.Mobileprovision,
			AddTime:         req.AddTime,
			Model:           req.Model,
			ExpireTime:      req.ExpireTime,
			RedeemCode:      req.RedeemCode,
			AccountType:     req.AccountType,
			WarrantyType:    req.WarrantyType,
			DeviceType:      req.DeviceType,
			Status:          req.Status,
			Pool:            req.Pool,
			ApiPlatform:     req.ApiPlatform,
			ApiWarrantyType: req.ApiWarrantyType,
			Active:          req.Active,
			UpdatedBy:       systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sSignDevice) Delete(ctx context.Context, ids []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignDevice.Ctx(ctx).Delete(dao.SignDevice.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// 设备管理禁用修改（状态）
func (s *sSignDevice) ChangeActive(ctx context.Context, id uint, active int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SignDevice.Ctx(ctx).WherePri(id).
			Update(do.SignDevice{
				Active: active,
			})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}
