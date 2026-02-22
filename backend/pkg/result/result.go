package result

import (
	"encoding/json"
	"net/http"

	"ms_tmdb/xerr"
)

// Response 统一响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// OkJSON 直接写入原始 JSON（兼容 TMDB 格式）
func OkJSON(w http.ResponseWriter, data json.RawMessage) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// Ok 成功响应
func Ok(w http.ResponseWriter, data interface{}) {
	resp := &Response{Code: xerr.OK, Msg: "ok", Data: data}
	writeJSON(w, http.StatusOK, resp)
}

// Error 错误响应
func Error(w http.ResponseWriter, err error) {
	code := xerr.ServerError
	msg := err.Error()

	if e, ok := err.(*xerr.CodeError); ok {
		code = e.Code
		msg = e.Msg
	}

	resp := &Response{Code: code, Msg: msg}
	writeJSON(w, http.StatusOK, resp)
}

func writeJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}
