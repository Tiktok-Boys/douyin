package model

import "time"

type Follow struct {
	follower_id int64
	followee_id int64
	created_at  time.Time
}

func (Follow) TableName() string {
	return "follow"
}
