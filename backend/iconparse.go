package main

import (
	"archive/zip"
	"fmt"
	"howett.net/plist"
	"io"
	"log"
	"os"
	"strings"
)

// InfoPlist 定义一个结构体来映射 Info.plist 中的字段
type InfoPlist struct {
	BundleIdentifier string                 `plist:"CFBundleIdentifier"`
	BundleVersion    string                 `plist:"CFBundleVersion"`
	DisplayName      string                 `plist:"CFBundleDisplayName"`
	IconFile         string                 `plist:"CFBundleIconFile"`
	IconFiles        []string               `plist:"CFBundlePrimaryIcon>CFBundleIconFiles"`
	PrimaryIcon      map[string]interface{} `plist:"CFBundleIcons>CFBundlePrimaryIcon"`
	AlternateIcons   map[string]interface{} `plist:"CFBundleIcons>CFBundleAlternateIcons"`
	OldIconFiles     []string               `plist:"CFBundleIconFiles"` // 旧版的图标字段
	OldIconFile      string                 `plist:"CFBundleIconFile"`  // 旧版的图标字段
}

func main() {
	// .ipa 文件路径
	//ipaFile := "./resource/public/upload_file/2025-02-19/d7wfyjcbm2j4d6qujy.ipa"
	ipaFile := "./resource/public/upload_file/2025-02-20/d7wxx4gxev0oquhgwf.ipa"

	// 打开 .ipa 文件
	r, err := zip.OpenReader(ipaFile)
	if err != nil {
		log.Fatalf("failed to open ipa file: %v", err)
	}
	defer r.Close()

	// 查找 .app 文件夹路径
	var appPath string
	for _, f := range r.File {
		// 确保路径包含 Payload/，并且文件夹名以 .app 结尾
		//fmt.Println("File:", f.Name) // 输出路径
		if strings.HasPrefix(f.Name, "Payload/") && strings.Contains(f.Name, ".app/") {
			// 找到 .app 文件夹路径
			appPath = strings.Split(f.Name, ".app/")[0] + ".app/"
			break
		}
	}

	if appPath == "" {
		log.Fatalf("No .app folder found in the IPA")
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
		log.Fatalf("failed to read plist file: %v", err)
	}

	// 使用 plist.Unmarshal 解析 plist 数据
	var infoPlist InfoPlist
	_, err = plist.Unmarshal(plistData, &infoPlist)
	if err != nil {
		log.Fatalf("failed to unmarshal plist: %v", err)
	}

	// 输出解析出的信息
	fmt.Println("Bundle Identifier:", infoPlist.BundleIdentifier)
	fmt.Println("Version:", infoPlist.BundleVersion)
	fmt.Println("Display Name:", infoPlist.DisplayName)

	// 处理图标信息
	iconPath := handleIconInfo(&infoPlist, &r.Reader, appPath)
	println(iconPath)
	//extractIconFile(&r.Reader, appPath, iconPath)
	extractIconFileToPath(&r.Reader, appPath, iconPath, ".")
}

// 处理图标信息
func handleIconInfo(infoPlist *InfoPlist, r *zip.Reader, appPath string) string {
	iconPath := ""
	// 情况 1: CFBundleIconFile (直接返回一个图标文件)
	if infoPlist.IconFile != "" {
		fmt.Println("Icon File:", infoPlist.IconFile)
		iconPath = infoPlist.IconFile
		if hasFile(r, appPath, iconPath) {
			return iconPath
		}

		if !strings.HasSuffix(iconPath, ".png") {
			// 如果没有后缀，假设为 PNG
			iconPath += ".png"

			// 再次检查文件是否存在
			if hasFile(r, appPath, iconPath) {
				return iconPath
			}
		}
	}

	// 情况 2: CFBundlePrimaryIcon > CFBundleIconFiles (如果存在则返回)
	if len(infoPlist.IconFiles) > 0 {
		fmt.Println("Primary Icon File:", infoPlist.IconFiles[0]) // 只获取第一个图标
		iconPath = infoPlist.IconFiles[0]
		if hasFile(r, appPath, iconPath) {
			return iconPath
		}

		if !strings.HasSuffix(iconPath, ".png") {
			// 如果没有后缀，假设为 PNG
			iconPath += ".png"

			// 再次检查文件是否存在
			if hasFile(r, appPath, iconPath) {
				return iconPath
			}
		}
	}

	// 情况 3: CFBundleIcons > CFBundlePrimaryIcon > CFBundleIconFiles (优先返回第一个图标)
	if infoPlist.PrimaryIcon != nil {
		if iconFiles, ok := infoPlist.PrimaryIcon["CFBundleIconFiles"].([]interface{}); ok {
			if len(iconFiles) > 0 {
				if str, ok := iconFiles[0].(string); ok {
					fmt.Println("Primary Icon File (from CFBundleIcons > CFBundlePrimaryIcon):", str)
					iconPath = str
					if hasFile(r, appPath, iconPath) {
						return iconPath
					}

					if !strings.HasSuffix(iconPath, ".png") {
						// 如果没有后缀，假设为 PNG
						iconPath += ".png"

						// 再次检查文件是否存在
						if hasFile(r, appPath, iconPath) {
							return iconPath
						}
					}
				}
			}
		}
	}

	// 情况 4: CFBundleIcons > CFBundleAlternateIcons (返回第一个备选图标)
	if infoPlist.AlternateIcons != nil {
		for altIconName, iconFiles := range infoPlist.AlternateIcons {
			if iconFilesArray, ok := iconFiles.([]interface{}); ok {
				if len(iconFilesArray) > 0 {
					if str, ok := iconFilesArray[0].(string); ok {
						fmt.Printf("Alternate Icon File (%s): %s\n", altIconName, str)
						iconPath = str
						if hasFile(r, appPath, iconPath) {
							return iconPath
						}
						if !strings.HasSuffix(iconPath, ".png") {
							// 如果没有后缀，假设为 PNG
							iconPath += ".png"

							// 再次检查文件是否存在
							if hasFile(r, appPath, iconPath) {
								return iconPath
							}
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
		if hasFile(r, appPath, iconPath) {
			return iconPath
		}

		if !strings.HasSuffix(iconPath, ".png") {
			// 如果没有后缀，假设为 PNG
			iconPath += ".png"

			// 再次检查文件是否存在
			if hasFile(r, appPath, iconPath) {
				return iconPath
			}
		}

		return iconPath
	}

	// 情况 6: CFBundleIconFile (旧版本图标，直接返回)
	if infoPlist.OldIconFile != "" {
		fmt.Println("Old Icon File:", infoPlist.OldIconFile)
		iconPath = infoPlist.OldIconFile
		if hasFile(r, appPath, iconPath) {
			return iconPath
		}

		if !strings.HasSuffix(iconPath, ".png") {
			// 如果没有后缀，假设为 PNG
			iconPath += ".png"

			// 再次检查文件是否存在
			if hasFile(r, appPath, iconPath) {
				return iconPath
			}
		}

		return iconPath
	}

	// 如果没有找到任何图标信息
	fmt.Println("No icon information found.")
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
