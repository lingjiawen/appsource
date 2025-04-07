package sysConfig

import (
	"context"
	"mangosmithy/internal/app/install/dao"
	"mangosmithy/internal/app/install/model"
	"mangosmithy/internal/app/install/service"
	"mangosmithy/library/liberr"

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
		m := dao.SysConfig.Ctx(ctx).WhereIn(dao.SysConfig.Columns().Group, []string{
			"install.frontend", "install.free", "install.download",
		})
		order := "config_id desc"
		var res []*model.SysConfigListRes
		err = m.Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = g.Map{}
		for _, v := range res {
			// 把configKey去除group作为key,configValue作为value
			key := v.ConfigKey
			if v.Group != "" {
				key = v.ConfigKey[len(v.Group)+1:]
			}
			listRes.List[key] = v.ConfigValue
		}
	})
	return
}
