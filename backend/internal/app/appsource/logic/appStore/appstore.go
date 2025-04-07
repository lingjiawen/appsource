package appStore

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"mangosmithy/internal/app/appsource/dao"
	"mangosmithy/internal/app/appsource/model"
	"mangosmithy/internal/app/appsource/service"
	commonConsts "mangosmithy/internal/app/common/consts"
	commonEntity "mangosmithy/internal/app/common/model/entity"
	"mangosmithy/library/libUtils"
	"mangosmithy/library/liberr"
	"strconv"
	"strings"
	"time"
)

func init() {
	service.RegisterAppStore(New())
}

func New() service.IAppStore {
	return &sAppStore{}
}

type sAppStore struct{}

func (s *sAppStore) AppStore(ctx context.Context, req *model.AppStoreReq) (listRes *model.AppStoreRes, err error) {
	listRes = new(model.AppStoreRes)
	err = g.Try(ctx, func(ctx context.Context) {
		// 获取配置
		configKeys := g.MapStrStr{
			"appsource.name":         "源名称",
			"appsource.identifier":   "源标识",
			"appsource.purchaseSite": "发卡网",
			"appsource.iconUrl":      "源站图标",
			"appsource.notification": "通知公告",
			"appsource.encryption":   "加密",
		}

		// 获取 configKeys 的 keys（用作数据库查询）
		configKeyList := make([]string, 0, len(configKeys))
		for key := range configKeys {
			configKeyList = append(configKeyList, key)
		}

		var configs []*commonEntity.SysConfig
		err = dao.SysConfig.Ctx(ctx).Where("config_key", configKeyList).Scan(&configs)
		if err != nil {
			return
		}

		// 将 configs 转换为 map
		configMap := make(map[string]string)
		for _, config := range configs {
			configMap[config.ConfigKey] = config.ConfigValue
		}

		// 检查哪些 key 不存在
		for key, value := range configKeys {
			if _, exists := configMap[key]; !exists {
				liberr.ErrIsNil(ctx, fmt.Errorf("%s未配置", value))
			}
		}

		sourceName := configMap["appsource.name"]
		sourceIdentifier := configMap["appsource.identifier"]
		purchaseSite := configMap["appsource.purchaseSite"]
		iconUrl := configMap["appsource.iconUrl"]
		notification := configMap["appsource.notification"]
		encryption := configMap["appsource.encryption"] == "1"

		udid := req.Udid
		code := req.Code

		var redeemCodeRecord *model.AsRedeemCodeInfoRes

		// 没有卡密
		if code == "" {
			// udid搜索
			if udid != "" {
				if !(len(udid) == 25 || len(udid) == 40) {
					liberr.ErrIsNil(ctx, fmt.Errorf("UDID错误"))
				}
				err := dao.AsRedeemCode.Ctx(ctx).Where("udid", udid).Where("active", 1).Scan(&redeemCodeRecord)
				if err != nil {
					return
				}
			}

			validated := isRedeemCodeNotExpired(redeemCodeRecord)
			m := dao.AsApplication.Ctx(ctx).WithAll()

			listCount, err := m.Count()
			liberr.ErrIsNil(ctx, err, "获取总行数失败")

			if listCount == 0 {
				liberr.ErrIsNil(ctx, errors.New("暂无app数据"))
			}

			order := "weigh desc"
			var appResults []*model.AsApplicationListRes
			err = m.Order(order).Scan(&appResults)

			apps := make([]model.App, 0)
			for _, v := range appResults {
				fileUrl := ""
				if validated {
					// 有效
					fileUrl = v.FileUrl
				} else if v.Lock == 0 {
					// 未加密
					fileUrl = v.FileUrl
				}

				uploadPath := strings.Trim(commonConsts.UploadPath, "/")
				if strings.HasPrefix(fileUrl, uploadPath) {
					fileUrl = libUtils.GetDomain(ctx, true) + "/" + fileUrl
				}

				appIconUrl := v.IconUrl
				if strings.HasPrefix(appIconUrl, uploadPath) {
					appIconUrl = libUtils.GetDomain(ctx, true) + "/" + appIconUrl
				}

				apps = append(apps, model.App{
					Name:               v.Name,
					Version:            v.Version,
					Type:               v.Type,
					VersionDate:        v.UpdatedAt.Layout("2006-01-02T15:04:05+08:00"),
					VersionDescription: strings.Replace(v.Description, "\\n", "\n", -1),
					Lock:               strconv.Itoa(v.Lock),
					DownloadURL:        fileUrl,
					IsLanZouCloud:      strconv.Itoa(v.Lanzou),
					IconURL:            appIconUrl,
					TintColor:          "",
					Size:               strconv.Itoa(v.Size),
				})
			}

			listRes = &model.AppStoreRes{
				Name:       sourceName,
				Message:    notification,
				Identifier: sourceIdentifier,
				SourceURL:  libUtils.GetDomain(ctx, true) + "/appstore",
				SourceIcon: iconUrl,
				PayURL:     purchaseSite,
				UnlockURL:  libUtils.GetDomain(ctx, true) + "/appstore",
				Apps:       apps,
			}

			if encryption {
				// 将 listRes 序列化成 JSON 字符串
				data, err := json.Marshal(listRes)
				if err != nil {
					liberr.ErrIsNil(ctx, fmt.Errorf("序列化JSON失败: %v", err))
					return
				}

				// 使用 AES CBC 模式加密数据
				encryptedData, ivBase64, err := encryptAES(data)
				if err != nil {
					liberr.ErrIsNil(ctx, fmt.Errorf("加密失败: %v", err))
					return
				}

				listRes.Identifier = "encrypted"

				// 将加密后的数据替换 listRes
				listRes.Name = encryptedData
				listRes.Message = ivBase64
			}
		} else {
			// 解锁

			// 验证UDID有效性
			if udid == "" || (len(udid) != 25 && len(udid) != 40) {
				liberr.ErrIsNil(ctx, fmt.Errorf("UDID错误"))
			}

			// 搜索解锁码
			err := dao.AsRedeemCode.Ctx(ctx).Where(dao.AsRedeemCode.Columns().Code, code).Scan(&redeemCodeRecord)
			liberr.ErrIsNil(ctx, err)

			if redeemCodeRecord == nil {
				liberr.ErrIsNil(ctx, fmt.Errorf("解锁码不存在"))
			}

			if redeemCodeRecord.Active != 0 {
				liberr.ErrIsNil(ctx, fmt.Errorf("解锁码已使用"))
			}

			// 删除已过期解锁码
			var usedRedeemCode []model.AsRedeemCodeInfoRes
			err = dao.AsRedeemCode.Ctx(ctx).Where("udid", udid).Scan(&usedRedeemCode)
			liberr.ErrIsNil(ctx, err)
			if len(usedRedeemCode) != 0 {
				for _, v := range usedRedeemCode {
					validated := isRedeemCodeNotExpired(&v)
					if !validated {
						_, err = dao.AsRedeemCode.Ctx(ctx).WherePri("id", v.Id).Delete()
						liberr.ErrIsNil(ctx, err)
					}
				}
			}

			// 激活
			_, err = dao.AsRedeemCode.Ctx(ctx).WherePri(redeemCodeRecord.Id).Update(g.Map{
				"udid":      udid,
				"active_at": gtime.Now(),
				"active":    true,
			})

			if err != nil {
				return
			}

			listRes.Identifier = "activeSuccess"
		}

	})
	return
}

