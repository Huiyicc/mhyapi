package EnkaNetwork

type AvatarBasisData struct {
	HP       DataSet //生命值
	Attack   DataSet //攻击力
	Defense  DataSet //防御力
	BasisMap map[string]DataSet
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
			req.HP.BasisNumber += getWeaponValue(avatarInfo.EquipList[i].Flat.WeaponStats, "FIGHT_PROP_HP_PERCENT") //生命值百分比
			//计算附加攻击力
			req.Attack.BasisNumber += getWeaponValue(avatarInfo.EquipList[i].Flat.WeaponStats, "FIGHT_PROP_ATTACK_PERCENT") //攻击力百分比
			//计算附加防御力
			req.Defense.BasisNumber += getWeaponValue(avatarInfo.EquipList[i].Flat.WeaponStats, "FIGHT_PROP_DEFENSE_PERCENT") //防御力百分比
		} else if avatarInfo.EquipList[i].Flat.ItemType == "ITEM_RELIQUARY" {
			//圣遗物
			/*reliquaryECInfo, err := GetReliquaryMainExcelConfigByID(avatarInfo.EquipList[i].Reliquary.MainPropId)
			if err != nil {
				continue
			}*/

		}

	}
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
