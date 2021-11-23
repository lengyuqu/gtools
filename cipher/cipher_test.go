package cipher

import "testing"

func TestCipher(t *testing.T) {
	str := "this is a boy"
	t.Log("MD5  摘要结果为：", Md5(str))
	t.Log("SHA1 摘要结果为：", Sha1(str))
}
