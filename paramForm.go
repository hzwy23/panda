package panda

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

// 解析用户提交的请求
func ParseForm(r *http.Request, rst interface{}, jsonName ...string) error {
	r.ParseForm()

	if len(jsonName) != 0 {
		// 解析json数据，json名字必须是可变参数的第一个
		js := r.FormValue(jsonName[0])
		if len(js) == 0 {
			return errors.New("empty")
		}
		err := json.Unmarshal([]byte(js), rst)
		return err
	}

	obj := reflect.ValueOf(rst)
	if obj.Kind() != reflect.Ptr {
		return errors.New("参数值接收对方不是指针类型")
	}

	obj = obj.Elem()
	if obj.Kind() != reflect.Struct {
		return errors.New("参数接收对象不是struct类型")
	}

	// 查看字段个数
	var size = obj.NumField()

	for i := 0; i < size; i++ {
		field := obj.Type().Field(i)
		name := field.Tag.Get("param")
		if len(name) == 0 {
			name = field.Tag.Get("json")
			if len(name) == 0 {
				name = field.Name
			}
		}

		val := r.FormValue(name)
		switch obj.Field(i).Kind() {
		case reflect.String:
			obj.Field(i).SetString(val)
		case reflect.Int64, reflect.Int, reflect.Int32, reflect.Int16, reflect.Int8:
			ret, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("无效的数字类型", val)
			} else {
				obj.Field(i).SetInt(int64(ret))
			}
		case reflect.Uint64, reflect.Uint32, reflect.Uint, reflect.Uint16, reflect.Uint8:
			ret, err := strconv.ParseUint(val, 10, 32)
			if err != nil {
				fmt.Println("无效的数字类型", val)
			} else {
				obj.Field(i).SetUint(ret)

			}
		case reflect.Float64, reflect.Float32:
			ret, err := strconv.ParseFloat(val, 32)
			if err != nil {
				fmt.Println("无效的浮点数", val)
			} else {
				obj.Field(i).SetFloat(ret)
			}
		case reflect.Bool:
			ret, err := strconv.ParseBool(val)
			if err != nil {
				fmt.Println("无效的布尔类型", val)
			} else {
				obj.Field(i).SetBool(ret)
			}
		default:
			fmt.Println("类型无法识别，使用byte填充", val)
			obj.Field(i).SetBytes([]byte(val))
		}
	}
	return nil
}
