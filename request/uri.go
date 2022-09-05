package request

import json "github.com/json-iterator/go"

type RequestStruct struct {
	Url   string
	Query string
	Sign  bool
	Body  BodyMap
}

func (t RequestStruct) Copy() RequestStruct {
	return RequestStruct{
		Url:   t.Url,
		Query: t.Query,
		Sign:  t.Sign,
		Body:  t.Body.Copy(),
	}
}

type BodyMap map[string]interface{}

func (t BodyMap) Copy() map[string]interface{} {
	r := make(map[string]interface{})
	for k, v := range t {
		r[k] = v
	}
	return r
}

func (t BodyMap) Get() string {
	data, _ := json.Marshal(t)
	return string(data)
}

var (
	// MYSINFO_API_INDEX 首页宝箱
	MYSINFO_API_INDEX = RequestStruct{
		Url:   "https://api-takumi-record.mihoyo.com/game_record/app/genshin/api/index",
		Query: "role_id=%s&server=%s",
	}

	//MYSINFO_API_SPIRALABYSS 深渊
	MYSINFO_API_SPIRALABYSS = RequestStruct{
		Url:   "https://api-takumi-record.mihoyo.com/game_record/app/genshin/api/spiralAbyss",
		Query: "role_id=%s&server=%s",
	}
	//MYSINFO_API_CHARACTER 角色详情
	MYSINFO_API_CHARACTER = RequestStruct{
		Url: "https://api-takumi-record.mihoyo.com/game_record/app/genshin/api/character",
		Body: map[string]interface{}{
			"role_id": "",
			"server":  "",
		},
	}
	//MYSINFO_API_DAILYNOTE 树脂
	MYSINFO_API_DAILYNOTE = RequestStruct{
		Url:   "https://api-takumi-record.mihoyo.com/game_record/app/genshin/api/dailyNote",
		Query: "role_id=%s&server=%s",
	}
	// MYSINFO_API_BBSSIGN_INFO 签到信息
	MYSINFO_API_BBSSIGN_INFO = RequestStruct{
		Url:   "https://api-takumi.mihoyo.com/event/bbs_sign_reward/info",
		Query: "act_id=e202009291139501&region=%s&uid=%s",
		Sign:  true,
	}
	//MYSINFO_API_BBSSIGN_HOME 签到奖励
	MYSINFO_API_BBSSIGN_HOME = RequestStruct{
		Url:   "https://api-takumi.mihoyo.com/event/bbs_sign_reward/home",
		Query: "act_id=e202009291139501&region=%s&uid=%s",
		Sign:  true,
	}
	// MYSINFO_API_BBSSIGN 签到
	MYSINFO_API_BBSSIGN = RequestStruct{
		Url: "https://api-takumi.mihoyo.com/event/bbs_sign_reward/sign",
		Body: map[string]interface{}{
			"act_id": "e202009291139501",
			"region": "",
			"uid":    "",
		},
		Sign: true,
	}
	// MYSINFO_API_DETAIL 详情
	MYSINFO_API_DETAIL = RequestStruct{
		Url:   "https://api-takumi.mihoyo.com/event/e20200928calculate/v1/sync/avatar/detail",
		Query: "uid=%s&region=%s&avatar_id=%s",
	}
	// MYSINFO_API_YSLEDGER 札记
	MYSINFO_API_YSLEDGER = RequestStruct{
		Url:   "https://hk4e-api.mihoyo.com/event/ys_ledger/monthInfo",
		Query: "month=%s&bind_uid=%s&bind_region=%s",
	}
	// MYSINFO_API_COMPUTE 养成计算器
	MYSINFO_API_COMPUTE = RequestStruct{
		Url: "https://api-takumi.mihoyo.com/event/e20200928calculate/v2/compute",
	}
	// MYSINFO_API_AVATARSKILL 角色技能
	MYSINFO_API_AVATARSKILL = RequestStruct{
		Url:   "https://api-takumi.mihoyo.com/event/e20200928calculate/v1/avatarSkill/list",
		Query: "avatar_id=%s",
	}
	// MIHOYOAPP_API_BINDINGO 获取绑定的游戏信息
	MIHOYOAPP_API_BINDINGO = RequestStruct{
		Url:   "https://api-takumi.mihoyo.com/binding/api/getUserGameRolesByCookie",
		Query: "game_biz=%s",
		Sign:  true,
	}
	// MIHOYOAPP_API_LOGINA app登陆第一阶段
	MIHOYOAPP_API_LOGINA = RequestStruct{
		Url:   "https://webapi.account.mihoyo.com/Api/cookie_accountinfo_by_loginticket",
		Query: "login_ticket=%s",
		Sign:  true,
	}
	// MIHOYOAPP_API_LOGINB app登陆第二阶段
	MIHOYOAPP_API_LOGINB = RequestStruct{
		Url:   "https://api-takumi.mihoyo.com/auth/api/getMultiTokenByLoginTicket",
		Query: "login_ticket=%s&token_types=3&uid=%s",
		Sign:  true,
	}
	// MIHOYOAPP_API_TASKS_LIST app任务列表
	MIHOYOAPP_API_TASKS_LIST = RequestStruct{
		Url: "https://bbs-api.mihoyo.com/apihub/sapi/getUserMissionsState",
		//Url:   "https://api-takumi.mihoyo.com/apihub/wapi/getUserMissionsState",
		Query: "point_sn=myb",
		Sign:  true,
	}
	// MIHOYOAPP_API_SIGN app内讨论区签到
	MIHOYOAPP_API_SIGN = RequestStruct{
		Url:  "https://bbs-api.mihoyo.com/apihub/app/api/signIn",
		Body: make(map[string]interface{}),
	}
	// MIHOYOAPP_API_POSTS_LIST 获取app内某讨论区帖子列表
	MIHOYOAPP_API_POSTS_LIST = RequestStruct{
		Url:   "https://bbs-api.mihoyo.com/post/api/getForumPostList",
		Query: "forum_id=%s&is_good=false&is_hot=false&page_size=%d&sort_type=1",
	}
	// MIHOYOAPP_API_POSTS_DETAIL 看帖
	MIHOYOAPP_API_POSTS_DETAIL = RequestStruct{
		Url:   "https://bbs-api.mihoyo.com/post/api/getPostFull",
		Query: "post_id=%s",
	}
	// MIHOYOAPP_API_POSTS_SHARE 分享帖子
	MIHOYOAPP_API_POSTS_SHARE = RequestStruct{
		Url:   "https://bbs-api.mihoyo.com/apihub/api/getShareConf",
		Query: "entity_id=%s&s&entity_type=1",
		Sign:  true,
	}
	// MIHOYOAPP_API_POSTS_LIKE 点赞帖子
	MIHOYOAPP_API_POSTS_LIKE = RequestStruct{
		Url: "https://bbs-api.mihoyo.com/apihub/sapi/upvotePost",
		Body: map[string]interface{}{
			"post_id":   "",
			"is_cancel": false,
		},
		Sign: true,
	}
)
