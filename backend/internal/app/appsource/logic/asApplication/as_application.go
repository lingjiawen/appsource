// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2025-02-21 18:37:33
// 生成路径: internal/app/appsource/logic/as_application.go
// 生成人：smithy
// desc:应用管理
// company:云南奇讯科技有限公司
// ==========================================================================

package asApplication

import (
	"archive/zip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
	"howett.net/plist"
	"io"
	"log"
	"mangosmithy/internal/app/appsource/common"
	appsourceConsts "mangosmithy/internal/app/appsource/consts"
	"mangosmithy/internal/app/appsource/dao"
	"mangosmithy/internal/app/appsource/model"
	"mangosmithy/internal/app/appsource/model/do"
	"mangosmithy/internal/app/appsource/service"
	commonConsts "mangosmithy/internal/app/common/consts"
	"mangosmithy/internal/app/system/consts"
	systemServer "mangosmithy/internal/app/system/service"
	systemService "mangosmithy/internal/app/system/service"
	"mangosmithy/library/libWebsocket"
	"mangosmithy/library/liberr"
	"math"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

func init() {
	service.RegisterAsApplication(New())
}

func New() service.IAsApplication {
	return &sAsApplication{}
}

type sAsApplication struct{}

func (s *sAsApplication) List(ctx context.Context, req *model.AsApplicationSearchReq) (listRes *model.AsApplicationSearchRes, err error) {
	listRes = new(model.AsApplicationSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.AsApplication.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.AsApplication.Columns().Id+" = ?", req.Id)
		}
		if req.FileUrl != "" {
			m = m.Where(dao.AsApplication.Columns().FileUrl+" = ?", req.FileUrl)
		}
		if req.Name != "" {
			m = m.Where(dao.AsApplication.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.BundleId != "" {
			m = m.Where(dao.AsApplication.Columns().BundleId+" = ?", req.BundleId)
		}
		if req.Version != "" {
			m = m.Where(dao.AsApplication.Columns().Version+" = ?", req.Version)
		}
		if req.Size != "" {
			m = m.Where(dao.AsApplication.Columns().Size+" = ?", gconv.Uint64(req.Size))
		}
		if req.Type != "" {
			m = m.Where(dao.AsApplication.Columns().Type+" = ?", gconv.Int(req.Type))
		}
		if req.Description != "" {
			m = m.Where(dao.AsApplication.Columns().Description+" = ?", req.Description)
		}
		if req.IconUrl != "" {
			m = m.Where(dao.AsApplication.Columns().IconUrl+" = ?", req.IconUrl)
		}
		if req.Lock != "" {
			m = m.Where(dao.AsApplication.Columns().Lock+" = ?", gconv.Int(req.Lock))
		}
		if req.Lanzou != "" {
			m = m.Where(dao.AsApplication.Columns().Lanzou+" = ?", gconv.Int(req.Lanzou))
		}
		if req.Weigh != "" {
			m = m.Where(dao.AsApplication.Columns().Weigh+" = ?", gconv.Int(req.Weigh))
		}
		if req.Active != "" {
			m = m.Where(dao.AsApplication.Columns().Active+" = ?", gconv.Int(req.Active))
		}
		if req.Note != "" {
			m = m.Where(dao.AsApplication.Columns().Note+" = ?", req.Note)
		}
		if req.CreatedBy != "" {
			m = m.Where(dao.AsApplication.Columns().CreatedBy+" = ?", gconv.Uint(req.CreatedBy))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.AsApplication.Columns().CreatedAt+" >=? AND "+dao.AsApplication.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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
		var res []*model.AsApplicationListRes
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.AsApplicationListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.AsApplicationListRes{
				Id:          v.Id,
				FileUrl:     v.FileUrl,
				Name:        v.Name,
				BundleId:    v.BundleId,
				Version:     v.Version,
				Size:        v.Size,
				Type:        v.Type,
				Description: v.Description,
				IconUrl:     v.IconUrl,
				Lock:        v.Lock,
				Lanzou:      v.Lanzou,
				Weigh:       v.Weigh,
				Active:      v.Active,
				Note:        v.Note,
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
			}
		}
	})
	return
}

