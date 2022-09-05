package mhyapp

import (
	"errors"
	"fmt"
	"github.com/Huiyicc/mhyapi/request"
	"github.com/Huiyicc/mhyapi/tools"
	json "github.com/json-iterator/go"
)

// GetPostsList 用于获取某分区帖子列表
func (t *AppCore) GetPostsList(forumID string, pageSize int) ([]AppForumInfo, error) {
	req := request.MIHOYOAPP_API_POSTS_LIST.Copy()
	req.Query = fmt.Sprintf(req.Query, forumID, pageSize)
	data, err := request.HttpGet(req, t.getHeaders())
	if err != nil {
		return nil, err
	}
	var resp getPostsListRequest
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	if err = resp.verify(); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}

// PostDetail 看帖
func (t *AppCore) PostDetail(postID string) (*AppForumInfo, error) {
	req := request.MIHOYOAPP_API_POSTS_DETAIL.Copy()
	req.Query = fmt.Sprintf(req.Query, postID)
	data, err := request.HttpGet(req, t.getHeaders())
	if err != nil {
		return nil, err
	}
	var resp getPostsInfoRequest
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, errors.New(string(data))
	}
	return &resp.Data.Post, resp.verify()
}

// PostVote 点赞帖子
// TODO: 未完成
func (t *AppCore) PostVote(postID string, isCancel bool) error {
	req := request.MIHOYOAPP_API_POSTS_LIKE.Copy()
	req.Body["post_id"] = postID
	req.Body["is_cancel"] = isCancel
	header := t.getHeaders().Clone()
	header["DS"] = []string{tools.GetDs(false)}
	data, err := request.HttpPost(req, t.headers)
	if err != nil {
		return err
	}
	var resp forumLikeResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return errors.New(string(data))
	}
	if resp.Retcode != 0 {
		return errors.New(resp.Message)
	}
	return nil
}

type forumLikeResponse struct {
	Retcode int
	Message string
}

type getPostsListRequest struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List     []AppForumInfo `json:"list"`
		LastId   string         `json:"last_id"`
		IsLast   bool           `json:"is_last"`
		IsOrigin bool           `json:"is_origin"`
	} `json:"data"`
}

type getPostsInfoRequest struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Post AppForumInfo `json:"post"`
	} `json:"data"`
}

func (t *getPostsInfoRequest) verify() error {
	return tools.Ifs(t.Retcode == 0, nil, errors.New(t.Message))
}

func (t *getPostsListRequest) verify() error {
	return tools.Ifs(t.Retcode == 0, nil, errors.New(t.Message))
}

