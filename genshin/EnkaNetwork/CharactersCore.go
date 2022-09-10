package EnkaNetwork

import (
	"strconv"
)

type CharactersCore CharactersInfo

type CharactersInfo struct {
	Ttl            int                    `json:"ttl"`
	Uid            string                 `json:"uid"`
	PlayerInfo     CharactersPlayerInfo   `json:"playerInfo"`
	AvatarInfoList []CharactersAvatarInfo `json:"avatarInfoList"`
}

// CharactersPlayerInfo 为获取到的玩家信息结构体
type CharactersPlayerInfo struct {
	Signature            string `json:"signature"`          //签名
	WorldLevel           int    `json:"worldLevel"`         //世界等级
	NameCardId           int    `json:"nameCardId"`         //资料名片 ID
	ShowNameCardIdList   []int  `json:"showNameCardIdList"` //正在展示的名片 ID 列表
	Nickname             string `json:"nickname"`           //名称
	Level                int    `json:"level"`
	FinishAchievementNum int    `json:"finishAchievementNum"` //已解锁成就数
	TowerFloorIndex      int    `json:"towerFloorIndex"`      //本期深境螺旋层数
	TowerLevelIndex      int    `json:"towerLevelIndex"`      //本期深境螺旋间数
	ShowAvatarInfoList   []struct {
		AvatarId  int `json:"avatarId"`  //角色 ID
		Level     int `json:"level"`     //角色等级
		CostumeId int `json:"costumeId"` //角色衣装 ID
	} `json:"showAvatarInfoList"` //角色 ID 与等级的列表
	ProfilePicture struct {
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
				MainPropId string  `json:"mainPropId"`
				StatValue  float64 `json:"statValue"`
			} `json:"reliquaryMainstat,omitempty"`
			ReliquarySubstats []struct {
				AppendPropId string  `json:"appendPropId"`
				StatValue    float64 `json:"statValue"`
			} `json:"reliquarySubstats,omitempty"`
			ItemType           string `json:"itemType"`
			Icon               string `json:"icon"`
			EquipType          string `json:"equipType,omitempty"`
			NameTextMapHash    string `json:"nameTextMapHash"`
			SetNameTextMapHash string `json:"setNameTextMapHash,omitempty"`
			RankLevel          int    `json:"rankLevel"`
			WeaponStats        []struct {
				StatValue    float64 `json:"statValue"`
				AppendPropId string  `json:"appendPropId"`
			} `json:"weaponStats,omitempty"`
		} `json:"flat"` //装备详细信息
		Weapon struct {
			Level        int `json:"level"`
			PromoteLevel int `json:"promoteLevel"`
			AffixMap     struct {
				Field1 int `json:"114502,omitempty"`
				Field2 int `json:"113303,omitempty"`
				Field3 int `json:"113509,omitempty"`
				Field4 int `json:"111503,omitempty"`
				Field5 int `json:"113507,omitempty"`
				Field6 int `json:"112503,omitempty"`
				Field7 int `json:"115402,omitempty"`
				Field8 int `json:"111505,omitempty"`
			} `json:"affixMap"` //武器基本信息
		} `json:"weapon,omitempty"`
	} `json:"equipList"` //装备列表：武器、圣遗物
	TalentIdList []int `json:"talentIdList,omitempty"`
}

func (t CharactersCore) GetPlayerInfo() {

}

// GetAvatarIDList 取展示角色的ID列表
func (t CharactersCore) GetAvatarIDList() (ret []int, err error) {

	/*	if t == nil {
			err = errors.New("characters is nil")
			return
		}
		avatarInfoList := t["avatarInfoList"]
		if avatarInfoList == nil {
			err = errors.New("avatar is nil")
			return
		}
		avatarInfoListArr, _ := avatarInfoList.([]any)
		if avatarInfoListArr == nil {
			err = errors.New("avatarInfoList is nil")
			return
		}
		ret = make([]int, 0, len(avatarInfoListArr))
		for i := 0; i < len(avatarInfoListArr); i++ {
			avatarInfo, _ := avatarInfoListArr[i].(map[string]any)
			if avatarInfo == nil {
				err = errors.New("avatarInfo is nil")
				return
			}
			aid, _ := avatarInfo["avatarId"]
			if aid == nil {
				err = errors.New("avatar[" + strconv.Itoa(i) + "]ID is nil")
				return
			}
			ad, _ := aid.(float64)
			if ad == 0 {
				ad1, _ := aid.(int64)
				ad = float64(ad1)
			}
			ret = append(ret, int(ad))
		}*/
	return
}

// GetAvatarInfoRow 取展示角色的元素
func (t CharactersCore) GetAvatarInfoRow(roleID int) (CharactersMapInfoRaw, error) {
	return charactersMap.GetInfoByID(strconv.Itoa(roleID))
}

// GetAvatarInfoTranslate 取展示角色的元素,自动翻译
func (t CharactersCore) GetAvatarInfoTranslate(roleID int) (CharactersMapInfoRaw, error) {
	return charactersMap.GetInfoByID(strconv.Itoa(roleID))
}
