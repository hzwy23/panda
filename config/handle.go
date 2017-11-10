package config

import (
	"errors"
)

// Handle是配置文件操作接口，对于不同的配置文件，读取和写入的方法不一样，
// 但所有的配置文件操作实现了Handle接口，所以，在操作配置文件时，只需要指定具体的配置文件类型，
// Load函数将会按照这个指定的类型来加载配置文件。
type Handle interface {
	// 读取配置文件中key对应的value,
	// 如果key值存在，则返回key对应的value，error为nil
	// 如果key不存在，则error显示错误信息
	Get(key string) (string, error)

	// 修改key对应的value值，如果key不存在，则新增key，value键值对，
	// 如果key存在，则修改key对应的value值，操作成功，error返回值是nil，
	// 否则error返回错误信息
	Set(key string, value string) error
}

type ConfType string

const (
	INI  ConfType = "INI"
	YAML ConfType = "YAML"
	JSON ConfType = "JSON"
)

// 加载配置文件，
// filePath 表示配置文件路径，
// typ 是可变参数，当可变参数为空时，表示默认读取INI类型配置文件，
// 如果typ不为空，则使用指定的配置文件读取，配置文件类型可以是:
//    YARM
//    JSON
//    INI
func Load(filePath string, typ ...ConfType) (Handle, error) {
	if len(typ) == 0 {
		return createINIConfig(filePath)
	}
	switch typ[0] {
	case INI:
		return createINIConfig(filePath)
	case YAML:
		// TODO
		// YAML 配置文件读取方式
		return createINIConfig(filePath)
	case JSON:
		// TODO
		// JSON 类型配置文件读取方式
		return createINIConfig(filePath)
	}

	// 无效的配置文件类型
	return newEmptyConfig()
}

type emptyConfig struct {
}

func newEmptyConfig() (Handle, error) {
	return nil, errors.New("empty")
}

func (r *emptyConfig) Get(key string) (string, error) {
	return "", errors.New("empty")
}

func (r *emptyConfig) Set(key string, value string) error {
	return errors.New("empty object")
}
