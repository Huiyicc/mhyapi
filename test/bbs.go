package main

import (
	_ "embed"
	"fmt"
	"github.com/Huiyicc/mhyapi/genshin"
	"github.com/Huiyicc/mhyapi/mhyapp"
)

//go:embed usercookies.p
var userCookiesStr string

//go:embed appcookies.p
var appCookiesStr string

func main() {
	//test_login()
	//test_posts()
	test_game()
}

func test_login() {
	app := mhyapp.AppCore{}
	if err := app.LoginToCookies(userCookiesStr); err != nil {
		panic(err)
	}
	fmt.Println(app.Cookies.GetStr())
}

func test_posts() {
	var (
		app mhyapp.AppCore
		err error
	)
	//登录
	if err = app.SetCookies(appCookiesStr); err != nil {
		panic(err)
	}
	/* //获取帖子列表
	var list []mhyapp.AppForumInfo
	if list, err = app.GetPostsList(define.FORUMID_GENSHI, 10); err != nil {
		panic(err)
	}
	fmt.Println(list)
	*/
	/* //看帖
	var postsInfo *mhyapp.AppForumInfo
	if postsInfo, err = app.PostDetail("28520829"); err != nil {
		panic(err)
	}
	fmt.Println(postsInfo)
	*/

}
func test_game() {
	gameCore, err := genshin.NewCore(appCookiesStr)
	if err != nil {
		panic(err)
	}
	gameCore.GameRecordIndex("112075042")
}
