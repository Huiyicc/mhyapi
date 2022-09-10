package EnkaNetwork

type dictionaries struct {
	init       bool
	Prop       map[string]string `json:"Prop"`
	FightProp  map[string]string `json:"FightProp"`
	EquipType  map[string]string `json:"EquipType"`
	AppendProp map[string]string `json:"AppendProp"`
}
