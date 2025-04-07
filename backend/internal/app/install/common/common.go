package common

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"mangosmithy/internal/app/install/model"
	"mangosmithy/library/liberr"
	"strings"
)

// GetFullDirPath 获取文件路径
func GetFullDirPath(ctx context.Context, fileUrl string) string {
	staticPath, _ := g.Cfg().Get(ctx, "server.serverRoot")
	if staticPath.IsEmpty() {
		liberr.ErrIsNil(ctx, fmt.Errorf("获取静态路经失败"))
	}

	sp := staticPath.String()
	if sp != "" {
		sp = strings.TrimRight(sp, "/")
	}

	return sp + "/" + fileUrl
}

// SplitAndTrim splits a string by the given separator and trims spaces from each part.
// It returns a slice of non-empty, trimmed strings.
func SplitAndTrim(input string, separator string) []string {
	rawValues := strings.Split(input, separator)
	var values []string
	for _, v := range rawValues {
		trimmedValue := strings.TrimSpace(v)
		if trimmedValue != "" {
			values = append(values, trimmedValue)
		}
	}
	return values
}

func MakeInstallPlist(plistInfo model.InstallPlist) string {
	// 生成安装plist文件
	iosPlistTem := `<?xml version="1.0" encoding="UTF-8"?>
<plist version="1.0"><dict>
  <key>items</key>
  <array>
    <dict>
      <key>assets</key>
      <array>
        <dict>
          <key>kind</key>
          <string>software-package</string>
          <key>url</key>
          <string><![CDATA[` + plistInfo.AppFileUrl + `]]></string>
        </dict>
        <dict>
          <key>kind</key>
          <string>display-image</string>
          <key>needs-shine</key>
          <integer>0</integer>
          <key>url</key>
          <string><![CDATA[` + plistInfo.IconUrl + `]]></string>
        </dict>
        <dict>
          <key>kind</key>
          <string>full-size-image</string>
          <key>needs-shine</key>
          <true/>
          <key>url</key>
          <string><![CDATA[` + plistInfo.IconUrl + `]]></string>
        </dict>
      </array>
      <key>metadata</key>
      <dict>
        <key>bundle-identifier</key>
        <string>` + plistInfo.BundleID + `</string>
        <key>bundle-version</key>
        <string><![CDATA[` + plistInfo.Version + `]]></string>
        <key>kind</key>
        <string>software</string>
        <key>title</key>
        <string><![CDATA[` + plistInfo.Name + `]]></string>
      </dict>
    </dict>
  </array>
</dict>
</plist>`
	return iosPlistTem
}
