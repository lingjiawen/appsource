// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-03-25 21:39:11
// 生成路径: internal/app/install/controller/sign_platform.go
// 生成人：smithy
// desc:平台
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"
	"errors"
	"mangosmithy/api/v1/install"
	"mangosmithy/internal/app/install/model"
	"mangosmithy/internal/app/install/service"
	systemController "mangosmithy/internal/app/system/controller"
	systemService "mangosmithy/internal/app/system/service"
	"mangosmithy/library/libUtils"

	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/xuri/excelize/v2"
)

type signPlatformController struct {
	systemController.BaseController
}

var SignPlatform = new(signPlatformController)

// List 列表
func (c *signPlatformController) List(ctx context.Context, req *install.SignPlatformSearchReq) (res *install.SignPlatformSearchRes, err error) {
	res = new(install.SignPlatformSearchRes)
	res.SignPlatformSearchRes, err = service.SignPlatform().List(ctx, &req.SignPlatformSearchReq)
	return
}

// Export 导出excel
func (c *signPlatformController) Export(ctx context.Context, req *install.SignPlatformExportReq) (res *install.SignPlatformExportRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		listData []*model.SignPlatformInfoRes
		//表头
		tableHead = []interface{}{"ID", "平台名", "平台标识", "域名", "开启SSL", "启用", "对接Token", "创建时间", "修改时间"}
		excelData [][]interface{}
		//字典选项处理
	)
	req.PageNum = 1
	req.PageSize = 500
	//获取字典数据
	excelData = append(excelData, tableHead)
	for {
		listData, err = service.SignPlatform().GetExportData(ctx, &req.SignPlatformSearchReq)
		if err != nil {
			return
		}
		if listData == nil {
			break
		}
		for _, v := range listData {
			var ()
			dt := []interface{}{
				v.Id,
				v.Name,
				v.Code,
				v.BaseUrl,
				v.OpenSsl,
				v.Status,
				v.Token,
				v.CreatedAt.Format("Y-m-d H:i:s"),
				v.UpdatedAt.Format("Y-m-d H:i:s"),
			}
			excelData = append(excelData, dt)
		}
		req.PageNum++
	}
	//创建excel处理对象
	excel := new(libUtils.ExcelHelper).CreateFile()
	defer excel.Close()
	excel.ArrToExcel("Sheet1", "A1", excelData)
	col, _ := excelize.ColumnNumberToName(len(tableHead))
	row := len(excelData)
	cr, _ := excelize.JoinCellName(col, row)
	excel.SetCellBorder("Sheet1", "A1", cr)
	_, err = excel.WriteTo(r.Response.BufferWriter)
	if err != nil {
		return
	}
	r.Response.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	r.Response.Header().Set("Accept-Ranges", "bytes")
	r.Response.Header().Set("Access-Control-Expose-Headers", "*")
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("平台")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *signPlatformController) ExcelTemplate(ctx context.Context, req *install.SignPlatformExcelTemplateReq) (res *install.SignPlatformExcelTemplateRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
		//表头
		tableHead = []interface{}{"平台名", "平台标识", "域名", "开启SSL", "启用", "对接Token", "创建时间", "修改时间"}
		excelData = [][]interface{}{tableHead}
	)
	//创建excel处理对象
	excel := new(libUtils.ExcelHelper).CreateFile()
	defer excel.Close()
	excel.ArrToExcel("Sheet1", "A1", excelData)
	col, _ := excelize.ColumnNumberToName(len(tableHead))
	row := len(excelData)
	cr, _ := excelize.JoinCellName(col, row)
	excel.SetCellBorder("Sheet1", "A1", cr)
	_, err = excel.WriteTo(r.Response.BufferWriter)
	if err != nil {
		return
	}
	r.Response.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	r.Response.Header().Set("Accept-Ranges", "bytes")
	r.Response.Header().Set("Access-Control-Expose-Headers", "*")
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("平台模板")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *signPlatformController) Import(ctx context.Context, req *install.SignPlatformImportReq) (res *install.SignPlatformImportRes, err error) {
	err = service.SignPlatform().Import(ctx, req.File)
	return
}

// Get 获取平台
func (c *signPlatformController) Get(ctx context.Context, req *install.SignPlatformGetReq) (res *install.SignPlatformGetRes, err error) {
	res = new(install.SignPlatformGetRes)
	res.SignPlatformInfoRes, err = service.SignPlatform().GetById(ctx, req.Id)
	return
}

// Add 添加平台
func (c *signPlatformController) Add(ctx context.Context, req *install.SignPlatformAddReq) (res *install.SignPlatformAddRes, err error) {
	err = service.SignPlatform().Add(ctx, req.SignPlatformAddReq)
	return
}

// Edit 修改平台
func (c *signPlatformController) Edit(ctx context.Context, req *install.SignPlatformEditReq) (res *install.SignPlatformEditRes, err error) {
	err = service.SignPlatform().Edit(ctx, req.SignPlatformEditReq)
	return
}

// Delete 删除平台
func (c *signPlatformController) Delete(ctx context.Context, req *install.SignPlatformDeleteReq) (res *install.SignPlatformDeleteRes, err error) {
	err = service.SignPlatform().Delete(ctx, req.Ids)
	return
}

// 平台开启SSL修改（状态）
func (c *signPlatformController) ChangeOpenSsl(ctx context.Context, req *install.SignPlatformOpenSslSwitchReq) (res *install.SignPlatformOpenSslSwitchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/install/signPlatform/edit") {
		err = errors.New("没有修改权限")
		return
	}
	err = service.SignPlatform().ChangeOpenSsl(ctx, req.Id, req.OpenSsl)
	return
}

// 平台启用修改（状态）
func (c *signPlatformController) ChangeStatus(ctx context.Context, req *install.SignPlatformStatusSwitchReq) (res *install.SignPlatformStatusSwitchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/install/signPlatform/edit") {
		err = errors.New("没有修改权限")
		return
	}
	err = service.SignPlatform().ChangeStatus(ctx, req.Id, req.Status)
	return
}
