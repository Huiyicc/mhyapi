package genshin

import (
	"errors"
	"fmt"
	"github.com/Huiyicc/mhyapi/define"
	"github.com/Huiyicc/mhyapi/request"
	"github.com/Huiyicc/mhyapi/tools"
	json "github.com/json-iterator/go"
)

// GetSignRewardsList 用于获取签到奖励列表
func (t *GenShinCore) GetSignRewardsList() ([]RewardInfo, error) {
	req := request.MYSINFO_API_GAMESIGN_HOME.Copy()
	req.Query = fmt.Sprintf(req.Query, define.ACTID_GENSHIN, t.GameInfo.Region, t.GameInfo.GameUid)
	data, err := request.HttpGet(req, t.getHeaders())
	if err != nil {
		return nil, err
	}
	var resp signRewardsListResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, errors.New(string(data))
	}
	if err = resp.verify(); err != nil {
		return nil, err
	}
	return resp.Data.Awards, nil
}

// SignInfo 用于获取签到信息
func (t *GenShinCore) SignInfo() (SignInfo, error) {
	req := request.MYSINFO_API_GAMESIGN_INFO.Copy()
	req.Query = fmt.Sprintf(req.Query, define.ACTID_GENSHIN, t.GameInfo.Region, t.GameInfo.GameUid)
	data, err := request.HttpGet(req, t.getHeaders())
	if err != nil {
		return SignInfo{}, err
	}
	var resp isSignResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return SignInfo{}, errors.New(string(data))
	}
	if err = resp.verify(); err != nil {
		return SignInfo{}, err
	}
	return resp.Data, err
}

// Sign 签到
func (t *GenShinCore) Sign() (SignInfo, error) {
	var (
		signInfo SignInfo
		err      error
	)
	//先判断是否已经签到
	if signInfo, err = t.SignInfo(); err != nil {
		return SignInfo{}, err
	}
	//已经签到就直接签到成功
	if signInfo.IsSign {
		return signInfo, nil
	}
	req := request.MYSINFO_API_GAMESIGN.Copy()
	req.Body["act_id"] = define.ACTID_GENSHIN
	req.Body["region"] = t.GameInfo.Region
	req.Body["uid"] = t.GameInfo.GameUid
	data, err := request.HttpPost(req, t.getHeaders())
	if err != nil {
		return signInfo, err
	}
	var resp signResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return signInfo, errors.New(string(data))
	}
	if err = resp.verify(); err != nil {
		return signInfo, err
	}
	if resp.Data.Success == 1 {
		return signInfo, errors.New("触发验证码,请手动签到")
	}
	return signInfo, nil
}

type isSignResponse struct {
	Retcode int      `json:"retcode"`
	Message string   `json:"message"`
	Data    SignInfo `json:"data"`
}

//验证器
func (t *isSignResponse) verify() error {
	return tools.Ifs(t.Retcode == 0, nil, errors.New(t.Message))
}

type SignInfo struct {
	TotalSignDay  int    `json:"total_sign_day"`  //总签到天数
	Today         string `json:"today"`           //今天
	IsSign        bool   `json:"is_sign"`         //是否已经签到
	FirstBind     bool   `json:"first_bind"`      //是否第一次绑定
	IsSub         bool   `json:"is_sub"`          //
	MonthFirst    bool   `json:"month_first"`     //第一个月
	SignCntMissed int    `json:"sign_cnt_missed"` //漏签天数
}

type signResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Code      string `json:"code"`
		RiskCode  int    `json:"risk_code"`
		Gt        string `json:"gt"`
		Challenge string `json:"challenge"`
		Success   int    `json:"success"`
	} `json:"data"`
}

//验证器
func (t *signResponse) verify() error {
	return tools.Ifs(t.Retcode == 0, nil, errors.New(t.Message))
}

type signRewardsListResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Month  int          `json:"month"`
		Awards []RewardInfo `json:"awards"`
		Resign bool         `json:"resign"`
	} `json:"data"`
}

// RewardInfo 签到奖励详细信息
type RewardInfo struct {
	Icon string `json:"icon"`
	Name string `json:"name"`
	Cnt  int    `json:"cnt"`
}

//验证器
func (t *signRewardsListResponse) verify() error {
	return tools.Ifs(t.Retcode == 0, nil, errors.New(t.Message))
}
