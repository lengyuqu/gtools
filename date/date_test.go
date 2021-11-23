package date

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestCurrentMicros(t *testing.T) {
	fmt.Printf("微秒级16位时间戳：%d\n", CurrentMicros())
	fmt.Printf("毫秒级13位时间戳：%d\n", CurrentMillis())
	fmt.Printf("秒级10位时间戳：%d\n", CurrentSeconds())
	fmt.Printf("当前时间：%s \n", Now())
	fmt.Printf("当前日期：%s \n", Today())
}

func TestGtime(t *testing.T) {
	array := make([]map[string]interface{}, 2)
	array[0] = map[string]interface{}{
		"no":    "n1",
		"ctime": Gtime(time.Now()),
	}

	array[1] = map[string]interface{}{
		"no":    "n2",
		"ctime": Gtime(time.Now()),
	}

	jsonBytes, err := json.Marshal(array)
	if err != nil {
		t.Errorf("Gtime test error")
	}

	t.Log("Gtime 序列化值为：", string(jsonBytes))
}
