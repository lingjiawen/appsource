package public

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	commonDao "mangosmithy/internal/app/common/dao"
	"mangosmithy/internal/app/common/model/entity"
	commonService "mangosmithy/internal/app/common/service"
	"mangosmithy/internal/app/install/common"
	"mangosmithy/internal/app/install/dao"
	"mangosmithy/internal/app/install/logic/zsign"
	"mangosmithy/internal/app/install/model"
	"mangosmithy/internal/app/install/service"
	"mangosmithy/internal/app/install/tools"
	"mangosmithy/library/libUtils"
	"mangosmithy/library/liberr"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"
)

func init() {
	service.RegisterPublic(New())
}

func New() service.IPublic {
	return &sPublic{}
}

type sPublic struct{}

// RequestPost 发起Post请求
func RequestPost(url string, postData g.Map) ([]byte, error) {
	// 将请求数据编码为 JSON
	jsonData, err := json.Marshal(postData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal postData: %w", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 设置请求头信息
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")

	// 发起请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// 检查 Content-Encoding 是否是 gzip
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	// 读取响应体
	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

func (s *sPublic) Install(ctx context.Context, req *model.ApplicationInstallReq) (res *model.ApplicationInstallRes, err error) {
	res = new(model.ApplicationInstallRes)
	err = g.Try(ctx, func(ctx context.Context) {
		// UDID必填, 判断UDID是25位或者40位
		if len(req.UDID) != 25 && len(req.UDID) != 40 {
			liberr.ErrIsNil(ctx, errors.New("UDID格式错误"))
		}

		p12 := ""
		mobileProvision := ""

		var device *model.SignDeviceInfoRes

		// 判断是否传了卡密
		if req.Code == "" {
			// 没有卡密, 查询历史设备
			err := dao.SignDevice.Ctx(ctx).
				Where("udid", req.UDID).
				Where("active", 1).
				Where("status", "normal").
				Order("id desc").
				Scan(&device)
			liberr.ErrIsNil(ctx, err, "查询设备失败")
			// 没有设备返回错误
			if device == nil {
				// 查询接口是否有历史设备
				pool := grpool.New(3)
				var mu sync.Mutex     // 保护共享数据的互斥锁
				var wg sync.WaitGroup // 用于等待所有任务完成

				var results []*model.GetDeviceData
				var resultErrors []error

				enabledPlatforms, err := getEnabledPlatforms(ctx)
				if err != nil {
					liberr.ErrIsNil(ctx, err, "获取平台信息失败")
				}
				for _, enabledPlatform := range enabledPlatforms {
					wg.Add(1)
					pool.Add(ctx, func(ctx context.Context) {
						defer wg.Done()
						// 获取平台信息
						apiUrl := gconv.String(enabledPlatform["url"])
						platformToken := gconv.String(enabledPlatform["token"])
						platformID := gconv.Int(enabledPlatform["platform"])

						requestData := g.Map{
							"udid":  req.UDID,
							"token": platformToken,
						}

						// 调用接口
						response, err := RequestPost(apiUrl, requestData)
						if err != nil {
							mu.Lock()
							resultErrors = append(resultErrors, fmt.Errorf("请求接口失败: %s", err.Error()))
							mu.Unlock()
							return
						}

						// 校验 JSON
						if !json.Valid(response) {
							mu.Lock()
							resultErrors = append(resultErrors, fmt.Errorf("无效的 JSON 响应: %s", string(response)))
							mu.Unlock()
							return
						}

						// 解析 JSON
						var result *model.GetDeviceDataResp
						err = json.Unmarshal(response, &result)
						if err != nil {
							mu.Lock()
							resultErrors = append(resultErrors, fmt.Errorf("解析 JSON 失败: %s", err.Error()))
							mu.Unlock()
							return
						}

						// 判断是否成功
						resultCode := gconv.Int(result.Code)
						if resultCode != 0 {
							resultErrors = append(resultErrors, fmt.Errorf("平台获取设备失败: %s", result.Msg))
							return
						}

						if result.Data.Status == "normal" && result.Data.Mobileprovision != "" {
							mu.Lock()
							result.Data.PlatformId = platformID
							results = append(results, result.Data)
							mu.Unlock()
						}
					})
				}

				wg.Wait()

				// 打印请求错误
				if len(resultErrors) > 0 {
					for _, err := range resultErrors {
						g.Log().Error(ctx, err.Error())
					}
				}

				// 打印请求结果
				g.Log().Debug(ctx, "请求结果: ", results)

				// 检查是否有可用结果
				if len(results) == 0 {
					liberr.ErrIsNil(ctx, errors.New("该UDID暂无可用证书, 请填写卡密"))
				} else {
					// 遍历结果, 选择第一个可用的
					for _, result := range results {
						if result.Status == "normal" && result.Mobileprovision != "" {
							device = &model.SignDeviceInfoRes{
								Udid:            req.UDID,
								ApiPlatform:     result.PlatformId,
								Name:            gconv.String(result.Name),
								CertId:          gconv.String(result.Id),
								P12:             gconv.String(result.P12),
								Mobileprovision: gconv.String(result.Mobileprovision),
								AddTime:         gconv.Uint64(result.AddTime),
								P12Password:     gconv.String(result.P12Password),
								Model:           gconv.String(result.Model),
								ExpireTime:      gconv.Uint64(result.ExpireTime),
								SerialNumber:    "", // 通过证书检测
								AccountType:     gconv.Int(result.Type),
								WarrantyType:    result.Warranty,
								DeviceType:      "iphone",
								Status:          result.Status,
								Active:          1,
								Pool:            gconv.Int(result.Pool),
							}
							// 插入数据库
							_, err = dao.SignDevice.Ctx(ctx).Insert(device)
							break
						}
					}

					if device == nil {
						liberr.ErrIsNil(ctx, errors.New("该UDID暂无可用证书, 请填写卡密"))
					}
				}
			}

			// 获取最新证书
			p12 = device.P12
			mobileProvision = device.Mobileprovision
		} else {
			// 有卡密, 查询卡密
			var redeemCode *model.SignRedeemCodeInfoRes
			err := dao.SignRedeemCode.Ctx(ctx).
				Where("code", req.Code).
				Scan(&redeemCode)
			liberr.ErrIsNil(ctx, err, "查询卡密失败")
			// 判断卡密是否存在
			if redeemCode == nil {
				liberr.ErrIsNil(ctx, errors.New("卡密不存在"))
			}

			// 判断卡密是否停用
			if redeemCode.Banned == 1 {
				liberr.ErrIsNil(ctx, errors.New("卡密已停用"))
			}

			// 判断卡密是否已使用
			if redeemCode.Active == 1 || redeemCode.Udid != "" {
				if redeemCode.Udid != req.UDID {
					liberr.ErrIsNil(ctx, errors.New("卡密已被使用"))
				}
				// 有可能用户误填了卡密, 实际已经绑定了UDID
				// 查询设备
				err := dao.SignDevice.Ctx(ctx).
					Where("udid", req.UDID).
					Where("active", 1).
					Where("status", "normal").
					Order("id desc").
					Scan(&device)
				if err != nil {
					liberr.ErrIsNil(ctx, err, fmt.Sprintf("查询设备失败: %s", err.Error()))
				}
				// 没有设备返回错误
				if device == nil {
					// 本地没有设备, 但是卡密已经使用
					// 这种适合手动填写, 指定udid可用
					apiPlatform := redeemCode.ApiPlatform
					// 获取平台信息
					apiUrl, platformToken := getPlatformUrlAndToken(ctx, apiPlatform, "/api/v1/public/device/add")
					accountType := redeemCode.AccountType
					deviceType := redeemCode.DeviceType
					pool := redeemCode.Pool
					apiWarrantyType := redeemCode.ApiWarrantyType
					warrantyType := redeemCode.WarrantyType
					force := redeemCode.Force
					note := redeemCode.Note

					requestData := g.Map{
						"udid":       req.UDID,
						"pool":       pool,
						"warranty":   apiWarrantyType,
						"deviceType": deviceType,
						"force":      force,
						"note":       note,
						"isBook":     accountType,
						"token":      platformToken,
					}

					// 调用接口
					response, err := RequestPost(apiUrl, requestData)
					if err != nil {
						liberr.ErrIsNil(ctx, err, fmt.Sprintf("请求接口失败: %s", err.Error()))
					}

					// 校验 JSON
					if !json.Valid(response) {
						liberr.ErrIsNil(ctx, fmt.Errorf("无效的 JSON 响应: %s", string(response)))
					}

					// 解析 JSON
					var result g.Map
					err = json.Unmarshal(response, &result)
					if err != nil {
						liberr.ErrIsNil(ctx, err, fmt.Sprintf("解析 JSON 失败: %s", err.Error()))
					}

					// 判断是否成功
					resultCode := gconv.Int(result["code"])
					if resultCode != 0 && resultCode != 1 {
						liberr.ErrIsNil(ctx, errors.New(result["message"].(string)))
					}

					// 保存设备信息
					resultData := result["data"].(g.Map)

					// 开启事务
					err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
						// 保存设备
						device = &model.SignDeviceInfoRes{
							Udid:   req.UDID,
							Name:   gconv.String(resultData["name"]),
							CertId: gconv.String(resultData["id"]),
							// DeviceId:        "",
							P12:             gconv.String(resultData["p12"]),
							Mobileprovision: gconv.String(resultData["mobileprovision"]),
							AddTime:         gconv.Uint64(resultData["add_time"]),
							P12Password:     gconv.String(resultData["p12_password"]),
							Model:           gconv.String(resultData["model"]),
							ExpireTime:      gconv.Uint64(resultData["expire_time"]),
							SerialNumber:    "", // 通过证书检测
							RedeemCode:      req.Code,
							AccountType:     accountType,
							WarrantyType:    warrantyType,
							DeviceType:      deviceType,
							Status:          "normal",
							Pool:            gconv.Int(resultData["pool"]),
							ApiPlatform:     apiPlatform,
							ApiWarrantyType: apiWarrantyType,
							Active:          1,
						}
						_, err = dao.SignDevice.Ctx(ctx).TX(tx).Insert(&device)
						if err != nil {
							liberr.ErrIsNil(ctx, err, fmt.Sprintf("保存设备失败: %s", err.Error()))
						}

						// 修改激活码
						_, err = dao.SignRedeemCode.Ctx(ctx).TX(tx).
							WherePri(redeemCode.Id).
							Update(g.Map{
								"udid":      req.UDID,
								"cert_id":   resultData["id"].(string),
								"active":    1,
								"active_at": gtime.Now(),
							})
						if err != nil {
							liberr.ErrIsNil(ctx, err, fmt.Sprintf("激活失败: %s", err.Error()))
						}
						return nil
					})
					liberr.ErrIsNil(ctx, err, fmt.Sprintf("事务失败: %s", err.Error()))

					p12 = resultData["p12"].(string)
					mobileProvision = resultData["mobileprovision"].(string)
				} else {
					p12 = device.P12
					mobileProvision = device.Mobileprovision
				}
			} else {
				// 新激活证书
				apiPlatform := redeemCode.ApiPlatform

				// 获取平台信息
				apiUrl, platformToken := getPlatformUrlAndToken(ctx, apiPlatform, "/api/v1/public/device/add")

				accountType := redeemCode.AccountType
				deviceType := redeemCode.DeviceType
				pool := redeemCode.Pool
				apiWarrantyType := redeemCode.ApiWarrantyType
				warrantyType := redeemCode.WarrantyType
				note := redeemCode.Note
				force := redeemCode.Force

				requestData := g.Map{
					"udid":       req.UDID,
					"pool":       pool,
					"warranty":   apiWarrantyType,
					"deviceType": deviceType,
					"force":      force,
					"note":       note,
					"isBook":     accountType,
					"token":      platformToken,
				}

				// 调用接口
				response, err := RequestPost(apiUrl, requestData)
				if err != nil {
					liberr.ErrIsNil(ctx, err, fmt.Sprintf("请求接口失败: %s", err.Error()))
				}

				// 校验 JSON
				if !json.Valid(response) {
					liberr.ErrIsNil(ctx, fmt.Errorf("无效的 JSON 响应: %s", string(response)))
				}

				// 解析 JSON
				var result g.Map
				err = json.Unmarshal(response, &result)
				if err != nil {
					liberr.ErrIsNil(ctx, err, fmt.Sprintf("解析 JSON 失败: %s", err.Error()))
				}

				// 判断是否成功
				resultCode := gconv.Int(result["code"])
				if resultCode != 0 && resultCode != 1 {
					liberr.ErrIsNil(ctx, errors.New(result["message"].(string)))
				}

				// 保存设备信息
				resultData := result["data"].(g.Map)

				// 开启事务
				err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
					// 保存设备
					device = &model.SignDeviceInfoRes{
						Udid:   req.UDID,
						Name:   gconv.String(resultData["name"]),
						CertId: gconv.String(resultData["id"]),
						// DeviceId:        "",
						P12:             gconv.String(resultData["p12"]),
						Mobileprovision: gconv.String(resultData["mobileprovision"]),
						AddTime:         gconv.Uint64(resultData["add_time"]),
						P12Password:     gconv.String(resultData["p12_password"]),
						Model:           gconv.String(resultData["model"]),
						ExpireTime:      gconv.Uint64(resultData["expire_time"]),
						SerialNumber:    "", // 通过证书检测
						RedeemCode:      req.Code,
						AccountType:     accountType,
						WarrantyType:    warrantyType,
						DeviceType:      deviceType,
						Status:          "normal",
						Pool:            gconv.Int(resultData["pool"]),
						ApiPlatform:     apiPlatform,
						ApiWarrantyType: apiWarrantyType,
						Active:          1,
					}
					_, err = dao.SignDevice.Ctx(ctx).TX(tx).Insert(&device)
					if err != nil {
						liberr.ErrIsNil(ctx, err, fmt.Sprintf("保存设备失败: %s", err.Error()))
					}

					// 修改激活码
					_, err = dao.SignRedeemCode.Ctx(ctx).TX(tx).
						WherePri(redeemCode.Id).
						Update(g.Map{
							"udid":      req.UDID,
							"cert_id":   resultData["id"].(string),
							"active":    1,
							"active_at": gtime.Now(),
						})
					if err != nil {
						liberr.ErrIsNil(ctx, err, fmt.Sprintf("激活失败: %s", err.Error()))
					}
					return nil
				})
				if err != nil {
					liberr.ErrIsNil(ctx, err, fmt.Sprintf("事务失败: %s", err.Error()))
				}

				p12 = resultData["p12"].(string)
				mobileProvision = resultData["mobileprovision"].(string)
			}
		}

		// 判断是否预约
		if mobileProvision == "" {

			if device == nil {
				liberr.ErrIsNil(ctx, errors.New("设备信息不存在1"))
			}

			apiPlatform := device.ApiPlatform
			// 获取平台信息
			apiUrl, platformToken := getPlatformUrlAndToken(ctx, apiPlatform, "/api/v1/public/device/get")
			requestData := g.Map{
				"cert_id": device.CertId,
				"token":   platformToken,
			}

			// 调用接口
			response, err := RequestPost(apiUrl, requestData)
			if err != nil {
				liberr.ErrIsNil(ctx, err, fmt.Sprintf("请求接口失败: %s", err.Error()))
			}

			// 校验 JSON
			if !json.Valid(response) {
				liberr.ErrIsNil(ctx, fmt.Errorf("无效的 JSON 响应: %s", string(response)))
			}

			// 解析 JSON
			var result *model.GetDeviceDataResp
			err = json.Unmarshal(response, &result)
			if err != nil {
				liberr.ErrIsNil(ctx, err, fmt.Sprintf("解析 JSON 失败: %s", err.Error()))
			}

			// 判断是否成功
			resultCode := gconv.Int(result.Code)
			if resultCode != 0 {
				liberr.ErrIsNil(ctx, errors.New(result.Msg))
			}

			if result.Data.Mobileprovision == "" {
				liberr.ErrIsNil(ctx, errors.New("设备审核中"))
			} else {
				// 更新设备信息
				_, err = dao.SignDevice.Ctx(ctx).
					WherePri(device.Id).
					Update(g.Map{
						"mobileprovision": result.Data.Mobileprovision,
					})

				if err != nil {
					liberr.ErrIsNil(ctx, err, fmt.Sprintf("更新设备信息失败: %s", err.Error()))
				}
			}
			mobileProvision = result.Data.Mobileprovision
		}

		plistUrl, err := signIpa(ctx, req.UDID, p12, mobileProvision)
		if err != nil {
			liberr.ErrIsNil(ctx, err, fmt.Sprintf("签名失败: %s", err.Error()))
		}
		// 返回安装地址
		res.Url = fmt.Sprintf("itms-services://?action=download-manifest&url=%s", plistUrl)
	})
	return
}

