package date

import "time"

// CurrentMicros 获取当前微秒 16位
func CurrentMicros() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// CurrentMillis 当前毫秒 13位
func CurrentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// CurrentSeconds 当前秒级时间戳
func CurrentSeconds() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}

// Now 获取当前时间字符串，格式为：yyyy-MM-dd HH:mm:ss
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Today 获取当前日期字符串，格式为：yyyy-MM-dd
func Today() string {
	return time.Now().Format("2006-01-02")
}
