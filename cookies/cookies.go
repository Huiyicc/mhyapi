package cookies

import (
	"errors"
	"fmt"
	"strings"
)

type CookiesCore struct {
	KeysMap map[string]string
}

// NewCookiesCore 用于从cookies文本创建一个对象
func NewCookiesCore(cookies string) (*CookiesCore, error) {
	r := &CookiesCore{}
	return r, r.Parse(cookies)
}

func (t *CookiesCore) Parse(cookies string) error {
	t.KeysMap = make(map[string]string)
	list := strings.Split(cookies, "; ")
	for _, v := range list {
		index := strings.Index(v, "=")
		key := v[:index]
		value := v[index+1:]
		t.KeysMap[key] = value
	}
	if t.KeysMap["login_ticket"] == "" {
		return errors.New("cookies不完整")
	}
	return nil
}

// GetStr 用于获取所有cookies拼接成的文本
func (t *CookiesCore) GetStr() string {
	var r string
	for k, v := range t.KeysMap {
		r += k + "=" + v + "; "
	}
	return r
}

func (t *CookiesCore) GetStoken() string {
	return fmt.Sprintf("stoken=%s; stuid=%s;", t.KeysMap["stoken"], t.KeysMap["stoken"])
}

func (t *CookiesCore) Set(key, value string) {
	if t.KeysMap == nil {
		t.KeysMap = make(map[string]string)
	}
	t.KeysMap[key] = value
}

func (t *CookiesCore) Get(key string) string {
	return t.KeysMap[key]
}
