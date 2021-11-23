package convert

import "strconv"

// BoolToString 布尔类型转字符串类型
//
// 基于 strconv.FormatBool 实现
func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

// StringToBool 字符串转布尔
//
// 基于 strconv.FormatBool 实现
func StringToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// Int64ToString 整数类型转字符串类型
//
// 基于 strconv.FormatBool 实现
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// StringToInt64 字符串转整数类型
//
// 基于 strconv.ParseInt 实现
//
// 默认转为十进制，其它进制请直接使用 strconv.ParseInt
func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// Uint64ToString 无符号整型类型转字符串类型
//
// 基于 strconv.FormatBool 实现
func Uint64ToString(i uint64) string {
	return strconv.FormatUint(i, 10)
}

// StringToUint64 字符串转无符号整数类型
//
// 基于 strconv.ParseUint 实现
//
// 默认转为十进制，其它进制请直接使用 strconv.ParseUint
func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

// IntToString 整型转字符串类型
//
// 基于 strconv.Itoa 实现
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// StringToInt 字符串转整型
//
// 基于 strconv.Atoi 实现
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// FloatToString 浮点类型转字符串类型
//
// 基于 strconv.FormatFloat 实现
func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// StringToFloat 字符串转浮点数类型
//
// 基于 strconv.ParseFloat 实现
func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
