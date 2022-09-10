package EnkaNetwork

import (
	"errors"
	"fmt"
)

type charactersMapCore map[string]CharactersMapInfoRaw

type CharactersMapInfoRaw struct {
	Element         string            `json:"Element"`
	Consts          []string          `json:"Consts"`
	SkillOrder      []int             `json:"SkillOrder"`
	Skills          map[string]string `json:"Skills"`
	ProudMap        map[string]int64  `json:"ProudMap"`
	NameTextMapHash int               `json:"NameTextMapHash"`
	SideIconName    string            `json:"SideIconName"`
	QualityType     string            `json:"QualityType"`
}

func (t charactersMapCore) GetInfoByID(roleID string) (r CharactersMapInfoRaw, err error) {
	l, isOk := t[roleID]
	if !isOk {
		err = errors.New(fmt.Sprintf("角色[%s]不存在", roleID))
		return
	}
	smp := make(map[string]string)
	for k, v := range l.Skills {
		smp[k] = v
	}
	pmp := make(map[string]int64)
	for k, v := range l.ProudMap {
		pmp[k] = v
	}
	r = CharactersMapInfoRaw{
		Element:         l.Element,
		Consts:          l.Consts,
		SkillOrder:      l.SkillOrder,
		Skills:          smp,
		ProudMap:        pmp,
		NameTextMapHash: l.NameTextMapHash,
		SideIconName:    l.SideIconName,
		QualityType:     l.QualityType,
	}
	return
}
