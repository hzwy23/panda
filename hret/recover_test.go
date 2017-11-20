package hret_test

import (
	"fmt"
	"testing"
	"github.com/hzwy23/panda/hret"
)

func TestHttpPanic(t *testing.T){
	defer hret.RecvPanic(func(){
		fmt.Println("捕获异常")
	})
	panic("抛出异常")
}