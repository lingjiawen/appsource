package common

import (
	"archive/zip"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"mangosmithy/internal/app/appsource/model"
	commonConsts "mangosmithy/internal/app/common/consts"
	commonService "mangosmithy/internal/app/common/service"
	"mangosmithy/internal/app/system/consts"
	systemService "mangosmithy/internal/app/system/service"
	"os"
	"strings"
)

// ExtractIconFileAndUpload 提取图标文件并上传
func ExtractIconFileAndUpload(ctx context.Context, r *zip.Reader, appPath, iconPath string) (filePath string, err error) {
	// 查找图标文件在 zip 中的路径
	for _, f := range r.File {
		if f.Name == appPath+iconPath {
			// 找到图标文件，打开并读取
			iconFileReader, err := f.Open()
			if err != nil {
				return "", fmt.Errorf("failed to open icon file: %v", err)
			}
			defer iconFileReader.Close()

			// 上传文件
			v, _ := g.Cfg().Get(ctx, "upload.default")
			info, err := commonService.Upload().UploadFileInner(ctx, iconFileReader, iconPath, commonConsts.CheckFileTypeImg, v.Int(), systemService.Context().Get(ctx).User.Id, consts.UploadAppId)
			if err != nil {
				return "", fmt.Errorf("failed to upload icon file: %v", err)
			}

			filePath = info.Path
			return filePath, nil
		}
	}

	// 如果没有找到图标文件
	return "", fmt.Errorf("icon file not found in the IPA")
}

// HandleIconInfo 处理图标的函数
func HandleIconInfo(infoPlist *model.InfoPlist, r *zip.Reader, appPath string) string {
	iconPath := ""

	// 情况 1: CFBundleIconFile (直接返回一个图标文件)
	if infoPlist.IconFile != "" {
		fmt.Println("Icon File:", infoPlist.IconFile)
		iconPath = infoPlist.IconFile
		// 调用封装的函数检查文件是否存在
		result := checkIconPaths(r, appPath, iconPath)
		if result != "" {
			return result
		}
	}

	// 情况 2: CFBundlePrimaryIcon > CFBundleIconFiles (如果存在则返回)
	if len(infoPlist.PrimaryIcon.CFBundleIconFiles) > 0 {
		fmt.Println("Primary Icon File:", infoPlist.PrimaryIcon.CFBundleIconFiles[0]) // 只获取第一个图标
		iconPath = infoPlist.PrimaryIcon.CFBundleIconFiles[0]
		// 调用封装的函数检查文件是否存在
		result := checkIconPaths(r, appPath, iconPath)
		if result != "" {
			return result
		}
	}

	// 情况 3: CFBundleIcons > CFBundlePrimaryIcon > CFBundleIconFiles (优先返回第一个图标)
	if infoPlist.CFBundleIcons != nil {
		if primaryIcon, ok := infoPlist.CFBundleIcons["CFBundlePrimaryIcon"].(map[string]interface{}); ok {
			if iconFiles, ok := primaryIcon["CFBundleIconFiles"].([]interface{}); ok && len(iconFiles) > 0 {
				if str, ok := iconFiles[0].(string); ok {
					fmt.Println("Primary Icon File (from CFBundleIcons > CFBundlePrimaryIcon):", str)
					iconPath = str
					result := checkIconPaths(r, appPath, iconPath)
					if result != "" {
						return result
					}
				}
			}
		}
	}

	// 情况 4: CFBundleIcons > CFBundleAlternateIcons (返回第一个备选图标)
	if infoPlist.CFBundleIcons != nil {
		if alternateIcons, ok := infoPlist.CFBundleIcons["CFBundleAlternateIcons"].(map[string]interface{}); ok {
			for altIconName, iconFiles := range alternateIcons {
				if iconFilesArray, ok := iconFiles.([]interface{}); ok && len(iconFilesArray) > 0 {
					if str, ok := iconFilesArray[0].(string); ok {
						fmt.Printf("Alternate Icon File (%s): %s\n", altIconName, str)
						iconPath = str
						result := checkIconPaths(r, appPath, iconPath)
						if result != "" {
							return result
						}
					}
				}
			}
		}
	}

	// 情况 5: CFBundleIconFiles (旧版图标，直接获取第一个)
	if len(infoPlist.OldIconFiles) > 0 {
		fmt.Println("Old Icon File:", infoPlist.OldIconFiles[0])
		iconPath = infoPlist.OldIconFiles[0]
		// 调用封装的函数检查文件是否存在
		result := checkIconPaths(r, appPath, iconPath)
		if result != "" {
			return result
		}
		return iconPath
	}

	// 情况 6: CFBundleIconFile (旧版本图标，直接返回)
	if infoPlist.OldIconFile != "" {
		fmt.Println("Old Icon File:", infoPlist.OldIconFile)
		iconPath = infoPlist.OldIconFile
		// 调用封装的函数检查文件是否存在
		result := checkIconPaths(r, appPath, iconPath)
		if result != "" {
			return result
		}
		return iconPath
	}

	// 如果没有找到任何图标信息
	fmt.Println("No icon information found.")
	return ""
}