func isRedeemCodeNotExpired(redeemCodeRecord *model.AsRedeemCodeInfoRes) bool {
	validated := false // 是否有效
	if redeemCodeRecord != nil {
		activeTime := redeemCodeRecord.ActiveAt
		if activeTime != nil {
			// 已激活
			// 获取当前时间
			currentTime := gtime.Now()
			var validPeriod int // 记录有效期天数
			switch redeemCodeRecord.Type {
			case 1:
				validPeriod = 30 // 30天
			case 2:
				validPeriod = 90 // 90天
			case 3:
				validPeriod = 365 // 365天
			case 4:
				validPeriod = 1 // 1天
			case 5:
				validPeriod = 7 // 7天
			default:
				// 如果类型未知，可以设置为无效
				validPeriod = 0
			}

			// 判断当前时间是否在有效期内
			if validPeriod > 0 {
				expirationTime := activeTime.Add(time.Duration(validPeriod) * time.Hour * 24) // 增加有效期天数
				if currentTime.Before(expirationTime) {
					validated = true
				}
			}
		}
	}
	return validated
}

func encryptAES(data []byte) (string, string, error) {
	var key = []byte("VUjvK4iEnUMk1niSAAxVN5o92oz44S7s") // AES密钥（32字节用于 AES-256）
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	// 生成随机的初始化向量（IV）
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		return "", "", err
	}

	// 补充数据到Block大小的整数倍
	padding := aes.BlockSize - len(data)%aes.BlockSize
	paddedData := append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)

	// CBC模式加密
	mode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(paddedData))
	mode.CryptBlocks(encrypted, paddedData)

	// 返回加密后的数据和IV的Base64编码
	return base64.StdEncoding.EncodeToString(encrypted), base64.StdEncoding.EncodeToString(iv), nil
}
