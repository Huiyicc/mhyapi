package EnkaNetwork

import (
	_ "embed"
	"errors"
)

var (
	//go:embed res_son/ReliquaryMainPropExcelConfigData.json
	reliquaryMainPropExcelConfigData []byte

	reliquaryMainPropExcelConfigList []ReliquaryMainExcelConfig
)

// ReliquaryMainExcelConfig 为圣遗物配置信息
type ReliquaryMainExcelConfig struct {
	Id          int    `json:"id"`
	PropDepotId int    `json:"propDepotId"`
	PropType    string `json:"propType"`
	AffixName   string `json:"affixName"`
}

type ReliquaryAffixExcelConfig struct {
	Id        int     `json:"id"`
	DepotId   int     `json:"depotId"`
	GroupId   int     `json:"groupId"`
	PropType  string  `json:"propType"`
	PropValue float64 `json:"propValue"`
}

// GetReliquaryMainExcelConfigByID id获取圣遗物主词条定义信息
func GetReliquaryMainExcelConfigByID(id int) (ReliquaryMainExcelConfig, error) {
	for i := 0; i < len(reliquaryMainPropExcelConfigList); i++ {
		if reliquaryMainPropExcelConfigList[i].Id == id {
			return reliquaryMainPropExcelConfigList[i], nil
		}
	}
	return ReliquaryMainExcelConfig{}, errors.New("ID Is Not Set")
}

// GetReliquaryAffixExcelConfigByID id获取圣遗物副词条定义信息
func GetReliquaryAffixExcelConfigByID(id int) (ReliquaryMainExcelConfig, error) {
	for i := 0; i < len(reliquaryMainPropExcelConfigList); i++ {
		if reliquaryMainPropExcelConfigList[i].Id == id {
			return reliquaryMainPropExcelConfigList[i], nil
		}
	}
	return ReliquaryMainExcelConfig{}, errors.New("ID Is Not Set")
}
