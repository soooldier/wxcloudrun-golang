package service

import (
	"net/http"
)

// LoginHandler 微信登录
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	w.Write([]byte(code))
}
