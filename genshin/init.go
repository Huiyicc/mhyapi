package genshin

import (
	"github.com/Huiyicc/mhyapi/cookies"
	"github.com/Huiyicc/mhyapi/define"
	"github.com/Huiyicc/mhyapi/tools"
	"math/rand"
	"net/http"
	"time"
)

func NewCore(cookie string) (*GenShinCore, error) {
	var (
		err error
		rt  GenShinCore
		ck  string
	)
	if rt.Cookies, err = cookies.NewCookiesCore(cookie); err != nil {
		return nil, err
	}
	if ck, err = rt.UpdateCookiesToken(); err != nil {
		return nil, err
	}
	rt.Cookies.Set("cookie_token", ck)
	if _, err = rt.UpdateBindInfo(); err != nil {
		return nil, err
	}
	return &rt, nil
}

type GenShinCore struct {
	Cookies  *cookies.CookiesCore
	headers  http.Header
	GameInfo GameInfo
}

func (t *GenShinCore) getHeaders() http.Header {
	t.updateHeader()
	return t.headers
}

func (t *GenShinCore) updateHeader() {
	if t.headers == nil {
		t.headers = make(map[string][]string)
	}
	rand.Seed(time.Now().UnixNano())
	t.headers["Accept"] = []string{"application/json, text/plain, */*"}
	t.headers["DS"] = []string{tools.GetDs(true)}
	t.headers["x-rpc-channel"] = []string{"miyousheluodi"}
	t.headers["Origin"] = []string{"https://webstatic.mihoyo.com"}
	t.headers["x-rpc-app_version"] = []string{define.APPCLIENT_VERSIONS}
	t.headers["User-Agent"] = []string{"Mozilla/5.0 (Linux; Android 12; Unspecified Device) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/103.0.5060.129 Mobile Safari/537.36 miHoYoBBS/" + define.APPCLIENT_VERSIONS}
	t.headers["x-rpc-client_type"] = []string{define.APPCLIENT_TYPE_WEBMOBILE}
	t.headers["X-Requested-With"] = []string{"com.mihoyo.hyperion"}
	t.headers["x-rpc-device_id"] = []string{tools.GetDevicesID()}
	t.headers["Cookie"] = []string{t.Cookies.GetStr()}
	t.headers["Referer"] = []string{"https://webstatic.mihoyo.com/bbs/event/signin-ys/index.html?bbs_auth_required=true&act_id=" + define.ACTID_GENSHIN + "&utm_source=bbs&utm_medium=mys&utm_campaign=icon"}
}

func (t *GenShinCore) getGameHeaders(q, b string) http.Header {
	headers := t.getHeaders().Clone()
	headers["Accept"] = []string{"*/*"}
	headers["DS"] = []string{tools.GetDs2(q, b)}
	headers["x-rpc-client_type"] = []string{define.APPCLIENT_TYPE_ANDROID}
	headers["x-rpc-app_version"] = []string{define.APPCLIENT_VERSIONS}
	headers["x-rpc-sys_version"] = []string{"6.0.1"}
	headers["x-rpc-channel"] = []string{"miyousheluodi"}
	headers["x-rpc-device_id"] = []string{tools.GetDevicesID()}
	headers["x-rpc-device_name"] = []string{tools.RandString(rand.Intn(5) + 5)}
	headers["x-rpc-device_model"] = []string{"Mi 10"}
	headers["Referer"] = []string{"https://app.mihoyo.com"}
	headers["User-Agent"] = []string{"okhttp/4.8.0"}
	headers["cookie"] = []string{t.Cookies.GetStr()}
	headers["Content-Type"] = []string{"application/json"}
	return headers
}