func (s *sPublic) PrivateInstall(ctx context.Context, req *model.ApplicationPrivateInstallReq) (res *model.ApplicationPrivateInstallRes, err error) {
	res = new(model.ApplicationPrivateInstallRes)
	err = g.Try(ctx, func(ctx context.Context) {
		// 判断文件是否存在
		if req.P12 == "" || req.MobileProvision == "" {
			liberr.ErrIsNil(ctx, errors.New("p12或mobileprovision不能为空"))
		}

		// 检测证书
		checkInfo, err := tools.CheckCert(req.P12, req.MobileProvision, req.Password)
		if err != nil {
			liberr.ErrIsNil(ctx, fmt.Errorf("证书检测出错: %v", err))
		}

		if !checkInfo.CertMatch {
			liberr.ErrIsNil(ctx, errors.New("证书与描述文件不匹配"))
		}

		if checkInfo.Status == "hidden" {
			liberr.ErrIsNil(ctx, errors.New("证书已掉签"))
		}

		if checkInfo.Status == "banned" {
			liberr.ErrIsNil(ctx, errors.New("证书已被封禁"))
		}

		uuid := checkInfo.ProvisionUUID

		plistUrl, err := signIpa(ctx, uuid, req.P12, req.MobileProvision)
		if err != nil {
			liberr.ErrIsNil(ctx, fmt.Errorf("签名失败: %s", err.Error()))
		}

		// 返回安装地址
		res.Url = fmt.Sprintf("itms-services://?action=download-manifest&url=%s", plistUrl)

	})
	return
}

