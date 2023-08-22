package model

import (
	"time"
)

type User struct {
	Id               int64
	Name             string
	Avatar           string
	Background_image string
	Signature        string
	Created_at       time.Time
	Updated_at       time.Time
	Deleted_at       time.Time
}
