package minio_api

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinioClient() *minio.Client {
	// 基本的配置信息
	endpoint := "kasperxms.xyz:9000"
	accessKeyID := "IiGUOYoJPQntPHuipHSx"
	secretAccessKey := "fMpdpvGqEtHs5KKAdyugpZyIRi674X4PD0Y0zSJy"

	// 初始化一个minio客户端对象
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalf("初始化MinioClient错误：%s", err.Error())
	}
	return minioClient
}

func ListBuckets(minioClient *minio.Client) {
	bucketInfos, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		fmt.Println("List Buckets err：", err.Error())
		return
	}
	for index, bucketInfo := range bucketInfos {
		fmt.Printf("List Bucket No {%d}----filename{%s}-----createTime{%s}\n", index+1, bucketInfo.Name, bucketInfo.CreationDate.Format("2006-01-02 15:04:05"))
	}
}

func CheckBuckets(minioClient *minio.Client, bucketName01, bucketName02 string) {
	isExist, err := minioClient.BucketExists(context.Background(), bucketName01)
	if err != nil {
		fmt.Printf("Check %s err：%s", bucketName01, err.Error())
		return
	}
	if isExist {
		fmt.Printf("%s exists!\n", bucketName01)
	} else {
		fmt.Printf("%s not exists!\n", bucketName01)
	}

	isExist, err = minioClient.BucketExists(context.Background(), bucketName02)
	if err != nil {
		fmt.Printf("Check %s err：%s", bucketName02, err.Error())
		return
	}
	if isExist {
		fmt.Printf("%s exists!\n", bucketName02)
	} else {
		fmt.Printf("%s not exists!\n", bucketName02)
	}
}

func RemoveBucket(minioClient *minio.Client, bucketName01 string) {
	isExist, err := minioClient.BucketExists(context.Background(), bucketName01)
	if err != nil {
		fmt.Printf("Check %s err：%s", bucketName01, err.Error())
		return
	}
	if isExist {
		fmt.Printf("%s exists! Start delete....\n", bucketName01)
		// 开始删除逻辑
		err = minioClient.RemoveBucket(context.Background(), bucketName01)
		if err != nil {
			fmt.Printf("Fail to remove %s:%s\n", bucketName01, err.Error())
			return
		}
		fmt.Printf("Success to remove %s\n", bucketName01)
	} else {
		fmt.Printf("%s not exists!\n", bucketName01)
	}
}

func GetObjects(minioClient *minio.Client, bucketName, objectName string) {
	object, err := minioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(object *minio.Object) {
		err := object.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(object)

	localFile, err := os.Create(objectName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(localFile *os.File) {
		err := localFile.Close()
		if err != nil {
			return
		}
	}(localFile)

	if _, err = io.Copy(localFile, object); err != nil {
		fmt.Println(err)
		return
	}
}

func PutObjects(minioClient *minio.Client, bucketName, filePath, objectName string) {
	// 检查bucket是否存在
	isExist, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		fmt.Printf("Check %s err：%s", bucketName, err.Error())
		return
	}
	if !isExist {
		fmt.Printf("%s not exists!\n", bucketName)
	}

	// 对象信息
	contentType := "multipart/form-data"
	fPath := filepath.Join(filePath, objectName)

	// 读取对象流
	fileInfo, err := os.Stat(fPath)
	if err == os.ErrNotExist {
		log.Printf("%s目标文件不存在", fPath)
	}
	f, err := os.Open(fPath)
	if err != nil {
		log.Printf("%s打开目标文件", fPath)
		return
	}

	// 上传文件
	uploadInfo, err := minioClient.PutObject(context.Background(), bucketName,
		objectName, f, fileInfo.Size(),
		minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, uploadInfo.Size)
}

func CopyObjects(minioClient *minio.Client, srcBucket, srcObject, destBucket, destObject string) {
	// Source object
	srcOpts := minio.CopySrcOptions{
		Bucket: srcBucket,
		Object: srcObject,
	}

	// Destination object
	dstOpts := minio.CopyDestOptions{
		Bucket: destBucket,
		Object: destObject,
	}

	// copy
	uploadInfo, err := minioClient.CopyObject(context.Background(), dstOpts, srcOpts)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully copied object:", uploadInfo)
}

func StateObjects(minioClient *minio.Client, bucketName, objectName string) {
	ObjInfo, err := minioClient.StatObject(context.Background(), bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("LastModified:%s\tETag:%s\tContentType:%s\tSize:%d\n",
		ObjInfo.LastModified.Format("2006-01-02 03:04:05"),
		ObjInfo.ETag, ObjInfo.ContentType, ObjInfo.Size)
}

func RemoveObject(minioClient *minio.Client, bucketName, objectName string) {
	opts := minio.RemoveObjectOptions{}
	err := minioClient.RemoveObject(context.Background(), bucketName, objectName, opts)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func RemoveObjects(minioClient *minio.Client, bucketName, objectName string) {
	objectsCh := make(chan minio.ObjectInfo)

	// 注意一般不要自己来构造，直接选择从bucket中查询，查询到的对象放入objectsCh
	for object := range minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{}) {
		if object.Key == objectName {
			objectsCh <- object
		}
	}
	defer close(objectsCh)

	// 删除
	for rErr := range minioClient.RemoveObjects(context.Background(), bucketName, objectsCh, minio.RemoveObjectsOptions{}) {
		fmt.Println("Delete err:", rErr.Err.Error())
	}
}

func UploadLargeFileObjects(minioClient *minio.Client, bucketName, filePath, objectName string) {
	uploadInfo, err := minioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: "application/csv",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded object: ", uploadInfo)
}

func DownloadLargeFileObjects(minioClient *minio.Client, bucketName, filePath, objectName string) {
	err := minioClient.FGetObject(context.Background(), bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
