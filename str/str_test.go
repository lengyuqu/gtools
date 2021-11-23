package str

import "testing"

func TestString(t *testing.T) {
	str := "abcdefg"

	c := Contains(str, "bcd")
	if !c {
		t.Error("function Contains() is error")
	}

	t.Log("function Contains() pass ")

	t.Log("4位验证码：", NumCaptcha(1))
	t.Log("5位验证码：", NumCaptcha(5))
	t.Log("6位验证码：", NumCaptcha(6))
	t.Log("7位验证码：", NumCaptcha(7))
	t.Log("8位验证码：", NumCaptcha(8))

}
