package jwt

import (
	"net/http"
	"sync"
)

var (
	handleLock    = new(sync.RWMutex)
	defaultHandle *handle
)

// 校验token信息是否有效
func ValidToken(token string) bool {
	handleLock.RLock()
	defer handleLock.RUnlock()
	return defaultHandle.ValidToken(token)
}

func ValidHttp(req *http.Request) bool {
	handleLock.RLock()
	defer handleLock.RUnlock()
	return defaultHandle.ValidHttp(req)
}

// 解析token是否有效，如果有效，则返回customClaims实例对象
func ParseToken(token string) (*customClaims, error) {
	handleLock.RLock()
	defer handleLock.RUnlock()
	return defaultHandle.ParseToken(token)
}

func ParseHttp(req *http.Request) (*customClaims, error) {
	handleLock.RLock()
	defer handleLock.RUnlock()
	return defaultHandle.ParseHttp(req)
}

// 修改默认的Handle方法
func SetHandle(jwtHandle *handle) {
	handleLock.Lock()
	defaultHandle = jwtHandle
	handleLock.Unlock()
}

func GenToken(private *Private) (string, error) {
	handleLock.RLock()
	defer handleLock.RUnlock()
	return defaultHandle.GenToken(private)
}

func init() {
	defaultHandle = NewHandle(defaultConfig)
}