func (s *sAsApplication) GetExportData(ctx context.Context, req *model.AsApplicationSearchReq) (listRes []*model.AsApplicationInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.AsApplication.Ctx(ctx).WithAll()
		if req.Id != "" {
			m = m.Where(dao.AsApplication.Columns().Id+" = ?", req.Id)
		}
		if req.FileUrl != "" {
			m = m.Where(dao.AsApplication.Columns().FileUrl+" = ?", req.FileUrl)
		}
		if req.Name != "" {
			m = m.Where(dao.AsApplication.Columns().Name+" like ?", "%"+req.Name+"%")
		}
		if req.BundleId != "" {
			m = m.Where(dao.AsApplication.Columns().BundleId+" = ?", req.BundleId)
		}
		if req.Version != "" {
			m = m.Where(dao.AsApplication.Columns().Version+" = ?", req.Version)
		}
		if req.Size != "" {
			m = m.Where(dao.AsApplication.Columns().Size+" = ?", gconv.Uint64(req.Size))
		}
		if req.Type != "" {
			m = m.Where(dao.AsApplication.Columns().Type+" = ?", gconv.Int(req.Type))
		}
		if req.Description != "" {
			m = m.Where(dao.AsApplication.Columns().Description+" = ?", req.Description)
		}
		if req.IconUrl != "" {
			m = m.Where(dao.AsApplication.Columns().IconUrl+" = ?", req.IconUrl)
		}
		if req.Lock != "" {
			m = m.Where(dao.AsApplication.Columns().Lock+" = ?", gconv.Int(req.Lock))
		}
		if req.Lanzou != "" {
			m = m.Where(dao.AsApplication.Columns().Lanzou+" = ?", gconv.Int(req.Lanzou))
		}
		if req.Weigh != "" {
			m = m.Where(dao.AsApplication.Columns().Weigh+" = ?", gconv.Int(req.Weigh))
		}
		if req.Active != "" {
			m = m.Where(dao.AsApplication.Columns().Active+" = ?", gconv.Int(req.Active))
		}
		if req.Note != "" {
			m = m.Where(dao.AsApplication.Columns().Note+" = ?", req.Note)
		}
		if req.CreatedBy != "" {
			m = m.Where(dao.AsApplication.Columns().CreatedBy+" = ?", gconv.Uint(req.CreatedBy))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.AsApplication.Columns().CreatedAt+" >=? AND "+dao.AsApplication.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
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

func (s *sAsApplication) Import(ctx context.Context, file *ghttp.UploadFile) (err error) {
	if file == nil {
		err = errors.New("请上传数据文件")
		return
	}
	var data []do.AsApplication
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
		data = make([]do.AsApplication, len(rows)-1)
		for k, v := range rows {
			if k == 0 {
				continue
			}
			for kv, vv := range v {
				d[kv] = vv
			}
			data[k-1] = do.AsApplication{
				FileUrl:     d[0],
				Name:        d[1],
				BundleId:    d[2],
				Version:     d[3],
				Size:        gconv.Int64(d[4]),
				Type:        gconv.Int64(d[5]),
				Description: d[6],
				IconUrl:     d[7],
				Lock:        gconv.Int64(d[8]),
				Lanzou:      gconv.Int64(d[9]),
				Weigh:       gconv.Int64(d[10]),
				Active:      gconv.Int64(d[11]),
				Note:        d[12],
				CreatedBy:   gconv.Int64(d[13]),
				UpdatedBy:   gconv.Int64(d[14]),
				CreatedAt:   gconv.GTime(d[15]),
				UpdatedAt:   gconv.GTime(d[16]),
			}
		}
		if len(data) > 0 {
			err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				_, err = dao.AsApplication.Ctx(ctx).Batch(500).Insert(data)
				return err
			})
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (s *sAsApplication) GetById(ctx context.Context, id uint) (res *model.AsApplicationInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.AsApplication.Ctx(ctx).WithAll().Where(dao.AsApplication.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sAsApplication) Add(ctx context.Context, req *model.AsApplicationAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AsApplication.Ctx(ctx).Insert(do.AsApplication{
			FileUrl:     req.FileUrl,
			Name:        req.Name,
			BundleId:    req.BundleId,
			Version:     req.Version,
			Size:        req.Size,
			Type:        req.Type,
			Description: req.Description,
			IconUrl:     req.IconUrl,
			Lock:        req.Lock,
			Lanzou:      req.Lanzou,
			Weigh:       req.Weigh,
			Active:      req.Active,
			Note:        req.Note,
			CreatedBy:   systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sAsApplication) Edit(ctx context.Context, req *model.AsApplicationEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AsApplication.Ctx(ctx).WherePri(req.Id).Update(do.AsApplication{
			FileUrl:     req.FileUrl,
			Name:        req.Name,
			BundleId:    req.BundleId,
			Version:     req.Version,
			Size:        req.Size,
			Type:        req.Type,
			Description: req.Description,
			IconUrl:     req.IconUrl,
			Lock:        req.Lock,
			Lanzou:      req.Lanzou,
			Weigh:       req.Weigh,
			Active:      req.Active,
			Note:        req.Note,
			UpdatedBy:   systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sAsApplication) Delete(ctx context.Context, ids []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AsApplication.Ctx(ctx).Delete(dao.AsApplication.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

func (s *sAsApplication) Extract(ctx context.Context, req *model.AsApplicationExtractReq) (res *model.AsApplicationExtractRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		res = &model.AsApplicationExtractRes{}
		uploadPath := strings.Trim(commonConsts.UploadPath, "/")

		var fullDirPath string
		if strings.HasPrefix(req.FileUrl, uploadPath) {
			// 本地文件处理
			fullDirPath = common.GetFullDirPath(ctx, req.FileUrl)
		} else {
			// 非本地文件，临时下载
			tempFile, err := downloadFileToTemp(ctx, req.FileUrl)
			if err != nil {
				liberr.ErrIsNil(ctx, fmt.Errorf("下载文件失败: %v", err))
				return
			}
			defer os.Remove(tempFile.Name()) // 确保临时文件删除

			fullDirPath = tempFile.Name() // 使用临时文件的路径
		}

		// 获取文件的大小
		fileInfo, err := os.Stat(fullDirPath)
		if err != nil {
			liberr.ErrIsNil(ctx, fmt.Errorf("无法获取文件大小: %v", err))
			return
		}
		// 直接计算并设置文件大小
		res.AppSize = fileInfo.Size()

		// 解压并处理文件
		r, err := zip.OpenReader(fullDirPath)
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

		res.BundleId = infoPlist.BundleIdentifier
		res.Version = infoPlist.BundleShortVersion
		res.Name = infoPlist.DisplayName

		// 处理图标信息
		iconPath := common.HandleIconInfo(&infoPlist, &r.Reader, appPath)
		var iconError error
		res.IconUrl, iconError = common.ExtractIconFileAndUpload(ctx, &r.Reader, appPath, iconPath)
		if iconError != nil {
			liberr.ErrIsNil(ctx, fmt.Errorf("上传图片失败: %v", iconError.Error()))
		}
		//err = fmt.Errorf("上传图片失败: %v", iconError)
	})
	return
}

// IconSearch 搜索图标并返回最快的结果
func (s *sAsApplication) IconSearch(ctx context.Context, req *model.AsApplicationIconSearchReq) (res g.Map, err error) {
	// 需要查询的 API URLs
	apiUrls := []string{
		"https://itunes.apple.com/search?term=" + url.QueryEscape(req.Name) + "&entity=software&country=" + req.Country + "&limit=20",
		"https://itunes.apple.com/cn/search?term=" + url.QueryEscape(req.Name) + "&entity=software&country=" + req.Country + "&limit=20",
		"http://itunes.apple.com/search?term=" + url.QueryEscape(req.Name) + "&entity=software&country=" + req.Country + "&limit=20",
	}

	// 创建协程池，指定并发数
	pool := grpool.New(3) // 根据需求调整并发数
	defer pool.Close()

	// 用于存储返回的结果和错误
	var result []byte
	var finalErr error
	var resultMutex sync.Mutex // 用于同步更新结果

	// 用于记录请求是否已完成
	completed := make(chan bool, 1) // 一个用于标记请求是否完成

	// 遍历所有 URL，发起并发请求
	for _, apiUrl := range apiUrls {
		u := apiUrl // 避免闭包问题

		// 将任务添加到协程池
		err := pool.Add(ctx, func(ctx context.Context) {
			// 发起 HTTP 请求
			resp, err := http.Get(u)
			if err != nil {
				// 如果请求失败，直接返回
				return
			}
			defer resp.Body.Close()

			// 读取响应体
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				// 如果读取失败，直接返回
				return
			}

			// 如果响应数据非空，表示请求成功且有数据
			if len(body) > 0 {
				// 只取第一个成功的响应，并更新结果
				resultMutex.Lock()
				if result == nil { // 确保只取第一个返回的有效结果
					result = body
				}
				resultMutex.Unlock()

				// 标记请求完成
				completed <- true
			}
		})

		// 如果添加任务失败，退出
		if err != nil {
			finalErr = err
			break
		}
	}

	// 使用 select 等待最早的完成请求或者超时
	select {
	case <-completed:
		// 如果有请求成功并返回数据，继续处理
		if len(result) > 0 {
			// 解析返回的 JSON 数据
			err = json.Unmarshal(result, &res)
			if err != nil {
				finalErr = fmt.Errorf("解析响应失败: %v", err)
			}
		} else {
			// 如果响应为空，表示没有数据
			finalErr = fmt.Errorf("未找到相关应用")
		}
	case <-time.After(30 * time.Second):
		// 如果30秒内没有成功返回，表示超时
		finalErr = fmt.Errorf("请求超时，所有服务不可用")
	}

	// 返回最终错误
	return res, finalErr
}

// SourceImport 软件源导入
//func (s *sAsApplication) SourceImport(ctx context.Context, req *model.AsApplicationSourceImportReq) (err error) {
//	err = g.Try(ctx, func(ctx context.Context) {
//		// 发起 HTTP 请求
//		//liberr.ErrIsNil(ctx, fmt.Errorf("测试报错"))
//		resp, err := http.Get(req.SourceUrl)
//		if err != nil {
//			// 如果请求失败，直接返回
//			return
//		}
//		defer resp.Body.Close()
//
//		// 读取响应体
//		body, err := io.ReadAll(resp.Body)
//		liberr.ErrIsNil(ctx, err)
//
//		// 如果响应数据非空，表示请求成功且有数据
//		var result []byte
//		if len(body) > 0 {
//			result = body
//		}
//
//		var appSource *model.AppStoreRes
//		if len(result) > 0 {
//			// 解析返回的 JSON 数据
//			err = json.Unmarshal(result, &appSource)
//			liberr.ErrIsNil(ctx, err, fmt.Sprintf("解析响应失败: %v", err))
//			if req.SettingImport {
//				// 导入配置项
//				name := appSource.Name
//				notification := appSource.Message
//				identifier := appSource.Identifier
//				sourceIcon := appSource.SourceIcon
//				payUrl := appSource.PayURL
//
//				_, err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigKey, "appsource.name").Update(g.Map{
//					dao.SysConfig.Columns().ConfigValue: name,
//				})
//
//				_, err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigKey, "appsource.identifier").Update(g.Map{
//					dao.SysConfig.Columns().ConfigValue: identifier,
//				})
//
//				_, err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigKey, "appsource.purchaseSite").Update(g.Map{
//					dao.SysConfig.Columns().ConfigValue: payUrl,
//				})
//
//				_, err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigKey, "appsource.iconUrl").Update(g.Map{
//					dao.SysConfig.Columns().ConfigValue: sourceIcon,
//				})
//
//				_, err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigKey, "appsource.notification").Update(g.Map{
//					dao.SysConfig.Columns().ConfigValue: notification,
//				})
//
//			}
//
//			apps := appSource.Apps
//			appsToInsert := make([]do.AsApplication, len(apps))
//			for index, app := range apps {
//				weigh := len(apps) - index
//				appsToInsert[index] = do.AsApplication{
//					FileUrl:     app.DownloadURL,
//					Name:        app.Name,
//					BundleId:    "",
//					Version:     app.Version,
//					Size:        math.Round(gconv.Float64(app.Size)),
//					Type:        gconv.Int64(app.Type),
//					Description: app.VersionDescription,
//					IconUrl:     app.IconURL,
//					Lock:        gconv.Int64(app.Lock),
//					Lanzou:      gconv.Int64(app.IsLanZouCloud),
//					Weigh:       weigh,
//					Active:      1,
//					Note:        gtime.Now().String() + "导入",
//					CreatedBy:   1,
//					UpdatedBy:   1,
//					CreatedAt:   gtime.NewFromStrLayout(app.VersionDate, "2006-01-02T15:04:05+08:00"),
//					UpdatedAt:   gtime.NewFromStrLayout(app.VersionDate, "2006-01-02T15:04:05+08:00"),
//				}
//			}
//			if len(appsToInsert) > 0 {
//				err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
//					_, err = dao.AsApplication.Ctx(ctx).Batch(500).Insert(appsToInsert)
//					return err
//				})
//				liberr.ErrIsNil(ctx, err)
//			}
//		} else {
//			// 如果响应为空，表示没有数据
//			liberr.ErrIsNil(ctx, fmt.Errorf("无法获取软件源"))
//		}
//	})
//	return
//}

func (s *sAsApplication) SourceImport(ctx context.Context, req *model.AsApplicationSourceImportReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 发起 HTTP 请求
		resp, err := http.Get(req.SourceUrl)
		if err != nil {
			// 如果请求失败，直接返回
			return
		}
		defer resp.Body.Close()

		// 读取响应体
		body, err := io.ReadAll(resp.Body)
		liberr.ErrIsNil(ctx, err)

		// 如果响应数据非空，表示请求成功且有数据
		var result []byte
		if len(body) > 0 {
			result = body
		}

		var appSource *model.AppStoreRes
		if len(result) > 0 {
			// 解析返回的 JSON 数据
			err = json.Unmarshal(result, &appSource)
			liberr.ErrIsNil(ctx, err, fmt.Sprintf("解析响应失败: %v", err))
			if req.SettingImport {
				// 导入配置项
				name := appSource.Name
				notification := appSource.Message
				identifier := appSource.Identifier
				sourceIcon := appSource.SourceIcon
				payUrl := appSource.PayURL

				_, err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigKey, "appsource.name").Update(g.Map{
					dao.SysConfig.Columns().ConfigValue: name,
				})

				_, err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigKey, "appsource.identifier").Update(g.Map{
					dao.SysConfig.Columns().ConfigValue: identifier,
				})

				_, err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigKey, "appsource.purchaseSite").Update(g.Map{
					dao.SysConfig.Columns().ConfigValue: payUrl,
				})

				_, err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigKey, "appsource.iconUrl").Update(g.Map{
					dao.SysConfig.Columns().ConfigValue: sourceIcon,
				})

				_, err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigKey, "appsource.notification").Update(g.Map{
					dao.SysConfig.Columns().ConfigValue: notification,
				})
			}

			apps := appSource.Apps
			appsToInsert := make([]do.AsApplication, len(apps))

			// 进度更新定时器，每 0.1 秒更新一次
			ticker := time.NewTicker(100 * time.Millisecond)
			defer ticker.Stop()

			// 用于保存当前进度的变量
			var processedCount int

			// 每处理一个应用，就更新进度
			for index, app := range apps {
				weigh := len(apps) - index
				appsToInsert[index] = do.AsApplication{
					FileUrl:     app.DownloadURL,
					Name:        app.Name,
					BundleId:    "",
					Version:     app.Version,
					Size:        math.Round(gconv.Float64(app.Size)),
					Type:        gconv.Int64(app.Type),
					Description: app.VersionDescription,
					IconUrl:     app.IconURL,
					Lock:        gconv.Int64(app.Lock),
					Lanzou:      gconv.Int64(app.IsLanZouCloud),
					Weigh:       weigh,
					Active:      1,
					Note:        gtime.Now().String() + "导入",
					CreatedBy:   1,
					UpdatedBy:   1,
					CreatedAt:   gtime.NewFromStrLayout(app.VersionDate, "2006-01-02T15:04:05+08:00"),
					UpdatedAt:   gtime.NewFromStrLayout(app.VersionDate, "2006-01-02T15:04:05+08:00"),
				}

				// 每 0.1 秒更新一次进度
				select {
				case <-ticker.C:
					progress := float64(processedCount) / float64(len(apps)) * 100
					log.Printf("Data Processing Progress: %.2f%%", progress)
					sendMessage(ctx, fmt.Sprintf("处理数据: %.2f%%", progress), appsourceConsts.WebsocketTypeSourceImport)
					processedCount++
				default:
					// 默认情况下继续执行
				}
			}

			// 处理完所有应用后，做一次最后的进度更新
			sendMessage(ctx, "开始保存数据", appsourceConsts.WebsocketTypeSourceImport)

			// 将应用信息批量插入数据库（分批次插入）
			if len(appsToInsert) > 0 {
				batchSize := 100 // 每批插入 100 条数据
				totalCount := len(appsToInsert)
				insertedCount := 0

				err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
					for i := 0; i < totalCount; i += batchSize {
						end := i + batchSize
						if end > totalCount {
							end = totalCount
						}

						// 插入当前批次的数据
						batch := appsToInsert[i:end]
						_, err = dao.AsApplication.Ctx(ctx).Batch(len(batch)).Insert(batch)
						if err != nil {
							return err
						}

						insertedCount += len(batch)

						// 每 0.1 秒更新一次进度
						select {
						case <-ticker.C:
							progress := float64(insertedCount) / float64(totalCount) * 100
							log.Printf("Database Insert Progress: %.2f%%", progress)
							sendMessage(ctx, fmt.Sprintf("正在保存数据: %.2f%%", progress), appsourceConsts.WebsocketTypeSourceImport)
						default:
						}
					}
					return nil
				})

				liberr.ErrIsNil(ctx, err)
			}
		} else {
			// 如果响应为空，表示没有数据
			liberr.ErrIsNil(ctx, fmt.Errorf("无法获取软件源"))
		}
	})
	return
}

