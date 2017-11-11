// Package jwt 提供 json web token 签名字符串生成与解析服务.
//
// 示例代码：
//
//	package jwt_test
//
//	import (
//		"testing"
//
//		"fmt"
//
//		"github.com/hzwy23/panda/jwt"
//	)
//
//	func TestToken(t *testing.T) {
//		token, err := jwt.GenToken(jwt.NewUserdata().SetUserId("huang").SetOrgunitId("100"))
//		fmt.Println(token, err)
//		claims, err := jwt.ParseToken(token)
//		fmt.Println(claims, err)
//
//		flag := jwt.ValidToken(token)
//		fmt.Println("true or false:", flag)
//
//		handle := jwt.NewHandle(nil)
//		flag = handle.ValidToken(token)
//		fmt.Println("true or false:", flag)
//		token, err = handle.GenToken(jwt.NewUserdata().SetUserId("huang").SetOrgunitId("100"))
//		fmt.Println(token, err)
//		claims, err = jwt.ParseToken(token)
//		fmt.Println(claims, err)
//
//		flag = jwt.ValidToken(token)
//		fmt.Println("true or false:", flag)
//	}
package jwt
