package wisdom_portal

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

// 获取项目路径
func BaseDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("get project location failed, err: %v \n", err))
	}
	return dir
}

// md5
func String2md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
