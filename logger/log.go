package logger

import (
	"fmt"
	"log"
)

var (
	// TRACE 记录所有日志
	TRACE *log.Logger

	// INFO 记录重要信息日志
	INFO *log.Logger

	// DEBUG 记录调试信息日志
	DEBUG *log.Logger

	// ERROR 记录错误信息日志
	ERROR *log.Logger
)

// CreateLog 创建日志记录
func CreateLog() error {
	tConf := DefaultConf()
	tConf.FileName = "trace"
	tWriter, err := NewGwriterWithStdout(tConf)
	if err != nil {
		return fmt.Errorf(" TRACE log create fail, reason: %s", err.Error())
	}

	TRACE = log.New(tWriter, "[TRACE] ", log.LstdFlags|log.LUTC|log.Lshortfile)

	iConf := DefaultConf()
	iConf.FileName = "info"
	iWriter, err := NewGwriterWithStdout(iConf)
	if err != nil {
		return fmt.Errorf(" INFO log create fail, reason: %s", err.Error())
	}
	INFO = log.New(iWriter, "[INFO] ", log.LstdFlags|log.LUTC|log.Lshortfile)

	dConf := DefaultConf()
	dConf.FileName = "debug"
	dConf.Compress = true
	dWriter, _ := NewGwriterWithStdout(dConf)
	if err != nil {
		return fmt.Errorf(" DEBUG log create fail, reason: %s", err.Error())
	}
	DEBUG = log.New(dWriter, "[DEBUG] ", log.LstdFlags|log.LUTC|log.Lshortfile)

	eConf := DefaultConf()
	eConf.FileName = "error"
	eConf.Compress = true
	eWriter, _ := NewGwriterWithStdout(eConf)
	if err != nil {
		return fmt.Errorf(" ERROR log create fail, reason: %s", err.Error())
	}
	ERROR = log.New(eWriter, "[ERROR] ", log.LstdFlags|log.LUTC|log.Lshortfile)

	return nil
}
