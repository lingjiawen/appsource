/*
* @desc:上传处理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/28 9:37
 */

package upload

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"io"
	"mangosmithy/internal/app/common/consts"
	"mangosmithy/internal/app/common/model"
	"mangosmithy/internal/app/common/model/entity"
	"mangosmithy/internal/app/common/service"
	"mangosmithy/library/libUtils"
	"mangosmithy/library/liberr"
	"mangosmithy/library/upload"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	service.RegisterUpload(New())
}

func New() service.IUpload {
	return &sUpload{}
}

type sUpload struct{}

// UploadFiles 上传多文件
func (s *sUpload) UploadFiles(ctx context.Context, files []*ghttp.UploadFile, checkFileType string, source int, userId uint64, appId string) (result []*model.UploadResponse, err error) {
	for _, item := range files {
		f, e := s.UploadFile(ctx, item, checkFileType, source, userId, appId)
		if e != nil {
			return
		}
		result = append(result, f)
	}
	return
}

// UploadFile 上传单文件
func (s *sUpload) UploadFile(ctx context.Context,
	file *ghttp.UploadFile, checkFileType string, source int,
	userId uint64, appId string) (result *model.UploadResponse, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 检查文件类型
		err = s.CheckType(ctx, checkFileType, file.Filename)
		liberr.ErrIsNil(ctx, err)

		// 检查文件大小
		err = s.CheckSize(ctx, checkFileType, file.Size)
		liberr.ErrIsNil(ctx, err)
		//判断该文件是否已经上传过，上传过则不再上传
		var (
			md5        = ""
			existsFile *model.SysAttachmentInfoRes
		)
		md5, err = s.computeMD5(file)
		liberr.ErrIsNil(ctx, err)
		existsFile, err = service.SysAttachment().GetByMd5(ctx, md5)
		liberr.ErrIsNil(ctx, err, "获取文件信息失败")
		if existsFile != nil {
			result = &model.UploadResponse{
				Size:     existsFile.Size,
				Path:     existsFile.Path,
				FullPath: libUtils.GetDomain(ctx, true) + "/" + existsFile.Path,
				Name:     existsFile.Name,
				Type:     existsFile.MimeType,
			}
			return
		}
		uploader := upload.GetUploader(upload.UploaderType(source))
		if uploader == nil {
			liberr.ErrIsNil(ctx, errors.New("没有找到上传适配器"))
		}
		result, err = uploader.Upload(ctx, file)
		liberr.ErrIsNil(ctx, err)
		//保存上传文件到数据库
		err = service.SysAttachment().AddUpload(ctx, result, &model.SysAttachmentAddAttribute{
			Md5:    md5,
			Driver: gconv.Uint(source),
			UserId: userId,
			AppId:  appId,
		})
		liberr.ErrIsNil(ctx, err)
	})
	return
}

// CheckSize 检查上传文件大小
func (s *sUpload) CheckSize(ctx context.Context, checkFileType string, fileSize int64) (err error) {

	var (
		configSize *entity.SysConfig
	)

	if checkFileType == consts.CheckFileTypeFile {

		//获取上传大小配置
		configSize, err = s.getUpConfig(ctx, consts.FileSizeKey)
		if err != nil {
			return
		}
	} else if checkFileType == consts.CheckFileTypeImg {

		//获取上传大小配置
		configSize, err = s.getUpConfig(ctx, consts.ImgSizeKey)
		if err != nil {
			return
		}
	} else {
		return errors.New(fmt.Sprintf("文件检查类型错误:%s|%s", consts.CheckFileTypeFile, consts.CheckFileTypeImg))
	}

	var rightSize bool
	rightSize, err = s.checkSize(configSize.ConfigValue, fileSize)
	if err != nil {
		return
	}
	if !rightSize {
		err = gerror.New("上传文件超过最大尺寸：" + configSize.ConfigValue)
		return
	}
	return
}

// CheckType 检查上传文件类型
func (s *sUpload) CheckType(ctx context.Context, checkFileType string, fileName string) (err error) {

	var (
		configType *entity.SysConfig
	)

	if checkFileType == consts.CheckFileTypeFile {
		//获取上传类型配置
		configType, err = s.getUpConfig(ctx, consts.FileTypeKey)
		if err != nil {
			return
		}

	} else if checkFileType == consts.CheckFileTypeImg {
		//获取上传类型配置
		configType, err = s.getUpConfig(ctx, consts.ImgTypeKey)
		if err != nil {
			return
		}
	} else {
		return errors.New(fmt.Sprintf("文件检查类型错误:%s|%s", consts.CheckFileTypeFile, consts.CheckFileTypeImg))
	}

	rightType := s.checkFileType(fileName, configType.ConfigValue)
	if !rightType {
		err = gerror.New("上传文件类型错误，只能包含后缀为：" + configType.ConfigValue + "的文件。")
		return
	}
	return
}

