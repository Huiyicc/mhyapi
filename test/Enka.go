package main

import (
	"fmt"
	"github.com/Huiyicc/mhyapi/genshin/EnkaNetwork"
)

func main() {
	test_Enka()
}

func test_Enka() {
	enka, err := EnkaNetwork.NewEnkaCore("./cache")
	if err != nil {
		panic(err)
	}
	//获取原始数据
	var rawCore *EnkaNetwork.UserData
	if rawCore, err = enka.GetUserDataRaw("112075042", false); err != nil {
		panic(err)
	}
	//获取第1个展柜角色信息
	var Avatar0 EnkaNetwork.ShowAvatarInfo
	if Avatar0, err = rawCore.GetPlayerAvatarInfoForIndex(0); err != nil {
		panic(err)
	}
	//获取角色id对应信息
	var linfor *EnkaNetwork.CharactersInfoLoc
	if linfor, err = EnkaNetwork.GetCharactersInfoLocByID(EnkaNetwork.HASHTEXT_LANGUAGE_ZHCN, Avatar0.AvatarId); err != nil {
		panic(err)
	}
	fmt.Println(linfor)
	//获取角色立绘
	_, err = EnkaNetwork.GetResourcesData(enka, linfor.SideImgName)
	if err != nil {
		panic(err)
	}
	/*//获取展柜id列表
	var AvatarIDList []int
	if AvatarIDList, err = rawCore.GetAvatarIDList(); err != nil {
		panic(err)
	}
	//获取第一个角色信息
	var Avatarinfo EnkaNetwork.CharactersMapInfoRaw
	if Avatarinfo, err = rawCore.GetAvatarInfoRow(AvatarIDList[0]); err != nil {
		panic(err)
	}
	fmt.Println(Avatarinfo)*/
}
