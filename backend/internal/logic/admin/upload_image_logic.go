package admin

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"ms_tmdb/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	maxUploadImageSize = 10 << 20
	uploadDirName      = "uploads"
)

var allowedImageExt = map[string]struct{}{
	".jpg":  {},
	".jpeg": {},
	".png":  {},
	".webp": {},
	".gif":  {},
}

type UploadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageLogic {
	return &UploadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImageLogic) UploadImage(file multipart.File, header *multipart.FileHeader) (string, error) {
	if header == nil {
		return "", errors.New("文件不能为空")
	}
	if header.Size > maxUploadImageSize {
		return "", errors.New("图片大小不能超过 10MB")
	}

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if _, ok := allowedImageExt[ext]; !ok {
		return "", errors.New("仅支持 jpg/jpeg/png/webp/gif 图片")
	}

	contentType := strings.TrimSpace(header.Header.Get("Content-Type"))
	if contentType != "" && !strings.HasPrefix(strings.ToLower(contentType), "image/") {
		return "", errors.New("仅支持图片文件上传")
	}

	uploadDir := filepath.Join(".", uploadDirName)
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	savePath := filepath.Join(uploadDir, fileName)

	dst, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return "/uploads/" + fileName, nil
}
