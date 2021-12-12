package str

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

// Contains 判断字符串s中是否包含字串sub
//
// 基于strings包实现 strings.Contains
func Contains(s, sub string) bool {
	return strings.Contains(s, sub)
}

// Join 将多个字符串通过连接符拼接起来，例：["abc","def","ghi"] 通过-连接 "abc-def-ghi"
func Join(s []string, sep string) string {
	return strings.Join(s, sep)
}

// Split 将字符串按分隔符分割为多个字符串 例："abc-def-ghi" 按-分割 ["abc","def","ghi"]
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// Trim 在s字符串的头部和尾部去除cutset指定的字符串
func Trim(s string, cutset string) string {
	return strings.Trim(s, cutset)
}

// NumCaptcha 生成数字验证码
//
// digit 验证码位数，取值4-8，默认4位
func NumCaptcha(digit int) string {
	if digit < 4 || digit > 8 {
		digit = 4
	}
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < digit; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// FromBytes 将字节转字符串
func FromBytes(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// ToBytes 将字符串转字节
func ToBytes(str string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// ContainsIn 判断str是否在str数组中
func ContainsIn(str string, strs []string) bool {
	for _, v := range strs {
		if v == str {
			return true
		}
	}
	return false
}
