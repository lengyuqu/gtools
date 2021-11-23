package logger

import (
	"sync"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	println(time.Now().Unix())
	if err := CreateLog(); err != nil {
		t.Error("CreateLog test error")
	}

	var w sync.WaitGroup
	w.Add(2)
	go recordDebugLog(&w)
	go recordErrorLog(&w)
	w.Wait()

	conf := DefaultConf()
	conf.FilePath = "D:/testlog/test"

	wg, err := NewGwriter(conf)
	if err != nil {
		t.Error("NewGwriter test error")
	}

	ERROR.SetOutput(wg)
	ERROR.Println("日志存储绝对路径D:/testlog/test")
}

func recordErrorLog(wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		time.Sleep(50 * time.Microsecond)
		ERROR.Println("尝试一下记录[错误]日志", "看看多参数怎么样", "多个几个参数", "房价肯定撒", "范德萨发大水", "范德萨范德萨发达发范德萨范德萨")
	}
	wg.Done()
}

func recordDebugLog(wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		time.Sleep(30 * time.Microsecond)
		DEBUG.Println("尝试一下记录[调试]日志", "没有多少参数", "多个几个参数", "房价肯定撒", "范德萨发大水", "范德萨范德萨发达发范德萨范德萨")
	}
	wg.Done()
}
