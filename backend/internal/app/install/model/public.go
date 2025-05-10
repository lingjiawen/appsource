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

type GetDeviceDataResp struct {
	Code int
	Msg  string
	Data *GetDeviceData
}

type GetDeviceData struct {
	Id                     string `json:"id" p:"id" dc:"证书编号"`
	Name                   string `json:"name" p:"name" dc:"证书名"`
	Udid                   string `json:"udid" p:"udid" dc:"UDID"`
	DeviceId               string `json:"device_id" p:"device_id" dc:"设备编号"`
	Pool                   string `json:"pool" p:"pool" dc:"池类型"`
	Model                  string `json:"model" p:"model" dc:"设备型号"`
	Type                   string `json:"type" p:"type" dc:"设备类型:审核1秒出0"`
	Mobileprovision        string `json:"mobileprovision" p:"mobileprovision" dc:"描述文件"`
	P12                    string `json:"p12" p:"P12" dc:"证书文件"`
	P12Password            string `json:"p12_password" p:"p12_password" dc:"证书密码"`
	Status                 string `json:"status" p:"status" dc:"证书状态"`
	State                  bool   `json:"state" p:"state" dc:"证书有效"`
	AddTime                int64  `json:"add_time" p:"add_time" dc:"添加时间"`
	ExpireTime             int64  `json:"expire_time" p:"expire_time" dc:"过期时间"`
	Warranty               int    `json:"warranty" p:"warranty" dc:"质保类型"`
	WarrantyExpirationTime int64  `json:"warranty_expiration_time" p:"warranty_expiration_time" dc:"质保过期时间"`
	PlatformId             int    `json:"platform_id" p:"platform_id" dc:"平台ID"`
}

type GetHelpRes struct {
	List []*SignHelpListRes `json:"list"`
}

type GetAboutRes struct {
	List [][]*SignAboutListRes `json:"list"`
}
