package tools

import (
	"bytes"
	"context"
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"howett.net/plist"
	"io"
	"log"
	"mangosmithy/internal/app/install/model"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/crypto/ocsp"
	"software.sslmate.com/src/go-pkcs12"
)

const (
	rootCAG2URL        = "https://www.apple.com/certificateauthority/AppleRootCA-G2.cer"
	rootDeveloperG2URL = "https://www.apple.com/certificateauthority/AppleWWDRCAG2.cer"
	rootDeveloperG3URL = "https://www.apple.com/certificateauthority/AppleWWDRCAG3.cer"
	// rootCAG3URL             = "https://www.apple.com/certificateauthority/AppleRootCA-G3.cer"
	intermediateWWDRCAG6URL = "https://www.apple.com/certificateauthority/AppleWWDRCAG6.cer"
	rootCACertPEM           = `-----BEGIN CERTIFICATE-----
MIIEUTCCAzmgAwIBAgIQfK9pCiW3Of57m0R6wXjF7jANBgkqhkiG9w0BAQsFADBi
MQswCQYDVQQGEwJVUzETMBEGA1UEChMKQXBwbGUgSW5jLjEmMCQGA1UECxMdQXBw
bGUgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkxFjAUBgNVBAMTDUFwcGxlIFJvb3Qg
Q0EwHhcNMjAwMjE5MTgxMzQ3WhcNMzAwMjIwMDAwMDAwWjB1MUQwQgYDVQQDDDtB
cHBsZSBXb3JsZHdpZGUgRGV2ZWxvcGVyIFJlbGF0aW9ucyBDZXJ0aWZpY2F0aW9u
IEF1dGhvcml0eTELMAkGA1UECwwCRzMxEzARBgNVBAoMCkFwcGxlIEluYy4xCzAJ
BgNVBAYTAlVTMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2PWJ/KhZ
C4fHTJEuLVaQ03gdpDDppUjvC0O/LYT7JF1FG+XrWTYSXFRknmxiLbTGl8rMPPbW
BpH85QKmHGq0edVny6zpPwcR4YS8Rx1mjjmi6LRJ7TrS4RBgeo6TjMrA2gzAg9Dj
+ZHWp4zIwXPirkbRYp2SqJBgN31ols2N4Pyb+ni743uvLRfdW/6AWSN1F7gSwe0b
5TTO/iK1nkmw5VW/j4SiPKi6xYaVFuQAyZ8D0MyzOhZ71gVcnetHrg21LYwOaU1A
0EtMOwSejSGxrC5DVDDOwYqGlJhL32oNP/77HK6XF8J4CjDgXx9UO0m3JQAaN4LS
VpelUkl8YDib7wIDAQABo4HvMIHsMBIGA1UdEwEB/wQIMAYBAf8CAQAwHwYDVR0j
BBgwFoAUK9BpR5R2Cf70a40uQKb3R01/CF4wRAYIKwYBBQUHAQEEODA2MDQGCCsG
AQUFBzABhihodHRwOi8vb2NzcC5hcHBsZS5jb20vb2NzcDAzLWFwcGxlcm9vdGNh
MC4GA1UdHwQnMCUwI6AhoB+GHWh0dHA6Ly9jcmwuYXBwbGUuY29tL3Jvb3QuY3Js
MB0GA1UdDgQWBBQJ/sAVkPmvZAqSErkmKGMMl+ynsjAOBgNVHQ8BAf8EBAMCAQYw
EAYKKoZIhvdjZAYCAQQCBQAwDQYJKoZIhvcNAQELBQADggEBAK1lE+j24IF3RAJH
Qr5fpTkg6mKp/cWQyXMT1Z6b0KoPjY3L7QHPbChAW8dVJEH4/M/BtSPp3Ozxb8qA
HXfCxGFJJWevD8o5Ja3T43rMMygNDi6hV0Bz+uZcrgZRKe3jhQxPYdwyFot30ETK
XXIDMUacrptAGvr04NM++i+MZp+XxFRZ79JI9AeZSWBZGcfdlNHAwWx/eCHvDOs7
bJmCS1JgOLU5gm3sUjFTvg+RTElJdI+mUcuER04ddSduvfnSXPN/wmwLCTbiZOTC
NwMUGdXqapSqqdv+9poIZ4vvK7iqF0mDr8/LvOnP6pVxsLRFoszlh6oKw0E6eVza
UDSdlTs=
-----END CERTIFICATE-----`
)

const (
	MessageLine   = 0
	MessageUpdate = 1
)

