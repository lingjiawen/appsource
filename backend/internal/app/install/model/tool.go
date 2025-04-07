package model

// CheckCertRes 证书检测返回
type CheckCertRes struct {
	Name                    string   `json:"name" dc:"证书名称"`
	ValidFrom               string   `json:"valid_from" dc:"签发时间"`
	ValidTo                 string   `json:"valid_to" dc:"有效期"`
	SerialNumber            string   `json:"serial_number" dc:"序列号"`
	Country                 string   `json:"country" dc:"国家"`
	Organization            string   `json:"organization" dc:"签发组织"`
	Status                  string   `json:"status" dc:"状态"`
	RevokedAt               string   `json:"revoked_at,omitempty" dc:"撤销时间"`
	CertMatch               bool     `json:"cert_match" dc:"与证书匹配"`
	DeviceMatch             bool     `json:"device_match" dc:"与设备匹配"`
	ProvisionName           string   `json:"provision_name" dc:"描述名称"`
	ProvisionUUID           string   `json:"provision_uuid" dc:"描述uuid"`
	ProductBundleIdentifier string   `json:"product_bundle_identifier" dc:"应用标识"`
	DevelopmentTeam         string   `json:"development_team" dc:"团队ID"`
	Devices                 []string `json:"devices" dc:"匹配设备列表"`
	Capabilities            []string `json:"capabilities" dc:"能力列表"`
	CapabilityCount         int      `json:"capability_count" dc:"权限数量"`
}

// CheckP12Res P12检测返回
type CheckP12Res struct {
	Name         string `json:"name" dc:"证书名称"`
	ValidFrom    string `json:"valid_from" dc:"签发时间"`
	ValidTo      string `json:"valid_to" dc:"有效期"`
	SerialNumber string `json:"serial_number" dc:"序列号"`
	Country      string `json:"country" dc:"国家"`
	Organization string `json:"organization" dc:"签发组织"`
	Status       string `json:"status" dc:"状态"`
	RevokedAt    string `json:"revoked_at,omitempty" dc:"撤销时间"`

	P8Active        bool     `json:"p8_active" dc:"p8有效"`
	P8Message       string   `json:"p8_message" dc:"p8报错内容"`
	Capabilities    []string `json:"capabilities" dc:"能力列表"`
	CapabilityCount int      `json:"capability_count" dc:"权限数量"`
}

// Capability 结构体
type Capability struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Key     string `json:"key"`
	ID      string `json:"id,omitempty"`
}
