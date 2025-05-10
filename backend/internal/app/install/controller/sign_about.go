// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-04-14 19:08:27
// 生成路径: internal/app/install/controller/sign_about.go
// 生成人：smithy
// desc:关于我们
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"
	"errors"
	"mangosmithy/api/v1/install"
	systemApi "mangosmithy/api/v1/system"
	commonService "mangosmithy/internal/app/common/service"
	"mangosmithy/internal/app/install/model"
	"mangosmithy/internal/app/install/service"
	systemController "mangosmithy/internal/app/system/controller"
	systemService "mangosmithy/internal/app/system/service"
	"mangosmithy/library/libUtils"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
)

type signAboutController struct {
	systemController.BaseController
}

var SignAbout = new(signAboutController)

// List 列表
func (c *signAboutController) List(ctx context.Context, req *install.SignAboutSearchReq) (res *install.SignAboutSearchRes, err error) {
	res = new(install.SignAboutSearchRes)
	res.SignAboutSearchRes, err = service.SignAbout().List(ctx, &req.SignAboutSearchReq)
	return
}

// Export 导出excel
func (c *signAboutController) Export(ctx context.Context, req *install.SignAboutExportReq) (res *install.SignAboutExportRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		listData []*model.SignAboutInfoRes
		//表头
		tableHead = []interface{}{"ID", "图标", "标题", "内容", "是否链接", "链接", "分组", "权重", "创建日期", "修改日期"}
		excelData [][]interface{}
		//字典选项处理
		signAboutGroup    *systemApi.GetDictRes
		signAboutGroupMap = gmap.New()
	)
	req.PageNum = 1
	req.PageSize = 500
	//获取字典数据
	signAboutGroup, err = commonService.SysDictData().GetDictWithDataByType(ctx, "sign_about_group", "")
	if err != nil {
		return
	}
	for _, v := range signAboutGroup.Values {
		signAboutGroupMap.Set(v.DictValue, v.DictLabel)
	}
	excelData = append(excelData, tableHead)
	for {
		listData, err = service.SignAbout().GetExportData(ctx, &req.SignAboutSearchReq)
		if err != nil {
			return
		}
		if listData == nil {
			break
		}
		for _, v := range listData {
			var (
				//单选-分组
				group interface{}
			)
			//字典-单选-分组
			group = signAboutGroupMap.Get(gconv.String(v.Group))
			dt := []interface{}{
				v.Id,
				v.Icon,
				v.Title,
				v.Subtitle,
				v.IsLink,
				v.Url,
				//字典-单选-分组
				group,
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("关于我们")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *signAboutController) ExcelTemplate(ctx context.Context, req *install.SignAboutExcelTemplateReq) (res *install.SignAboutExcelTemplateRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
		//表头
		tableHead = []interface{}{"图标", "标题", "内容", "是否链接", "链接", "分组", "权重", "创建日期", "修改日期"}
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("关于我们模板")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *signAboutController) Import(ctx context.Context, req *install.SignAboutImportReq) (res *install.SignAboutImportRes, err error) {
	err = service.SignAbout().Import(ctx, req.File)
	return
}

// Get 获取关于我们
func (c *signAboutController) Get(ctx context.Context, req *install.SignAboutGetReq) (res *install.SignAboutGetRes, err error) {
	res = new(install.SignAboutGetRes)
	res.SignAboutInfoRes, err = service.SignAbout().GetById(ctx, req.Id)
	return
}

// Add 添加关于我们
func (c *signAboutController) Add(ctx context.Context, req *install.SignAboutAddReq) (res *install.SignAboutAddRes, err error) {
	err = service.SignAbout().Add(ctx, req.SignAboutAddReq)
	return
}

// Edit 修改关于我们
func (c *signAboutController) Edit(ctx context.Context, req *install.SignAboutEditReq) (res *install.SignAboutEditRes, err error) {
	err = service.SignAbout().Edit(ctx, req.SignAboutEditReq)
	return
}

// Delete 删除关于我们
func (c *signAboutController) Delete(ctx context.Context, req *install.SignAboutDeleteReq) (res *install.SignAboutDeleteRes, err error) {
	err = service.SignAbout().Delete(ctx, req.Ids)
	return
}

// 关于我们是否链接修改（状态）
func (c *signAboutController) ChangeIsLink(ctx context.Context, req *install.SignAboutIsLinkSwitchReq) (res *install.SignAboutIsLinkSwitchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/install/signAbout/edit") {
		err = errors.New("没有修改权限")
		return
	}
	err = service.SignAbout().ChangeIsLink(ctx, req.Id, req.IsLink)
	return
}
