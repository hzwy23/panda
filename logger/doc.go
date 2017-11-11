// Package logger 是日志服务组件
//
// 采用uber开源的日志服务组件, 通过文件配置的方式，实现日志文件输出管理.
//
// 示例代码：
//
//	package logger_test
//
//	import (
//		"testing"
//		"github.com/hzwy23/panda/logger"
//	)
//
//	func TestNewLogger(t *testing.T) {
//
//		logger.Info("hello world abcd")
//		logger.Info("my name is huang zhan wei")
//
// 		conf := logger.NewConfig()
//		conf.SetName("newLogName.log")
//		lg := logger.NewLogger(conf)
//		lg.Error("hello world this is new logger")
//
//		conf2 := logger.NewConfig("log.conf")
//		ll := logger.NewLogger(conf2)
//		ll.Info("hello world")
//	}
package logger
