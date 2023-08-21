package handler

import (
	"context"
	"feed/dal/db"
	"fmt"
	"time"
)

func Feed(ctx context.Context) error {
	CurrentTime := time.Now().Unix()
	//日期转换
	format := "2006-01-02 15:04:05"
	t := time.Unix(CurrentTime/1000, 0)
	searchTime := t.Format(format)

	//调用数据库方法，查询最近上传的5条视频
	videos := db.QueryVideo(&searchTime, 5)

	fmt.Println(videos)
	return nil
}
