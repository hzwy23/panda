//+build go1.9
package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger = zap.SugaredLogger

func NewLogger(conf *Config) *Logger {
	if conf == nil {
		conf = NewConfig()
	}
	// 创建日志参数配置对象
	cfg := zap.NewProductionConfig()

	//判断日志所在目录,是否存在
	_, err := os.Stat(conf.GetLogFile())
	if err != nil {
		if os.IsNotExist(err) {
			// 创建日志目录
			err := os.MkdirAll(conf.logOutputDir, os.ModePerm)
			if err != nil {
				// 日志文件无法创建
				// 使用console作为日志输出
				fmt.Println("创建日志文件目录失败，将会使用控制台输出日志文件信息")
			} else {
				cfg.OutputPaths = []string{conf.GetLogFile()}
				cfg.ErrorOutputPaths = []string{conf.GetLogFile()}
			}
		}
		// 如果日志文件存在,但是无法获取Stat信息
		// 将日志输出到console上
	} else {
		cfg.OutputPaths = []string{conf.GetLogFile()}
		cfg.ErrorOutputPaths = []string{conf.GetLogFile()}
	}

	cfg.EncoderConfig.EncodeTime = iso8601TimeEncoder
	cfg.DisableStacktrace = true

	cfg.Level.UnmarshalText([]byte(conf.logLevel))
	lo, err := cfg.Build()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	log := lo.WithOptions(zap.AddCallerSkip(1)).Sugar()

	// 将上边的错误信息写入日志文件中
	for _, info := range conf.msg {
		log.Warn(info)
	}
	return log
}

func iso8601TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
