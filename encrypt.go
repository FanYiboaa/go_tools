package go_tools

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"strings"
)

func Sha1(d []byte) string {
	str := fmt.Sprintf("%x", sha1.Sum(d))
	str_list := make([]string, 0, 5)
	for i := 0; i < len(str); i += 8 {
		str_list = append(str_list, str[i:i+8])
	}
	return strings.Join(str_list, "-")
}

func Md5(d []byte) string {
	return fmt.Sprintf("%x", md5.Sum(d))
}
