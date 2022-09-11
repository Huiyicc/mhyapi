package EnkaNetwork

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// 人物数据表map封装
type charactersMapCore map[string]CharactersMapInfoRaw

// CharactersMapInfoRaw 为人物数据表原始结构体
type CharactersMapInfoRaw struct {
	Element         string            `json:"Element"`
	Consts          []string          `json:"Consts"`
	SkillOrder      []int             `json:"SkillOrder"`
	Skills          map[string]string `json:"Skills"`
	ProudMap        map[string]int64  `json:"ProudMap"`
	NameTextMapHash int64             `json:"NameTextMapHash"`
	SideIconName    string            `json:"SideIconName"`
	QualityType     string            `json:"QualityType"`
}

// CharactersInfoLoc 本地化后的人物信息数据
type CharactersInfoLoc struct {
	Element         LocInfo           //元素属性
	Consts          []LocUrl          //固有天赋
	SkillOrder      []int             //技能列表
	Skills          map[string]LocUrl //技能图标列表
	ProudMap        map[string]int64  //
	NameTextMapHash LocInfoInt64      //名称信息
	SideIconName    LocUrl            //头像图标
	SideImgName     LocUrl            //抽卡立绘
	QualityType     string
}

type LocInfo struct {
	Name    string
	LocName string
}
type LocInfoInt64 struct {
	Number  int64
	LocName string
}
type LocUrl struct {
	Name string
	Url  string
}

func rowListToLocalInfos(list []string) []LocInfo {
	r := make([]LocInfo, 0, len(list))
	for i := 0; i < len(list); i++ {
		r = append(r, LocInfo{
			Name: list[i],
			//Url: ,
		})
	}
	return r
}
func rowListToLocalUrls(list []string) []LocUrl {
	r := make([]LocUrl, 0, len(list))
	for i := 0; i < len(list); i++ {
		r = append(r, LocUrl{
			Name: list[i],
			Url:  "https://enka.network/ui/" + list[i] + ".png",
		})
	}
	return r
}
func rowMapToLocalUrls(maps map[string]string) map[string]LocUrl {
	r := make(map[string]LocUrl)
	for k, v := range maps {
		r[k] = LocUrl{
			Name: v,
			Url:  "https://enka.network/ui/" + v + ".png",
		}
	}
	return r
}

// GetCharactersInfoRawByID 用角色ID获得角色信息原型数据
func GetCharactersInfoRawByID(RoleID int64) (*CharactersMapInfoRaw, error) {
	rmap, isSet := charactersMap[strconv.FormatInt(RoleID, 10)]
	if !isSet {
		return nil, errors.New("角色查无信息")
	}
	r := CharactersMapInfoRaw{
		Element:         rmap.Element,
		NameTextMapHash: rmap.NameTextMapHash,
		SideIconName:    rmap.SideIconName,
		QualityType:     rmap.QualityType,
	}
	r.Consts = make([]string, len(rmap.Consts))
	copy(r.Consts, rmap.Consts)
	r.SkillOrder = make([]int, len(rmap.SkillOrder))
	copy(r.SkillOrder, rmap.SkillOrder)
	r.Skills = make(map[string]string)
	for k, v := range rmap.Skills {
		r.Skills[k] = v
	}
	r.ProudMap = make(map[string]int64)
	for k, v := range rmap.ProudMap {
		r.ProudMap[k] = v
	}
	return &r, nil
}

// GetCharactersInfoLocByID 用角色ID获得本地化之后的角色信息
func GetCharactersInfoLocByID(lang string, RoleID int64) (*CharactersInfoLoc, error) {
	raw, err := GetCharactersInfoRawByID(RoleID)
	if err != nil {
		return nil, err
	}
	hs, err := GetHashStr(lang, raw.NameTextMapHash)
	if err != nil {
		return nil, err
	}
	si := strings.Replace(raw.SideIconName, "UI_AvatarIcon_Side", "UI_Gacha_AvatarImg", 1)
	r := CharactersInfoLoc{
		Element: LocInfo{
			Name:    raw.Element,
			LocName: TranslateRoleElement(raw.Element),
		},
		Consts:     rowListToLocalUrls(raw.Consts),
		SkillOrder: raw.SkillOrder,
		Skills:     rowMapToLocalUrls(raw.Skills),
		ProudMap:   raw.ProudMap,
		NameTextMapHash: LocInfoInt64{
			Number:  raw.NameTextMapHash,
			LocName: hs,
		},
		SideIconName: LocUrl{
			Name: raw.SideIconName,
			Url:  "https://enka.network/ui/" + raw.SideIconName + ".png",
		},
		SideImgName: LocUrl{
			Name: si,
			Url:  "https://enka.network/ui/" + si + ".png",
		},
		QualityType: raw.QualityType,
	}
	return &r, nil
}

// GetResourcesData 用于获取资源
func GetResourcesData(core *EnkaCore, LocInfo LocUrl, cache ...bool) ([]byte, error) {
	addCache := true
	if len(cache) > 0 {
		addCache = cache[0]
	}
	cachePath := core.cachePath + "Resources"
	os.MkdirAll(cachePath, 0655)
	cacheFileName := core.getUrlCachePath(LocInfo.Url, cachePath) + ".png"
	if !addCache {
		os.Remove(cacheFileName)
	}
	data, err := core.httpGetIncidentalCache(LocInfo.Url, cacheFileName)
	return data, err
}