// MobileProvision represents the structure of a parsed mobileprovision file.
type MobileProvision struct {
	CreationDate          time.Time              `plist:"CreationDate"`
	ExpirationDate        time.Time              `plist:"ExpirationDate"`
	Entitlements          map[string]interface{} `plist:"Entitlements"`
	ProvisionedDevices    []string               `plist:"ProvisionedDevices"`
	DeveloperCertificates [][]byte               `plist:"DeveloperCertificates"`
	UUID                  string                 `plist:"UUID"`
	Name                  string                 `plist:"Name"`
	TeamIdentifier        []string               `plist:"TeamIdentifier"`
}

// CertificateDetails holds the details of a developer certificate.
type CertificateDetails struct {
	Name            string `json:"name"`
	CreationDate    string `json:"creation_date"`
	SerialNumberDec string `json:"serial_number_dec"`
	SerialNumberHex string `json:"serial_number_hex"`
	SHA1            string `json:"sha1"`
}

func verifyTest() {
	// 正常
	// p12Path := "/Users/camen/Library/Containers/com.tencent.xinWeChat/Data/Library/Application Support/com.tencent.xinWeChat/2.0b4.0.9/aa0ba35726e928dd4f5dfdf115324240/Message/MessageTemp/fbe105618b579054f7e5c071f55a01e8/File/cert/cert.p12"
	// 撤销
	p12Path := "/Users/camen/Documents/撤销证书/TU5R30_证书文件（密码1）.p12"
	// 封号
	// p12Path := "/Users/camen/Documents/封号证书/2Y427J_证书文件（密码1）.p12"
	p12Data, err := os.ReadFile(p12Path)
	if err != nil {
		log.Fatalf("无法读取 p12 文件: %v", err)
	}
	password := "1"
	result, err := OnlyVerifyP12Status(base64.StdEncoding.EncodeToString(p12Data), password)

	if err != nil {
		log.Fatalf("检测出错: %v", err)
	}

	log.Println("Resulst:", result)
}

func infoTest() {
	httpClient := &http.Client{Timeout: 5 * time.Second}

	// 从苹果服务器获取根 CA 和中间 CA 证书
	// rootCACert, err := fetchCACert(httpClient, rootDeveloperG3URL)
	// if err != nil {
	// 	log.Fatalf("获取根 CA 证书失败: %v", err)
	// }

	// 解析 PEM 根证书
	block, _ := pem.Decode([]byte(rootCACertPEM))
	if block == nil {
		log.Fatalf("无法解码 PEM 根证书")
	}

	rootCACert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatalf("解析根证书失败: %v", err)
	}

	// 正常
	// p12Path := "/Users/camen/Library/Containers/com.tencent.xinWeChat/Data/Library/Application Support/com.tencent.xinWeChat/2.0b4.0.9/aa0ba35726e928dd4f5dfdf115324240/Message/MessageTemp/fbe105618b579054f7e5c071f55a01e8/File/cert/cert.p12"
	// 撤销
	p12Path := "/Users/camen/Documents/撤销证书/test.p12"
	// 封号
	// p12Path := "/Users/camen/Documents/封号证书/2Y427J_证书文件（密码1）.p12"
	p12Data, err := os.ReadFile(p12Path)
	if err != nil {
		log.Fatalf("无法读取 p12 文件: %v", err)
	}

	password := "1"

	// 将 Base64 编码的P12数据解码为二进制
	// p12Data, err := base64.StdEncoding.DecodeString(base64Data)
	// if err != nil {
	// 	log.Fatalf("无法解码 Base64 数据: %v", err)
	// }

	// 加载 p12 文件
	// 使用 pkcs12.Decode 解码 p12 文件，提取证书和私钥
	// privateKey, cert, err := pkcs12.Decode(p12Data, password)

	_, cert, err := pkcs12.Decode(p12Data, password)
	if err != nil {
		log.Fatalf("无法解码 p12 文件: %v", err)
	}

	log.Printf("加载的证书: %s", cert.Subject.CommonName)

	// x509Cert, err := x509.ParseCertificate(cert.Raw)
	// if err != nil {
	// log.Fatalf("无法解析 X.509 证书: %v", err)
	// }

	// 转换为 PEM 格式并打印
	// certPEM := certToPEM(cert)
	// keyPEM, err := keyToPEM(privateKey)
	// if err != nil {
	// 	log.Fatalf("转换私钥失败: %v", err)
	// }

	// fmt.Printf("证书 PEM:\n%s\n", certPEM)
	// fmt.Printf("私钥 PEM:\n%s\n", keyPEM)

	log.Println("Root CA Subject:", rootCACert.Subject)

	ctx := context.Background()
	verifier := newOCSPVerifier(httpClient)

	// 对中间 CA 证书进行 OCSP 验证
	verifyResult, err := verifier.verify(ctx, cert, rootCACert)
	if err != nil {
		log.Fatalf("OCSP 验证失败: %v", err)
	}

	// 打印证书的所有信息
	printCertAllInfo(cert, verifyResult)

	log.Println("OCSP 验证成功")
}

