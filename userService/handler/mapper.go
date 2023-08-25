package handler

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"strconv"
	"userService/db"
	"userService/model"
	userService "userService/proto"
	"userService/utils"
)

type Mapper struct {
	db   *gorm.DB
	rdb1 *redis.Client
	rdb2 *redis.Client
	rdb3 *redis.Client
	rdb4 *redis.Client
	rdb5 *redis.Client
	rdb6 *redis.Client
}

func NewMapper() *Mapper {
	return &Mapper{
		db:   db.MysqlDB,
		rdb1: db.RedisDB1,
		rdb2: db.RedisDB2,
		rdb3: db.RedisDB3,
		rdb4: db.RedisDB4,
		rdb5: db.RedisDB5,
		rdb6: db.RedisDB6,
	}
}

func (m *Mapper) getUserInfo(ctx context.Context, targetUserId int64, myUserId int64, token string) (*userService.User, error) {
	var user_ model.User
	err := m.db.First(&user_, targetUserId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	followCountStr, _ := m.rdb2.Get(ctx, strconv.FormatInt(targetUserId, 10)).Result()
	followerCountStr, _ := m.rdb3.Get(ctx, strconv.FormatInt(targetUserId, 10)).Result()
	favoritedCountStr, _ := m.rdb4.Get(ctx, strconv.FormatInt(targetUserId, 10)).Result()
	workCountStr, _ := m.rdb5.Get(ctx, strconv.FormatInt(targetUserId, 10)).Result()
	favCountStr, _ := m.rdb6.Get(ctx, strconv.FormatInt(targetUserId, 10)).Result()

	followCount, _ := strconv.ParseInt(followCountStr, 10, 64)
	followerCount, _ := strconv.ParseInt(followerCountStr, 10, 64)
	favoritedCount, _ := strconv.ParseInt(favoritedCountStr, 10, 64)
	workCount, _ := strconv.ParseInt(workCountStr, 10, 64)
	favCount, _ := strconv.ParseInt(favCountStr, 10, 64)

	isFollow := true
	var follow model.Follow
	err = m.db.Where("follower_id = ? AND followee_id = ?", myUserId, targetUserId).First(&follow).Error
	if err != nil {
		isFollow = false
	}

	user := &userService.User{
		Id:              user_.Id,
		Name:            user_.Name,
		FollowCount:     followCount,
		FollowerCount:   followerCount,
		IsFollow:        isFollow,
		Avatar:          user_.Avatar,
		BackgroundImage: user_.Background_image,
		Signature:       user_.Signature,
		TotalFavorited:  favoritedCount,
		WorkCount:       workCount,
		FavoriteCount:   favCount,
	}
	return user, nil
}

func (m *Mapper) getIdByUsername(ctx context.Context, username string) (int64, error) {
	var user model.User
	err := m.db.Where("name = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return -1, nil
	}
	return user.Id, err
}

func (m *Mapper) createUser(ctx context.Context, username string, password string) (int64, error) {
	user := model.User{
		Name:             username,
		Avatar:           "http://kasperxms.xyz:9000/avatar/c.jpg",
		Background_image: "http://kasperxms.xyz:9000/avatar/IMG_5067.jpg",
		Signature:        "这个用户还没有个人签名",
	}
	cipheredPassword := utils.Sha256(password)
	err := m.rdb1.Set(ctx, username, cipheredPassword, 0).Err()
	if err != nil {
		return -1, err
	}
	result := m.db.Select("name", "avatar", "background", "signature").Create(&user)
	return user.Id, result.Error
}

func (m *Mapper) initUserCounts(ctx context.Context, userId int64) error {
	userIdStr := strconv.FormatInt(userId, 10)
	err := m.rdb2.Set(ctx, userIdStr, "0", 0).Err()
	err = m.rdb3.Set(ctx, userIdStr, "0", 0).Err()
	err = m.rdb4.Set(ctx, userIdStr, "0", 0).Err()
	err = m.rdb5.Set(ctx, userIdStr, "0", 0).Err()
	err = m.rdb6.Set(ctx, userIdStr, "0", 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (m *Mapper) isPasswordCorrect(ctx context.Context, username string, password string) (bool, error) {
	cipheredPassword := utils.Sha256(password)
	realPassword, err := m.rdb1.Get(ctx, username).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	if realPassword == cipheredPassword {
		return true, nil
	}
	return false, nil
}
