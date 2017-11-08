package config

import (
	"errors"
	"fmt"
)

type Handle interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

const (
	INI  string = "INI"
	YAML string = "YAML"
	JSON string = "JSON"
)

// 加载配置文件，
// filePath 表示配置文件路径
// typ 是可变参数，当可变参数为空时，表示默认读取INI类型配置文件
// 如果typ不为空，则使用指定的配置文件读取，配置文件类型可以是：
// YARM,JSON,INI
func Load(filePath string, typ ...string) (Handle, error) {
	if len(typ) == 0 {
		fmt.Println("default ini")
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
