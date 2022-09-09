package tools

import (
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Ifs 三目运算的函数
func Ifs[T any](a bool, b, c T) T {
	if a {
		return b
	}
	return c
}

// RandString 生成随机长度字符串
func RandString(strlen int) string {
	var alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = alphanum[rand.Intn(len(alphanum))]
	}
	return string(result)
}

func GetMd5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return strings.ToLower(hex.EncodeToString(m.Sum(nil)))
}

func GetDevicesID() string {
	u := uuid.NewV4()
	return u.String()
}

// PathExists 判断一个文件或文件夹是否存在
// 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
