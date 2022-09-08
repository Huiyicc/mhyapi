package genshin

import (
	"errors"
	"fmt"
	"github.com/Huiyicc/mhyapi/request"
	"github.com/Huiyicc/mhyapi/tools"
	json "github.com/json-iterator/go"
)

// GetSpiralAbyssInfo 用于获取深渊相关数据
// scheduleType:1 为本期数据,2为上期数据
func (t *GenShinCore) GetSpiralAbyssInfo(scheduleType int, gameID, server string) (*SpiralAbyssInfo, error) {
	req := request.MYSINFO_API_SPIRALABYSS.Copy()
	req.Query = fmt.Sprintf(req.Query, gameID, scheduleType, server)
	/*	headers := make(http.Header)
		headers["DS"] = []string{tools.GetDs3(req.Query, "")}
		headers["x-rpc-app_version"] = []string{"2.34.1"}
		headers["x-rpc-client_type"] = []string{"5"}
		headers["User-Agent"] = []string{"Mozilla/5.0 (Linux; Android 12; Unspecified Device) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.73 Mobile Safari/537.36 miHoYoBBS/2.35.2"}
		request.ProxyUrl = "http://127.0.0.1:8080"*/
	headers := t.getGameHeaders(req.Query, "")
	data, err := request.HttpGet(req, headers)
	if err != nil {
		return nil, err
	}
	var resp spiralAbyssInfoResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, errors.New(string(data))
	}
	if err = resp.verify(); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

type spiralAbyssInfoResponse struct {
	Retcode int             `json:"retcode"`
	Message string          `json:"message"`
	Data    SpiralAbyssInfo `json:"data"`
}

// 验证器
func (t *spiralAbyssInfoResponse) verify() error {
	return tools.Ifs(t.Retcode == 0, nil, errors.New(t.Message))
}

type SpiralAbyssInfo struct {
	ScheduleId       int               `json:"schedule_id"`
	StartTime        string            `json:"start_time"`         //统计开始时间
	EndTime          string            `json:"end_time"`           //统计结束时间
	TotalBattleTimes int               `json:"total_battle_times"` //战斗次数
	TotalWinTimes    int               `json:"total_win_times"`    //
	MaxFloor         string            `json:"max_floor"`          //最深抵达层数
	RevealRank       []SpiralAbyssRank `json:"reveal_rank"`        //出战次数展示
	DefeatRank       []SpiralAbyssRank `json:"defeat_rank"`        //最多击破数
	DamageRank       []SpiralAbyssRank `json:"damage_rank"`        //最强一击
	TakeDamageRank   []SpiralAbyssRank `json:"take_damage_rank"`   //最大承受伤害
	NormalSkillRank  []SpiralAbyssRank `json:"normal_skill_rank"`  //元素战技释放次数
	EnergySkillRank  []SpiralAbyssRank `json:"energy_skill_rank"`  //元素爆发释放次数
	Floors           []struct {
		Index      int    `json:"index"`       //层数
		Icon       string `json:"icon"`        //
		IsUnlock   bool   `json:"is_unlock"`   //是否解锁
		SettleTime string `json:"settle_time"` //
		Star       int    `json:"star"`        //已获取星
		MaxStar    int    `json:"max_star"`    //最大可获取星
		Levels     []struct {
			Index   int           `json:"index"`    //第x间
			Star    int           `json:"star"`     //已获取星
			MaxStar int           `json:"max_star"` //最大可获取星
			Battles []interface{} `json:"battles"`
		} `json:"levels"` //房间
	} `json:"floors"` //层级
	TotalStar int  `json:"total_star"` //本期已获取星数量
	IsUnlock  bool `json:"is_unlock"`  //是否解锁深渊
}

type SpiralAbyssRank struct {
	AvatarId   int    `json:"avatar_id"`   //角色ID
	AvatarIcon string `json:"avatar_icon"` //角色图像
	Value      int    `json:"value"`       //值
	Rarity     int    `json:"rarity"`      //星级
}
