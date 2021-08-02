package utils

import (
	"crypto/md5"
	"fmt"
)

// 签名算法
func CreateSign(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)
}