// ChangeLock 应用管理付费修改（状态）
func (s *sAsApplication) ChangeLock(ctx context.Context, id uint, lock int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AsApplication.Ctx(ctx).WherePri(id).
			Update(do.AsApplication{
				Lock: lock,
			})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

// ChangeActive 应用管理是否启用修改（状态）
func (s *sAsApplication) ChangeActive(ctx context.Context, id uint, active int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.AsApplication.Ctx(ctx).WherePri(id).
			Update(do.AsApplication{
				Active: active,
			})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

// ProgressWriter 用于实时显示下载进度
type ProgressWriter struct {
	Download int64
	Total    int64
	Ctx      context.Context // 添加 ctx 字段
	lastSent time.Time       // 上次发送的时间
	interval time.Duration   // 发送间隔
}

// 格式化字节数，转换为适当的单位（B/KB/MB/GB）
func formatBytes(bytes int64) string {
	if bytes < 1024 {
		return fmt.Sprintf("%d B", bytes)
	} else if bytes < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(bytes)/1024)
	} else if bytes < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(bytes)/(1024*1024))
	} else {
		return fmt.Sprintf("%.2f GB", float64(bytes)/(1024*1024*1024))
	}
}

// Write 实现了 io.Writer 接口
func (p *ProgressWriter) Write(data []byte) (int, error) {
	n := len(data)
	p.Download += int64(n)

	// 计算并格式化输出进度
	percent := float64(p.Download) / float64(p.Total) * 100
	downloaded := formatBytes(p.Download)
	total := formatBytes(p.Total)
	currentTime := time.Now()

	// 检查是否超过设定的发送间隔（例如 0.5 秒）
	if currentTime.Sub(p.lastSent) >= p.interval {
		// 更新发送时间
		p.lastSent = currentTime
		// 打印下载进度
		fmt.Printf("\rDownloading: %.2f%% (%s/%s)", percent, downloaded, total)
		// 发送下载进度通知
		sendMessage(p.Ctx, fmt.Sprintf("下载中: %.2f%% (%s/%s)", percent, downloaded, total), appsourceConsts.WebsocketTypeDownload)
	}

	return n, nil
}

