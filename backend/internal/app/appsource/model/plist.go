package model

// InfoPlist 定义一个结构体来映射 Info.plist 中的字段
type InfoPlist struct {
	BundleIdentifier   string                 `plist:"CFBundleIdentifier"`
	BundleShortVersion string                 `plist:"CFBundleShortVersionString"`
	BundleVersion      string                 `plist:"CFBundleVersion"`
	DisplayName        string                 `plist:"CFBundleDisplayName"`
	IconFile           string                 `plist:"CFBundleIconFile"`
	PrimaryIcon        CFBundlePrimaryIcon    `plist:"CFBundlePrimaryIcon"`
	AlternateIcons     map[string]interface{} `plist:"CFBundleAlternateIcons"`
	OldIconFiles       []string               `plist:"CFBundleIconFiles"` // 旧版的图标字段
	OldIconFile        string                 `plist:"CFBundleIconFile"`  // 旧版的图标字段
	CFBundleIcons      map[string]interface{} `plist:"CFBundleIcons"`     // 新增 CFBundleIcons 字段
}

type CFBundlePrimaryIcon struct {
	CFBundleIconFiles []string `plist:"CFBundleIconFiles"`
	CFBundleIconName  string   `plist:"CFBundleIconName"`
}