// OnlyVerifyP12Status Public: 单独验证证书是否掉签
func OnlyVerifyP12Status(base64P12 string, password string) (result g.Map, err error) {
	// http Client
	httpClient := &http.Client{Timeout: 5 * time.Second}
	// 解析 PEM 根证书
	block, _ := pem.Decode([]byte(rootCACertPEM))
	if block == nil {
		return nil, fmt.Errorf("无法解码 PEM 根证书")
	}

	rootCACert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析根证书失败: %v", err)
	}

	// log.Println("Root CA Subject:", rootCACert.Subject)

	// 将 Base64 编码的P12数据解码为二进制
	p12Data, err := base64.StdEncoding.DecodeString(base64P12)
	if err != nil {
		return nil, fmt.Errorf("无法解码 Base64 数据: %v", err)
	}

	_, cert, err := pkcs12.Decode(p12Data, password)
	if err != nil {
		// 检测错误是否是因为密码错误
		if strings.Contains(err.Error(), "password incorrect") {
			return nil, fmt.Errorf("证书密码错误")
		}
		return nil, fmt.Errorf("无法解码 p12 文件: %v", err)
	}

	// log.Printf("加载的证书: %s", cert.Subject.CommonName)

	ctx := context.Background()
	verifier := newOCSPVerifier(httpClient)

	// 对中间 CA 证书进行 OCSP 验证
	verifyResult, err := verifier.verify(ctx, cert, rootCACert)
	if err != nil {
		return nil, fmt.Errorf("OCSP 验证失败: %v", err)
	}

	return verifyResult, nil
}

// GetCertBaseInfo Public: 单独获取证书信息
func GetCertBaseInfo(base64P12, password string) (result g.Map, err error) {
	p12Data, err := base64.StdEncoding.DecodeString(base64P12)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("无法解码 Base64 数据: %v", err))
	}

	_, cert, err := pkcs12.Decode(p12Data, password)
	if err != nil {
		// 检测错误是否是因为密码错误
		if strings.Contains(err.Error(), "password incorrect") {
			return nil, fmt.Errorf("证书密码错误")
		}
		return nil, errors.New(fmt.Sprintf("无法解码 p12 文件: %v", err))
	}

	//log.Printf("加载的证书: %s", cert.Subject.CommonName)

	return getCertBaseInfo(cert)
}

// GetCertAllInfo Public: 单独获取证书所有信息
func GetCertAllInfo(base64P12, password string) (result g.Map, err error) {
	p12Data, err := base64.StdEncoding.DecodeString(base64P12)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("无法解码 Base64 数据: %v", err))
	}

	_, cert, err := pkcs12.Decode(p12Data, password)
	if err != nil {
		// 检测错误是否是因为密码错误
		if strings.Contains(err.Error(), "password incorrect") {
			return nil, fmt.Errorf("证书密码错误")
		}
		return nil, errors.New(fmt.Sprintf("无法解码 p12 文件: %v", err))
	}

	log.Printf("加载的证书: %s", cert.Subject.CommonName)

	p12Result, err := getCertAllInfo(cert)
	if err != nil {
		return nil, err
	}
	return p12Result, nil
}