type AppForumInfo struct {
	Post struct {
		GameId     int      `json:"game_id"`
		PostId     string   `json:"post_id"`
		FForumId   int      `json:"f_forum_id"`
		Uid        string   `json:"uid"`
		Subject    string   `json:"subject"`
		Content    string   `json:"content"`
		Cover      string   `json:"cover"`
		ViewType   int      `json:"view_type"`
		CreatedAt  int      `json:"created_at"`
		Images     []string `json:"images"`
		PostStatus struct {
			IsTop      bool `json:"is_top"`
			IsGood     bool `json:"is_good"`
			IsOfficial bool `json:"is_official"`
		} `json:"post_status"`
		TopicIds               []int         `json:"topic_ids"`
		ViewStatus             int           `json:"view_status"`
		MaxFloor               int           `json:"max_floor"`
		IsOriginal             int           `json:"is_original"`
		RepublishAuthorization int           `json:"republish_authorization"`
		ReplyTime              string        `json:"reply_time"`
		IsDeleted              int           `json:"is_deleted"`
		IsInteractive          bool          `json:"is_interactive"`
		StructuredContent      string        `json:"structured_content"`
		StructuredContentRows  []interface{} `json:"structured_content_rows"`
		ReviewId               int           `json:"review_id"`
		IsProfit               bool          `json:"is_profit"`
		IsInProfit             bool          `json:"is_in_profit"`
		UpdatedAt              int           `json:"updated_at"`
		DeletedAt              int           `json:"deleted_at"`
		PrePubStatus           int           `json:"pre_pub_status"`
		CateId                 int           `json:"cate_id"`
	} `json:"post"`
	Forum struct {
		Id        int         `json:"id"`
		Name      string      `json:"name"`
		Icon      string      `json:"icon"`
		GameId    int         `json:"game_id"`
		ForumCate interface{} `json:"forum_cate"`
	} `json:"forum"`
	Topics []struct {
		Id            int    `json:"id"`
		Name          string `json:"name"`
		Cover         string `json:"cover"`
		IsTop         bool   `json:"is_top"`
		IsGood        bool   `json:"is_good"`
		IsInteractive bool   `json:"is_interactive"`
		GameId        int    `json:"game_id"`
		ContentType   int    `json:"content_type"`
	} `json:"topics"`
	User struct {
		Uid           string `json:"uid"`
		Nickname      string `json:"nickname"`
		Introduce     string `json:"introduce"`
		Avatar        string `json:"avatar"`
		Gender        int    `json:"gender"`
		Certification struct {
			Type  int    `json:"type"`
			Label string `json:"label"`
		} `json:"certification"`
		LevelExp struct {
			Level int `json:"level"`
			Exp   int `json:"exp"`
		} `json:"level_exp"`
		IsFollowing bool   `json:"is_following"`
		IsFollowed  bool   `json:"is_followed"`
		AvatarUrl   string `json:"avatar_url"`
		Pendant     string `json:"pendant"`
	} `json:"user"`
	SelfOperation struct {
		Attitude    int  `json:"attitude"`
		IsCollected bool `json:"is_collected"`
	} `json:"self_operation"`
	Stat struct {
		ViewNum     int `json:"view_num"`
		ReplyNum    int `json:"reply_num"`
		LikeNum     int `json:"like_num"`
		BookmarkNum int `json:"bookmark_num"`
		ForwardNum  int `json:"forward_num"`
	} `json:"stat"`
	HelpSys struct {
		TopUp     interface{}   `json:"top_up"`
		TopN      []interface{} `json:"top_n"`
		AnswerNum int           `json:"answer_num"`
	} `json:"help_sys"`
	Cover *struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
		Format string `json:"format"`
		Size   string `json:"size"`
		Crop   *struct {
			X   int    `json:"x"`
			Y   int    `json:"y"`
			W   int    `json:"w"`
			H   int    `json:"h"`
			Url string `json:"url"`
		} `json:"crop"`
		IsUserSetCover bool   `json:"is_user_set_cover"`
		ImageId        string `json:"image_id"`
		EntityType     string `json:"entity_type"`
		EntityId       string `json:"entity_id"`
	} `json:"cover"`
	ImageList []struct {
		Url            string      `json:"url"`
		Height         int         `json:"height"`
		Width          int         `json:"width"`
		Format         string      `json:"format"`
		Size           string      `json:"size"`
		Crop           interface{} `json:"crop"`
		IsUserSetCover bool        `json:"is_user_set_cover"`
		ImageId        string      `json:"image_id"`
		EntityType     string      `json:"entity_type"`
		EntityId       string      `json:"entity_id"`
	} `json:"image_list"`
	IsOfficialMaster bool        `json:"is_official_master"`
	IsUserMaster     bool        `json:"is_user_master"`
	HotReplyExist    bool        `json:"hot_reply_exist"`
	VoteCount        int         `json:"vote_count"`
	LastModifyTime   int         `json:"last_modify_time"`
	RecommendType    string      `json:"recommend_type"`
	Collection       interface{} `json:"collection"`
	VodList          []struct {
		Id          string `json:"id"`
		Duration    int    `json:"duration"`
		Cover       string `json:"cover"`
		Resolutions []struct {
			Url        string `json:"url"`
			Definition string `json:"definition"`
			Height     int    `json:"height"`
			Width      int    `json:"width"`
			Bitrate    int    `json:"bitrate"`
			Size       string `json:"size"`
			Format     string `json:"format"`
			Label      string `json:"label"`
		} `json:"resolutions"`
		ViewNum           int `json:"view_num"`
		TranscodingStatus int `json:"transcoding_status"`
		ReviewStatus      int `json:"review_status"`
	} `json:"vod_list"`
	IsBlockOn     bool          `json:"is_block_on"`
	ForumRankInfo interface{}   `json:"forum_rank_info"`
	LinkCardList  []interface{} `json:"link_card_list"`
}
