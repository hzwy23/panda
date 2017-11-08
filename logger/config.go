package logger

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/hzwy23/utils/config"
)

// 日志文件配置类
type LogConfig struct {
	confPath    string
	logDirPath  string
	logFilePath string
	logLevel    string
	logName     string
	lock        *sync.RWMutex
	msg         []string
}

// 从环境变量中获取日志配置信息的基准目录
const prefixPath = "HBIGDATA_HOME"

// 创建日志文件配置实例对象
// 配置文件路径可以通过环境变量配置，环境变量 HBIGDATA_HOME
// 配置文件在环境变量所指向的目录中，conf目录下
func NewLogConfig() *LogConfig {
	prefix := os.Getenv(prefixPath)
	conf := &LogConfig{
		confPath:    filepath.Join(prefix, "conf", "log.conf"),
		logDirPath:  filepath.Join("temp"),
		logFilePath: filepath.Join("temp", "wisrc.log"),
		logLevel:    "info",
		logName:     "wisrc.log",
		lock:        new(sync.RWMutex),
	}
	return conf.LoadConfigFile()
}

func (r *LogConfig) SetLogLevel(level string) {
	r.lock.Lock()
	r.logLevel = level
	r.lock.Unlock()
}

func (r *LogConfig) SetLogName(name string) {
	r.lock.Lock()
	r.logFilePath = filepath.Join(r.logDirPath, name)
	r.logName = name
	r.lock.Unlock()
}

func (r *LogConfig) SetConfPath(p string) {
	r.lock.Lock()
	r.confPath = p
	r.lock.Unlock()
}

func (r *LogConfig) setLogFilePath(p string) {
	r.lock.Lock()
	r.logFilePath = p
	r.lock.Unlock()
}

func (r *LogConfig) SetLogDirPath(p string) {
	if len(p) == 0 {
		p = "."
	}
	r.lock.Lock()
	r.logDirPath = p
	r.logFilePath = filepath.Join(p, r.logName)
	r.lock.Unlock()
}

func (r *LogConfig) LoadConfigFile() *LogConfig {

	c, err := config.Load(r.confPath)
	if err != nil {
		r.msg = append(r.msg, "读取日志配置信息失败，将会使用默认配置来初始化日志文件，读取日志错误信息是：", err.Error())
	} else {
		logLevel, err := c.Get("level")
		if err != nil {
			r.msg = append(r.msg, "日志级别配置不存在，使用默认日志级别：info")
		} else {
			r.SetLogLevel(logLevel)
		}

		logName, err := c.Get("name")
		if err != nil {
			r.msg = append(r.msg, "日志级别配置不存在，使用默认日志级别：info")
		} else {
			r.setLogFilePath(filepath.Join(r.logDirPath, logName))
			r.SetLogName(logName)
		}
	}
	return r
}

var defaultConfig *LogConfig

func init() {
	prefix := os.Getenv(prefixPath)
	defaultConfig = &LogConfig{
		confPath:    filepath.Join(prefix, "conf", "log.conf"),
		logDirPath:  filepath.Join("temp"),
		logFilePath: filepath.Join("temp", "wisrc.log"),
		logLevel:    "info",
		logName:     "wisrc.log",
		lock:        new(sync.RWMutex),
	}
}
