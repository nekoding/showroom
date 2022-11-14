package showroom_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nekoding/showroom"
	"github.com/stretchr/testify/assert"
)

func TestGetLiveInfo(t *testing.T) {
	fakedata := `{
		"enquete_gift_num": 0,
		"is_enquete": false,
		"is_recording_prohibited": false,
		"live_id": 0,
		"is_enquete_result": false,
		"room_name": "Gracia/グラシア（JKT48）",
		"background_image_url": "https://image.showroom-cdn.com/showroom-prod/assets/img/room/background/cinderella_2022.png",
		"age_verification_status": 0,
		"bcsvr_port": 8080,
		"video_type": 0,
		"banner_image_url": "https://image.showroom-cdn.com/showroom-prod/image/room_banner/5fac628e0e8746b270413bc40b2951dc.png",
		"banner_destination_url": "https://campaign.showroom-live.com/akb48_sr/",
		"live_type": 0,
		"is_free_gift_only": false,
		"premium_room_type": 0,
		"bcsvr_host": "online.showroom-live.com",
		"live_status": 1,
		"bcsvr_key": "",
		"room_id": 318208
		}`

	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fakedata))
	}))

	defer mock.Close()

	var liveinfo showroom.LiveInfo
	_ = json.Unmarshal([]byte(fakedata), liveinfo)
	res := showroom.GetLiveinfo(mock.URL)

	assert.Equal(t, liveinfo.RoomName, res.RoomName, "room name must be same")
}

func TestGetProfile(t *testing.T) {
	fakedata := `{
		"prev_league_id": null,
		"image_list": [],
		"banner_list": null,
		"is_talk_online": false,
		"award_list": null,
		"push_send_status": {},
		"performer_name": "",
		"follower_num": 77925,
		"live_continuous_days": 0,
		"next_league_id": null,
		"live_id": 0,
		"league_id": 0,
		"is_official": true,
		"is_follow": false,
		"voice_list": [],
		"show_rank_subdivided": "Excluded",
		"event": null,
		"is_birthday": false,
		"description": "\"Name: Shania Gracia\r\nBirthday: 31 August 1999\r\nBirthplace: Jakarta\r\nBlood type: A\r\nZodiac signs:  Virgo\r\nHobby:Traveling, Photoshoot, Playing ukulele\r\n\r\nTwitter: S_GraciaJKT48\r\nInstagram: jkt48gracia\"",
		"live_tags": [],
		"genre_id": 102,
		"prev_score": 0,
		"youtube_id": "eq0s1atl_K0",
		"visit_count": 0,
		"recommend_comment_list": [
		{
		"created_at": 1667794786,
		"comment": "😍",
		"user": {
		"name": "riska desnita",
		"image": "https://image.showroom-cdn.com/showroom-prod/image/avatar/3.png?v=92"
		}
		},
		{
		"created_at": 1667667279,
		"comment": "Kamu nenyeee? Kamu bertanya-tanya? ",
		"user": {
		"name": "Cha Eun Yu",
		"image": "https://image.showroom-cdn.com/showroom-prod/image/avatar/1038966.png?v=92"
		}
		},
		{
		"created_at": 1667649799,
		"comment": "tetap jadi diri kamu sendiri yg selalu baik sama orang lain, aku selalu dukung kamu",
		"user": {
		"name": "Kitot Asya",
		"image": "https://image.showroom-cdn.com/showroom-prod/image/avatar/1038966.png?v=92"
		}
		}
		],
		"current_live_started_at": 0,
		"next_show_rank_subdivided": "",
		"share_text_live": "Gracia/グラシア（JKT48） Broadcasting!\nhttps://www.showroom-live.com/JKT48_Gracia?t=1668416254",
		"sns_list": [
		{
		"icon": "https://image.showroom-cdn.com/showroom-prod/assets/img/icon/sns/twitter.png",
		"url": "https://www.showroom-live.com/social/twitter/redirect_to_twitter?room_id=318208&user_id=4427473",
		"name": "Twitter"
		},
		{
		"icon": "https://image.showroom-cdn.com/showroom-prod/assets/img/icon/sns/youtube.png",
		"url": "https://www.youtube.com/watch?v=eq0s1atl_K0",
		"name": "Youtube"
		}
		],
		"recommend_comments_url": "https://www.showroom-live.com/room/recommend_comments?room_id=318208",
		"share_url": "https://www.showroom-live.com/room/profile?room_id=318208",
		"room_url_key": "JKT48_Gracia",
		"league_label": "",
		"is_live_tag_campaign_opened": false,
		"avatar": {
		"description": "Send gifts to this performer, and you will receive an original avatar!",
		"list": [
		"https://image.showroom-cdn.com/showroom-prod/image/avatar/1038966.png?v=92"
		]
		},
		"share_url_live": "https://www.showroom-live.com/JKT48_Gracia?t=1668416254",
		"prev_show_rank_subdivided": "",
		"is_talk_opened": false,
		"image_square": "https://image.showroom-cdn.com/showroom-prod/image/room/cover/3026b77b664e3f6acc416ed606d20c23356fe91d18cb4345d53715ac50541912_square_m.png?v=1617298208",
		"recommend_comment_post_url": "https://www.showroom-live.com/room/recommend_comments?room_id=318208#post",
		"genre_name": "Idol",
		"room_name": "Gracia/グラシア（JKT48）",
		"birthday": 84034800,
		"room_level": 628,
		"party_live_status": 0,
		"party": {},
		"ec_config": {
		"sales_available": 0,
		"is_external_ec": 0,
		"links": []
		},
		"image": "https://image.showroom-cdn.com/showroom-prod/image/room/cover/3026b77b664e3f6acc416ed606d20c23356fe91d18cb4345d53715ac50541912_m.jpeg?v=1658303198",
		"recommend_comment_open_status": 1,
		"main_name": "Gracia/グラシア（JKT48）",
		"view_num": 0,
		"has_more_recommend_comment": true,
		"is_party_enabled": true,
		"premium_room_type": 0,
		"next_score": 0,
		"is_onlive": false,
		"room_id": 318208
		}`

	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fakedata))
	}))

	defer mock.Close()

	var profile showroom.Profile
	_ = json.Unmarshal([]byte(fakedata), profile)
	res := showroom.GetProfile(mock.URL)

	assert.Equal(t, profile.MainName, res.MainName, "main name must be same")
}

func TestSummaryRank(t *testing.T) {
	fakedata := `{
		"ranking": [
		{
		"avatar_id": 1057739,
		"avatar_url": "https://placehold.co/200x200",
		"name": "John doe",
		"point": 6163606,
		"order": 1,
		"visit_count": 199,
		"user_id": 4742415,
		"rank": 1
		},
		{
		"avatar_id": 1038222,
		"avatar_url": "https://placehold.co/200x200",
		"name": "Enggar",
		"point": 4210606,
		"order": 2,
		"visit_count": 16,
		"user_id": 4909993,
		"rank": 2
		}
		]
		}`

	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fakedata))
	}))

	defer mock.Close()

	var rank showroom.Ranking
	_ = json.Unmarshal([]byte(fakedata), rank)
	res := showroom.GetSummaryRank(mock.URL)

	for index, data := range res.Ranking {
		t.Run(data.Name, func(t *testing.T) {
			assert.Equal(t, data.Name, rank.Ranking[index].Name, "name must be same")
		})
	}
}
