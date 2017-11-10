package jwt

import (
	"net/http"
	"sync"
)

var (
	handleLock    = new(sync.RWMutex)
	defaultHandle *Handle
)

// 校验token信息是否有效
func ValidToken(token string) bool {
	handleLock.RLock()
	defer handleLock.RUnlock()
	return defaultHandle.ValidToken(token)
}

// 校验http请求是否有有效的授权认证信息
func ValidHttp(req *http.Request) bool {
	handleLock.RLock()
	defer handleLock.RUnlock()
	return defaultHandle.ValidHttp(req)
}

//解析token是否有效，如果有效，则返回userdata实例对象
func ParseToken(token string) (*Userdata, error) {
	handleLock.RLock()
	defer handleLock.RUnlock()
	return defaultHandle.ParseToken(token)
}

// 从http请求中获取用户信息，如果用户已经获得授权，则返回用户信息，error为nil，
// 如果用户没有授权信息，或者授权信息已经过期，则errors为错误信息
func ParseHttp(req *http.Request) (*Userdata, error) {
	handleLock.RLock()
	defer handleLock.RUnlock()
	return defaultHandle.ParseHttp(req)
}

// 修改默认的Handle方法
func SetHandle(jwtHandle *Handle) {
	handleLock.Lock()
	defaultHandle = jwtHandle
	handleLock.Unlock()
}

// 生成授权信息
func GenToken(private *Userdata) (string, error) {
	handleLock.RLock()
	defer handleLock.RUnlock()
	return defaultHandle.GenToken(private)
}

func init() {
	defaultHandle = NewHandle(defaultConfig)
}