// 检查多个可能的文件路径是否存在，返回第一个存在的路径
func checkIconPaths(r *zip.Reader, appPath string, iconPath string) string {
	// 先检查没有后缀的情况（.png 后缀）
	if !strings.HasSuffix(iconPath, ".png") {
		// 检查 iconPath、@2x、@3x 路径
		iconPaths := []string{
			iconPath + ".png",    // 添加 .png 后缀
			iconPath + "@2x.png", // 检查 @2x.png
			iconPath + "@3x.png", // 检查 @3x.png
		}

		// 检查是否存在任一图标文件
		for _, path := range iconPaths {
			if hasFile(r, appPath, path) {
				return path
			}
		}
	} else {
		// 如果已经有后缀，直接检查原始路径
		if hasFile(r, appPath, iconPath) {
			return iconPath
		}
	}

	// 如果都不存在，返回空字符串
	return ""
}

// 检查文件是否存在
func hasFile(r *zip.Reader, appPath, fileName string) bool {
	iconPath := appPath + fileName
	for _, f := range r.File {
		if f.Name == iconPath {
			return true
		}
	}
	return false
}

// 提取图标文件
func extractIconFile(r *zip.Reader, appPath, iconFile string) {
	for _, f := range r.File {
		if f.Name == appPath+iconFile {
			fmt.Printf("Found icon: %s\n", f.Name)
			// 此处可以添加代码提取图标文件，保存或进行其他处理
			return
		}
	}
}

// 提取图标文件并保存到指定目录
func extractIconFileToPath(r *zip.Reader, appPath, iconFile, saveDir string) {
	// 查找图标文件在 zip 中的路径
	for _, f := range r.File {
		if f.Name == appPath+iconFile {
			fmt.Printf("Found icon: %s\n", f.Name)

			// 打开图标文件
			iconFileReader, err := f.Open()
			if err != nil {
				fmt.Printf("Failed to open icon file: %v\n", err)
				return
			}
			defer iconFileReader.Close()

			// 创建保存图标的文件路径
			savePath := saveDir + "/" + iconFile
			if !strings.HasSuffix(savePath, ".png") {
				// 如果没有后缀，尝试假设为 PNG
				savePath += ".png"
			}

			// 创建目标目录，如果它不存在
			err = os.MkdirAll(saveDir, os.ModePerm)
			if err != nil {
				fmt.Printf("Failed to create directory: %v\n", err)
				return
			}

			// 创建目标文件
			outFile, err := os.Create(savePath)
			if err != nil {
				fmt.Printf("Failed to create file: %v\n", err)
				return
			}
			defer outFile.Close()

			// 将图标文件内容写入目标文件
			_, err = io.Copy(outFile, iconFileReader)
			if err != nil {
				fmt.Printf("Failed to write icon file: %v\n", err)
				return
			}

			fmt.Printf("Icon saved to: %s\n", savePath)
			return
		}
	}

	// 如果没有找到图标文件
	fmt.Println("Icon file not found in the IPA.")
}
