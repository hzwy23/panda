package panda_test

import (
	"fmt"
	"github.com/hzwy23/panda"
	"testing"
)

func TestCurTime(t *testing.T) {
	fmt.Println(panda.CurTime())
	fmt.Println(panda.CurDate())
}