func (s *sPublic) GetPlist(ctx context.Context, req *model.GetPlistReq) (res *model.GetPlistRes, err error) {
	res = &model.GetPlistRes{}
	// 获取 plist 信息
	cacheKey := req.UUID
	plistInfoJSON, err := g.Redis().Do(ctx, "GET", cacheKey)
	if err != nil {
		return nil, fmt.Errorf("获取 Redis 缓存失败: %v", err)
	}

	// 没有数据
	if plistInfoJSON == nil || plistInfoJSON.IsNil() {
		return nil, errors.New("安装信息已过有效期")
	}

	// 解析 plist 信息
	var plistInfo model.InstallPlist
	err = json.Unmarshal(gconv.Bytes(plistInfoJSON), &plistInfo)
	if err != nil {
		return nil, fmt.Errorf("解析 plist 信息失败: %v", err)
	}

	// 获取安装链接
	res.PlistContent = common.MakeInstallPlist(plistInfo)
	return
}

func (s *sPublic) CheckDevices(ctx context.Context, req *model.CheckDevicesReq) (res *model.CheckDevicesRes, err error) {
	res = &model.CheckDevicesRes{}
	err = g.Try(ctx, func(ctx context.Context) {

		var devices []*model.SignDeviceInfoRes

		if req.SearchType == "code" {
			// 有卡密, 查询卡密
			code := req.SearchValue
			err := dao.SignDevice.Ctx(ctx).
				Where("redeem_code", code).
				Where("active", 1).
				Order("id desc").
				Scan(&devices)
			liberr.ErrIsNil(ctx, err, "查询设备失败")
		} else if req.SearchType == "udid" {
			udid := req.SearchValue
			// UDID必填, 判断UDID是25位或者40位
			if len(udid) != 25 && len(udid) != 40 {
				liberr.ErrIsNil(ctx, errors.New("UDID格式错误"))
			}

			// 查询设备
			err := dao.SignDevice.Ctx(ctx).
				Where("udid", udid).
				Where("active", 1).
				//Where("status", "normal").
				Order("id desc").
				Scan(&devices)
			liberr.ErrIsNil(ctx, err, "查询设备失败")
		} else if req.SearchType == "cert_id" {
			certId := req.SearchValue
			err := dao.SignDevice.Ctx(ctx).
				Where("cert_id", certId).
				Where("active", 1).
				//Where("status", "normal").
				Order("id desc").
				Scan(&devices)
			liberr.ErrIsNil(ctx, err, "查询设备失败")
		} else {
			liberr.ErrIsNil(ctx, errors.New("搜索类型错误"))
		}

		// 没有设备
		if len(devices) == 0 {
			liberr.ErrIsNil(ctx, errors.New("未找到可用设备"))
		}

		var accountTypeData []*entity.SysDictData
		err = commonDao.SysDictData.Ctx(ctx).Where(commonDao.SysDictData.Columns().DictType, "apple_account_type").Scan(&accountTypeData)
		// 获取所有DictLabel: DictValue的Map
		accountTypeMap := g.MapStrStr{}
		for _, v := range accountTypeData {
			accountTypeMap[v.DictValue] = v.DictLabel
		}

		var certStatusData []*entity.SysDictData
		err = commonDao.SysDictData.Ctx(ctx).Where(commonDao.SysDictData.Columns().DictType, "cert_status").Scan(&certStatusData)
		// 获取所有DictLabel: DictValue的Map
		certStatusMap := g.MapStrStr{}
		for _, v := range certStatusData {
			certStatusMap[v.DictValue] = v.DictLabel
		}

		var devicesInfo []*model.DeviceInfo
		for _, device := range devices {
			expireTimeParsed := gtime.NewFromTimeStamp(int64(device.ExpireTime))
			expireTime := expireTimeParsed.Format("Y-m-d H:i:s")

			addTimeParsed := gtime.NewFromTimeStamp(int64(device.AddTime))
			addTime := addTimeParsed.Format("Y-m-d H:i:s")
			devicesInfo = append(devicesInfo, &model.DeviceInfo{
				Name:        device.Name,
				Udid:        device.Udid,
				CertId:      device.CertId,
				RedeemCode:  device.RedeemCode,
				AddTime:     addTime,
				ExpireTime:  expireTime,
				Status:      certStatusMap[device.Status],
				AccountType: accountTypeMap[strconv.Itoa(device.AccountType)],
			})

			// 判断是否审核中
			if device.Mobileprovision == "" {
				// 查询接口是否审核通过
				apiPlatform := device.ApiPlatform
				// 获取平台信息
				apiUrl, platformToken := getPlatformUrlAndToken(ctx, apiPlatform, "/api/v1/public/device/get")
				requestData := g.Map{
					"cert_id": device.CertId,
					"token":   platformToken,
				}

				// 调用接口
				response, err := RequestPost(apiUrl, requestData)
				if err != nil {
					liberr.ErrIsNil(ctx, err, fmt.Sprintf("请求接口失败: %s", err.Error()))
				}

				// 校验 JSON
				if !json.Valid(response) {
					liberr.ErrIsNil(ctx, fmt.Errorf("无效的 JSON 响应: %s", string(response)))
				}

				// 解析 JSON
				var result *model.GetDeviceDataResp
				err = json.Unmarshal(response, &result)
				if err != nil {
					liberr.ErrIsNil(ctx, err, fmt.Sprintf("解析 JSON 失败: %s", err.Error()))
				}

				// 判断是否成功
				resultCode := gconv.Int(result.Code)
				if resultCode != 0 {
					liberr.ErrIsNil(ctx, errors.New(result.Msg))
				}

				if result.Data.Mobileprovision == "" {
					devicesInfo[len(devicesInfo)-1].Status = result.Msg
				} else {
					// 更新设备信息
					_, err = dao.SignDevice.Ctx(ctx).
						WherePri(device.Id).
						Update(g.Map{
							"mobileprovision": result.Data.Mobileprovision,
						})

					if err != nil {
						liberr.ErrIsNil(ctx, err, fmt.Sprintf("更新设备信息失败: %s", err.Error()))
					}
				}
			}
		}

		// 有设备
		res.Devices = devicesInfo
	})
	return
}

