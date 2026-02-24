package admin

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUploadedFileHandler(_ *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		fileName := strings.TrimSpace(req.Filename)
		if fileName == "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("文件名不能为空"))
			return
		}
		if strings.Contains(fileName, "/") || strings.Contains(fileName, "\\") {
			httpx.ErrorCtx(r.Context(), w, errors.New("非法文件名"))
			return
		}
		if filepath.Base(fileName) != fileName {
			httpx.ErrorCtx(r.Context(), w, errors.New("非法文件名"))
			return
		}

		fullPath := filepath.Join(".", "uploads", fileName)
		if _, err := os.Stat(fullPath); err != nil {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, fullPath)
	}
}
