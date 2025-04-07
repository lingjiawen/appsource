package controller

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/net/html/charset"
	"mangosmithy/api/v1/install"
	systemController "mangosmithy/internal/app/system/controller"
	"mangosmithy/library/libUtils"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"go.mozilla.org/pkcs7"
)

type getUDIDController struct {
	systemController.BaseController
}

var GetUDIDController = new(getUDIDController)

// 获取设备UDID请求
const (
	mobileConfigTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.appapple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>PayloadContent</key>
        <dict>
            <key>URL</key>
            <string>{{.ReturnUrl}}</string>
            <key>DeviceAttributes</key>
            <array>
                <string>UDID</string>
                <string>IMEI</string>
                <string>ICCID</string>
                <string>VERSION</string>
                <string>PRODUCT</string>
            </array>
        </dict>
        <key>PayloadOrganization</key>
        <string>{{.Domain}}</string>
        <key>PayloadDisplayName</key>
        <string>查询设备UDID</string>
        <key>PayloadVersion</key>
        <integer>1</integer>
        <key>PayloadUUID</key>
        <string>3C4DC7D2-E475-3375-489C-0BB8D737A653</string>
        <key>PayloadIdentifier</key>
        <string>{{.Domain}}</string>
        <key>PayloadDescription</key>
        <string>本文件仅用来获取设备ID</string>
        <key>PayloadType</key>
        <string>Profile Service</string>
    </dict>
</plist>`
)

func (c *getUDIDController) GetUDID(ctx context.Context, req *install.GetUDIDReq) (res *install.GetUDIDRes, err error) {
	res = new(install.GetUDIDRes)

	r := g.RequestFromCtx(ctx)
	domainURL := libUtils.GetDomain(ctx, false)
	// 安全处理domain参数
	domain, _ := extractDomain(domainURL)
	if domain == "" {
		return nil, gerror.New("invalid domain parameter")
	}

	safeDomain := strings.ReplaceAll(domain, ":", "_") // 将冒号转为下划线

	// 检查文件是否存在
	publicPathConfig, _ := g.Cfg().Get(ctx, "server.serverRoot")
	publicPath := publicPathConfig.String()
	getUDIDPath := filepath.Join(publicPath, "getudid")
	configPath := filepath.Join(getUDIDPath, fmt.Sprintf("/%s.mobileconfig", safeDomain))
	if fileExists(configPath) {
		r.Response.ServeFileDownload(configPath)
		return res, nil
	}

	// 生成MobileConfig内容
	returnURL, _ := url.JoinPath(domainURL, "/udidResult")
	// TODO:跳转前端页面, 这里要获取配置的前端页面地址
	redirectURL, _ := url.JoinPath(domainURL, "/sign")
	// 解析 returnURL
	u, err := url.Parse(returnURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// 添加查询参数 ?site=redirectURL
	query := u.Query()
	query.Set("site", redirectURL)
	u.RawQuery = query.Encode()

	// 获取最终的 URL
	finalReturnURL := u.String()
	content, err := generateMobileConfig(finalReturnURL, domain)
	if err != nil {
		return nil, err
	}

	// 创建目录
	if err := os.MkdirAll(getUDIDPath, 0755); err != nil {
		return nil, gerror.Wrap(err, "create directory failed")
	}

	// 写入临时文件
	tempPath := filepath.Join(getUDIDPath, fmt.Sprintf("unsigned_%s.mobileconfig", safeDomain))
	if err := os.WriteFile(tempPath, []byte(content), 0644); err != nil {
		return nil, gerror.Wrap(err, "write temp file failed")
	}

	// 执行签名命令
	if err := signMobileConfig(ctx, safeDomain); err != nil {
		return nil, err
	}

	// 返回文件
	r.Response.ServeFileDownload(getUDIDPath, fmt.Sprintf("%s.mobileconfig", safeDomain))
	return res, nil
}

func extractDomain(doURL string) (string, error) {
	parsedURL, err := url.Parse(doURL)
	if err != nil {
		return "", err
	}
	return parsedURL.Host, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func generateMobileConfig(returnUrl, domain string) (string, error) {
	tmpl, err := template.New("mobileconfig").Parse(mobileConfigTemplate)
	if err != nil {
		return "", gerror.Wrap(err, "parse template failed")
	}

	var buf strings.Builder
	data := map[string]string{
		"ReturnUrl": returnUrl,
		"Domain":    domain,
	}
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", gerror.Wrap(err, "execute template failed")
	}

	return buf.String(), nil
}

func signMobileConfig(ctx context.Context, safeDomain string) error {
	basePathConfig, _ := g.Cfg().Get(ctx, "server.serverRoot")
	basePath := basePathConfig.String()
	getUDIDPath := filepath.Join(basePath, "getudid")
	oldFile := filepath.Join(getUDIDPath, fmt.Sprintf("unsigned_%s.mobileconfig", safeDomain))
	newFile := filepath.Join(getUDIDPath, fmt.Sprintf("%s.mobileconfig", safeDomain))

	// 1. 加载签名证书和私钥
	signerCert, privateKey, err := loadSignerCredentials(
		filepath.Join(basePath, "discer.pem"),
		filepath.Join(basePath, "pkey.key"),
	)
	if err != nil {
		return gerror.Wrap(err, "load credentials failed")
	}

	// 2. 加载CA证书链
	caCerts, err := loadCACertificates(
		filepath.Join(basePath, "rootcer.pem"),
	)
	if err != nil {
		return gerror.Wrap(err, "load CA certificates failed")
	}

	// 3. 读取待签名内容
	content, err := os.ReadFile(oldFile)
	if err != nil {
		return gerror.Wrap(err, "read content failed")
	}

	// 4. 创建PKCS7签名
	signedData, err := pkcs7.NewSignedData(content)
	if err != nil {
		return gerror.Wrap(err, "create signed data failed")
	}

	// 5. 添加签名者
	if err = signedData.AddSigner(signerCert, privateKey, pkcs7.SignerInfoConfig{}); err != nil {
		return gerror.Wrap(err, "add signer failed")
	}

	// 6. 添加CA证书
	for _, cert := range caCerts {
		signedData.AddCertificate(cert)
	}

	// 7. 生成DER格式签名
	derBytes, err := signedData.Finish()
	if err != nil {
		return gerror.Wrap(err, "generate signature failed")
	}

	// 8. 写入目标文件
	if err = os.WriteFile(newFile, derBytes, 0644); err != nil {
		return gerror.Wrap(err, "write signed file failed")
	}

	// 9. 清理临时文件
	if err = os.Remove(oldFile); err != nil {
		return gerror.Wrap(err, "remove temp file failed")
	}

	return nil
}

func loadSignerCredentials(certPath, keyPath string) (*x509.Certificate, any, error) {
	// 加载PEM证书
	certData, err := os.ReadFile(certPath)
	if err != nil {
		return nil, nil, err
	}
	block, _ := pem.Decode(certData)
	if block == nil {
		return nil, nil, gerror.New("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, nil, err
	}

	// 加载私钥（支持PKCS1和PKCS8）
	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, nil, err
	}
	block, _ = pem.Decode(keyData)
	if block == nil {
		return nil, nil, gerror.New("failed to parse key PEM")
	}

	var privateKey any
	switch block.Type {
	case "RSA PRIVATE KEY":
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	case "PRIVATE KEY":
		privateKey, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	default:
		return nil, nil, gerror.Newf("unsupported key type: %s", block.Type)
	}
	if err != nil {
		return nil, nil, err
	}

	return cert, privateKey, nil
}

func loadCACertificates(caPath string) ([]*x509.Certificate, error) {
	caData, err := os.ReadFile(caPath)
	if err != nil {
		return nil, err
	}

	var certs []*x509.Certificate
	for {
		block, rest := pem.Decode(caData)
		if block == nil {
			break
		}
		if block.Type == "CERTIFICATE" {
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, err
			}
			certs = append(certs, cert)
		}
		caData = rest
	}

	if len(certs) == 0 {
		return nil, gerror.New("no CA certificates found")
	}
	return certs, nil
}

// UDID数据苹果返回

type UDIDPayload struct {
	XMLName xml.Name `xml:"plist"`
	Version string   `xml:"version,attr"`
	Dict    struct {
		Keys   []string `xml:"key"`
		Values []string `xml:"string"`
	} `xml:"dict"`
}

func (c *getUDIDController) UDIDResult(ctx context.Context, req *install.GetUDIDResultReq) (res *install.GetUDIDResultRes, err error) {
	r := g.RequestFromCtx(ctx)
	res = &install.GetUDIDResultRes{}

	// 1. 获取原始请求体
	rawBody := r.GetBody()
	if len(rawBody) == 0 {
		return nil, gerror.New("empty request body")
	}

	// 2. 提取有效XML片段
	xmlContent, err := extractXMLContent(rawBody)
	if err != nil {
		return nil, gerror.Wrap(err, "XML内容提取失败")
	}

	// 3. 编码转换
	utf8Reader, err := charset.NewReader(
		bytes.NewReader(xmlContent),
		"application/xml; charset=utf-8",
	)
	if err != nil {
		return nil, gerror.Wrap(err, "编码转换失败")
	}

	var payload UDIDPayload
	decoder := xml.NewDecoder(utf8Reader)
	decoder.Strict = false
	decoder.Entity = xml.HTMLEntity
	if err := decoder.Decode(&payload); err != nil {
		g.Log().Errorf(ctx, "XML解析结构体: %+v", payload) // 新增调试日志
		return nil, gerror.Wrap(err, "XML结构解析失败")
	}

	// 提取设备信息的改进方法
	deviceInfo := make(map[string]string)
	for i, key := range payload.Dict.Keys {
		if i < len(payload.Dict.Values) {
			switch key {
			case "UDID", "IMEI", "PRODUCT", "VERSION":
				deviceInfo[key] = payload.Dict.Values[i]
			}
		}
	}

	// 验证必要字段
	if deviceInfo["UDID"] == "" {
		g.Log().Errorf(ctx, "解析后的设备信息: %+v", deviceInfo)
		return nil, gerror.New("UDID字段缺失")
	}

	// 7. 构造重定向URL
	targetSite := r.GetQuery("site").String()
	if targetSite == "" {
		return nil, gerror.New("缺少站点参数")
	}

	// 安全验证域名格式
	if !isValidDomain(targetSite) {
		return nil, gerror.New("非法域名格式")
	}

	// 8. 执行重定向
	params := g.Map{
		"udid": deviceInfo["UDID"],
	}
	redirectURL := buildRedirectURL(targetSite, params)
	r.Response.Header().Set("Location", redirectURL)
	r.Response.WriteHeader(301)
	return
}

// 安全验证域名或 IP
func isValidDomain(targetSite string) bool {
	// 解析 URL，避免直接判断
	parsedURL, err := url.Parse(targetSite)
	if err != nil {
		return false // 解析失败，说明格式非法
	}

	host := parsedURL.Host
	if host == "" {
		host = parsedURL.Path // 处理裸域名（无 http/https）
	}

	// 允许 localhost
	if strings.HasPrefix(host, "localhost") {
		return true
	}

	// 解析端口号，去掉 ":8080" 之类的端口部分
	hostWithoutPort, _, err := net.SplitHostPort(host)
	if err != nil {
		hostWithoutPort = host // 没有端口号，直接使用原 host
	}

	// 检查是否是 IP 地址
	if net.ParseIP(hostWithoutPort) != nil {
		return true
	}

	// 允许合法的域名（必须包含 "."，且不能包含特殊字符）
	if strings.Contains(hostWithoutPort, ".") && !strings.ContainsAny(hostWithoutPort, " /\\#%&") {
		return true
	}

	return false
}

// 构造重定向URL
func buildRedirectURL(site string, params g.Map) string {
	queryParams := make([]string, 0, len(params))
	for k, v := range params {
		if strVal, ok := v.(string); ok && strVal != "" {
			queryParams = append(queryParams, fmt.Sprintf("%s=%s", k, strVal))
		}
	}

	// 自动添加协议头
	if !strings.HasPrefix(site, "http") {
		site = "https://" + site
	}

	return fmt.Sprintf("%s?%s", site, strings.Join(queryParams, "&"))
}

// 提取XML核心内容（兼容签名数据）
func extractXMLContent(raw []byte) ([]byte, error) {
	// 查找XML起始位置
	start := bytes.Index(raw, []byte("<?xml"))
	if start == -1 {
		return nil, fmt.Errorf("未找到XML起始标记")
	}

	// 查找XML结束位置
	end := bytes.LastIndex(raw, []byte("</plist>"))
	if end == -1 {
		return nil, fmt.Errorf("未找到XML结束标记")
	}
	end += len("</plist>")

	// 截取有效内容
	content := raw[start:end]

	// 清理非法字符
	return bytes.Map(func(r rune) rune {
		if r == 0 {
			return -1
		}
		return r
	}, content), nil
}