func (s *sPublic) GetHelp(ctx context.Context) (listRes *model.GetHelpRes, err error) {
	listRes = new(model.GetHelpRes)
	err = g.Try(ctx, func(ctx context.Context) {
		order := "weigh desc"
		var res []*model.SignHelpListRes
		err = dao.SignHelp.Ctx(ctx).WithAll().Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.SignHelpListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.SignHelpListRes{
				Id:       v.Id,
				Title:    v.Title,
				Content:  v.Content,
				IsExpand: v.IsExpand,
				Weigh:    v.Weigh,
			}
		}
	})
	return
}

func (s *sPublic) GetAbout(ctx context.Context) (listRes *model.GetAboutRes, err error) {
	listRes = new(model.GetAboutRes)
	err = g.Try(ctx, func(ctx context.Context) {
		order := "weigh desc"
		// 1. 先查询全部数据（不分组）
		var allItems []*model.SignAboutListRes
		err = dao.SignAbout.Ctx(ctx).WithAll().Order(order).Scan(&allItems)

		// 2. 手动按 group 字段分类
		groupMap := make(map[string][]*model.SignAboutListRes)
		for _, item := range allItems {
			groupMap[strconv.Itoa(item.Group)] = append(groupMap[strconv.Itoa(item.Group)], item)
		}

		// 3. 转换为二维切片
		listRes.List = make([][]*model.SignAboutListRes, 0, len(groupMap))
		for _, groupItems := range groupMap {
			listRes.List = append(listRes.List, groupItems)
		}
	})
	return
}

