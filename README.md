# 米哈游相关api

米游社app的api封装的go实现

## 说明

支持米游社获取帖子数据,可实现自动化做米游币任务,原神相关api的实现等

## 包说明

### cookies 包

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

### define 包

> 定义了全局常量

### genshin 包

> 定义了原神相关api

### mhyapp 包

> 定义了米游币等相关api

### request 包

> 定义了内部请求封装函数

### tools 包

> 内部杂项函数

### 未完成
