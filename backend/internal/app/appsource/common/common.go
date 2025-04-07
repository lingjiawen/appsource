package common

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
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
