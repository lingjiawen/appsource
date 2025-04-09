package model

type GetUDIDResultReq struct {
	UDID string `p:"udid" dc:"UDID"`
}

type ApplicationInstallReq struct {
	UDID string `p:"udid" dc:"UDID"`
	Code string `p:"code" dc:"兑换码"`
}

type ApplicationInstallRes struct {
	Url string `json:"url" dc:"安装地址"`
}

type ApplicationPrivateInstallReq struct {
	P12             string `p:"base64p12" dc:"p12文件base64"`
	MobileProvision string `p:"base64mp" dc:"mobileprovision文件base64"`
	Password        string `p:"privatePassword" dc:"p12密码"`
}

type ApplicationPrivateInstallRes struct {
	Url string `json:"url" dc:"安装地址"`
}

type GetPlistReq struct {
	UUID string `p:"uuid" dc:"uuid"`
}

type GetPlistRes struct {
	PlistContent string `json:"plistContent" dc:"plist内容"`
}

type CheckDevicesReq struct {
	SearchType  string `p:"search_type" dc:"搜索类型"`
	SearchValue string `p:"search_value" dc:"搜索值"`
}

type DeviceInfo struct {
	Name        string `json:"name" dc:"设备名称"`
	Udid        string `json:"udid" dc:"UDID"`
	RedeemCode  string `json:"redeem_code" dc:"兑换码"`
	CertId      string `json:"cert_id" dc:"证书ID"`
	AddTime     string `json:"add_time" dc:"添加时间"`
	ExpireTime  string `json:"expire_time" dc:"过期时间"`
	Status      string `json:"status" dc:"状态"`
	AccountType string `json:"account_type" dc:"账户类型"`
}

type CheckDevicesRes struct {
	Devices []*DeviceInfo `json:"devices" dc:"设备列表"`
}
