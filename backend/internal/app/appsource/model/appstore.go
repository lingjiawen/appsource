package model

// AppStoreReq 分页请求参数
type AppStoreReq struct {
	Udid string `p:"udid" dc:"UDID"` //UDID
	Code string `p:"code" dc:"兑换码"`  //兑换码
}

// AppStoreRes 列表返回结果
type AppStoreRes struct {
	Name       string `json:"name"`
	Message    string `json:"message"`
	Identifier string `json:"identifier"`
	SourceURL  string `json:"sourceURL"`
	SourceIcon string `json:"sourceicon"`
	PayURL     string `json:"payURL"`
	UnlockURL  string `json:"unlockURL"`
	Apps       []App  `json:"apps"`
}

type App struct {
	Name               string      `json:"name"`
	Version            string      `json:"version"`
	Type               interface{} `json:"type"`
	VersionDate        string      `json:"versionDate"`
	VersionDescription string      `json:"versionDescription"`
	Lock               string      `json:"lock"`
	DownloadURL        string      `json:"downloadURL"`
	IsLanZouCloud      string      `json:"isLanZouCloud"`
	IconURL            string      `json:"iconURL"`
	TintColor          string      `json:"tintColor"`
	Size               string      `json:"size"`
}
