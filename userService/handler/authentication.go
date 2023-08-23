package handler

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/redis/go-redis/v9"
	"go-micro.dev/v4/errors"
	"strconv"
	"time"
	"userService/db"
)

var jwtSecret = []byte("setting.JwtSecret")
var validTime = 24 * time.Hour

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type Authentication struct {
	rdb *redis.Client
}

func NewAuthentication() *Authentication {
	return &Authentication{rdb: db.RedisDB0}
}

func (auth *Authentication) ValidateToken(ctx context.Context, token string) (int64, error) {
	auth = NewAuthentication()
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
	auth = NewAuthentication()
	succeed, err := auth.rdb.Expire(ctx, token, validTime).Result()
	if err != nil {
		return err
	}
	if !succeed {
		return errors.InternalServerError("500", "update failed")
	}
	return nil
}

func (auth *Authentication) GenerateToken(ctx context.Context, username, password string, userId int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(validTime)

	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "tiktok",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	auth = NewAuthentication()
	auth.rdb.Set(ctx, token, userId, validTime)
	return token, err
}
