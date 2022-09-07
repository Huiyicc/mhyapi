package mhyapp

import (
	"errors"
	"fmt"
	"github.com/Huiyicc/mhyapi/cookies"
	"github.com/Huiyicc/mhyapi/request"
	"github.com/Huiyicc/mhyapi/tools"
	json "github.com/json-iterator/go"
	"strconv"
)

// Login 用于登陆app,需要实现初始化cookies
func (t *AppCore) Login() error {
	if err := t.login1(); err != nil {
		return err
	}
	if err := t.login2(); err != nil {
		return err
	}
	t.Cookies.Set("stuid", t.Cookies.Get("account_id"))
	return nil
}

// LoginToCookies 使用cookies登陆app
func (t *AppCore) LoginToCookies(str string) error {
	var err error
	t.Cookies, err = cookies.NewCookiesCore(str)
	if err != nil {
		return err
	}
	return t.Login()
}

// SetCookies 设置内部cookies
func (t *AppCore) SetCookies(str string) error {
	var err error
	t.Cookies, err = cookies.NewCookiesCore(str)
	return err
}

// 第一阶段登陆
func (t *AppCore) login1() error {
	req := request.MIHOYOAPP_API_LOGINA.Copy()
	req.Query = fmt.Sprintf(req.Query, t.Cookies.Get("login_ticket"))
	data, err := request.HttpGet(req, nil)
	if err != nil {
		return err
	}
	var resp loginARequest
	if err = json.Unmarshal(data, &resp); err != nil {
		return errors.New(string(data))
	}
	if err = resp.verify(); err != nil {
		return err
	}
	if resp.Data.CookieInfo.AccountId == 0 {
		return errors.New("服务器错误")
	}
	t.Cookies.Set("account_id", strconv.Itoa(resp.Data.CookieInfo.AccountId))
	return nil
}

type loginARequest struct {
	Code int `json:"code"`
	Data struct {
		CookieInfo struct {
			AccountId   int    `json:"account_id"`
			CookieToken string `json:"cookie_token"`
			CreateTime  int    `json:"create_time"`
			CurTime     int    `json:"cur_time"`
			Email       string `json:"email"`
			IsAdult     int    `json:"is_adult"`
			IsRealname  int    `json:"is_realname"`
			Mobile      string `json:"mobile"`
		} `json:"cookie_info"`
		Msg    string `json:"msg"`
		Sign   string `json:"sign"`
		Status int    `json:"status"`
	} `json:"data"`
}

//验证器
func (t *loginARequest) verify() error {
	if t.Code != 200 {
		return errors.New(t.Data.Msg)
	}
	if t.Data.Status != 1 {
		return errors.New(t.Data.Msg)
	}
	return nil
}

// login2 第二阶段登陆
func (t *AppCore) login2() error {
	req := request.MIHOYOAPP_API_LOGINB.Copy()
	req.Query = fmt.Sprintf(req.Query, t.Cookies.Get("login_ticket"), t.Cookies.Get("account_id"))
	data, err := request.HttpGet(req, nil)
	if err != nil {
		return err
	}
	var resp loginBRequest
	if err = json.Unmarshal(data, &resp); err != nil {
		return errors.New(string(data))
	}
	if err = resp.verify(); err != nil {
		return err
	}
	for _, v := range resp.Data.List {
		t.Cookies.Set(v.Name, v.Token)
	}
	return nil
}

type loginBRequest struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			Name  string `json:"name"`
			Token string `json:"token"`
		} `json:"list"`
	} `json:"data"`
}

//验证器
func (t *loginBRequest) verify() error {
	return tools.Ifs(t.Retcode == 0, nil, errors.New(t.Message))
}
