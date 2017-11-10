package logger

import (
	"os"
	"path/filepath"
	"sync"
	"github.com/hzwy23/panda/config"
)

const (
	// 环境变量中，设置程序运行所需要文件的基准路径，
	// 如果环境变量中，WI_HOME变量值为空，
	// 则ApplicationBase值为当前程序运行目录。
	ApplicationBase = "WI_HOME"
	// 配置文件相对于环境变量WI_HOME的相对路径
	DefaultConfigFile = "conf/log.conf"
)

// 日志文件配置类
type Config struct {
	logOutputDir  string
	logName string
	logLevel    string
	lock        *sync.RWMutex
	msg         []string
}

// 创建日志文件配置实例对象
// 配置文件路径可以通过环境变量配置，环境变量 WI_HOME
// 配置文件在环境变量所指向的目录中，conf目录下
func NewConfig(file ...string) *Config {
	fp := DefaultConfigFile
	if len(file) != 0{
		fp = file[0]
	} else if len(file) > 1{
		panic("NewConfig最多只能接收一个参数")
	}

	prefix := os.Getenv(ApplicationBase)
	configFile:=filepath.Join(prefix,fp)

	conf := &Config{
		logOutputDir:  filepath.Join("temp"),
		logLevel:    "info",
		logName:     "wisrc.log",
		lock:        new(sync.RWMutex),
	}
	return conf.Load(configFile)
}

// 设置日志文件级别
// level 值可以是：
// info,debug,warn,error
func (r *Config) SetLevel(level string) {
	r.lock.Lock()
	r.logLevel = level
	r.lock.Unlock()
}

// 设置日志文件名称
func (r *Config) SetName(name string) {
	r.lock.Lock()
	r.logName = name
	r.lock.Unlock()
}

// 设置日志文件输出目录，当目录为空时，则指向当前运行程序所在的目录
func (r *Config) SetLogOutputDir(p string) {
	if len(p) == 0 {
		p = "."
	}
	r.lock.Lock()
	r.logOutputDir = p
	r.lock.Unlock()
}

// 获取日志输出文件地址以及文件名称
func (r *Config)GetLogFile() string{
	return filepath.Join(r.logOutputDir,r.logName)
}

// 加载配置文件，Load将会指定文件中配置信息，覆盖当前实例对象中的日志配置信息
func (r *Config) Load(file string) *Config {
	c, err := config.Load(file)
	if err != nil {
		r.msg = append(r.msg, "读取日志配置信息失败，将会使用默认配置来初始化日志文件，读取日志错误信息是：", err.Error())
	} else {
		logLevel, err := c.Get("level")
		if err != nil {
			r.msg = append(r.msg, "日志级别配置不存在，使用默认日志级别：info")
		} else {
			r.SetLevel(logLevel)
		}

		output, err := c.Get("output")
		if err != nil {
			r.msg = append(r.msg, "没有指定日志文件输出路径，采用采用默认输出路径：./tmp")
		} else {
			r.SetLogOutputDir(output)
		}

		logName, err := c.Get("name")
		if err != nil {
			r.msg = append(r.msg, "没有指定日志文件名称，将会采用默认的日志文件名称：wisrc.log")
		} else {
			r.SetName(logName)
		}
	}
	return r
}

// 默认日志配置
var defaultConfig *Config

func init() {
	defaultConfig = NewConfig()
}
