package date

import (
	"time"
)

// Gtime 是对time.Time类型的扩展，在序列化与反序列化时支持 yyyy-MM-dd HH:mm:ss 时间格式
// golang 中 time.Time类型序列化 rfc3339 标准时间格式进行序列化的，此方式序列化出的时间字符串不方便阅读
// 因此Gtime重写了 MarshalJSON，UnmarshalJSON，MarshalText，UnmarshalText，String 方法来支持 yyyy-MM-dd HH:mm:ss 的格式
type Gtime time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

// UnmarshalJSON time类型json反序列化
func (t *Gtime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	//fmt.Println(now)
	*t = Gtime(now)
	return
}

// MarshalJSON time类型json序列化
func (t *Gtime) MarshalJSON() ([]byte, error) {
	//时间为空默认格式
	if time.Time(*t).IsZero() {
		return []byte(`""`), nil
	}

	return []byte(`"` + time.Time(*t).Format(timeFormart) + `"`), nil
}

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in RFC 3339 format, with sub-second precision added if present.
func (t Gtime) MarshalText() ([]byte, error) {

	//时间为空默认格式
	if time.Time(t).IsZero() {
		return []byte(`""`), nil
	}

	return []byte(time.Time(t).Format(timeFormart)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time is expected to be in RFC 3339 format.
func (t *Gtime) UnmarshalText(data []byte) error {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	//fmt.Println(now)
	*t = Gtime(now)
	return err
}

func (t Gtime) String() string {
	return time.Time(t).Format(timeFormart)
}
