// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-02-18 20:01:49
// 生成路径: internal/app/appsource/logic/sys_config.go
// 生成人：smithy
// desc:配置项
// company:云南奇讯科技有限公司
// ==========================================================================

package sysConfig

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"howett.net/plist"
	"io"
	"log"
	"mangosmithy/internal/app/appsource/common"
	"mangosmithy/internal/app/appsource/dao"
	"mangosmithy/internal/app/appsource/model"
	"mangosmithy/internal/app/appsource/model/do"
	"mangosmithy/internal/app/appsource/service"
	commonService "mangosmithy/internal/app/common/service"
	"mangosmithy/library/libUtils"
	"mangosmithy/library/liberr"
	"net/url"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterSysConfig(New())
}

func New() service.ISysConfig {
	return &sSysConfig{}
}

type sSysConfig struct{}

func (s *sSysConfig) List(ctx context.Context, req *model.SysConfigSearchReq) (listRes *model.SysConfigSearchRes, err error) {
	listRes = new(model.SysConfigSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysConfig.Ctx(ctx).WithAll()
		if req.Group != "" {
			m = m.Where(dao.SysConfig.Columns().Group+" = ?", req.Group)
		}
		order := "config_id asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.SysConfigListRes
		err = m.Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")

		listRes.List = make(map[string][]model.SysConfigListRes, len(res))
		groupSet := make(map[string]struct{}) // 用于去重
		groups := make([]string, 0, len(res))

		for _, v := range res {
			group := v.Group
			// 如果 group 还没有添加到 groupSet，则添加到 Groups
			if _, exists := groupSet[group]; !exists {
				groupSet[group] = struct{}{} // 标记这个 group 已添加
				groups = append(groups, group)
			}

			// 按照 group 分类存入 List
			listRes.List[group] = append(listRes.List[group], model.SysConfigListRes{
				ConfigId:    v.ConfigId,
				ConfigName:  v.ConfigName,
				ConfigKey:   v.ConfigKey,
				ConfigValue: v.ConfigValue,
				ConfigType:  v.ConfigType,
				AllowDel:    v.AllowDel,
				ElementType: v.ElementType,
				Group:       v.Group,
				Remark:      v.Remark,
				CreatedAt:   v.CreatedAt,
			})
		}

		// 查询group
		var linkedSysConfigSysConfigGroup []*model.LinkedSysConfigSysConfigGroup
		err = dao.SysConfigGroup.Ctx(ctx).Scan(&linkedSysConfigSysConfigGroup)

		listRes.Groups = g.Map{}
		for _, v := range linkedSysConfigSysConfigGroup {
			listRes.Groups[v.GroupKey] = v.GroupName
		}
	})
	return
}

// LinkedSysConfigDataSearch 相关连表查询数据
func (s *sSysConfig) LinkedSysConfigDataSearch(ctx context.Context) (res *model.LinkedSysConfigDataSearchRes, err error) {
	res = new(model.LinkedSysConfigDataSearchRes)
	res.LinkedSysConfigSysConfigGroup, err = s.ListSysConfigSysConfigGroup(ctx)
	liberr.ErrIsNil(ctx, err, "获取关联表信息失败")
	return
}

func (s *sSysConfig) GetByConfigId(ctx context.Context, configId uint) (res *model.SysConfigInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysConfig.Ctx(ctx).WithAll().Where(dao.SysConfig.Columns().ConfigId, configId).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sSysConfig) Add(ctx context.Context, req *model.SysConfigAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysConfig.Ctx(ctx).Insert(do.SysConfig{
			ConfigName:  req.ConfigName,
			ConfigKey:   req.ConfigKey,
			ConfigValue: req.ConfigValue,
			ConfigType:  req.ConfigType,
			Group:       req.Group,
			Remark:      req.Remark,
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sSysConfig) Edit(ctx context.Context, req *model.SysConfigEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		configs := req.Configs
		if len(configs) == 0 {
			liberr.ErrIsNil(ctx, errors.New("未提交数据"), "未提交数据")
		}

		// 获取所有要更新的 ConfigId
		var configIds []interface{}
		for _, config := range configs {
			configIds = append(configIds, config.ConfigId)
		}

		// 查询数据库中的原始数据
		var existingConfigs []do.SysConfig
		err := dao.SysConfig.Ctx(ctx).
			WhereIn(dao.SysConfig.Columns().ConfigId, configIds).
			Scan(&existingConfigs)
		liberr.ErrIsNil(ctx, err, "查询配置失败")

		// 构造 ConfigId 到新数据的映射
		configMap := make(map[int]*model.SysConfigReqData)
		for _, config := range configs {
			configMap[int(config.ConfigId)] = config
		}

		var updateData []do.SysConfig
		for _, existingConfig := range existingConfigs {
			// 如果新数据中包含这个 ID，就更新字段
			if newConfig, ok := configMap[gconv.Int(existingConfig.ConfigId)]; ok {
				updatedConfig := existingConfig // **拷贝原始结构体**
				updatedConfig.ConfigValue = newConfig.ConfigValue
				updatedConfig.Remark = newConfig.Remark
				updateData = append(updateData, updatedConfig)
			}
		}

		// 批量保存
		_, err = dao.SysConfig.Ctx(ctx).Save(updateData)
		liberr.ErrIsNil(ctx, err, "批量修改失败")
	})
	return
}