// 获取最终的下载链接及文件大小
func getFinalURL(url string) (string, error) {
	// 使用 http.Head 发送请求来检查最终的 URL
	client := &http.Client{
		// 在这里处理重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 打印重定向过程
			fmt.Println("Redirected to:", req.URL)
			return nil // 继续重定向
		},
	}

	// 发送请求获取最终的 URL
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to follow redirect: %v", err)
	}
	defer resp.Body.Close()

	// 返回最终 URL
	return resp.Request.URL.String(), nil
}

func downloadFileToTemp(ctx context.Context, url string) (*os.File, error) {
	// 获取最终的下载链接
	finalURL, err := getFinalURL(url)
	if err != nil {
		return nil, err
	}

	// 创建临时文件
	tempFile, err := os.CreateTemp("", "ipa_*.zip")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %v", err)
	}
	defer tempFile.Close()

	// 获取最终的响应
	resp, err := http.Get(finalURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()

	// 获取文件大小
	contentLength := resp.ContentLength
	if contentLength == -1 {
		return nil, fmt.Errorf("Content-Length header not found")
	}

	// 使用 ProgressWriter 追踪下载进度
	progress := &ProgressWriter{
		Total:    contentLength,
		Ctx:      ctx,
		interval: 500 * time.Millisecond, // 设置发送进度的间隔为 0.5 秒
	}

	// 通过 ProgressWriter 包装 tempFile，跟踪进度
	_, err = io.Copy(tempFile, io.TeeReader(resp.Body, progress))
	if err != nil {
		return nil, fmt.Errorf("error downloading file: %v", err)
	}

	// 下载完成
	fmt.Printf("\nDownload complete.\n")
	return tempFile, nil
}

func sendMessage(ctx context.Context, message string, event string) {
	userId := systemServer.Context().GetUserId(ctx)
	libWebsocket.SendToUser(gconv.Uint64(userId), &libWebsocket.WResponse{
		Event: event,
		Data: g.Map{
			"type": 1,
			"msg":  message,
		},
	})
}