// GetP12AndMPInfo Public: 获取所有证书和描述文件信息
func GetP12AndMPInfo(base64P12, base64mp, password string) (result g.Map, err error) {
	result = g.Map{}
	p12Result, err := GetCertAllInfo(base64P12, password)
	if err != nil {
		return nil, err
	}
	result["certInfo"] = p12Result
	/*
		{
			"name": "iPhone Distribution: Olwah Mutasem (VWRGL2WQ5T)", // 证书名称
			"validFrom": "2024-09-11 16:33:26 GMT+08:00",   // 签发时间
			"validTo": "2025-09-11 16:33:25 GMT+08:00",   // 有效期
			"issuer": "Apple Worldwide Developer Relations Certification Authority", // 机构
			"country": "US", // 国家
			"organization": "Olwah Mutasem", // 签发组织
			"serialNumberHex": "4D79690370FDB8324979790319C86698", // 序列号
			"serialNumberDec": "102980953520983156500250457409187243672",
			"status": "Revoked",  // or "Good" or "Suspended"  // 证书状态
			"revokedAt": "2024-09-12 16:04:01 GMT+08:00"  // if revoked // 撤销时间
		}
	*/
	mpResult, err := getProvisionInfo(base64mp)
	if err != nil {
		return nil, err
	}
	/*
		{
			name -> 00008220-001179121AE8201E16,  // 描述名称
			applicationIdentifier -> 57587UR6HN.com.neicexiahh2285.sign8950,
			uuid -> 94e8b4bf-9bcf-481c-83e8-8cd4aa866d11, // 描述UUID
			productBundleIdentifier -> com.neicexiahh2285.sign8950, // 应用标识
			developmentTeam -> 57587UR6HN,     // 团队ID
			certificates -> [{
				Name: CN=iPhone Distribution: Neerij Pamwani (57587UR6HN),OU=57587UR6HN,O=Neerij Pamwani,C=US,0.9.2342.19200300.100.1.1=#130a3537353837555236484e,
				CreationDate: 2024-11-11T16:01:28Z
				SerialNumberDec: 25651999809875930196301993886588512946
				SerialNumberHex: 134c653ca6e75246e90339021b42aab2
				SHA1: c50e50a3cf9b89c283454b5c001e98f88ca962666b89b76aa7b8922629111e0cd5215261a03bdeae357e4f36b5efc22d62e0adff3ff2fbfc2e53993c51c341fe73b5845370bc53f8ffe291bb946ad2c0f9fea319593461ed592ef814ba550b13b4fc1ac7aec8a2efaa5fd0b2079a76a1aac62b19d34d6ff0c57e4f91994ca0919895ed62fa464687b5ad62640314f125134cfcf48342819bf9e3419c01e7bf2b41739f1c1e023649bcdfd7f1047cbd465f8c1c22ae93e15b33b253fc29bb6b26905b821dc580fb20f65869b6b34203e450855ecf41a09837d00226f3160fc9db563280ef0e561003e8deed32499e16a40d8fdadb895e41f950dcd31cb0f0bd62
			}]
			capabilities -> [1. Wi-Fi 信息访问(ACCESS_WIFI_INFORMATION)：✅, 2. Apple ID 认证(APPLE_ID_AUTH)：❌...]
			devices -> ["00008220-001179121AE8201E"]
			creationDate -> 2024-11-12T10:52:38Z
			expirationDate -> 2025-11-11T16:01:27Z
		}
	*/
	result["mpInfo"] = mpResult
	return result, nil
}

// CheckCert Public: 证书检测
func CheckCert(base64P12, base64mp, password string) (res *model.CheckCertRes, err error) {
	checkInfo, err := GetP12AndMPInfo(base64P12, base64mp, password)
	if err != nil {
		return nil, err
	}

	certInfo := checkInfo["certInfo"].(g.Map)
	mpInfo := checkInfo["mpInfo"].(g.Map)

	// 使用安全的类型断言和默认值处理
	getStringValue := func(m g.Map, key string, defaultValue string) string {
		if value, ok := m[key].(string); ok {
			return value
		}
		return defaultValue
	}
	getStringArrayValue := func(m g.Map, key string, defaultValue []string) []string {
		if value, ok := m[key].([]string); ok {
			return value
		}
		return defaultValue
	}

	validFrom := getStringValue(certInfo, "validFrom", "")
	validTo := getStringValue(certInfo, "validTo", "")
	revokedAt := getStringValue(certInfo, "revokedAt", "")
	// 使用 gtime.NewFromStrFormat 解析字符串为时间对象
	parsedValidFrom := gtime.NewFromStrLayout(validFrom, "2006-01-02T15:04:05-07:00")
	parsedValidTo := gtime.NewFromStrLayout(validTo, "2006-01-02T15:04:05-07:00")

	// 使用 gtime.Format 格式化输出为 "Y-m-d H:i:s" 格式
	validFromFormatted := parsedValidFrom.Format("Y-m-d H:i:s")
	validToFormatted := parsedValidTo.Format("Y-m-d H:i:s")

	if revokedAt != "" {
		parsedRevokedAt := gtime.NewFromStrLayout(revokedAt, "2006-01-02T15:04:05-07:00")
		revokedAt = parsedRevokedAt.Format("Y-m-d H:i:s")
	}

	// 判断设备是否匹配
	devices := getStringArrayValue(mpInfo, "devices", []string{})

	// 获取序列号
	serialNumber := getStringValue(certInfo, "serialNumberHex", "")
	if serialNumber == "" {
		return nil, errors.New("获取serialNumber错误")
	}

	// 判断证书和描述文件是否匹配
	certificates, ok := mpInfo["certificates"].([]CertificateDetails)
	if !ok {
		return nil, errors.New("解析证书数据失败")
	}
	certMatch := false
	for _, cert := range certificates {
		if cert.SerialNumberHex == serialNumber {
			certMatch = true
			break
		}
	}

	// 证书状态修改
	certStatus := getStringValue(certInfo, "status", "Unknown")
	status := "normal"
	if certStatus == "Revoked" {
		status = "hidden"
	}
	if certStatus == "Suspended" {
		status = "banned"
	}
	res = &model.CheckCertRes{
		// 证书信息
		Name:         getStringValue(certInfo, "name", ""),
		ValidFrom:    validFromFormatted,
		ValidTo:      validToFormatted,
		SerialNumber: serialNumber,
		Country:      getStringValue(certInfo, "country", ""),
		Organization: getStringValue(certInfo, "organization", ""),
		Status:       status,
		RevokedAt:    revokedAt,
		CertMatch:    certMatch,
		// 描述文件信息
		ProvisionName:           getStringValue(mpInfo, "name", ""),
		ProvisionUUID:           getStringValue(mpInfo, "uuid", ""),
		ProductBundleIdentifier: getStringValue(mpInfo, "productBundleIdentifier", ""),
		DevelopmentTeam:         getStringValue(mpInfo, "developmentTeam", ""),
		Devices:                 devices,
	}
	return
}

