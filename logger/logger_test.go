package logger_test

import (
	"testing"

	"github.com/hzwy23/panda/logger"
)

func TestNewLogger(t *testing.T) {
	logger.Info("hello world abcd")
	logger.Info("my name is huang zhan wei")
	conf := logger.NewConfig()
	conf.SetName("newLogName.log")
	lg := logger.NewLogger(conf)
	lg.Error("hello world this is new logger")

	conf2 := logger.NewConfig("log.conf")
	ll := logger.NewLogger(conf2)
	ll.Info("hello world")
}
