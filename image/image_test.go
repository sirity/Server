package server

import (
	"fmt"
	"testing"
)

type PutExtra struct {
	Params map[string]string //可选，用户自定义参数，必须以 "x:" 开头
	//若不以x:开头，则忽略
	MimeType string //可选，当为 "" 时候，服务端自动判断
	Crc32    uint32
	CheckCrc uint32
	// CheckCrc == 0: 表示不进行 crc32 校验
	// CheckCrc == 1: 对于 Put 等同于 CheckCrc = 2；对于 PutFile 会自动计算 crc32 值
	// CheckCrc == 2: 表示进行 crc32 校验，且 crc32 值就是上面的 Crc32 变量
}

func UpFlie(localFile string) {
	var err error
	var ret io.PutRet
	var extra = &io.PutExtra{
	//Params:    params,
	//MimeType:  mieType,
	//Crc32:     crc32,
	//CheckCrc:  CheckCrc,
	}
	uptoken := GetUpToken("mytutu")
	// ret       变量用于存取返回的信息，详情见 io.PutRet
	// uptoken   为业务服务器生成的上传口令
	// key       为文件存储的标识
	// localFile 为本地文件名
	// extra     为上传文件的额外信息，详情见 io.PutExtra，可选
	err = io.PutFile(nil, &ret, uptoken, "test", localFile, extra)

	if err != nil {
		//上传产生错误
		log.Print("io.PutFile failed:", err)
		return
	}
	//上传成功，处理返回值
	log.Print(ret.Hash, ret.Key)
}

func TestUpImage(t *testing.T) {
	UpFlie("/Users/hong/Desktop/1.png")
}