// CheckP12 Public: P12检测
func CheckP12(base64P12, password string) (res *model.CheckP12Res, err error) {
	certInfo, err := GetCertAllInfo(base64P12, password)
	if err != nil {
		return nil, err
	}

	// 使用安全的类型断言和默认值处理
	getStringValue := func(m g.Map, key string, defaultValue string) string {
		if value, ok := m[key].(string); ok {
			return value
		}
		return defaultValue
	}

	validFrom := getStringValue(certInfo, "validFrom", "")
	validTo := getStringValue(certInfo, "validTo", "")
	revokedAt := getStringValue(certInfo, "revokedAt", "")
	// 使用 gtime.NewFromStrFormat 解析字符串为时间对象
	parsedValidFrom := gtime.NewFromStrLayout(validFrom, "2006-01-02T15:04:05-07:00")
	parsedValidTo := gtime.NewFromStrLayout(validTo, "2006-01-02T15:04:05-07:00")

	// 使用 gtime.Format 格式化输出为 "Y-m-d H:i:s" 格式
	validFromFormatted := parsedValidFrom.Format("Y-m-d H:i:s")
	validToFormatted := parsedValidTo.Format("Y-m-d H:i:s")

	if revokedAt != "" {
		parsedRevokedAt := gtime.NewFromStrLayout(revokedAt, "2006-01-02T15:04:05-07:00")
		revokedAt = parsedRevokedAt.Format("Y-m-d H:i:s")
	}

	// 获取序列号
	serialNumber := getStringValue(certInfo, "serialNumberHex", "")
	if serialNumber == "" {
		return nil, errors.New("获取serialNumber错误")
	}

	// 证书状态修改
	certStatus := getStringValue(certInfo, "status", "Unknown")
	status := "normal"
	if certStatus == "Revoked" {
		status = "hidden"
	}
	if certStatus == "Suspended" {
		status = "banned"
	}
	res = &model.CheckP12Res{
		// 证书信息
		Name:         getStringValue(certInfo, "name", ""),
		ValidFrom:    validFromFormatted,
		ValidTo:      validToFormatted,
		SerialNumber: serialNumber,
		Country:      getStringValue(certInfo, "country", ""),
		Organization: getStringValue(certInfo, "organization", ""),
		Status:       status,
		RevokedAt:    revokedAt,
	}
	return
}

// ChangeP12Password Public: 修改 P12 证书的密码
func ChangeP12Password(base64P12, currentPassword string) (string, error) {
	// 解码 base64P12
	p12Bytes, err := base64.StdEncoding.DecodeString(base64P12)
	if err != nil {
		return "", err
	}

	// 使用当前密码解码 P12
	privateKey, cert, caCerts, err := pkcs12.DecodeChain(p12Bytes, currentPassword)
	if err != nil {
		return "", err
	}

	// 重新编码 P12，设置新密码为 '1'
	newPassword := "1"
	newP12Bytes, err := pkcs12.Modern.Encode(privateKey, cert, caCerts, newPassword)
	if err != nil {
		return "", err
	}

	// 编码为 base64
	newBase64P12 := base64.StdEncoding.EncodeToString(newP12Bytes)
	return newBase64P12, nil
}

// 打印证书的 Base64 编码
func printCertBase64(cert *x509.Certificate) {
	// 将证书转换为 PEM 格式
	pemCert := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})

	// 打印 Base64 编码
	base64Cert := base64.StdEncoding.EncodeToString(pemCert)
	fmt.Println("Certificate Base64:")
	fmt.Println(base64Cert)
}

// 打印证书内容
func printCert(cert *x509.Certificate) {
	// 将证书转换为 PEM 格式
	pemCert := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})

	// 打印证书内容
	fmt.Println("Certificate PEM:")
	fmt.Println(string(pemCert))
}

