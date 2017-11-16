package panda

import (
	"strings"
	"errors"
	"strconv"
)

const (
	separator = 0x1f
)

func JoinKey(str ...string) string{
	size := len(str)
	result := ""
	for i := 0; i < size; i++ {
		result += str[i]+ string(separator)
	}
	return result
}

func GetKey(str string,index int) (string,error) {
   	s:=strings.Split(str,string(separator))
   	if len(s) == 0{
   		return "",errors.New("no value")
	}
	if index >= len(s) {
		return "",errors.New("index out of range, the length is :"+strconv.Itoa(len(s)-1))
	}
	if index == 0 {
		return s[0],nil
	}
	return s[index-1],nil
}
