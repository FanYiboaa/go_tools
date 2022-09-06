package tools

import (
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

// TimeIdentify 获取当前时间字符串 无符号隔开 到秒数后四位
func TimeIdentify() string {
	return strings.Replace(time.Now().Format("20060102150405.9999"), ".", "", -1)
}

// UUID 通过github.com/satori/go.uuid 生成uuid字符串
func UUID() string {
	return uuid.NewV4().String()
}

// RandomInt64 通过时间戳作为种子生成 n 长度的随机int64
func RandomInt64(n int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(n)
}