func printCertAllInfo(cert *x509.Certificate, verifyResult g.Map) {
	printCertBaseInfo(cert)
	if verifyResult["status"] == ocsp.Revoked {
		certStatus := "未知"
		if verifyResult["revocationReason"] == 1 {
			certStatus = "🔴撤销"
		} else {
			certStatus = "🔴封号"
		}
		revokedTime := verifyResult["revokedAt"]
		fmt.Printf("证书状态：%s\n", certStatus)
		// 加载东八区的时区信息
		location, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			log.Fatalf("无法加载时区信息: %v", err)
		}
		// 将证书时间转换为东八区
		revokedTimeFormatted := revokedTime.(time.Time).In(location).Format("2006-01-02T15:04:05-07:00")
		fmt.Printf("撤销时间：%s\n", revokedTimeFormatted)
	} else if verifyResult["status"] == ocsp.Good {
		fmt.Printf("证书状态：🟢正常\n")
	} else {
		fmt.Printf("证书状态：🔴未知\n")
	}

}

func printCertBaseInfo(cert *x509.Certificate) {
	// 获取证书信息
	certName := cert.Subject.CommonName

	// 加载东八区的时区信息
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatalf("无法加载时区信息: %v", err)
	}
	// 将证书时间转换为东八区
	validFrom := cert.NotBefore.In(location).Format("2006-01-02T15:04:05-07:00")
	validTo := cert.NotAfter.In(location).Format("2006-01-02T15:04:05-07:00")
	// validFrom := cert.NotBefore.Format("2006-01-02 15:04:05 GMT-07:00")
	// validTo := cert.NotAfter.Format("2006-01-02 15:04:05 GMT-07:00")
	issuer := cert.Issuer.CommonName
	country := cert.Subject.Country[0]
	organization := cert.Subject.Organization[0]
	serialHex := hex.EncodeToString(cert.SerialNumber.Bytes())
	serialDec := cert.SerialNumber.String()

	// 格式化输出
	fmt.Printf("证书名称：%s\n", certName)
	fmt.Printf("生效时间：%s\n", validFrom)
	fmt.Printf("过期时间：%s\n", validTo)
	fmt.Printf("签发机构：%s\n", issuer)
	fmt.Printf("所属国家：%s\n", country)
	fmt.Printf("所属组织：%s\n", organization)
	fmt.Printf("证书编号(十六进制)：%s\n", serialHex)
	fmt.Printf("证书编号(十进制)：%s\n", serialDec)
}

func getCertAllInfo(cert *x509.Certificate) (g.Map, error) {
	/*
		{
			"name": "iPhone Distribution: Olwah Mutasem (VWRGL2WQ5T)",
			"validFrom": "2024-09-11 16:33:26 GMT+08:00",
			"validTo": "2025-09-11 16:33:25 GMT+08:00",
			"issuer": "Apple Worldwide Developer Relations Certification Authority",
			"country": "US",
			"organization": "Olwah Mutasem",
			"serialNumberHex": "4D79690370FDB8324979790319C86698",
			"serialNumberDec": "102980953520983156500250457409187243672",
			"status": "Revoked",  // or "Good" or "Suspended"
			"revokedAt": "2024-09-12 16:04:01 GMT+08:00"  // if revoked
		}
	*/

	result, err := getCertBaseInfo(cert)
	if err != nil {
		return nil, err
	}

	// 解析 PEM 根证书
	block, _ := pem.Decode([]byte(rootCACertPEM))
	if block == nil {
		//log.Fatalf("无法解码 PEM 根证书")
		return nil, errors.New("无法解码 PEM 根证书")
	}

	rootCACert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析根证书失败: %v", err))
	}

	httpClient := &http.Client{Timeout: 5 * time.Second}
	ctx := context.Background()
	verifier := newOCSPVerifier(httpClient)

	// 对中间 CA 证书进行 OCSP 验证
	verifyResult, err := verifier.verify(ctx, cert, rootCACert)
	if err != nil {
		return nil, fmt.Errorf("OCSP 验证失败: %v", err)
	}

	if verifyResult["status"] == ocsp.Revoked {
		certStatus := "Unknown"
		if verifyResult["revocationReason"] == 1 {
			certStatus = "Revoked"
		} else {
			certStatus = "Suspended"
		}
		revokedTime := verifyResult["revokedAt"]

		// Load timezone info for Shanghai (UTC+8)
		location, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return nil, err
			//log.Fatalf("Failed to load timezone info: %v", err)
		}

		// Format the revocation time to the Shanghai timezone
		revokedTimeFormatted := revokedTime.(time.Time).In(location).Format("2006-01-02T15:04:05-07:00")
		result["status"] = certStatus
		result["revokedAt"] = revokedTimeFormatted
	} else if verifyResult["status"] == ocsp.Good {
		result["status"] = "Good"
	} else {
		result["status"] = "Unknown"
	}

	return result, nil
}

