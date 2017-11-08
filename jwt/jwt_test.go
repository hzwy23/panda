package jwt_test

import (
	"testing"

	"fmt"

	"github.com/hzwy23/utils/jwt"
)

func TestToken(t *testing.T) {
	// 创建token
	token, err := jwt.NewHandle(nil).SetUserId("owner").Build()
	fmt.Println(token, err)

	// 解析token
	parse, err := jwt.ParseToken(token)
	fmt.Println("claims is:", parse, err)

	// 使用不同的key解析token，结果显示解析失败
	j := jwt.NewJwtConfig([]byte("hello world"))
	token2, err := jwt.NewHandle(j).SetUserId("owner").Build()
	fmt.Println(jwt.ParseToken(token2))

	fmt.Println(jwt.ValidToken(token))
	fmt.Println(jwt.ValidToken(token2))

	newHandle := jwt.NewHandle(nil).SetUserId("hzwy23")
	newHandle.Build()

}
