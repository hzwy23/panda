package uuid

import (
	"github.com/satori/go.uuid"
)

// 采用随机数的方式生成uuid
func Random() string {
	return uuid.NewV4().String()
}

// 采用随机数和sha1组合的方式生成uuid
func UUID() string {
	uid1 := uuid.NewV1().String()
	uid2 := uuid.NewV4()
	return uuid.NewV5(uid2, uid1).String()
}