func getCertBaseInfo(cert *x509.Certificate) (g.Map, error) {
	// Create g.Map to store certificate info
	result := g.Map{}

	// Get certificate information
	certName := cert.Subject.CommonName

	// Load timezone info for Shanghai (UTC+8)
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to load timezone info: %v", err))
	}

	// Convert certificate times to the Shanghai timezone
	validFrom := cert.NotBefore.In(location).Format("2006-01-02T15:04:05-07:00")
	validTo := cert.NotAfter.In(location).Format("2006-01-02T15:04:05-07:00")
	issuer := cert.Issuer.CommonName
	country := cert.Subject.Country[0]
	organization := cert.Subject.Organization[0]
	serialHex := hex.EncodeToString(cert.SerialNumber.Bytes())
	serialDec := cert.SerialNumber.String()

	// Populate result map with the certificate data
	result["name"] = certName
	result["validFrom"] = validFrom
	result["validTo"] = validTo
	result["issuer"] = issuer
	result["country"] = country
	result["organization"] = organization
	result["serialNumberHex"] = serialHex
	result["serialNumberDec"] = serialDec

	return result, nil
}

func getProvisionInfo(base64mp string) (g.Map, error) {
	result := g.Map{}

	data, err := base64.StdEncoding.DecodeString(base64mp)
	if err != nil {
		return nil, err
	}
	start := []byte("<plist")
	startIdx := bytes.Index(data, start)
	if startIdx == -1 {
		return nil, errors.New("描述文件解析错误")
	}
	plistData := data[startIdx:]

	// Parse the plist data
	var mp MobileProvision
	decoder := plist.NewDecoder(bytes.NewReader(plistData))
	err = decoder.Decode(&mp)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析描述文件失败: %v", err))
	}

	result["creationDate"] = mp.CreationDate.Format(time.RFC3339)
	result["expirationDate"] = mp.ExpirationDate.Format(time.RFC3339)
	// Print CreationDate and ExpirationDate
	//fmt.Printf("CreationDate: %s\n", mp.CreationDate.Format(time.RFC3339))
	//fmt.Printf("ExpirationDate: %s\n", mp.ExpirationDate.Format(time.RFC3339))

	// Extract additional data
	developmentTeam := ""
	if len(mp.TeamIdentifier) > 0 {
		developmentTeam = mp.TeamIdentifier[0]
	}
	result["developmentTeam"] = developmentTeam

	applicationIdentifier := ""
	productBundleIdentifier := ""
	if appID, ok := mp.Entitlements["application-identifier"].(string); ok {
		applicationIdentifier = appID
		if parts := strings.SplitN(appID, ".", 2); len(parts) > 1 {
			productBundleIdentifier = parts[1]
		}
	}

	result["uuid"] = mp.UUID
	result["name"] = mp.Name
	result["applicationIdentifier"] = applicationIdentifier
	result["productBundleIdentifier"] = productBundleIdentifier

	// Print extracted data
	//fmt.Printf("DEVELOPMENT_TEAM: %s\n", developmentTeam)
	//fmt.Printf("PROVISIONING_PROFILE: %s\n", mp.UUID)
	//fmt.Printf("PROVISIONING_PROFILE_SPECIFIER: %s\n", mp.Name)
	//fmt.Printf("APPLICATION_IDENTIFIER: %s\n", applicationIdentifier)
	//fmt.Printf("PRODUCT_BUNDLE_IDENTIFIER: %s\n", productBundleIdentifier)

	// Print ProvisionedDevices
	//fmt.Println("Provisioned Devices:")
	devices := make([]string, 0, len(mp.ProvisionedDevices))
	for _, device := range mp.ProvisionedDevices {
		//fmt.Printf("  %s\n", device)
		devices = append(devices, device)
	}
	result["devices"] = devices

	// Parse and Print DeveloperCertificates
	var certificateDetails []CertificateDetails
	fmt.Println("Developer Certificates:")
	for _, certData := range mp.DeveloperCertificates {
		// Try decoding the certificate
		certBlock, _ := pem.Decode(certData)
		if certBlock == nil {
			// If the block is not PEM, try parsing it as a DER certificate
			cert, err := x509.ParseCertificate(certData)
			if err != nil {
				log.Printf("Failed to parse certificate: %v", err)
				continue
			}
			certDetails := extractCertificateDetails(cert)
			certificateDetails = append(certificateDetails, certDetails)
			// Clean up and print the Subject and Issuer
			cleanedSubject := parseSubject(cert.Subject.String())
			cleanedIssuer := parseSubject(cert.Issuer.String())
			fmt.Printf("  Subject: %s\n", cleanedSubject)
			fmt.Printf("  Issuer: %s\n", cleanedIssuer)
		} else {
			// Parse the PEM block
			cert, err := x509.ParseCertificate(certBlock.Bytes)
			if err != nil {
				log.Printf("Failed to parse certificate: %v", err)
				continue
			}
			certDetails := extractCertificateDetails(cert)
			certificateDetails = append(certificateDetails, certDetails)
			// Clean up and print the Subject and Issuer
			cleanedSubject := parseSubject(cert.Subject.String())
			cleanedIssuer := parseSubject(cert.Issuer.String())
			fmt.Printf("  Subject: %s\n", cleanedSubject)
			fmt.Printf("  Issuer: %s\n", cleanedIssuer)
		}
	}
	result["certificates"] = certificateDetails
	return result, nil
}

