package image

import (
	. "github.com/qiniu/api/conf"
	"github.com/qiniu/api/rs"
	"time"
)

const (
	BucketName = "mytutu"
)

var client rs.Client

func init() {
	ACCESS_KEY = "ULtAj7XU42HfRPxrlhRJM-PCNhh6QpcksUhTRmfk"
	SECRET_KEY = "m4L4vEwgUsR0sK6q0-I7845mdEkl3bfA97cXGnAz"
	client = rs.New(nil)
}

func GetUpToken(username string) string {
	//空间名称
	bucketName := BucketName
	// 过期时间 now + 1 h
	var deadline uint32 = uint32(time.Now().Second() + 1*3600)
	//主属表示
	endUser := "sirity-app"
	//回调
	// callbackUrl := ""
	//callbackHost := ""
	// callbackBody := ""
	// callbackBodyType := "application/x-www-form-urlencoded"
	//文件大小 400kb

	var fsizeLimit int64 = 1024 * 200
	//文件类型
	// var mimeLimit string = "image/*"
	//资源名
	var saveKey string = username + "-head"
	// 删除 可能存在的 key
	// DeteleImage(bucketName, saveKey)
	// time.Sleep(time.Second * 10)
	returnBody := `{"key": $(key), "bucket": $(bucket), "domain":".qiniu"}`
	putPolicy := rs.PutPolicy{
		Scope:      bucketName,
		Expires:    deadline,
		InsertOnly: 1,
		SaveKey:    saveKey,
		EndUser:    endUser,
		ReturnBody: returnBody,
		// CallbackUrl:  callbackUrl,
		// CallbackBody: callbackBody,
		FsizeLimit: fsizeLimit,
		// MimeLimit:  mimeLimit,
	}
	return putPolicy.Token(nil)
}

func DeteleImage(bucketName, key string) {
	client.Delete(nil, bucketName, key)
}
