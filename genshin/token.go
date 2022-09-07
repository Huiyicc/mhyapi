package genshin

import (
	"errors"
	"github.com/Huiyicc/mhyapi/define"
	"github.com/Huiyicc/mhyapi/request"
	"github.com/Huiyicc/mhyapi/tools"
	json "github.com/json-iterator/go"
)

// UpdateCookiesToken 更新cookies_token,返回获取到的cookies_token
func (t *GenShinCore) UpdateCookiesToken() (string, error) {
	req := request.MIHOYOAPP_API_COOKIESTOKEN.Copy()
	headers := t.getHeaders().Clone()
	headers["DS"] = []string{tools.GetDs(false)}
	headers["x-rpc-client_type"] = []string{define.APPCLIENT_TYPE_ANDROID}
	headers.Del("Referer")
	headers["User-Agent"] = []string{"okhttp/4.8.0"}
	data, err := request.HttpGet(req, headers)
	if err != nil {
		return "", err
	}
	var resp updateCookiesTokenResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return "", errors.New(string(data))
	}
	if err = resp.verify(); err != nil {
		return "", err
	}
	return resp.Data.CookieToken, nil
}

type updateCookiesTokenResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Uid         string `json:"uid"`
		CookieToken string `json:"cookie_token"`
	} `json:"data"`
}

func (t *updateCookiesTokenResponse) verify() error {
	return tools.Ifs(t.Retcode == 0, nil, errors.New(t.Message))
}