func savePlistInfo(ctx context.Context, plistInfo model.InstallPlist) (uuid *string, err error) {
	cacheKey := gconv.String(gtime.TimestampMilli())
	// 将 plistInfo 转换为 JSON 字符串
	plistInfoJSON, err := json.Marshal(plistInfo)
	if err != nil {
		return nil, fmt.Errorf("转换 plistInfo 为 JSON 字符串失败: %v", err)
	}

	// 保存到 Redis
	_, err = g.Redis().Do(ctx, "SET", cacheKey, plistInfoJSON, "EX", 60*60*2)
	if err != nil {
		return nil, fmt.Errorf("设置 Redis 缓存失败: %v", err)
	}

	return &cacheKey, nil
}

func getEnabledPlatforms(ctx context.Context) (platforms []g.Map, err error) {
	var platformsInfo []*model.SignPlatformInfoRes
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SignPlatform.Ctx(ctx).
			Where("status", 1).
			Order("weigh desc").
			Scan(&platformsInfo)
		liberr.ErrIsNil(ctx, err, "查询平台失败")

		if len(platformsInfo) == 0 {
			liberr.ErrIsNil(ctx, errors.New("没有可用的平台"))
		}

		// 获取平台信息
		platforms = []g.Map{}
		for _, platform := range platformsInfo {
			platformUrl, token := getPlatformUrlAndToken(ctx, int(platform.Id), "/api/v1/public/device/get")
			platforms = append(platforms, g.Map{
				"platform": platform.Id,
				"url":      platformUrl,
				"token":    token,
			})
		}
	})
	return
}

