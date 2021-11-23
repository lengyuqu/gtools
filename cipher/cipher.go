package cipher

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// Md5 摘要算法
//
// 基于 crypto/md5 封装
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1 摘要算法
//
// 基于 crypto/md5 封装
func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
