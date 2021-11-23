package convert

import (
	"testing"
)

func TestConvert(t *testing.T) {

	bl := true
	it := -8394
	uit := 283432
	ft := 2.34543
	in := 12345

	s := BoolToString(bl)
	t.Logf("布尔类型 true 通过 BoolToString 转为 %s", s)

	s = Int64ToString(int64(it))
	t.Logf("整型类型 -8394 通过 IntToString 转为 %s", s)

	s = Uint64ToString(uint64(uit))
	t.Logf("无符号整型类型 283432 通过 UintToString 转为 %s", s)

	s = FloatToString(ft)
	t.Logf("浮点类型 2.34543 通过 FloatToString 转为 %s", s)

	s = IntToString(in)
	t.Logf("整型 12345 通过 IntToString 转为 %s", s)

	floatString := "3.1415926"
	intString := "-23654"
	uintString := "12554"
	boolString := "false"

	f, e := StringToFloat(floatString)
	if e != nil {
		t.Errorf("StringToFloat test error")
	}

	t.Logf("字符串 3.1415926 转为浮点数为：%f \n", f)

	i, e := StringToInt64(intString)
	if e != nil {
		t.Errorf("StringToInt test error")
	}

	t.Logf("字符串 -23654 转为整型数为：%d \n", i)

	u, e := StringToUint64(uintString)
	if e != nil {
		t.Errorf("StringToUint test error")
	}

	t.Logf("字符串 12554 转为整型数为：%d \n", u)

	i1, e := StringToInt(intString)
	if e != nil {
		t.Errorf("StringToInt test error")
	}

	t.Logf("字符串 -23654 转为整型数为：%d \n", i1)

	b, e := StringToBool(boolString)
	if e != nil {
		t.Errorf("StringToBool test error")
	}

	t.Logf("字符串 false 转为整型数为：%t \n", b)

}
