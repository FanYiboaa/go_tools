package tools

import (
	"crypto/md5"
	"fmt"
	"os"
)

// IsFile 判断 path 是否为一个文件
func IsFile(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	} else {
		return !f.IsDir()
	}
}

// IsDir 判断 path 是否为目录
func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	} else {
		return f.IsDir()
	}
}

// IsExists 判断 path 是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	} else {
		return true
	}
}

// FileMd5 获取 path 的md5
func FileMd5(path string) (string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5.Sum(f)), nil
}
