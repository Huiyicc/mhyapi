package EnkaNetwork

import (
	"errors"
)

type CharactersCore CharactersInfo

type CharactersInfo struct {
	Ltime          int64                  `json:"ltime"`
	Ttl            int                    `json:"ttl"`
	Uid            string                 `json:"uid"`
	PlayerInfo     CharactersPlayerInfo   `json:"playerInfo"`
	AvatarInfoList []CharactersAvatarInfo `json:"avatarInfoList"`
}

// CharactersPlayerInfo 为获取到的玩家信息结构体
type CharactersPlayerInfo struct {
	Signature            string           `json:"signature"`            //签名
	WorldLevel           int              `json:"worldLevel"`           //世界等级
	NameCardId           int              `json:"nameCardId"`           //资料名片 ID
	ShowNameCardIdList   []int            `json:"showNameCardIdList"`   //正在展示的名片 ID 列表
	Nickname             string           `json:"nickname"`             //名称
	Level                int              `json:"level"`                //玩家等级
	FinishAchievementNum int              `json:"finishAchievementNum"` //已解锁成就数
	TowerFloorIndex      int              `json:"towerFloorIndex"`      //本期深境螺旋层数
	TowerLevelIndex      int              `json:"towerLevelIndex"`      //本期深境螺旋间数
	ShowAvatarInfoList   []ShowAvatarInfo `json:"showAvatarInfoList"`   //角色 ID 与等级的列表
	ProfilePicture       struct {
		AvatarId int `json:"avatarId"`
	} `json:"profilePicture"` //玩家头像的角色 ID
}

type CharactersAvatarInfo struct {
	FetterInfo struct {
		ExpLevel int `json:"expLevel"` //等级
	} `json:"fetterInfo"` //角色好感等级
	AvatarId int `json:"avatarId"` //角色ID
	PropMap  map[string]struct {
		Type int    `json:"type"` //属性类型,可使用 dictionariesMap.Prop[type] 获取
		Ival string `json:"ival"` //忽略它
		Val  string `json:"val"`  //属性值
	} `json:"propMap"` //角色属性列表
	FightPropMap           map[string]float64 `json:"fightPropMap"`           //角色战斗属性 Map,可使用 dictionariesMap.FightProp[kay] 获取
	SkillDepotId           int                `json:"skillDepotId"`           //角色天赋 ID
	InherentProudSkillList []int              `json:"inherentProudSkillList"` //固定天赋列表 ID
	SkillLevelMap          map[string]int     `json:"skillLevelMap"`          //天赋等级 Map
	EquipList              []struct {
		ItemId    int `json:"itemId"` // 装备/圣遗物 ID
		Reliquary struct {
			AppendPropIdList []int `json:"appendPropIdList"` //圣遗物副属性 ID 列表
			Level            int   `json:"level"`            //圣遗物等级 [1-21]
			MainPropId       int   `json:"mainPropId"`       //圣遗物主属性 ID
		} `json:"reliquary,omitempty"` //圣遗物基本信息
		Flat struct {
			ReliquaryMainstat struct {
				MainPropId string  `json:"mainPropId"` //属性名称
				StatValue  float64 `json:"statValue"`  //属性值
			} `json:"reliquaryMainstat,omitempty"` //圣遗物主属性
			ReliquarySubstats []struct {
				AppendPropId string  `json:"appendPropId"` //属性名称
				StatValue    float64 `json:"statValue"`    //属性值
			} `json:"reliquarySubstats,omitempty"` //圣遗物副属性列表
			ItemType           string `json:"itemType"`                     //装备类别：武器、圣遗物
			Icon               string `json:"icon"`                         //装备图标名称
			EquipType          string `json:"equipType,omitempty"`          //圣遗物类型
			NameTextMapHash    string `json:"nameTextMapHash"`              //装备名的哈希值
			SetNameTextMapHash string `json:"setNameTextMapHash,omitempty"` //圣遗物套装的名称的哈希值
			RankLevel          int    `json:"rankLevel"`                    //装备稀有度
			WeaponStats        []struct {
				StatValue    float64 `json:"statValue"`    //属性名称
				AppendPropId string  `json:"appendPropId"` //属性值
			} `json:"weaponStats,omitempty"` //武器属性列表：基础攻击力、副属性
		} `json:"flat"` //装备详细信息
		Weapon struct {
			Level        int            `json:"level"`        //武器等级
			PromoteLevel int            `json:"promoteLevel"` //武器突破星级
			AffixMap     map[string]int `json:"affixMap"`     //武器精炼信息
		} `json:"weapon,omitempty"` //武器基本信息
	} `json:"equipList"` //装备列表：武器、圣遗物
	TalentIdList []int `json:"talentIdList,omitempty"` //命之座 ID 列表,如果未解锁任何命之座则此数据不存在
}

// ShowAvatarInfo 为对外展示的基础角色信息
type ShowAvatarInfo struct {
	AvatarId  int64 `json:"avatarId"`  //角色 ID
	Level     int   `json:"level"`     //角色等级
	CostumeId int   `json:"costumeId"` //角色衣装 ID
}

// GetPlayerAvatarInfoForID 指定ID取玩家展柜角色数据
func (t CharactersCore) GetPlayerAvatarInfoForID(AvatarID int64) (ShowAvatarInfo, error) {
	for i := 0; i < len(t.PlayerInfo.ShowAvatarInfoList); i++ {
		if t.PlayerInfo.ShowAvatarInfoList[i].AvatarId == AvatarID {
			return t.PlayerInfo.ShowAvatarInfoList[i], nil
		}
	}
	return ShowAvatarInfo{}, errors.New("展柜中无此角色")
}

// GetPlayerAvatarInfoForIndex 指定索引取玩家展柜角色数据 [index:0-7]
func (t CharactersCore) GetPlayerAvatarInfoForIndex(index int) (ShowAvatarInfo, error) {
	if index < 0 || index > 7 || index >= len(t.PlayerInfo.ShowAvatarInfoList) {
		return ShowAvatarInfo{}, errors.New("索引超出范围")
	}
	return t.PlayerInfo.ShowAvatarInfoList[index], nil
}

// GetAvatarInfoList 取展示角色的信息列表
func (t CharactersCore) GetAvatarInfoList() ([]ShowAvatarInfo, error) {
	if len(t.PlayerInfo.ShowAvatarInfoList) == 0 {
		return nil, errors.New("展柜无数据或未开放展柜")
	}
	return t.PlayerInfo.ShowAvatarInfoList, nil
}

/*// GetAvatarInfoRow 取展示角色的元素
func (t CharactersCore) GetAvatarInfoRow(roleID int) (CharactersMapInfoRaw, error) {
	return charactersMap.GetInfoByID(strconv.Itoa(roleID))
}*/

/*// GetAvatarInfoTranslate 取展示角色的元素,自动翻译
func (t CharactersCore) GetAvatarInfoTranslate(roleID int) (CharactersMapInfoRaw, error) {
	return charactersMap.GetInfoByID(strconv.Itoa(roleID))
}*/
