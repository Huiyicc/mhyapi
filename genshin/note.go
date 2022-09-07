package genshin

import (
	"errors"
	"fmt"
	"github.com/Huiyicc/mhyapi/request"
	"github.com/Huiyicc/mhyapi/tools"
	json "github.com/json-iterator/go"
	"strconv"
	"time"
)

// GetNoteInfo 获取体力信息
func (t *GenShinCore) GetNoteInfo() (*NoteInfo, error) {
	req := request.MYSINFO_API_DAILYNOTE.Copy()
	req.Query = fmt.Sprintf(req.Query, t.GameInfo.GameUid, t.GameInfo.Region)
	headers := t.getGameHeaders(req.Query, "")
	data, err := request.HttpGet(req, headers)
	if err != nil {
		return nil, err
	}
	var resp mysNoteResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, errors.New(string(data))
	}
	if err = resp.verify(); err != nil {
		return nil, err
	}
	resp.Data.Region = t.GameInfo.Region
	resp.Data.GameUID = t.GameInfo.GameUid
	var r NoteInfo
	r.parse(resp)
	return &r, nil
}

type NoteInfo struct {
	GameInfo      NodeGameInfo          //游戏信息
	Resin         NodeResinInfo         //树脂相关
	Task          NodeTaskInfo          //委托相关
	ResinDiscount NodeResinDiscountInfo //周本折扣相关
	Expeditions   NodeExpeditionsInfo   //派遣相关
	HomeCoin      NodeHomeCoinInfo      //洞天宝钱相关
	Transformer   NodeTransformerInfo   //参变仪相关
}

type NodeGameInfo struct {
	Region  string //服务器
	GameUID string //游戏id
}

// NodeHomeCoinInfo 为洞天宝钱相关
type NodeHomeCoinInfo struct {
	CurrentHomeCoin          int   //洞天宝钱已有数量
	MaxHomeCoin              int   //洞天宝钱最大数量
	HomeCoinRecoveryTimeUnix int64 //洞天宝钱恢复时间(秒)
	HomeCoinRecoveryTime     int64 //洞天宝钱恢复时间(确切时间戳)
}

// NodeResinDiscountInfo 为周本折扣相关结构体
type NodeResinDiscountInfo struct {
	RemainResinDiscountNum int //周本折扣剩余
	ResinUnusedDiscountNum int //周本折扣已使用次数
	ResinDiscountNumLimit  int //周本折扣总次数
}

// NodeTaskInfo 为委托任务相关相关结构体
type NodeTaskInfo struct {
	FinishedTaskNum           int  //已完成任务数
	TotalTaskNum              int  //任务总数
	IsExtraTaskRewardReceived bool //是否已经领取委托奖励
}

// NodeResinInfo 为树脂相关结构体
type NodeResinInfo struct {
	CurrentResin          int   //当前树脂
	MaxResin              int   //最大树脂
	ResinRecoveryTimeUnix int64 //剩余恢复时间(秒)
	ResinRecoveryTime     int64 //剩余恢复时间(准确时间的时间戳)
}

// NodeExpeditionsInfo 为便签派遣结构体
type NodeExpeditionsInfo struct {
	CurrentExpeditionNum int                       //当前探索派遣次数
	MaxExpeditionNum     int                       //最大探索派遣次数
	Expeditions          []NodeExpeditionsRoleInfo //探索队伍
}

// NodeExpeditionsRoleInfo 为便签派遣角色信息结构体
type NodeExpeditionsRoleInfo struct {
	AvatarSideIcon string //角色头像
	Status         string //派遣状态
	RemainedTime   int64  //派遣剩余时间
	Deadline       int64  //派遣结束时间(准确时间的时间戳)
}

// NodeTransformerInfo 为便签参变仪结构体
type NodeTransformerInfo struct {
	Obtained     bool                            //是否拥有参变仪
	RecoveryTime NodeTransformerRecoveryTimeInfo //剩余恢复时间
}

// NodeTransformerRecoveryTimeInfo 为便签参变仪恢复时间子结构体
type NodeTransformerRecoveryTimeInfo struct {
	Day     int  //天
	Hour    int  //小时
	Minute  int  //分钟
	Second  int  //秒
	Reached bool //是否冷却完毕
}

