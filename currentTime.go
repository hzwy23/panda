package panda

import "time"
import "errors"

// CurTime获取系统当前时间，时间格式是：yyyy-mm-dd h24:mm:ss
func CurTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// CurDate获取系统当前日前，日期格式是： yyyy-mm-dd
func CurDate() string {
	return time.Now().Format("2006-01-02")
}

func DateFormat(date string, tag string) (string,error){
	if tag == "YYYY-MM-DD" {
		t,err:=time.Parse("2006-01-02",date)
		if err != nil {
			return "",err
		}
		return t.Format("2006-01-02"),nil
	} else if tag == "YYYY-MM-DD HH24:MM:SS" {
		t,err:=time.Parse("2006-01-02 15:04:05",date)
		if err != nil {
			return "",err
		}
		return t.Format("2006-01-02 15:04:05"),nil
	} else if tag == "YYYY-MM-DD HH:MM:SS" {
		t,err:=time.Parse("2006-01-02 03:04:05",date)
		if err != nil {
			return "",err
		}
		return t.Format("2006-01-02 03:04:05"),nil
	} else {
		return "",errors.New("format is wrong.")
	}
}