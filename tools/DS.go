package tools

import (
	"fmt"
	"github.com/Huiyicc/mhyapi/define"
	"math/rand"
	"time"
)

// GetDs 用于获取DS值，请求需要
func GetDs(web bool) string {
	n := Ifs(web, define.BBS_STALT_WEB, define.BBS_STALT_A)
	t := time.Now().Unix()
	r := RandString(6)
	c := GetMd5(fmt.Sprintf("salt=%s&t=%d&r=%s", n, t, r))
	return fmt.Sprintf("%d,%s,%s", t, r, c)
}

// GetDs2 用于获取DS值，请求需要
func GetDs2(q, b string) string {
	n := define.BBS_STALT_B
	t := time.Now().Unix()
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(99999) + 100000
	c := GetMd5(fmt.Sprintf("salt=%s&t=%d&r=%d&b=%s&q=%s", n, t, r, b, q))
	return fmt.Sprintf("%d,%d,%s", t, r, c)
}
