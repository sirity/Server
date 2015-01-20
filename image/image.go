package image

import (
	. "github.com/qiniu/api/conf"
	"github.com/qiniu/api/io"
	"github.com/qiniu/api/rs"
	"log"
	"time"
)

const (
	BucketName = "mytutu"
)

func init() {
	ACCESS_KEY = "ULtAj7XU42HfRPxrlhRJM-PCNhh6QpcksUhTRmfk"
	SECRET_KEY = "m4L4vEwgUsR0sK6q0-I7845mdEkl3bfA97cXGnAz"
}

func GetUpToken(username string) string {
	//空间名称
	bucketName := BucketName
	// 过期时间 now + 1 h
	var deadline int64 = time.Now().Second() + 1*3600
	//主属表示
	endUser := "sirity-app"
	//回调
	callbackUrl := ""
	callbackHost := ""
	callbackBody := ""
	callbackBodyType := ""
	//文件大小 400kb
	var fsizeLimit int64 = 1024 * 400
	//文件类型
	var mimeLimit string = "image/*"
	//资源名
	var saveKey string = username + "-head"

	putPolicy := rs.PutPolicy{
		Scope:            bucketName,
		DeadLine:         deadline,
		SaveKey:          saveKey,
		EndUser:          endUser,
		CallbackUrl:      callbackUrl,
		CallbackHost:     callbackHost,
		CallbackBody:     callbackBody,
		CallbackBodyType: callbackBodyType,
		FileSizeLimi:     fsizeLimit,
		MimeLimit:        mimeLimit,
	}
	return putPolicy.Token(nil)
}
