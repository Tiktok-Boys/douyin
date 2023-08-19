package model

import "time"

type Video struct {
	Id         int64
	User_id    int64
	Play_url   string
	Cover_url  string
	Title      string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
