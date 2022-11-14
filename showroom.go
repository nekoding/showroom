package showroom

import (
	"encoding/json"
	"net/http"
)

const BaseEndpoint string = "https://www.showroom-live.com/api/"

type LiveInfo struct {
	EnqueteGiftNum        int    `json:"enquete_gift_num"`
	IsEnquete             bool   `json:"is_enquete"`
	IsRecordingProhibited bool   `json:"is_recording_prohibited"`
	LiveID                int    `json:"live_id"`
	IsEnqueteResult       bool   `json:"is_enquete_result"`
	RoomName              string `json:"room_name"`
	BackgroundImageURL    string `json:"background_image_url"`
	AgeVerificationStatus int    `json:"age_verification_status"`
	BcsvrPort             int    `json:"bcsvr_port"`
	VideoType             int    `json:"video_type"`
	BannerImageURL        string `json:"banner_image_url"`
	BannerDestinationURL  string `json:"banner_destination_url"`
	LiveType              int    `json:"live_type"`
	IsFreeGiftOnly        bool   `json:"is_free_gift_only"`
	PremiumRoomType       int    `json:"premium_room_type"`
	BcsvrHost             string `json:"bcsvr_host"`
	LiveStatus            int    `json:"live_status"`
	BcsvrKey              string `json:"bcsvr_key"`
	RoomID                int    `json:"room_id"`
}

type Profile struct {
	PrevLeagueID   interface{}   `json:"prev_league_id"`
	ImageList      []interface{} `json:"image_list"`
	BannerList     interface{}   `json:"banner_list"`
	IsTalkOnline   bool          `json:"is_talk_online"`
	AwardList      interface{}   `json:"award_list"`
	PushSendStatus struct {
	} `json:"push_send_status"`
	PerformerName        string        `json:"performer_name"`
	FollowerNum          int           `json:"follower_num"`
	LiveContinuousDays   int           `json:"live_continuous_days"`
	NextLeagueID         interface{}   `json:"next_league_id"`
	LiveID               int           `json:"live_id"`
	LeagueID             int           `json:"league_id"`
	IsOfficial           bool          `json:"is_official"`
	IsFollow             bool          `json:"is_follow"`
	VoiceList            []interface{} `json:"voice_list"`
	ShowRankSubdivided   string        `json:"show_rank_subdivided"`
	Event                interface{}   `json:"event"`
	IsBirthday           bool          `json:"is_birthday"`
	Description          string        `json:"description"`
	LiveTags             []interface{} `json:"live_tags"`
	GenreID              int           `json:"genre_id"`
	PrevScore            int           `json:"prev_score"`
	YoutubeID            string        `json:"youtube_id"`
	VisitCount           int           `json:"visit_count"`
	RecommendCommentList []struct {
		CreatedAt int    `json:"created_at"`
		Comment   string `json:"comment"`
		User      struct {
			Name  string `json:"name"`
			Image string `json:"image"`
		} `json:"user"`
	} `json:"recommend_comment_list"`
	CurrentLiveStartedAt   int    `json:"current_live_started_at"`
	NextShowRankSubdivided string `json:"next_show_rank_subdivided"`
	ShareTextLive          string `json:"share_text_live"`
	SnsList                []struct {
		Icon string `json:"icon"`
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"sns_list"`
	RecommendCommentsURL    string `json:"recommend_comments_url"`
	ShareURL                string `json:"share_url"`
	RoomURLKey              string `json:"room_url_key"`
	LeagueLabel             string `json:"league_label"`
	IsLiveTagCampaignOpened bool   `json:"is_live_tag_campaign_opened"`
	Avatar                  struct {
		Description string   `json:"description"`
		List        []string `json:"list"`
	} `json:"avatar"`
	ShareURLLive            string `json:"share_url_live"`
	PrevShowRankSubdivided  string `json:"prev_show_rank_subdivided"`
	IsTalkOpened            bool   `json:"is_talk_opened"`
	ImageSquare             string `json:"image_square"`
	RecommendCommentPostURL string `json:"recommend_comment_post_url"`
	GenreName               string `json:"genre_name"`
	RoomName                string `json:"room_name"`
	Birthday                int    `json:"birthday"`
	RoomLevel               int    `json:"room_level"`
	PartyLiveStatus         int    `json:"party_live_status"`
	Party                   struct {
	} `json:"party"`
	EcConfig struct {
		SalesAvailable int           `json:"sales_available"`
		IsExternalEc   int           `json:"is_external_ec"`
		Links          []interface{} `json:"links"`
	} `json:"ec_config"`
	Image                      string `json:"image"`
	RecommendCommentOpenStatus int    `json:"recommend_comment_open_status"`
	MainName                   string `json:"main_name"`
	ViewNum                    int    `json:"view_num"`
	HasMoreRecommendComment    bool   `json:"has_more_recommend_comment"`
	IsPartyEnabled             bool   `json:"is_party_enabled"`
	PremiumRoomType            int    `json:"premium_room_type"`
	NextScore                  int    `json:"next_score"`
	IsOnlive                   bool   `json:"is_onlive"`
	RoomID                     int    `json:"room_id"`
}

type Ranking struct {
	Ranking []struct {
		AvatarID   int    `json:"avatar_id"`
		AvatarURL  string `json:"avatar_url"`
		Name       string `json:"name"`
		Point      int    `json:"point"`
		Order      int    `json:"order"`
		VisitCount int    `json:"visit_count"`
		UserID     int    `json:"user_id"`
		Rank       int    `json:"rank"`
	} `json:"ranking"`
}

func callShowroomApi[T LiveInfo | Profile | Ranking](path string) T {
	var result T
	res, err := http.Get(BaseEndpoint + path)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func GetLiveinfo(roomId string) LiveInfo {
	var path string = "live/live_info?room_id=" + roomId
	result := callShowroomApi[LiveInfo](path)
	return result
}

func GetProfile(roomId string) Profile {
	var path string = "room/profile?room_id=" + roomId
	result := callShowroomApi[Profile](path)

	return result
}

func GetSummaryRank(roomId string) Ranking {
	var path string = "summary_ranking?room_id=" + roomId
	result := callShowroomApi[Ranking](path)

	return result
}
