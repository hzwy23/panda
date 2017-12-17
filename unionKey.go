package panda

import (
	"strings"
	"errors"
	"strconv"
)

const (
	separator = 0x1f
)

// function JoinKey 联合主键拼接，将多个key拼接成一个字符串，默认分隔符是0x1f
func JoinKey(str ...string) string{
	size := len(str)
	result := ""
	for i := 0; i < size; i++ {
		result += str[i]+ string(separator)
	}
	return strings.TrimSuffix(result, string(separator))
}

// GetKey从联合主键字符串中获取指定顺序的字符串，
// index 为0时，获取第一个key，
// index 大于0且小于联合主键中key个数时，返回序号为index的key，index == 1 表示返回第一个key，
// 当index大于联合主键key个数时，error返回一场信息。
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
