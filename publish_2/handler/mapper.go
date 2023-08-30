package handler

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"os"
	"publish_2/db"
	"publish_2/model"
)

type Mapper struct {
	db          *gorm.DB
	minioClient *minio.Client
}

func NewMapper() *Mapper {
	return &Mapper{
		db:          db.MysqlDB,
		minioClient: db.MinioClient,
	}
}

func (m *Mapper) getUser(ctx context.Context) {

}

func (m *Mapper) UploadVideoBytes(ctx context.Context, videoByte os.File) (string, error) {
	byteReader := bytes.NewReader(videoByte)

}

// 存入一个新的Video，返回Video实例
func (m *Mapper) CreateVideo(ctx context.Context, video *model.Video) (*model.Video, error) {
	/*Video := Video{Name: Videoname, Password: password, FollowingCount: 0, FollowerCount: 0, CreateAt: time.Now()}*/
	result := m.db.Create(&video)

	if result.Error != nil {
		return nil, result.Error
	}

	return video, nil
}

// 根据videoid，查找video实体
func (m *Mapper) FindVideoById(ctx context.Context, videoId int64) (*model.Video, error) {
	video := model.Video{Id: videoId}

	result := m.db.Where("video_id = ?", videoId).First(&video)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &video, err
}
