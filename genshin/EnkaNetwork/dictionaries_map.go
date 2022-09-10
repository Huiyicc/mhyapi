package EnkaNetwork

import "strings"

type dictionariesMapCore struct {
	init       bool
	Prop       map[string]string `json:"Prop"`
	FightProp  map[string]string `json:"FightProp"`
	EquipType  map[string]string `json:"EquipType"`
	AppendProp map[string]string `json:"AppendProp"`
}

var ElementMap = map[string]string{
	"FIRE":  "火",
	"WATER": "水",
	"GRASS": "草",
	"ELEC":  "雷",
	"ROCK":  "岩",
	"WIND":  "风",
	"ICE":   "冰",
}

// TranslateRoleElement 翻译角色的属性
func TranslateRoleElement(Element string) string {
	s := ElementMap[strings.ToUpper(Element)]
	if s == "" {
		return Element
	}
	return s
}