// 获取上传配置
func (s *sUpload) getUpConfig(ctx context.Context, key string) (config *entity.SysConfig, err error) {
	config, err = service.SysConfig().GetConfigByKey(ctx, key)
	if err != nil {
		return
	}
	if config == nil {
		err = gerror.New("上传文件类型未设置，请在后台配置")
		return
	}
	return
}

// 判断上传文件类型是否合法
func (s *sUpload) checkFileType(fileName, typeString string) bool {
	suffix := gstr.SubStrRune(fileName, gstr.PosRRune(fileName, ".")+1, gstr.LenRune(fileName)-1)
	imageType := gstr.Split(typeString, ",")
	rightType := false
	for _, v := range imageType {
		if gstr.Equal(suffix, v) {
			rightType = true
			break
		}
	}
	return rightType
}

// 检查文件大小是否合法
func (s *sUpload) checkSize(configSize string, fileSize int64) (bool, error) {
	match, err := gregex.MatchString(`^([0-9]+)(?i:([a-z]*))$`, configSize)
	if err != nil {
		return false, err
	}
	if len(match) == 0 {
		err = gerror.New("上传文件大小未设置，请在后台配置，格式为（30M,30k,30MB）")
		return false, err
	}
	var cfSize int64
	switch gstr.ToUpper(match[2]) {
	case "MB", "M":
		cfSize = gconv.Int64(match[1]) * 1024 * 1024
	case "KB", "K":
		cfSize = gconv.Int64(match[1]) * 1024
	case "":
		cfSize = gconv.Int64(match[1])
	}
	if cfSize == 0 {
		err = gerror.New("上传文件大小未设置，请在后台配置，格式为（30M,30k,30MB），最大单位为MB")
		return false, err
	}
	return cfSize >= fileSize, nil
}

// 静态文件夹目录
func (s *sUpload) getStaticPath(ctx context.Context) string {
	value, _ := g.Cfg().Get(ctx, "server.serverRoot")
	if !value.IsEmpty() {
		return value.String()
	}
	return ""
}

// computeMD5 计算给定文件的 MD5 值
func (s *sUpload) computeMD5(upFile *ghttp.UploadFile) (string, error) {
	file, err := upFile.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 创建一个新的 MD5 哈希对象
	hash := md5.New()

	// 逐块读取文件内容并更新哈希
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// 将 MD5 值转换为十六进制字符串
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (s *sUpload) CheckMultipart(ctx context.Context, req *model.CheckMultipartReq) (res *model.CheckMultipartRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 检查文件类型
		err = s.CheckType(ctx, req.UploadType, req.FileName)
		liberr.ErrIsNil(ctx, err)

		// 检查文件大小
		err = s.CheckSize(ctx, req.UploadType, req.Size)
		liberr.ErrIsNil(ctx, err)
		//检查文件是否已经上传
		var existsFile *model.SysAttachmentInfoRes
		existsFile, err = service.SysAttachment().GetByMd5(ctx, req.Md5)
		liberr.ErrIsNil(ctx, err, "获取文件信息失败")
		res = new(model.CheckMultipartRes)
		if existsFile != nil {
			res.Attachment = &model.UploadResponse{
				Size:     existsFile.Size,
				Path:     existsFile.Path,
				FullPath: libUtils.GetDomain(ctx, true) + "/" + existsFile.Path,
				Name:     existsFile.Name,
				Type:     existsFile.MimeType,
			}
			return
		}

		uploader := upload.GetUploader(upload.UploaderType(req.DriverType))
		if uploader == nil {
			liberr.ErrIsNil(ctx, errors.New("没有找到上传适配器"))
		}
		res, err = uploader.CheckMultipart(ctx, req)
		liberr.ErrIsNil(ctx, err)
	})
	return
}

