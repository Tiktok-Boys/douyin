package handler

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-micro.dev/v4/errors"
	"publish_2/db"
	"strconv"
	"time"
)

var validTime = 24 * time.Hour

type Authentication struct {
	rdb *redis.Client
}

func NewAuthentication() *Authentication {
	return &Authentication{rdb: db.RedisDB0}
}

func (auth *Authentication) ValidateToken(ctx context.Context, token string) (int64, error) {
	val, err := auth.rdb.Get(ctx, token).Result()
	if err != nil || err == redis.Nil || val == "" {
		return -1, errors.NotFound("404", "token is not valid")
	}
	usr_id, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return -1, errors.InternalServerError("500", "convert failed")
	}
	return usr_id, nil
}

func (auth *Authentication) RefreshToken(ctx context.Context, token string) error {
	succeed, err := auth.rdb.Expire(ctx, token, validTime).Result()
	if err != nil {
		return err
	}
	if !succeed {
		return errors.InternalServerError("500", "update failed")
	}
	return nil
}
