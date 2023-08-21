package handler

import (
	"github.com/Tiktok-Boys/douyin/messageservice/dal/db"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

type FeedService struct{}

func (s *FeedService) Feed(ctx context.Context) error {
	var user_id int64
	user_id = -1
	if req.Token != "" {
		user_id, _ = rpc_server.GetIdByToken(req.Token)
	}

	CurrentTime := time.Now()
	//日期转换
	format := "2006-01-02 15:04:05"
	t := time.Unix(CurrentTime/1000, 0)
	searchTime := t.Format(format)

	//调用数据库方法，查询最近上传的5条视频
	videos := model.NewVideoDaoInstance().QueryVideo(&searchTime, 5)

	for video in videos{

	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