func getPlatformUrlAndToken(ctx context.Context, apiPlatform int, apiPath string) (string, string) {
	var platform *model.SignPlatformInfoRes
	err := dao.SignPlatform.Ctx(ctx).
		WherePri(apiPlatform).
		Scan(&platform)
	liberr.ErrIsNil(ctx, err, "查询平台失败")

	// 判断平台是否存在
	if platform == nil {
		liberr.ErrIsNil(ctx, errors.New("出书平台配置错误, 请联系管理员"))
	}

	// 判断平台是否启用
	if platform.Status == 0 {
		liberr.ErrIsNil(ctx, errors.New("出书平台未启用"))
	}

	// 解析基础 URL
	baseURL := platform.BaseUrl
	if !gstr.HasPrefix(baseURL, "http://") && !gstr.HasPrefix(baseURL, "https://") {
		// 如果用户没有填写协议，根据 OpenSsl 设置添加协议
		if gconv.Bool(platform.OpenSsl) {
			baseURL = "https://" + baseURL
		} else {
			baseURL = "http://" + baseURL
		}
	}
	u, err := url.Parse(baseURL)
	liberr.ErrIsNil(ctx, err, "解析 URL 失败")
	// 拼接路径
	u.Path = path.Join(u.Path, apiPath)
	// 最终得到的 URL
	apiUrl := u.String()
	return apiUrl, platform.Token
}

