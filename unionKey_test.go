package panda_test

import (
	"testing"
	"github.com/hzwy23/panda"
	"fmt"
	)

func TestGetKey(t *testing.T) {
	s:=panda.JoinKey("hello world","my name is","c huang zhan wei")
	fmt.Println(s)
	fmt.Println(panda.GetKey(s,4))
}
