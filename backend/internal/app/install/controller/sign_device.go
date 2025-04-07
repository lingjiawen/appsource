// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-04-01 19:38:11
// 生成路径: internal/app/install/controller/sign_device.go
// 生成人：smithy
// desc:设备管理
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

type signDeviceController struct {
	systemController.BaseController
}

var SignDevice = new(signDeviceController)

// List 列表
func (c *signDeviceController) List(ctx context.Context, req *install.SignDeviceSearchReq) (res *install.SignDeviceSearchRes, err error) {
	res = new(install.SignDeviceSearchRes)
	res.SignDeviceSearchRes, err = service.SignDevice().List(ctx, &req.SignDeviceSearchReq)
	return
}

// Export 导出excel
func (c *signDeviceController) Export(ctx context.Context, req *install.SignDeviceExportReq) (res *install.SignDeviceExportRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		listData []*model.SignDeviceInfoRes
		//表头
		tableHead = []interface{}{"ID", "设备码", "证书名", "证书标识", "设备标识", "证书文件", "描述文件", "添加时间", "证书密码", "设备型号", "过期时间", "序列号", "兑换卡密", "时效类型", "售后类型", "设备类型", "状态", "出书方式", "对接平台", "对接售后类型", "禁用", "创建人", "修改人", "创建时间", "修改时间", "删除时间"}
		excelData [][]interface{}
		//字典选项处理
		appleAccountType     *systemApi.GetDictRes
		appleAccountTypeMap  = gmap.New()
		appleWarrantyType    *systemApi.GetDictRes
		appleWarrantyTypeMap = gmap.New()
		appleDeviceType      *systemApi.GetDictRes
		appleDeviceTypeMap   = gmap.New()
		certStatus           *systemApi.GetDictRes
		certStatusMap        = gmap.New()
		applePoolType        *systemApi.GetDictRes
		applePoolTypeMap     = gmap.New()
		apiPlatformType      *systemApi.GetDictRes
		apiPlatformTypeMap   = gmap.New()
	)
	req.PageNum = 1
	req.PageSize = 500
	//获取字典数据
	appleAccountType, err = commonService.SysDictData().GetDictWithDataByType(ctx, "apple_account_type", "")
	if err != nil {
		return
	}
	appleWarrantyType, err = commonService.SysDictData().GetDictWithDataByType(ctx, "apple_warranty_type", "")
	if err != nil {
		return
	}
	appleDeviceType, err = commonService.SysDictData().GetDictWithDataByType(ctx, "apple_device_type", "")
	if err != nil {
		return
	}
	certStatus, err = commonService.SysDictData().GetDictWithDataByType(ctx, "cert_status", "")
	if err != nil {
		return
	}
	applePoolType, err = commonService.SysDictData().GetDictWithDataByType(ctx, "apple_pool_type", "")
	if err != nil {
		return
	}
	apiPlatformType, err = commonService.SysDictData().GetDictWithDataByType(ctx, "api_platform_type", "")
	if err != nil {
		return
	}
	for _, v := range appleAccountType.Values {
		appleAccountTypeMap.Set(v.DictValue, v.DictLabel)
	}
	for _, v := range appleWarrantyType.Values {
		appleWarrantyTypeMap.Set(v.DictValue, v.DictLabel)
	}
	for _, v := range appleDeviceType.Values {
		appleDeviceTypeMap.Set(v.DictValue, v.DictLabel)
	}
	for _, v := range certStatus.Values {
		certStatusMap.Set(v.DictValue, v.DictLabel)
	}
	for _, v := range applePoolType.Values {
		applePoolTypeMap.Set(v.DictValue, v.DictLabel)
	}
	for _, v := range apiPlatformType.Values {
		apiPlatformTypeMap.Set(v.DictValue, v.DictLabel)
	}
	excelData = append(excelData, tableHead)
	for {
		listData, err = service.SignDevice().GetExportData(ctx, &req.SignDeviceSearchReq)
		if err != nil {
			return
		}
		if listData == nil {
			break
		}
		for _, v := range listData {
			var (
				//单选-时效类型
				accountType interface{}
				//单选-售后类型
				warrantyType interface{}
				//单选-设备类型
				deviceType interface{}
				//单选-状态
				status interface{}
				//单选-出书方式
				pool interface{}
				//单选-对接平台
				apiPlatform interface{}
				//单选-对接售后类型
				apiWarrantyType interface{}
				//创建人
				createdBy string
				//修改人
				updatedBy string
			)
			//字典-单选-时效类型
			accountType = appleAccountTypeMap.Get(gconv.String(v.AccountType))
			//字典-单选-售后类型
			warrantyType = appleWarrantyTypeMap.Get(gconv.String(v.WarrantyType))
			//字典-单选-设备类型
			deviceType = appleDeviceTypeMap.Get(gconv.String(v.DeviceType))
			//字典-单选-状态
			status = certStatusMap.Get(gconv.String(v.Status))
			//字典-单选-出书方式
			pool = applePoolTypeMap.Get(gconv.String(v.Pool))
			//字典-单选-对接平台
			apiPlatform = apiPlatformTypeMap.Get(gconv.String(v.ApiPlatform))
			//字典-单选-对接售后类型
			apiWarrantyType = appleWarrantyTypeMap.Get(gconv.String(v.ApiWarrantyType))
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
				v.Udid,
				v.Name,
				v.CertId,
				v.DeviceId,
				v.P12,
				v.Mobileprovision,
				v.AddTime,
				v.P12Password,
				v.Model,
				v.ExpireTime,
				v.SerialNumber,
				v.RedeemCode,
				//字典-单选-时效类型
				accountType,
				//字典-单选-售后类型
				warrantyType,
				//字典-单选-设备类型
				deviceType,
				//字典-单选-状态
				status,
				//字典-单选-出书方式
				pool,
				//字典-单选-对接平台
				apiPlatform,
				//字典-单选-对接售后类型
				apiWarrantyType,
				v.Active,
				//创建人
				createdBy,
				//修改人
				updatedBy,
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("设备管理")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *signDeviceController) ExcelTemplate(ctx context.Context, req *install.SignDeviceExcelTemplateReq) (res *install.SignDeviceExcelTemplateRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
		//表头
		tableHead = []interface{}{"设备码", "证书名", "证书标识", "设备标识", "证书文件", "描述文件", "添加时间", "证书密码", "设备型号", "过期时间", "序列号", "兑换卡密", "时效类型", "售后类型", "设备类型", "状态", "出书方式", "对接平台", "对接售后类型", "禁用", "创建人", "修改人", "创建时间", "修改时间", "删除时间"}
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("设备管理模板")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *signDeviceController) Import(ctx context.Context, req *install.SignDeviceImportReq) (res *install.SignDeviceImportRes, err error) {
	err = service.SignDevice().Import(ctx, req.File)
	return
}

// Get 获取设备管理
func (c *signDeviceController) Get(ctx context.Context, req *install.SignDeviceGetReq) (res *install.SignDeviceGetRes, err error) {
	res = new(install.SignDeviceGetRes)
	res.SignDeviceInfoRes, err = service.SignDevice().GetById(ctx, req.Id)
	return
}

// Add 添加设备管理
func (c *signDeviceController) Add(ctx context.Context, req *install.SignDeviceAddReq) (res *install.SignDeviceAddRes, err error) {
	err = service.SignDevice().Add(ctx, req.SignDeviceAddReq)
	return
}

// Edit 修改设备管理
func (c *signDeviceController) Edit(ctx context.Context, req *install.SignDeviceEditReq) (res *install.SignDeviceEditRes, err error) {
	err = service.SignDevice().Edit(ctx, req.SignDeviceEditReq)
	return
}

// Delete 删除设备管理
func (c *signDeviceController) Delete(ctx context.Context, req *install.SignDeviceDeleteReq) (res *install.SignDeviceDeleteRes, err error) {
	err = service.SignDevice().Delete(ctx, req.Ids)
	return
}

// 设备管理禁用修改（状态）
func (c *signDeviceController) ChangeActive(ctx context.Context, req *install.SignDeviceActiveSwitchReq) (res *install.SignDeviceActiveSwitchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/install/signDevice/edit") {
		err = errors.New("没有修改权限")
		return
	}
	err = service.SignDevice().ChangeActive(ctx, req.Id, req.Active)
	return
}