func signIpa(ctx context.Context, udid, base64p12, base64mp string) (plistUrl string, err error) {
	plistUrl = ""
	err = g.Try(ctx, func(ctx context.Context) {

		// 先判断配置信息
		appNameConfig, err := commonService.SysConfig().GetConfigByKey(ctx, "install.download.name")
		liberr.ErrIsNil(ctx, err)
		appIconConfig, err := commonService.SysConfig().GetConfigByKey(ctx, "install.download.icon")
		liberr.ErrIsNil(ctx, err)
		appBundleIDConfig, err := commonService.SysConfig().GetConfigByKey(ctx, "install.download.bundleID")
		liberr.ErrIsNil(ctx, err)
		appVersionConfig, err := commonService.SysConfig().GetConfigByKey(ctx, "install.download.version")
		liberr.ErrIsNil(ctx, err)
		appPathConfig, err := commonService.SysConfig().GetConfigByKey(ctx, "install.download.url")
		liberr.ErrIsNil(ctx, err)

		// 获取配置信息
		appName := appNameConfig.ConfigValue
		appIcon := appIconConfig.ConfigValue
		appBundleID := appBundleIDConfig.ConfigValue
		appVersion := appVersionConfig.ConfigValue
		appPath := appPathConfig.ConfigValue
		if appName == "" || appIcon == "" || appBundleID == "" || appVersion == "" || appPath == "" {
			liberr.ErrIsNil(ctx, errors.New("请先配置安装信息"))
		}

		// 开始签名
		publicPathConfig, _ := g.Cfg().Get(ctx, "server.serverRoot")
		publicPath := publicPathConfig.String()
		ipaPath := filepath.Join(publicPath, appPath)
		signedIpaPath := filepath.Join(publicPath, "signed")
		if err := gfile.Mkdir(signedIpaPath); err != nil {
			liberr.ErrIsNil(ctx, err, "创建签名文件夹失败")
		}

		timeNowUnix := gtime.Now().Unix()

		fmt.Printf("开始签名: UDID为 %s, 对应时间戳为 %d\n", udid, timeNowUnix)
		signedPath := filepath.Join(signedIpaPath, udid)
		if err := gfile.Mkdir(signedPath); err != nil {
			liberr.ErrIsNil(ctx, err, "创建签名UDID文件夹失败")
		}

		outputPath := filepath.Join(signedPath, fmt.Sprintf("%d.ipa", timeNowUnix))
		p12Path := filepath.Join(signedPath, fmt.Sprintf("%s.p12", udid))
		mobileProvisionPath := filepath.Join(signedPath, fmt.Sprintf("%s.mobileprovision", udid))

		// 解码p12
		p12Decoded, err := base64.StdEncoding.DecodeString(base64p12)
		if err != nil {
			liberr.ErrIsNil(ctx, err, "解码p12失败")
		}

		// 解码mobileProvision
		mobileProvisionDecoded, err := base64.StdEncoding.DecodeString(base64mp)
		if err != nil {
			liberr.ErrIsNil(ctx, err, "解码mobileprovision失败")
		}

		// 保存解码后的p12文件
		if err := gfile.PutBytes(p12Path, p12Decoded); err != nil {
			liberr.ErrIsNil(ctx, err, "保存p12失败")
		}

		// 保存解码后的mobileProvision文件
		if err := gfile.PutBytes(mobileProvisionPath, mobileProvisionDecoded); err != nil {
			liberr.ErrIsNil(ctx, err, "保存mobileprovision失败")
		}

		//获取当前运行时目录
		curDir, err := os.Getwd()
		liberr.ErrIsNil(ctx, err, "获取本地路径失败")

		signRet, err := zsign.Sign(ctx, zsign.SignParam{
			IpaPath:       filepath.Join(curDir, ipaPath),
			P12Path:       filepath.Join(curDir, p12Path),
			ProvisionPath: filepath.Join(curDir, mobileProvisionPath),
			OutputPath:    filepath.Join(curDir, outputPath),
			CertPassword:  "1",
			Dylibs:        nil,
		})
		if err != nil {
			liberr.ErrIsNil(ctx, err, fmt.Sprintf("签名失败: %s", err.Error()))
		}

		if signRet["code"] != 0 {
			// 签名出错
			liberr.ErrIsNil(ctx, errors.New(signRet["message"].(string)))
		}

		// 保存 plist 信息

		// 去除outputPath的publicPath部分
		relativeAppPath := gstr.SubStr(outputPath, len(publicPath))
		urlPrefix := libUtils.GetDomain(ctx, false)
		appFileUrl, _ := url.JoinPath(urlPrefix, relativeAppPath)
		cacheKey, err := savePlistInfo(ctx, model.InstallPlist{
			Name:       appName,
			BundleID:   appBundleID,
			Version:    appVersion,
			IconUrl:    appIcon,
			AppFileUrl: appFileUrl,
		})

		// 拼接获取plist的URL
		plistUrl = fmt.Sprintf("%s/getInstallPlist/%s", urlPrefix, *cacheKey)
	})
	return
}
