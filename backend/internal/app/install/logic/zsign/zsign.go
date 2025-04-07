package zsign

/*
#cgo LDFLAGS: -L/usr/local/lib -lzsign
#cgo CFLAGS: -I/usr/local/include/zsign

#include "zsign.h"
#include <stdlib.h>
*/
import "C"
import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"mangosmithy/library/liberr"
	"os"
	"path/filepath"
	"unsafe"
)

// SignParam 签名参数
type SignParam struct {
	IpaPath          string   // IPA路径
	P12Path          string   // 证书路径
	ProvisionPath    string   // 描述文件路径
	OutputPath       string   // 输出路径
	CertPassword     string   // 证书密码
	BundleID         string   // Bundle ID
	BundleName       string   // Bundle名称
	BundleVersion    string   // 修改版本号
	EntitlementsPath string   // Entitlements文件路径
	TempFolder       string   // 临时文件夹路径
	ZipLevel         int      // 压缩级别
	Dylibs           []string // 动态库
}

func Sign(ctx context.Context, params SignParam) (g.Map, error) {
	var signResult C.int
	err := g.Try(ctx, func(ctx context.Context) {
		// 验证必需参数
		if params.IpaPath == "" || params.P12Path == "" || params.ProvisionPath == "" || params.CertPassword == "" || params.OutputPath == "" {
			liberr.ErrIsNil(ctx, errors.New("缺乏必要参数"))
		}

		// 默认压缩级别
		if params.ZipLevel <= 0 || params.ZipLevel > 9 {
			params.ZipLevel = 1
		}

		publicPathConfig, _ := g.Cfg().Get(ctx, "server.serverRoot")
		publicPath := publicPathConfig.String()
		// 默认临时文件路径
		if params.TempFolder == "" {
			tempDir := filepath.Join(publicPath, "signTemp")
			if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
				liberr.ErrIsNil(ctx, err, "创建临时文件夹失败")
			}
			params.TempFolder = tempDir
		}

		// 转换dylib参数
		var dylibPtrs []*C.char
		if params.Dylibs != nil {
			for _, dylibPath := range params.Dylibs {
				dylibPtrs = append(dylibPtrs, C.CString(dylibPath))
			}
		}

		// 构建C结构体
		opts := C.ZSignOptions{
			ipa_path:       stringToCChar(params.IpaPath),
			p12_path:       stringToCChar(params.P12Path),
			provision_path: stringToCChar(params.ProvisionPath),
			password:       stringToCChar(params.CertPassword),
			output_path:    stringToCChar(params.OutputPath),

			bundle_id:      stringToCChar(params.BundleID),
			bundle_name:    stringToCChar(params.BundleName),
			bundle_version: stringToCChar(params.BundleVersion),
			entitlements:   stringToCChar(params.EntitlementsPath),
			temp_folder:    stringToCChar(params.TempFolder),
			zip_level:      C.int(params.ZipLevel),

			//force:       boolToCInt(true),
			dylib_paths: nil,

			//debug: boolToCInt(true),
		}

		// 设置dylib路径
		if len(dylibPtrs) > 0 {
			opts.dylib_paths = (**C.char)(unsafe.Pointer(&dylibPtrs[0]))
		}

		// 释放内存
		defer func() {
			freeCString(opts.ipa_path)
			freeCString(opts.p12_path)
			freeCString(opts.provision_path)
			freeCString(opts.output_path)
			freeCString(opts.password)
			freeCString(opts.bundle_id)
			freeCString(opts.bundle_name)
			freeCString(opts.bundle_version)
			freeCString(opts.entitlements)
			freeCString(opts.temp_folder)

			for _, ptr := range dylibPtrs {
				freeCString(ptr)
			}
		}()

		// 调用签名函数
		signResult = C.ZSign(&opts)
	})

	if err != nil {
		return nil, err
	}

	return handleResult(signResult), nil
}

func boolToCInt(b bool) C.int {
	if b {
		return 1
	}
	return 0
}

func freeCString(ptr *C.char) {
	if ptr != nil {
		C.free(unsafe.Pointer(ptr))
	}
}

// 通用处理函数
func stringToCChar(s string) *C.char {
	if s == "" {
		return nil
	}
	return C.CString(s)
}

func handleResult(ret C.int) g.Map {
	switch ret {
	case 0:
		return g.Map{"code": 0, "message": "签名成功"}
	case 1:
		return g.Map{"code": 1, "message": "zsign 帮助信息"}
	case 2:
		return g.Map{"code": 2, "message": "zsign 版本信息"}
	default:
		// 使用 fmt.Sprintf 格式化错误信息并返回
		errorMessage := fmt.Sprintf("✗ 签名失败，错误码: %d", int(ret))
		return g.Map{"code": ret, "message": errorMessage}
	}
}
