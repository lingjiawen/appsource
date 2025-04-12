// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2025-03-25 13:14:15
// 生成路径: internal/app/install/controller/sign_redeem_code.go
// 生成人：smithy
// desc:签名卡密
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

type signRedeemCodeController struct {
	systemController.BaseController
}

var SignRedeemCode = new(signRedeemCodeController)

// List 列表
func (c *signRedeemCodeController) List(ctx context.Context, req *install.SignRedeemCodeSearchReq) (res *install.SignRedeemCodeSearchRes, err error) {
	res = new(install.SignRedeemCodeSearchRes)
	res.SignRedeemCodeSearchRes, err = service.SignRedeemCode().List(ctx, &req.SignRedeemCodeSearchReq)
	return
}

// Export 导出excel
func (c *signRedeemCodeController) Export(ctx context.Context, req *install.SignRedeemCodeExportReq) (res *install.SignRedeemCodeExportRes, err error) {
	var (
		r        = ghttp.RequestFromCtx(ctx)
		listData []*model.SignRedeemCodeInfoRes
		//表头
		tableHead = []interface{}{"ID", "兑换码", "设备码", "证书标识", "时效类型", "售后类型", "设备类型", "出书方式", "对接平台", "备注", "对接售后类型", "禁用", "激活", "激活时间", "创建人", "修改人", "创建时间", "修改时间", "删除时间"}
		excelData [][]interface{}
		//字典选项处理
		appleAccountType     *systemApi.GetDictRes
		appleAccountTypeMap  = gmap.New()
		appleWarrantyType    *systemApi.GetDictRes
		appleWarrantyTypeMap = gmap.New()
		appleDeviceType      *systemApi.GetDictRes
		appleDeviceTypeMap   = gmap.New()
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
	for _, v := range applePoolType.Values {
		applePoolTypeMap.Set(v.DictValue, v.DictLabel)
	}
	for _, v := range apiPlatformType.Values {
		apiPlatformTypeMap.Set(v.DictValue, v.DictLabel)
	}
	excelData = append(excelData, tableHead)
	for {
		listData, err = service.SignRedeemCode().GetExportData(ctx, &req.SignRedeemCodeSearchReq)
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
				v.Code,
				v.Udid,
				v.CertId,
				//字典-单选-时效类型
				accountType,
				//字典-单选-售后类型
				warrantyType,
				//字典-单选-设备类型
				deviceType,
				//字典-单选-出书方式
				pool,
				//字典-单选-对接平台
				apiPlatform,
				v.Note,
				//字典-单选-对接售后类型
				apiWarrantyType,
				v.Banned,
				v.Active,
				v.ActiveAt.Format("Y-m-d H:i:s"),
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("签名卡密")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *signRedeemCodeController) ExcelTemplate(ctx context.Context, req *install.SignRedeemCodeExcelTemplateReq) (res *install.SignRedeemCodeExcelTemplateRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
		//表头
		tableHead = []interface{}{"兑换码", "设备码", "证书标识", "时效类型", "售后类型", "设备类型", "出书方式", "对接平台", "备注", "对接售后类型", "禁用", "激活", "激活时间", "创建人", "修改人", "创建时间", "修改时间", "删除时间"}
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
	r.Response.Header().Set("Content-Disposition", "attachment; filename="+gurl.Encode("签名卡密模板")+".xlsx")
	r.Response.Buffer()
	r.Exit()
	return
}
func (c *signRedeemCodeController) Import(ctx context.Context, req *install.SignRedeemCodeImportReq) (res *install.SignRedeemCodeImportRes, err error) {
	err = service.SignRedeemCode().Import(ctx, req.File)
	return
}

// Get 获取签名卡密
func (c *signRedeemCodeController) Get(ctx context.Context, req *install.SignRedeemCodeGetReq) (res *install.SignRedeemCodeGetRes, err error) {
	res = new(install.SignRedeemCodeGetRes)
	res.SignRedeemCodeInfoRes, err = service.SignRedeemCode().GetById(ctx, req.Id)
	return
}

// Add 添加签名卡密
func (c *signRedeemCodeController) Add(ctx context.Context, req *install.SignRedeemCodeAddReq) (res *install.SignRedeemCodeAddRes, err error) {
	err = service.SignRedeemCode().Add(ctx, req.SignRedeemCodeAddReq)
	return
}

// Edit 修改签名卡密
func (c *signRedeemCodeController) Edit(ctx context.Context, req *install.SignRedeemCodeEditReq) (res *install.SignRedeemCodeEditRes, err error) {
	err = service.SignRedeemCode().Edit(ctx, req.SignRedeemCodeEditReq)
	return
}

// Delete 删除签名卡密
func (c *signRedeemCodeController) Delete(ctx context.Context, req *install.SignRedeemCodeDeleteReq) (res *install.SignRedeemCodeDeleteRes, err error) {
	err = service.SignRedeemCode().Delete(ctx, req.Ids)
	return
}

// ChangeForce 签名卡密强制修改（状态）
func (c *signRedeemCodeController) ChangeForce(ctx context.Context, req *install.SignRedeemCodeForceSwitchReq) (res *install.SignRedeemCodeForceSwitchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/install/signRedeemCode/edit") {
		err = errors.New("没有修改权限")
		return
	}
	err = service.SignRedeemCode().ChangeForce(ctx, req.Id, req.Force)
	return
}

// ChangeBanned 签名卡密禁用修改（状态）
func (c *signRedeemCodeController) ChangeBanned(ctx context.Context, req *install.SignRedeemCodeBannedSwitchReq) (res *install.SignRedeemCodeBannedSwitchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/install/signRedeemCode/edit") {
		err = errors.New("没有修改权限")
		return
	}
	err = service.SignRedeemCode().ChangeBanned(ctx, req.Id, req.Banned)
	return
}

// ChangeActive 签名卡密激活修改（状态）
func (c *signRedeemCodeController) ChangeActive(ctx context.Context, req *install.SignRedeemCodeActiveSwitchReq) (res *install.SignRedeemCodeActiveSwitchRes, err error) {
	if !systemService.SysUser().AccessRule(ctx, systemService.Context().GetUserId(ctx), "api/v1/install/signRedeemCode/edit") {
		err = errors.New("没有修改权限")
		return
	}
	err = service.SignRedeemCode().ChangeActive(ctx, req.Id, req.Active)
	return
}