func parseSubject(subject string) string {
	// Use regex to match OID fields (e.g., 0.9.2342.19200300.100.1.1=#...)
	oidRegex := regexp.MustCompile(`0\.\d+(\.\d+)*=#\w+`)
	return oidRegex.ReplaceAllString(subject, "") // Remove all matches
}

// extractCertificateDetails extracts the necessary details from a certificate.
func extractCertificateDetails(cert *x509.Certificate) CertificateDetails {
	// Extract certificate details
	serialDec := cert.SerialNumber.String()
	serialHex := hex.EncodeToString(cert.SerialNumber.Bytes())
	sha1 := fmt.Sprintf("%x", cert.Signature)

	// Get the creation date (NotBefore)
	creationDate := cert.NotBefore.Format(time.RFC3339)

	// Return the certificate details
	return CertificateDetails{
		Name:            cert.Subject.String(),
		CreationDate:    creationDate,
		SerialNumberDec: serialDec,
		SerialNumberHex: serialHex,
		SHA1:            sha1,
	}
}

// 从指定 URL 获取 CA 证书
func fetchCACert(httpClient *http.Client, url string) (*x509.Certificate, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("无法获取证书: %w", err)
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取证书内容失败: %w", err)
	}

	cert, err := x509.ParseCertificate(buf)
	if err != nil {
		return nil, fmt.Errorf("解析证书失败: %w", err)
	}

	if !cert.IsCA {
		return nil, errors.New("下载的证书不是 CA 证书")
	}

	return cert, nil
}

type ocspVerifier struct {
	client *http.Client
}

func newOCSPVerifier(httpClient *http.Client) *ocspVerifier {
	return &ocspVerifier{
		client: httpClient,
	}
}

func (v *ocspVerifier) verify(ctx context.Context, cert, issuer *x509.Certificate) (result g.Map, err error) {
	result = g.Map{}
	if len(cert.OCSPServer) == 0 {
		return nil, errors.New("证书中未包含 OCSP 服务器地址")
	}
	ocspServer := cert.OCSPServer[0]
	// log.Println("OCSP 服务器地址:", ocspServer)

	// 创建 OCSP 请求
	ocspRequest, err := ocsp.CreateRequest(cert, issuer, &ocsp.RequestOptions{Hash: crypto.SHA1})
	if err != nil {
		return nil, fmt.Errorf("创建 OCSP 请求失败: %w", err)
	}

	// 发送 OCSP 请求
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ocspServer, bytes.NewReader(ocspRequest))
	if err != nil {
		return nil, fmt.Errorf("创建 HTTP 请求失败: %w", err)
	}
	req.Header.Set("Content-Type", "application/ocsp-request")
	req.Header.Set("Accept", "application/ocsp-response")
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := v.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送 OCSP 请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OCSP 服务器返回非 200 状态码: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取 OCSP 响应失败: %w", err)
	}

	// 解析并验证 OCSP 响应
	ocspResponse, err := ocsp.ParseResponse(body, issuer)
	if err != nil {
		return nil, fmt.Errorf("解析 OCSP 响应失败: %w", err)
	}

	// 检查 OCSP 响应状态
	// log.Println("证书状态为", ocspResponse.Status)
	result["status"] = ocspResponse.Status

	if ocspResponse.Status == ocsp.Revoked {
		// log.Printf("证书已被撤销，撤销时间: %v，撤销原因: %v", ocspResponse.RevokedAt, ocspResponse.RevocationReason)
		result["revokedAt"] = ocspResponse.RevokedAt
		result["revocationReason"] = ocspResponse.RevocationReason
	} else if ocspResponse.Status == ocsp.Good {
		// log.Println("证书状态为 Good")
	} else {
		log.Printf("未知的证书状态: %v", ocspResponse.Status)
	}

	// 检查 OCSP 响应是否过期
	now := time.Now()
	if ocspResponse.NextUpdate.Before(now) {
		return result, errors.New("OCSP 响应已过期")
	}

	return result, nil
}
