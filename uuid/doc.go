// Package uuid 生成随机唯一标识符uuid.
//
// uuid包提供了两种生成uuid的方法，分别是：
//
//  1. 随机数生成uuid
//
//  2. 随机数加sha1生成uuid
//
// 示例代码：
//
//	package uuid_test
//
//	import (
//		"fmt"
//		"testing"
//
//		"github.com/hzwy23/panda/uuid"
//	)
//
//	func TestUuid(t *testing.T) {
//		fmt.Println(uuid.Random())
//		fmt.Println(uuid.Random())
//		fmt.Println(uuid.Random())
//		fmt.Println(uuid.Random())
//
//		fmt.Println(uuid.UUID())
//		fmt.Println(uuid.UUID())
//		fmt.Println(uuid.UUID())
//	}
package uuid