// parse NoteInfo.Info() 解析返回的结构为封装体,同时返回自身
func (t *NoteInfo) parse(r mysNoteResponse) *NoteInfo {
	//当前树脂
	t.Resin.CurrentResin = r.Data.CurrentResin
	//最大树脂
	t.Resin.MaxResin = r.Data.MaxResin
	//剩余恢复时间(秒)
	t.Resin.ResinRecoveryTimeUnix, _ = strconv.ParseInt(r.Data.ResinRecoveryTime, 10, 64)
	//剩余恢复时间(准确时间)
	t.Resin.ResinRecoveryTime = time.Now().Unix() + t.Resin.ResinRecoveryTimeUnix
	//已完成任务数
	t.Task.FinishedTaskNum = r.Data.FinishedTaskNum
	//任务总数
	t.Task.TotalTaskNum = r.Data.TotalTaskNum
	//是否已经领取委托奖励
	t.Task.IsExtraTaskRewardReceived = r.Data.IsExtraTaskRewardReceived
	//周本折扣剩余
	t.ResinDiscount.RemainResinDiscountNum = r.Data.RemainResinDiscountNum
	//周本折扣总次数
	t.ResinDiscount.ResinDiscountNumLimit = r.Data.ResinDiscountNumLimit
	//周本折扣已使用次数
	t.ResinDiscount.ResinUnusedDiscountNum = t.ResinDiscount.ResinDiscountNumLimit - t.ResinDiscount.RemainResinDiscountNum
	//当前探索派遣次数
	t.Expeditions.CurrentExpeditionNum = r.Data.CurrentExpeditionNum
	//最大探索派遣次数
	t.Expeditions.MaxExpeditionNum = r.Data.MaxExpeditionNum
	//提前定义切片
	t.Expeditions.Expeditions = make([]NodeExpeditionsRoleInfo, len(r.Data.Expeditions))
	//探索队伍
	for i, v := range r.Data.Expeditions {
		t.Expeditions.Expeditions[i].AvatarSideIcon = v.AvatarSideIcon
		t.Expeditions.Expeditions[i].Status = v.Status
		t.Expeditions.Expeditions[i].RemainedTime, _ = strconv.ParseInt(v.RemainedTime, 10, 64)
	}
	//洞天宝钱已有数量
	t.HomeCoin.CurrentHomeCoin = r.Data.CurrentHomeCoin
	//洞天宝钱最大数量
	t.HomeCoin.MaxHomeCoin = r.Data.MaxHomeCoin
	//洞天宝钱恢复时间(秒)
	t.HomeCoin.HomeCoinRecoveryTimeUnix, _ = strconv.ParseInt(r.Data.HomeCoinRecoveryTime, 10, 64)
	//洞天宝钱恢复时间(确切时间戳)
	t.HomeCoin.HomeCoinRecoveryTime = time.Now().Unix() + t.HomeCoin.HomeCoinRecoveryTimeUnix
	//参变仪相关
	t.Transformer = NodeTransformerInfo{
		Obtained: r.Data.Transformer.Obtained,
		RecoveryTime: NodeTransformerRecoveryTimeInfo{
			Day:     r.Data.Transformer.RecoveryTime.Day,
			Hour:    r.Data.Transformer.RecoveryTime.Hour,
			Minute:  r.Data.Transformer.RecoveryTime.Minute,
			Second:  r.Data.Transformer.RecoveryTime.Second,
			Reached: r.Data.Transformer.RecoveryTime.Reached,
		},
	}
	t.GameInfo.Region = r.Data.Region
	t.GameInfo.GameUID = r.Data.GameUID
	return t
}

// 验证器
func (t *mysNoteResponse) verify() error {
	return tools.Ifs(t.Retcode == 0, nil, errors.New(t.Message))
}

type mysNoteResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		CurrentResin              int    `json:"current_resin"`
		MaxResin                  int    `json:"max_resin"`
		ResinRecoveryTime         string `json:"resin_recovery_time"`
		FinishedTaskNum           int    `json:"finished_task_num"`
		TotalTaskNum              int    `json:"total_task_num"`
		IsExtraTaskRewardReceived bool   `json:"is_extra_task_reward_received"`
		RemainResinDiscountNum    int    `json:"remain_resin_discount_num"`
		ResinDiscountNumLimit     int    `json:"resin_discount_num_limit"`
		CurrentExpeditionNum      int    `json:"current_expedition_num"`
		MaxExpeditionNum          int    `json:"max_expedition_num"`
		Expeditions               []struct {
			AvatarSideIcon string `json:"avatar_side_icon"`
			Status         string `json:"status"`
			RemainedTime   string `json:"remained_time"`
		} `json:"expeditions"`
		CurrentHomeCoin      int    `json:"current_home_coin"`
		MaxHomeCoin          int    `json:"max_home_coin"`
		HomeCoinRecoveryTime string `json:"home_coin_recovery_time"`
		CalendarUrl          string `json:"calendar_url"`
		Transformer          struct {
			Obtained     bool `json:"obtained"`
			RecoveryTime struct {
				Day     int  `json:"Day"`
				Hour    int  `json:"Hour"`
				Minute  int  `json:"Minute"`
				Second  int  `json:"Second"`
				Reached bool `json:"reached"`
			} `json:"recovery_time"`
			Wiki        string `json:"wiki"`
			Noticed     bool   `json:"noticed"`
			LatestJobId string `json:"latest_job_id"`
		} `json:"transformer"`
		Region  string
		GameUID string
	} `json:"data"`
}