func (s *sSysConfig) Delete(ctx context.Context, configIds []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysConfig.Ctx(ctx).Delete(dao.SysConfig.Columns().ConfigId+" in (?)", configIds)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

func (s *sSysConfig) ListSysConfigSysConfigGroup(ctx context.Context) (linkedSysConfigSysConfigGroup []*model.LinkedSysConfigSysConfigGroup, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysConfigGroup.
			Ctx(ctx).
			Fields(model.LinkedSysConfigSysConfigGroup{}).Scan(&linkedSysConfigSysConfigGroup)
		liberr.ErrIsNil(ctx, err)
	})
	return
}

func (s *sSysConfig) ParseAppInfo(ctx context.Context, req *model.ParseAppInfoReq) (res *model.ParseAppInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		res = &model.ParseAppInfoRes{}
		fullAppPath := common.GetFullDirPath(ctx, req.AppPath)

		r, err := zip.OpenReader(fullAppPath)
		if err != nil {
			liberr.ErrIsNil(ctx, fmt.Errorf("failed to open ipa file: %v", err))
			return
		}
		defer r.Close()

		// 查找 .app 文件夹路径
		var appPath string
		for _, f := range r.File {
			if strings.HasPrefix(f.Name, "Payload/") && strings.Contains(f.Name, ".app/") {
				// 找到 .app 文件夹路径
				appPath = strings.Split(f.Name, ".app/")[0] + ".app/"
				break
			}
		}

		if appPath == "" {
			liberr.ErrIsNil(ctx, fmt.Errorf("no .app folder found in the IPA"))
			return
		}

		// Info.plist 文件路径
		plistPath := appPath + "Info.plist"

		// 查找 Info.plist 文件
		var plistReader io.Reader
		for _, f := range r.File {
			if f.Name == plistPath {
				// 找到 Info.plist 文件
				plistReader, err = f.Open()
				if err != nil {
					log.Fatalf("failed to open plist file: %v", err)
				}
				break
			}
		}

		if plistReader == nil {
			log.Fatalf("Info.plist not found in the IPA")
		}

		// 读取 Info.plist 文件的内容
		plistData, err := io.ReadAll(plistReader)
		if err != nil {
			liberr.ErrIsNil(ctx, fmt.Errorf("读取Plist文件失败: %v", err))
			return
		}

		// 使用 plist.Unmarshal 解析 plist 数据
		var infoPlist model.InfoPlist
		_, err = plist.Unmarshal(plistData, &infoPlist)
		if err != nil {
			liberr.ErrIsNil(ctx, fmt.Errorf("解析Plist文件失败: %v", err))
			return
		}

		res.BundleID = infoPlist.BundleIdentifier
		res.Version = infoPlist.BundleShortVersion
		res.Name = infoPlist.DisplayName

		iconUrlConfig, err := commonService.SysConfig().GetConfigByKey(ctx, "appsource.iconUrl")
		liberr.ErrIsNil(ctx, err)
		if iconUrlConfig == nil || iconUrlConfig.ConfigValue == "" {
			// 处理图标信息
			iconPath := common.HandleIconInfo(&infoPlist, &r.Reader, appPath)
			icon, iconError := common.ExtractIconFileAndUpload(ctx, &r.Reader, appPath, iconPath)
			if iconError != nil {
				liberr.ErrIsNil(ctx, fmt.Errorf("上传图片失败: %v", iconError.Error()))
			}

			urlPrefix := libUtils.GetDomain(ctx, false)
			res.Icon, _ = url.JoinPath(urlPrefix, icon)
		} else {
			res.Icon = iconUrlConfig.ConfigValue
		}

	})
	return
}
