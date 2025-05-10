// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-04-14 19:12:45
// 生成路径: internal/app/install/controller/sign_help.go
// 生成人：smithy
// desc:帮助中心
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

type signHelpController struct {
	systemController.BaseController
}

var SignHelp = new(signHelpController)

// List 列表
func (c *signHelpController) List(ctx context.Context, req *install.SignHelpSearchReq) (res *install.SignHelpSearchRes, err error) {
	res = new(install.SignHelpSearchRes)
	res.SignHelpSearchRes, err = service.SignHelp().List(ctx, &req.SignHelpSearchReq)
	return
}

// Export 导出excel
func (c *signHelpController) Export(ctx context.Context, req *install.SignHelpExportReq) (res *install.SignHelpExportRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		listData []*model.SignHelpInfoRes
		//表头
		tableHead = []interface{}{"ID", "标题", "默认展开", "权重", "创建日期", "修改日期"}
		excelData [][]interface{}
		//字典选项处理
	)
	req.PageNum = 1
	req.PageSize = 500
	//获取字典数据
	excelData = append(excelData, tableHead)
	for {
		listData, err = service.SignHelp().GetExportData(ctx, &req.SignHelpSearchReq)
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
				v.Title,
				v.IsExpand,
				v.Weigh,
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("帮助中心")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *signHelpController) ExcelTemplate(ctx context.Context, req *install.SignHelpExcelTemplateReq) (res *install.SignHelpExcelTemplateRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
		//表头
		tableHead = []interface{}{"标题", "内容", "默认展开", "权重", "创建日期", "修改日期"}
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("帮助中心模板")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *signHelpController) Import(ctx context.Context, req *install.SignHelpImportReq) (res *install.SignHelpImportRes, err error) {
	err = service.SignHelp().Import(ctx, req.File)
	return
}

// Get 获取帮助中心
func (c *signHelpController) Get(ctx context.Context, req *install.SignHelpGetReq) (res *install.SignHelpGetRes, err error) {
	res = new(install.SignHelpGetRes)
	res.SignHelpInfoRes, err = service.SignHelp().GetById(ctx, req.Id)
	return
}

// Add 添加帮助中心
func (c *signHelpController) Add(ctx context.Context, req *install.SignHelpAddReq) (res *install.SignHelpAddRes, err error) {
	err = service.SignHelp().Add(ctx, req.SignHelpAddReq)
	return
}

// Edit 修改帮助中心
func (c *signHelpController) Edit(ctx context.Context, req *install.SignHelpEditReq) (res *install.SignHelpEditRes, err error) {
	err = service.SignHelp().Edit(ctx, req.SignHelpEditReq)
	return
}

// Delete 删除帮助中心
func (c *signHelpController) Delete(ctx context.Context, req *install.SignHelpDeleteReq) (res *install.SignHelpDeleteRes, err error) {
	err = service.SignHelp().Delete(ctx, req.Ids)
	return
}

// 帮助中心默认展开修改（状态）
func (c *signHelpController) ChangeIsExpand(ctx context.Context, req *install.SignHelpIsExpandSwitchReq) (res *install.SignHelpIsExpandSwitchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/install/signHelp/edit") {
		err = errors.New("没有修改权限")
		return
	}
	err = service.SignHelp().ChangeIsExpand(ctx, req.Id, req.IsExpand)
	return
}
