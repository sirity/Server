package image

import (
	. "github.com/qiniu/api/conf"
	"github.com/qiniu/api/rs"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	BucketName = "sirityportrait"
)

var client rs.Client

func init() {
	ACCESS_KEY = "ecK6qDx0fNV3yrjdGzOY4_glpCZSyNhgoQ7LjvfH"
	SECRET_KEY = "utiFcMNl1dIeKv-XAcNzJp6dfgblOwiW_haC44AA"
	client = rs.New(nil)
}

func GetUpToken(username, id string) (string, string) {
	//空间名称
	bucketName := BucketName
	// 过期时间 now + 0.5 h
	var deadline uint32 = uint32(time.Now().Second() + 1*1800)
	//主属表示
	endUser := "sirity-app"
	//回调
	// callbackUrl := ""
	//callbackHost := ""
	// callbackBody := ""
	// callbackBodyType := "application/x-www-form-urlencoded"
	//文件大小 200kb

	var fsizeLimit int64 = 1024 * 200
	//文件类型
	// var mimeLimit string = "image/*"
	//资源名
	var saveKey string = GetHeadKey(username, id)
	// time.Sleep(time.Second * 10)
	returnBody := `{"key": $(key), "bucket": $(bucket), "domain":"7u2olr.com2.z0.glb.qiniucdn"}`
	putPolicy := rs.PutPolicy{
		Scope:      bucketName + ":" + saveKey,
		InsertOnly: 0,
		Expires:    deadline,
		SaveKey:    saveKey,
		EndUser:    endUser,
		ReturnBody: returnBody,
		// CallbackUrl:  callbackUrl,
		// CallbackBody: callbackBody,
		FsizeLimit: fsizeLimit,
		// MimeLimit:  mimeLimit,
	}
	return putPolicy.Token(nil), saveKey
}

func DeleteImage(bucketName, key string) {
	client.Delete(nil, bucketName, key)
}

func CallbackDeleteImage(username, url string) {
	arrs := strings.Split(url, ".")
	arrs1 := strings.Split(arrs[0], "//")
	tempBucket := arrs1[1]
	id := GetLastHeadImageId(url)
	if id != "-1" {
		key := GetHeadKey(username, id)
		DeleteImage(tempBucket, key)
		log.Printf("delete portrait bucket:%s key %s\n", tempBucket, key)
	}

}

// 获取图片的id病
func GetNextHeadImageId(url string) string {
	//http: //mytutu.qiniudn.com/portrait/lll-head-id-1.jpg
	if !strings.Contains(url, "-head-") {
		return "0"
	}
	arr := strings.Split(url, "-")
	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1] == "head" && arr[i+1] == "1.jpg" {
			id, err := strconv.Atoi(arr[i])
			if err == nil {
				return strconv.Itoa(id + 1)
			} else {
				return "0"
			}
		}
	}
	return "0"
}

// 获取图片的id病
func GetLastHeadImageId(url string) string {
	//http: //mytutu.qiniudn.com/portrait/lll-head-id-1.jpg
	if !strings.Contains(url, "-head-") {
		return "-1"
	}
	arr := strings.Split(url, "-")
	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1] == "head" && arr[i+1] == "1.jpg" {
			id, err := strconv.Atoi(arr[i])
			if id > 0 {
				if err == nil {
					return strconv.Itoa(id - 1)
				} else {
					return "-1"
				}
			} else {
				return "-1"
			}
		}
	}
	return "-1"
}

// 判断Url是否正确
func CheckImageUrl(url string) bool {
	//http: //mytutu.qiniudn.com/lll-head-id-1.jpg
	if !strings.Contains(url, "-head-") {
		return false
	}
	arr := strings.Split(url, "-")
	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1] == "head" && arr[i+1] == "1.jpg" {
			_, err := strconv.Atoi(arr[i])
			if err == nil {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func GetHeadKey(username, id string) string {
	return "portrait/" + username + "-head-" + id + "-1.jpg"
}
