package genshin

import (
	"errors"
	"fmt"
	"github.com/Huiyicc/mhyapi/request"
	"github.com/Huiyicc/mhyapi/tools"
	json "github.com/json-iterator/go"
)

// UpdateBindInfo 用于更新绑定信息,参数为账号索引,返回游戏绑定信息
func (t *GenShinCore) UpdateBindInfo(index ...int) (*GameInfo, error) {
	req := request.MIHOYOAPP_API_BINDINGO.Copy()
	req.Query = fmt.Sprintf(req.Query, "hk4e_cn")
	data, err := request.HttpGet(req, t.getHeaders())
	if err != nil {
		return nil, err
	}
	var inde int
	if len(index) > 0 {
		inde = index[0]
	}
	var resp updateBindInfoResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, errors.New(string(data))
	}
	if err = resp.verify(); err != nil {
		return nil, err
	}
	if inde > len(resp.Data.List) {
		return nil, errors.New("账号索引错误")
	}
	t.GameInfo = resp.Data.List[inde]
	return &t.GameInfo, nil
}

func (t *GenShinCore) GetBindInfo() *GameInfo {
	return &t.GameInfo
}

type GameInfo struct {
	GameBiz    string `json:"game_biz"`
	Region     string `json:"region"`
	GameUid    string `json:"game_uid"`
	Nickname   string `json:"nickname"`
	Level      int    `json:"level"`
	IsChosen   bool   `json:"is_chosen"`
	RegionName string `json:"region_name"`
	IsOfficial bool   `json:"is_official"`
}

type updateBindInfoResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []GameInfo `json:"list"`
	} `json:"data"`
}

func (t *updateBindInfoResponse) verify() error {
	return tools.Ifs(t.Retcode == 0, nil, errors.New(t.Message))
}
