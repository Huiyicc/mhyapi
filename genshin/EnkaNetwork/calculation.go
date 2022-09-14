package EnkaNetwork

import "github.com/Huiyicc/mhyapi/tools"

type AvatarBasisData struct {
	HP      DataSet //生命值
	Attack  DataSet //攻击力
	Defense DataSet //防御力
}

type DataSet struct {
	//原始基础数据
	BasisNumber float64
	//圣遗物加上的数据
	ReliquaryNumber float64
}

//FIGHT_PROP_HP_PERCENT
//FIGHT_PROP_ATTACK_PERCENT
//FIGHT_PROP_DEFENSE_PERCENT

// CalculationBasis 计算基础面板(生命值,攻击力,防御力)
func CalculationBasis(avatarInfo *AvatarInfo) AvatarBasisData {
	req := AvatarBasisData{
		HP: DataSet{
			BasisNumber:     avatarInfo.FightPropMap["1"],
			ReliquaryNumber: 0,
		}, Attack: DataSet{
			BasisNumber:     avatarInfo.FightPropMap["4"],
			ReliquaryNumber: 0,
		}, Defense: DataSet{
			BasisNumber:     avatarInfo.FightPropMap["7"],
			ReliquaryNumber: 0,
		},
	}
	for i := 0; i < len(avatarInfo.EquipList); i++ {
		if avatarInfo.EquipList[i].Flat.ItemType == "ITEM_WEAPON" {
			//武器
			//计算附加生命值
			req.HP.ReliquaryNumber += req.HP.BasisNumber * (getWeaponValue(avatarInfo.EquipList[i].Flat.WeaponStats, "FIGHT_PROP_HP_PERCENT") / 100) //生命值百分比
			//计算附加攻击力
			req.Attack.ReliquaryNumber += req.Attack.BasisNumber * (getWeaponValue(avatarInfo.EquipList[i].Flat.WeaponStats, "FIGHT_PROP_ATTACK_PERCENT") / 100) //攻击力百分比
			//计算附加防御力
			req.Defense.ReliquaryNumber += req.Defense.BasisNumber * (getWeaponValue(avatarInfo.EquipList[i].Flat.WeaponStats, "FIGHT_PROP_DEFENSE_PERCENT") / 100) //防御力百分比
		} else if avatarInfo.EquipList[i].Flat.ItemType == "ITEM_RELIQUARY" {
			//圣遗物
			vList := make([]ReliquaryStat, 0, 5)
			vList = append(vList, avatarInfo.EquipList[i].Flat.ReliquaryMainstat)
			vList = append(vList, avatarInfo.EquipList[i].Flat.ReliquarySubstats...)
			for _, stat := range vList {
				propId := tools.Ifs(stat.MainPropId == "", stat.AppendPropID, stat.MainPropId)
				sv := stat.StatValue
				switch propId {
				case "FIGHT_PROP_HP": //生命值
					req.HP.ReliquaryNumber += sv //生命值叠加
				case "FIGHT_PROP_HP_PERCENT": //生命值百分比
					req.HP.ReliquaryNumber += req.HP.BasisNumber * (sv / 100) //换成百分比相乘
				case "FIGHT_PROP_ATTACK": //攻击力
					req.Attack.ReliquaryNumber += sv //攻击力叠加
				case "FIGHT_PROP_ATTACK_PERCENT": //攻击力百分比
					req.Attack.ReliquaryNumber += req.Attack.BasisNumber * (sv / 100) //换成百分比相乘
				case "FIGHT_PROP_DEFENSE": //防御力
					req.Defense.ReliquaryNumber += sv //防御力叠加
				case "FIGHT_PROP_DEFENSE_PERCENT": //防御力百分比
					req.Defense.ReliquaryNumber += req.Defense.BasisNumber * (sv / 100) //换成百分比相乘
				}
			}
		}

	}
	req.Attack.ReliquaryNumber -= 0.5
	return req
}

func getWeaponValue(weaponStats []EquipWeaponStats, key string) float64 {
	for i := 0; i < len(weaponStats); i++ {
		if weaponStats[i].AppendPropId == key {
			return weaponStats[i].StatValue
		}
	}
	return 0
}
