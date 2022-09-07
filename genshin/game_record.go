package genshin

import (
	"errors"
	"fmt"
	"github.com/Huiyicc/mhyapi/define"
	"github.com/Huiyicc/mhyapi/request"
	"github.com/Huiyicc/mhyapi/tools"
	json "github.com/json-iterator/go"
)

// GameRecordIndex 用于获取游戏主页信息数据
func (t *GenShinCore) GameRecordIndex(gameID string) (*GameRecordIndexInfo, error) {
	req := request.MYSINFO_API_INDEX.Copy()
	req.Query = fmt.Sprintf(req.Query, gameID, define.GAMERSERVER_GENSHIN_TIANKONGDAO)
	herders := t.getGameHeaders(req.Query, "")
	data, err := request.HttpGet(req, herders)
	if err != nil {
		return nil, err
	}
	var resp gameRecordIndexResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, errors.New(string(data))
	}
	if err = resp.verify(); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// 验证器
func (t *gameRecordIndexResponse) verify() error {
	return tools.Ifs(t.Retcode == 0, nil, errors.New(t.Message))
}

type gameRecordIndexResponse struct {
	Retcode int                 `json:"retcode"`
	Message string              `json:"message"`
	Data    GameRecordIndexInfo `json:"data"`
}
type GameRecordIndexInfo struct {
	Role struct { //用户角色
		AvatarUrl string `json:"AvatarUrl"` //头像
		Nickname  string `json:"nickname"`  //昵称
		Region    string `json:"region"`    //服务器
		Level     int    `json:"level"`     //等级
	} `json:"role"`
	Avatars []struct { //游戏角色
		Id                      int    `json:"id"`                        //角色ID
		Image                   string `json:"image"`                     //角色证件照(1寸照片;)笑)
		Name                    string `json:"name"`                      //名称
		Element                 string `json:"element"`                   //元素属性
		Fetter                  int    `json:"fetter"`                    //好感度
		Level                   int    `json:"level"`                     //角色等级
		Rarity                  int    `json:"rarity"`                    //角色星级(4星或5星)
		ActivedConstellationNum int    `json:"actived_constellation_num"` //角色命座数据
		CardImage               string `json:"card_image"`                //卡片版图像
		IsChosen                bool   `json:"is_chosen"`                 //是否为主动展示
	} `json:"avatars"`
	Stats struct { //统计数据
		ActiveDayNumber      int    `json:"active_day_number"`      //活跃天数
		AchievementNumber    int    `json:"achievement_number"`     //成就数
		AnemoculusNumber     int    `json:"anemoculus_number"`      //风神瞳数量
		GeoculusNumber       int    `json:"geoculus_number"`        //岩神瞳数量
		AvatarNumber         int    `json:"avatar_number"`          //获得角色数
		WayPointNumber       int    `json:"way_point_number"`       //解锁传送点
		DomainNumber         int    `json:"domain_number"`          //解锁秘境数量
		SpiralAbyss          string `json:"spiral_abyss"`           //深境螺旋关卡
		PreciousChestNumber  int    `json:"precious_chest_number"`  //珍贵宝箱数
		LuxuriousChestNumber int    `json:"luxurious_chest_number"` //华丽宝箱数
		ExquisiteChestNumber int    `json:"exquisite_chest_number"` //精致宝箱数
		CommonChestNumber    int    `json:"common_chest_number"`    //普通宝箱数
		ElectroculusNumber   int    `json:"electroculus_number"`    //雷神瞳数量
		MagicChestNumber     int    `json:"magic_chest_number"`     //奇亏宝箱数
		DendroculusNumber    int    `json:"dendroculus_number"`     //草神瞳数量
	} `json:"stats"`
	CityExplorations  []interface{} `json:"city_explorations"` //城市探索度(城市是什么鬼玩意)
	WorldExplorations []struct {    //世界探索度
		Level                 int        `json:"level"`                  //声望等级
		ExplorationPercentage int        `json:"exploration_percentage"` //探索度(满点1000)
		Icon                  string     `json:"icon"`                   //区域图标
		Name                  string     `json:"name"`                   //区域名称
		Type                  string     `json:"type"`                   //
		Offerings             []struct { //贡献类型 忍冬树,神樱树,梦之树之类的
			Name  string `json:"name"`  //名称
			Level int    `json:"level"` //等级
			Icon  string `json:"icon"`  //图标
		} `json:"offerings"`
		Id              int    `json:"id"`
		ParentId        int    `json:"parent_id"`
		MapUrl          string `json:"map_url"`          //地图链接
		StrategyUrl     string `json:"strategy_url"`     //观测枢专区链接
		BackgroundImage string `json:"background_image"` //背景图像
		InnerIcon       string `json:"inner_icon"`       //中心图像
		Cover           string `json:"cover"`            //封面
	} `json:"world_explorations"`
	Homes []struct { //尘歌壶
		Level            int    `json:"level"`              //信任等级
		VisitNum         int    `json:"visit_num"`          //历史访客数
		ComfortNum       int    `json:"comfort_num"`        //最高洞天仙力
		ItemNum          int    `json:"item_num"`           //获得摆设数
		Name             string `json:"name"`               //名称
		Icon             string `json:"icon"`               //背景图像
		ComfortLevelName string `json:"comfort_level_name"` //阿圆的等级名称
		ComfortLevelIcon string `json:"comfort_level_icon"` //阿圆的图标
	} `json:"homes"`
}
