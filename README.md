# 米哈游相关api

米游社app的api封装的go实现

## 说明

支持米游社获取帖子数据,可实现自动化做米游币任务,原神相关api的实现等

## 包说明

## cookies 包

> 用于存储登陆后的cookies,大部分情况下是内部使用

```go
package main

import (
	"fmt"
	"github.com/Huiyicc/mhyapi/cookies"
)

func main() {
	cookiesStr := "cookiesStr"
	cookiesCore,err := cookies.NewCookiesCore(cookiesStr)
	if err != nil {
        panic(err)
	}
	fmt.Println(cookiesCore.GetStr())
}
```

## define 包

> 定义了全局常量


> 用于 AppCore.GetPostsList()

| 常量名            | 值   | 备注           |
|----------------|-----|--------------|
| FORUMID_GENSHI | 26  | 分区ID:原神      |
| FORUMID_BH3    | 1   | 分区ID:崩坏3     |
| FORUMID_BH2    | 30  | 分区ID:崩坏2     |
| FORUMID_WD     | 37  | 分区ID:未定事件簿   |
| FORUMID_DBY    | 34  | 分区ID:大别野     |
| FORUMID_SR     | 52  | 分区ID:崩坏：星穹铁道 |
| FORUMID_ZZZ    | 57  | 分区ID:绝区零     |

> 米游币任务ID

| 常量名                             | 值   | 备注            |
|---------------------------------|-----|---------------|
| TASKS_MISSION_ID_BBS_SIGN       | 58  | 米游币任务ID:讨论区签到 |
| TASKS_MISSION_ID_BBS_READ_POSTS | 59  | 米游币任务ID:看帖子   |
| TASKS_MISSION_ID_BBS_LIKE_POSTS | 60  | 米游币任务ID:给帖子点赞 |
| TASKS_MISSION_ID_BBS_SHARE      | 61  | 米游币任务ID:分享帖子  |

> 原神服务器类型

| 常量名                             | 值       | 备注               |
|---------------------------------|---------|------------------|
| GAMERSERVER_GENSHIN_TIANKONGDAO | cn_gf01 | 原神服务器类型:天空岛(官服)  |
| GAMERSERVER_GENSHIN_SHIJIESHU   | cn_qd01 | 原神服务器类型:世界树(渠道服) |

## genshin 包

> 定义了原神相关api


| 父结构         | 方法名                | 注释                        |
|-------------|--------------------|---------------------------|
| -           | NewCore            | 使用coookies创建一个GenShinCore |
| GenShinCore | GetBindInfo        | 用于获取绑定的角色信息               |
| GenShinCore | UpdateBindInfo     | 用于更新绑定的角色信息               |
| GenShinCore | GetNoteInfo        | 获取游戏体力信息                  |
| GenShinCore | Sign               | 米游社内游戏签到                  |
| GenShinCore | GetSignRewardsList | 获取签到奖励列表                  |
| GenShinCore | SignInfo           | 获取信息                      |
| GenShinCore | GameRecordIndex    | 获取宝箱,探索度,声望等数据            |
| GenShinCore | GetSpiralAbyssInfo | 获取深渊数据                    |


#### 初始化

```go
//先登录
app := mhyapp.AppCore{}
if err := app.LoginToCookies(userCookiesStr); err != nil {
    panic(err)
}
appCookiesStr := app.Cookies.GetStr()

//appCookiesStr为使用mhyapp包登陆之后的cookies
gameCore, err := genshin.NewCore(appCookiesStr)
if err != nil {
	panic(err)
}
```

## mhyapp 包

> 定义了米游币等相关api

### 方法列表


| 父结构     | 方法名                      | 注释                                                          |
|---------|--------------------------|-------------------------------------------------------------|
| AppCore | Login                    | 将user.mihoyo.com的cookies换成token<br />登陆后的cookies应自行保存,可重复使用 |
| AppCore | LoginToCookies           | 与Login方法一样,区别是需要自己给参数                                       |
| AppCore | SetCookies               | 直接设置内部cookies,一般用以跳过登陆阶段                                    |
| AppCore | GetTasksInfo             | 获取已做过的米游币任务,当前米游币,今日已获得米游币数量等                               |
| AppCore | GetTasksIncompleteIDList | 获取未完成的米游币每日任务列表                                             |
| AppCore | GetPostsList             | 用于获取某分区帖子列表                                                 |
| AppCore | PostDetail               | 看帖                                                          |
| AppCore | PostVote                 | 帖子点/取消赞                                                     |
| AppCore | PostShare                | 帖子分享                                                        |
| AppCore | BBSSign                  | 指定板块签到                                                      |

> cookies登陆 https://user.mihoyo.com/ 获取

```go
app := mhyapp.AppCore{}
if err := app.LoginToCookies(userCookiesStr); err != nil {
	panic(err)
}
//自行保存cookies
fmt.Println(app.Cookies.GetStr())
```

## request 包

> 定义了内部请求封装函数

## tools 包

> 内部杂项函数

## 更新中
