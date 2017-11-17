package panda

import "time"

// CurTime获取系统当前时间，时间格式是：yyyy-mm-dd h24:mm:ss
func CurTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// CurDate获取系统当前日前，日期格式是： yyyy-mm-dd
func CurDate() string {
	return time.Now().Format("2006-01-02")
}
