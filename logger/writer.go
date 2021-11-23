package logger

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

const (
	// RollNon 日志文件不轮转
	RollNon int = 0

	// RollBySize 按大小轮转
	RollBySize int = 1

	// RollByTime 按时间轮转
	RollByTime int = 2
)

// IGWriter 实现了 io.writer，可以将其传入任何框架以支持日志文件记录
type IGWriter interface {
	io.Writer
	Close() error
}

// GWriter 是一个线程安全的文件写入对象，它实现了IGWriter接口
type GWriter struct {

	// 配置文件指针
	conf *Conf

	// 当前日志文件指针
	f *os.File

	// 当前日志文件绝对路径
	absFile string

	// 写操作同步锁
	sync.Mutex

	// 日期文件创建时间戳
	stamp int64
}

// Write GWriter 实现 IGWriter 中的 io.Writer 接口
func (w *GWriter) Write(b []byte) (n int, err error) {
	w.Lock()
	defer w.Unlock()

	switch w.conf.RollPolicy {
	case RollBySize:
		if info, err := w.f.Stat(); err == nil && info.Size() > w.conf.RollSize {
			if err := w.Rolling(); err != nil {
				panic("GWriter 对象文件大小轮转失败, reason:" + err.Error())
			}
		}
	case RollByTime:
		if _, err := w.f.Stat(); err == nil && time.Now().Unix()-w.stamp > w.conf.RollTime {
			if err := w.Rolling(); err != nil {
				panic("GWriter 对象文件时间轮转失败, reason:" + err.Error())
			}
		}
	default:
	}

	return w.f.Write(b)
}

// Close GWriter 实现 IGWriter 中的 Close 接口
func (w *GWriter) Close() error {
	w.Lock()
	defer w.Unlock()
	return w.f.Close()
}

// NewGwriter 创建 Gwriter
func NewGwriter(c *Conf) (IGWriter, error) {
	if c.FilePath == "" || c.FileName == "" {
		return nil, errors.New("Gwriter config is error, invalid arguments")
	}

	if err := os.MkdirAll(c.FilePath, 0700); err != nil {
		return nil, fmt.Errorf("Gwriter config create [%s] error, reason: %s)", c.FilePath, err.Error())
	}

	absFilename := path.Join(c.FilePath, c.FileName) + ".log"
	file, err := os.OpenFile(absFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.FileMode(0644))
	if err != nil {
		return nil, err
	}

	var w IGWriter
	w = &GWriter{
		f:       file,
		absFile: absFilename,
		conf:    c,
		stamp:   time.Now().Unix(),
	}

	return w, nil

}

// NewGwriterWithStdout 创建带控制台输出的 MultiWriter
func NewGwriterWithStdout(c *Conf) (io.Writer, error) {
	if c.FilePath == "" || c.FileName == "" {
		return nil, errors.New("Gwriter config is error, invalid arguments")
	}

	if err := os.MkdirAll(c.FilePath, 0700); err != nil {
		return nil, fmt.Errorf("Gwriter config create [%s] error, reason: %s)", c.FilePath, err.Error())
	}

	absFilename := path.Join(c.FilePath, c.FileName) + ".log"
	file, err := os.OpenFile(absFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.FileMode(0644))
	if err != nil {
		return nil, err
	}

	var w IGWriter
	w = &GWriter{
		f:       file,
		absFile: absFilename,
		conf:    c,
		stamp:   time.Now().Unix(),
	}

	writers := []io.Writer{
		w,
		os.Stdout,
	}

	multiWriter := io.MultiWriter(writers...)

	return multiWriter, nil
}

// Rolling 轮换文件
func (w *GWriter) Rolling() error {
	var newName string
	// 不压缩文件名格式：filename.2007010215041517.log
	// 压缩文件名格式：filename.2007010215041517.log.gz
	if w.conf.Compress {
		newName = path.Join(w.conf.FilePath, w.conf.FileName+"."+time.Now().Format(w.conf.TimePatten)+".log.tmp")
		defer func() {
			go func() {
				if err := w.CompressFile(newName); err != nil {
					log.Println("error in compress log file", err)
					return
				}
			}()
		}()
	} else {
		newName = path.Join(w.conf.FilePath, w.conf.FileName+"."+time.Now().Format(w.conf.TimePatten)+".log")
	}

	w.f.Close()

	// 将文件已时间戳重命名
	if err := os.Rename(w.absFile, newName); err != nil {
		return err
	}

	// 已原文件名重新打开一个文件
	newfile, err := os.OpenFile(w.absFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.FileMode(0644))
	if err != nil {
		return err
	}

	w.f = newfile
	w.stamp = time.Now().Unix()

	return nil
}

// CompressFile 文件gz压缩
func (w *GWriter) CompressFile(srcFile string) error {
	f, err := os.Open(srcFile)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	gzTmpName := strings.TrimRight(srcFile, ".tmp") + ".gz"
	cmpfile, err := os.OpenFile(gzTmpName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.FileMode(0644))
	defer cmpfile.Close()
	if err != nil {
		return err
	}
	gw := gzip.NewWriter(cmpfile)
	defer gw.Close()

	if _, err = f.Seek(0, 0); err != nil {
		return err
	}

	if _, err = io.Copy(gw, f); err != nil {
		if errR := os.Remove(gzTmpName); errR != nil {
			return errR
		}
		return err
	}

	f.Close()
	return os.Remove(srcFile) // remove 源文件
}

// Conf 是 Gwriter 的配置文件
type Conf struct {

	// 文件路径 如: C:\\log
	FilePath string `json:"filePath"`

	// 命名日期时间格式：yyyyMMddHHmmss
	TimePatten string

	// 文件名 如：system.log
	FileName string `json:"fileName"`

	// 保留天数，超过后自动清除
	RemainDay int `json:"remainDay"`

	// 文件滚动策略，按大小，按时间等等
	// RollPolicy 策略支持 RollNon, RollBySize, RollByTime
	// RollTime 滚动的时间间隔，单位秒
	// RollSize 滚动的文件大小，单位Byte
	RollPolicy int   `json:"rollPolicy"`
	RollTime   int64 `json:"rollTime"`
	RollSize   int64 `json:"rollSize"`

	// 是否压缩
	Compress bool `json:"compress"`
}

// DefaultConf 创建默认GLog配置
func DefaultConf() *Conf {
	return &Conf{
		FilePath:   "./log",
		FileName:   "sys",
		RemainDay:  -1,
		RollPolicy: RollByTime,
		TimePatten: "200601021504",
		RollTime:   3600 * 24,
		RollSize:   1024 * 1024,
		Compress:   false,
	}
}
