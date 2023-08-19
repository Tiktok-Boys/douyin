package model

import "time"

type Like struct {
	User_id    int64
	Video_id   int64
	Created_at time.Time
}
