package logger_test

import (
	"testing"

	"github.com/hzwy23/utils/logger"
)

func TestNewLogger(t *testing.T) {
	logger.Info("hello world abcd")
	logger.Info("my name is huang zhan wei")
	conf := logger.NewLogConfig()
	conf.SetLogName("newLogName.log")
	lg := logger.NewLogger(conf)
	lg.Error("hello world this is new logger")

	conf2 := logger.NewLogConfig()
	conf2.SetLogName("wisrc.log")
	conf2.SetLogDirPath(".")
	ll := logger.NewLogger(conf2)
	ll.Info("hello world")
}
