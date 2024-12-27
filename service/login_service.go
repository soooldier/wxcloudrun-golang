package service

import (
	"encoding/json"
	"net/http"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
)

var (
	appId     = "wx9615e00ae1398eda"
	appSecret = "eaf1ea75602bb17551f789e63e04c1e2"
	token     = "TOKEN"
)

// LoginHandler 微信登录
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	user, err := GetWechatOfficialNickname(appId, appSecret, token, code, cache.NewMemory())
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	user4json, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(user4json)
}

func GetWechatOfficialNickname(appId, appSecret, token, code string, cache cache.Cache) (userinfo oauth.UserInfo, err error) {
	cfg := &config.Config{
		AppID:     appId,
		AppSecret: appSecret,
		Token:     token,
		Cache:     cache,
	}
	wc := wechat.NewWechat()
	officialAccount := wc.GetOfficialAccount(cfg)
	oauth := officialAccount.GetOauth()
	accessToken, err := oauth.GetUserAccessToken(code)
	if err != nil {
		return
	}
	userinfo, err = oauth.GetUserInfo(accessToken.AccessToken, accessToken.OpenID, "")
	if err != nil {
		return
	}
	return
}