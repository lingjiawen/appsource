package controller

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gurl"
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

type asRedeemCodeController struct {
	systemController.BaseController
}

var AsRedeemCode = new(asRedeemCodeController)

// List 列表
func (c *asRedeemCodeController) List(ctx context.Context, req *appsource.AsRedeemCodeSearchReq) (res *appsource.AsRedeemCodeSearchRes, err error) {
	res = new(appsource.AsRedeemCodeSearchRes)
	res.AsRedeemCodeSearchRes, err = service.AsRedeemCode().List(ctx, &req.AsRedeemCodeSearchReq)
	return
}

// Export 导出excel
func (c *asRedeemCodeController) Export(ctx context.Context, req *appsource.AsRedeemCodeExportReq) (res *appsource.AsRedeemCodeExportRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		listData []*model.AsRedeemCodeInfoRes
		//表头
		tableHead = []interface{}{"ID", "兑换码", "设备码", "类型", "激活", "激活时间", "创建人", "修改人", "创建时间", "修改时间", "删除时间"}
		excelData [][]interface{}
		//字典选项处理
		sysShowHide    *systemApi.GetDictRes
		sysShowHideMap = gmap.New()
	)
	req.PageNum = 1
	req.PageSize = 500
	//获取字典数据
	sysShowHide, err = commonService.SysDictData().GetDictWithDataByType(ctx, "sys_show_hide", "")
	if err != nil {
		return
	}
	for _, v := range sysShowHide.Values {
		sysShowHideMap.Set(v.DictValue, v.DictLabel)
	}
	excelData = append(excelData, tableHead)
	for {
		listData, err = service.AsRedeemCode().GetExportData(ctx, &req.AsRedeemCodeSearchReq)
		if err != nil {
			return
		}
		if listData == nil {
			break
		}
		for _, v := range listData {
			var (
				//单选-类型
				codeType interface{}
				//创建人
				createdBy string
				//修改人
				updatedBy string
			)
			//字典-单选-类型
			codeType = sysShowHideMap.Get(gconv.String(v.Type))
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
				v.Code,
				v.Udid,
				//字典-单选-类型
				codeType,
				v.Active,
				v.ActiveAt.Format("Y-m-d H:i:s"),
				//创建人
				createdBy,
				//修改人
				updatedBy,
				v.Note,
				v.CreatedAt.Format("Y-m-d H:i:s"),
				v.UpdatedAt.Format("Y-m-d H:i:s"),
				v.DeletedAt.Format("Y-m-d H:i:s"),
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("卡密管理")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *asRedeemCodeController) ExcelTemplate(ctx context.Context, req *appsource.AsRedeemCodeExcelTemplateReq) (res *appsource.AsRedeemCodeExcelTemplateRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
		//表头
		tableHead = []interface{}{"ID", "兑换码", "设备码", "类型", "激活", "激活时间", "创建人", "修改人", "创建时间", "修改时间", "删除时间"}
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("卡密管理模板")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *asRedeemCodeController) Import(ctx context.Context, req *appsource.AsRedeemCodeImportReq) (res *appsource.AsRedeemCodeImportRes, err error) {
	err = service.AsRedeemCode().Import(ctx, req.File)
	return
}

// Get 获取卡密管理
func (c *asRedeemCodeController) Get(ctx context.Context, req *appsource.AsRedeemCodeGetReq) (res *appsource.AsRedeemCodeGetRes, err error) {
	res = new(appsource.AsRedeemCodeGetRes)
	res.AsRedeemCodeInfoRes, err = service.AsRedeemCode().GetById(ctx, req.Id)
	return
}

// Add 添加卡密管理
func (c *asRedeemCodeController) Add(ctx context.Context, req *appsource.AsRedeemCodeAddReq) (res *appsource.AsRedeemCodeAddRes, err error) {
	err = service.AsRedeemCode().Add(ctx, req.AsRedeemCodeAddReq)
	return
}

// Edit 修改卡密管理
func (c *asRedeemCodeController) Edit(ctx context.Context, req *appsource.AsRedeemCodeEditReq) (res *appsource.AsRedeemCodeEditRes, err error) {
	err = service.AsRedeemCode().Edit(ctx, req.AsRedeemCodeEditReq)
	return
}

// Delete 删除卡密管理
func (c *asRedeemCodeController) Delete(ctx context.Context, req *appsource.AsRedeemCodeDeleteReq) (res *appsource.AsRedeemCodeDeleteRes, err error) {
	err = service.AsRedeemCode().Delete(ctx, req.Ids)
	return
}

// ChangeActive 卡密管理是否激活修改（状态）
func (c *asRedeemCodeController) ChangeActive(ctx context.Context, req *appsource.AsRedeemCodeActiveSwitchReq) (res *appsource.AsRedeemCodeActiveSwitchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/appsource/asRedeemCode/edit") {
		err = errors.New("没有修改权限")
		return
	}
	err = service.AsRedeemCode().ChangeActive(ctx, req.Id, req.Active)
	return
}
