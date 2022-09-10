package main

import (
	"github.com/Huiyicc/mhyapi/genshin/EnkaNetwork"
)

func main() {
	test_Enka()
}

func test_Enka() {
	enka, err := EnkaNetwork.NewEnkaCore("./cache", EnkaNetwork.NODE_URL_CN1)
	if err != nil {
		panic(err)
	}
	var rawMap EnkaNetwork.CharactersCore
	if rawMap, err = enka.GetUserDataRaw("112075042", false); err != nil {
		panic(err)
	}

	/*//获取展柜id列表
	var AvatarIDList []int
	if AvatarIDList, err = rawMap.GetAvatarIDList(); err != nil {
		panic(err)
	}
	//获取第一个角色信息
	var Avatarinfo EnkaNetwork.CharactersMapInfoRaw
	if Avatarinfo, err = rawMap.GetAvatarInfoRow(AvatarIDList[0]); err != nil {
		panic(err)
	}
	fmt.Println(Avatarinfo)*/
}
