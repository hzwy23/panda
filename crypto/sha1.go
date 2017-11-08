package crypto

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"sync"
)

// Sha1加密
type sha1Handle struct {
	sep  string
	lock *sync.RWMutex
}

// 创建Sha1实例对象
// 在创建对象时，可以指定多个key的分割符，如果不指定分隔符，将会使用默认的分隔符
func NewSHA1(sep ...string) *sha1Handle {
	r := &sha1Handle{
		sep:  "hzwy23",
		lock: new(sync.RWMutex),
	}

	// 如果指定了分隔符，则使用指定的分隔符
	if len(sep) != 0 {
		r.sep = sep[0]
	}
	return r
}

// 使用Sha1加密明文信息
// 第二个是可变参数，可以对多个key进行组合加密
// 多个key使用分隔符进行区分隔离。
func (r *sha1Handle) Sha1(plainText string, keys ...string) string {
	s := sha1.New()
	s.Write([]byte(plainText))

	// 在加密过程中，在明文后边追加值，实现对多个关键字加密
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, str := range keys {
		s.Write([]byte(r.sep + str))
	}
	return fmt.Sprintf("%x", s.Sum(nil))
}

// 修改多个key之间的分割符
func (r *sha1Handle) SetSeparator(sep string) error {
	if len(sep) == 0 {
		return errors.New("separator is empty")
	}
	r.lock.Lock()
	r.sep = sep
	r.lock.Unlock()
	return nil
}

var defaultSHA1 *sha1Handle

func Sha1(plainText string, keys ...string) string {
	return defaultSHA1.Sha1(plainText, keys...)
}

func init() {
	defaultSHA1 = &sha1Handle{
		sep:  "hzwy23",
		lock: new(sync.RWMutex),
	}
}