func (s *sUpload) UploadPart(ctx context.Context, req *model.UploadPartReq) (res *model.UploadPartRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		uploader := upload.GetUploader(upload.UploaderType(req.DriverType))
		if uploader == nil {
			liberr.ErrIsNil(ctx, errors.New("没有找到上传适配器"))
		}
		res, err = uploader.UploadPart(ctx, req)
		liberr.ErrIsNil(ctx, err)
		//保存上传文件到数据库
		if res != nil && res.Finish {
			err = service.SysAttachment().AddUpload(ctx, res.Attachment, &model.SysAttachmentAddAttribute{
				Md5:    req.Md5,
				Driver: gconv.Uint(req.DriverType),
				UserId: req.UserId,
				AppId:  req.AppId,
			})
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

// UploadFileInner 内部文件上传方法
func (s *sUpload) UploadFileInner(ctx context.Context, fileReader io.Reader, filename string, checkFileType string, source int, userId uint64, appId string) (result *model.UploadResponse, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 检查文件类型
		err = s.CheckType(ctx, checkFileType, filename)
		liberr.ErrIsNil(ctx, err)

		// 获取文件大小和 MD5，同时缓存文件内容
		fileSize, md5FromReader, buffer, err := getFileSizeAndMD5(fileReader)
		if err != nil {
			return
		}

		// 检查文件是否已经上传过
		existsFile, err := service.SysAttachment().GetByMd5(ctx, md5FromReader)
		liberr.ErrIsNil(ctx, err, "获取文件信息失败")
		if existsFile != nil {
			result = &model.UploadResponse{
				Size:     existsFile.Size,
				Path:     existsFile.Path,
				FullPath: libUtils.GetDomain(ctx, true) + "/" + existsFile.Path,
				Name:     existsFile.Name,
				Type:     existsFile.MimeType,
			}
			return
		}

		// 调用 Local 上传方法，使用缓存的 buffer 作为文件内容
		result, err = localUpload(ctx, buffer, filename, fileSize)
		liberr.ErrIsNil(ctx, err)

		// 保存上传文件到数据库
		err = service.SysAttachment().AddUpload(ctx, result, &model.SysAttachmentAddAttribute{
			Md5:    md5FromReader,
			Driver: gconv.Uint(source),
			UserId: userId,
			AppId:  appId,
		})
		liberr.ErrIsNil(ctx, err)
	})
	return
}

// computeMD5FromReader 计算文件的 MD5（通过 io.Reader）
func (s *sUpload) computeMD5FromReader(reader io.Reader) (string, error) {
	hash := md5.New()
	_, err := io.Copy(hash, reader)
	if err != nil {
		return "", fmt.Errorf("failed to compute MD5: %v", err)
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func getFileSizeAndMD5(fileReader io.Reader) (fileSize int64, md5Checksum string, buffer *bytes.Buffer, err error) {
	// 使用一个 bytes.Buffer 缓存读取的数据
	buffer = &bytes.Buffer{}

	// 创建一个 TeeReader，它会同时将数据写入 buffer 和计算 MD5
	hash := md5.New()
	teeReader := io.TeeReader(fileReader, io.MultiWriter(hash, buffer))

	// 读取数据并计算文件大小
	_, err = io.Copy(io.Discard, teeReader) // 使用 io.Discard 丢弃数据，只关心读取
	if err != nil {
		return 0, "", nil, err
	}

	// 获取 MD5 校验码
	md5Checksum = fmt.Sprintf("%x", hash.Sum(nil))
	fileSize = int64(buffer.Len()) // 获取文件大小（从缓冲区的长度）
	return
}

func localUpload(ctx context.Context, fileReader io.Reader, filename string, fileSize int64) (result *model.UploadResponse, err error) {
	if fileReader == nil {
		err = errors.New("文件不能为空")
		return
	}

	// 用于存储上传文件的路径和其它信息
	urlPrefix := libUtils.GetDomain(ctx, true)
	p := strings.Trim(consts.UploadPath, "/")
	value, _ := g.Cfg().Get(ctx, "server.serverRoot")
	sp := value.String()

	if sp != "" {
		sp = strings.TrimRight(sp, "/")
	}
	nowData := time.Now().Format("2006-01-02")
	// 包含静态文件夹的路径
	fullDirPath := sp + "/" + p + "/" + nowData
	filename = gfile.Basename(filename)
	name := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	fileName := name + gfile.Ext(filename)
	fullPath := p + "/" + nowData + "/" + fileName

	if !gfile.Exists(fullDirPath) {
		if err = gfile.Mkdir(fullDirPath); err != nil {
			return nil, err
		}
	} else if !gfile.IsDir(fullDirPath) {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, `parameter "fullDirPath" should be a directory path`)
	}

	// 保存文件
	file, err := os.Create(fullDirPath + "/" + fileName)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			err = fmt.Errorf("failed to close file: %v", err)
		}
	}(file)

	// 将文件内容写入磁盘
	_, err = io.Copy(file, fileReader)
	if err != nil {
		return nil, fmt.Errorf("failed to save file: %v", err)
	}

	result = &model.UploadResponse{
		Size:     fileSize,
		Path:     fullPath,
		FullPath: urlPrefix + "/" + fullPath,
		Name:     filename,
		Type:     "application/octet-stream", // 或者根据需要设置 MIME 类型
	}
	return
}
