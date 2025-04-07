// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-02-21 18:37:33
// 生成路径: internal/app/appsource/controller/as_application.go
// 生成人：smithy
// desc:应用管理
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
	"mangosmithy/api/v1/appsource"
	systemApi "mangosmithy/api/v1/system"
	"mangosmithy/internal/app/appsource/model"
	"mangosmithy/internal/app/appsource/service"
	commonService "mangosmithy/internal/app/common/service"
	systemController "mangosmithy/internal/app/system/controller"
	systemService "mangosmithy/internal/app/system/service"
	"mangosmithy/library/libUtils"
)

type asApplicationController struct {
	systemController.BaseController
}

var AsApplication = new(asApplicationController)

// List 列表
func (c *asApplicationController) List(ctx context.Context, req *appsource.AsApplicationSearchReq) (res *appsource.AsApplicationSearchRes, err error) {
	res = new(appsource.AsApplicationSearchRes)
	res.AsApplicationSearchRes, err = service.AsApplication().List(ctx, &req.AsApplicationSearchReq)
	return
}

// Export 导出excel
func (c *asApplicationController) Export(ctx context.Context, req *appsource.AsApplicationExportReq) (res *appsource.AsApplicationExportRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		listData []*model.AsApplicationInfoRes
		//表头
		tableHead = []interface{}{"ID", "文件", "应用名称", "包名", "版本号", "文件大小", "类型", "描述", "图标", "付费", "是否蓝奏云链接", "权重", "是否启用", "备注", "创建人", "修改人", "创建时间", "修改时间"}
		excelData [][]interface{}
		//字典选项处理
		appsourceAppType    *systemApi.GetDictRes
		appsourceAppTypeMap = gmap.New()
		sysNormalDisable    *systemApi.GetDictRes
		sysNormalDisableMap = gmap.New()
		sysYesNo            *systemApi.GetDictRes
		sysYesNoMap         = gmap.New()
	)
	req.PageNum = 1
	req.PageSize = 500
	//获取字典数据
	appsourceAppType, err = commonService.SysDictData().GetDictWithDataByType(ctx, "appsource_app_type", "")
	if err != nil {
		return
	}
	sysNormalDisable, err = commonService.SysDictData().GetDictWithDataByType(ctx, "sys_normal_disable", "")
	if err != nil {
		return
	}
	sysYesNo, err = commonService.SysDictData().GetDictWithDataByType(ctx, "sys_yes_no", "")
	if err != nil {
		return
	}
	for _, v := range appsourceAppType.Values {
		appsourceAppTypeMap.Set(v.DictValue, v.DictLabel)
	}
	for _, v := range sysNormalDisable.Values {
		sysNormalDisableMap.Set(v.DictValue, v.DictLabel)
	}
	for _, v := range sysYesNo.Values {
		sysYesNoMap.Set(v.DictValue, v.DictLabel)
	}
	excelData = append(excelData, tableHead)
	for {
		listData, err = service.AsApplication().GetExportData(ctx, &req.AsApplicationSearchReq)
		if err != nil {
			return
		}
		if listData == nil {
			break
		}
		for _, v := range listData {
			var (
				//单选-类型
				applicationType interface{}
				//单选-付费
				lock interface{}
				//单选-是否蓝奏云链接
				lanzou interface{}
				//单选-是否启用
				active interface{}
				//创建人
				createdBy string
				//修改人
				updatedBy string
			)
			//字典-单选-类型
			applicationType = appsourceAppTypeMap.Get(gconv.String(v.Type))
			//字典-单选-付费
			lock = sysNormalDisableMap.Get(gconv.String(v.Lock))
			//字典-单选-是否蓝奏云链接
			lanzou = sysYesNoMap.Get(gconv.String(v.Lanzou))
			//字典-单选-是否启用
			active = sysNormalDisableMap.Get(gconv.String(v.Active))
			//创建人
			if v.CreatedUser != nil {
				createdBy = v.CreatedUser.UserNickname
			}
			//修改人
			if v.UpdatedUser != nil {
				updatedBy = v.UpdatedUser.UserNickname
			}
			dt := []interface{}{
				v.Id,
				v.FileUrl,
				v.Name,
				v.BundleId,
				v.Version,
				v.Size,
				//字典-单选-类型
				applicationType,
				v.Description,
				v.IconUrl,
				//字典-单选-付费
				lock,
				//字典-单选-是否蓝奏云链接
				lanzou,
				v.Weigh,
				//字典-单选-是否启用
				active,
				v.Note,
				//创建人
				createdBy,
				//修改人
				updatedBy,
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("应用管理")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *asApplicationController) ExcelTemplate(ctx context.Context, req *appsource.AsApplicationExcelTemplateReq) (res *appsource.AsApplicationExcelTemplateRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
		//表头
		tableHead = []interface{}{"文件", "应用名称", "包名", "版本号", "文件大小", "类型", "描述", "图标", "付费", "是否蓝奏云链接", "权重", "是否启用", "备注", "创建人", "修改人", "创建时间", "修改时间"}
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("应用管理模板")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *asApplicationController) Import(ctx context.Context, req *appsource.AsApplicationImportReq) (res *appsource.AsApplicationImportRes, err error) {
	err = service.AsApplication().Import(ctx, req.File)
	return
}

// Get 获取应用管理
func (c *asApplicationController) Get(ctx context.Context, req *appsource.AsApplicationGetReq) (res *appsource.AsApplicationGetRes, err error) {
	res = new(appsource.AsApplicationGetRes)
	res.AsApplicationInfoRes, err = service.AsApplication().GetById(ctx, req.Id)
	return
}

// Add 添加应用管理
func (c *asApplicationController) Add(ctx context.Context, req *appsource.AsApplicationAddReq) (res *appsource.AsApplicationAddRes, err error) {
	err = service.AsApplication().Add(ctx, req.AsApplicationAddReq)
	return
}

// Edit 修改应用管理
func (c *asApplicationController) Edit(ctx context.Context, req *appsource.AsApplicationEditReq) (res *appsource.AsApplicationEditRes, err error) {
	err = service.AsApplication().Edit(ctx, req.AsApplicationEditReq)
	return
}

// Delete 删除应用管理
func (c *asApplicationController) Delete(ctx context.Context, req *appsource.AsApplicationDeleteReq) (res *appsource.AsApplicationDeleteRes, err error) {
	err = service.AsApplication().Delete(ctx, req.Ids)
	return
}

// Extract 解析应用
func (c *asApplicationController) Extract(ctx context.Context, req *appsource.AsApplicationExtractReq) (res *appsource.AsApplicationExtractRes, err error) {
	res = new(appsource.AsApplicationExtractRes)
	res.AsApplicationExtractRes, err = service.AsApplication().Extract(ctx, req.AsApplicationExtractReq)
	return
}

// IconSearch 搜索应用图标
func (c *asApplicationController) IconSearch(ctx context.Context, req *appsource.AsApplicationIconSearchReq) (res *appsource.AsApplicationIconSearchRes, err error) {
	resData, err := service.AsApplication().IconSearch(ctx, req.AsApplicationIconSearchReq)
	r := ghttp.RequestFromCtx(ctx)

	if err != nil {
		return nil, err
	} else {
		r.Response.WriteJsonExit(g.Map{
			"code":    0,
			"message": "获取成功",
			"data":    resData,
		})
	}
	return
}

// SourceImport 软件源导入
func (c *asApplicationController) SourceImport(ctx context.Context, req *appsource.AsApplicationSourceImportReq) (res *appsource.AsApplicationSourceImportRes, err error) {
	res = new(appsource.AsApplicationSourceImportRes)
	err = service.AsApplication().SourceImport(ctx, req.AsApplicationSourceImportReq)
	return
}

// ChangeLock 应用管理付费修改（状态）
func (c *asApplicationController) ChangeLock(ctx context.Context, req *appsource.AsApplicationLockSwitchReq) (res *appsource.AsApplicationLockSwitchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/appsource/asApplication/edit") {
		err = errors.New("没有修改权限")
		return
	}
	err = service.AsApplication().ChangeLock(ctx, req.Id, req.Lock)
	return
}

// ChangeActive 应用管理是否启用修改（状态）
func (c *asApplicationController) ChangeActive(ctx context.Context, req *appsource.AsApplicationActiveSwitchReq) (res *appsource.AsApplicationActiveSwitchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/appsource/asApplication/edit") {
		err = errors.New("没有修改权限")
		return
	}
	err = service.AsApplication().ChangeActive(ctx, req.Id, req.Active)
	return
}
