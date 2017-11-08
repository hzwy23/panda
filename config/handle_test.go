package config_test

import (
	"fmt"
	"testing"

	"github.com/hzwy23/utils/config"
)

func TestLoad(t *testing.T) {
	c, err := config.Load("./testData/data.txt")
	fmt.Println(c, err)
	c, err = config.Load("./testData/data.txt", config.INI)
	fmt.Println(c, err)
	c, err = config.Load("./testData/data.txt", config.YAML)
	fmt.Println(c, err)
	c, err = config.Load("./testData/data.txt", config.JSON)
	fmt.Println(c, err)

	fmt.Println(c.Set("myyfwe", "demfdsdfo"))
	fmt.Println(c.Get("abc"))
	fmt.Println(c.Get("hello"))
}
