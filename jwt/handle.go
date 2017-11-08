package jwt

import (
	"errors"
	"net/http"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

const (
	// 默认情况下从http请求的中读取的key名称
	header string = "Authorization"
)

// jwt操作
type handle struct {
	*jwtConfig
	lock *sync.RWMutex
}

// 根据token校验token是否有效
// 如果token有效，则返回true，否则返回false
func (r *handle) ValidToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return r.key, nil
	})
	return err == nil
}

// 解析token，返回claims信息，
// 如果token无效，则error将会展示错误信息，
// 如果token有效，customClaims将会返回连接用户的信息
func (r *handle) ParseToken(token string) (*customClaims, error) {
	var jclaim = &customClaims{}
	_, err := jwt.ParseWithClaims(token, jclaim, func(*jwt.Token) (interface{}, error) {
		return r.key, nil
	})
	if err != nil {
		return nil, errors.New("parase with claims failed.")
	}
	return jclaim, nil
}

// 从Http的请求中获取Token，校验token是否有效
// 如果token有效，则返回true，如果无效，则返回false
func (r *handle) ValidHttp(req *http.Request) bool {
	token := r.httpToken(req)
	return r.ValidToken(token)
}

// 从http请求中获取token，然后解析token
func (r *handle) ParseHttp(req *http.Request) (*customClaims, error) {
	token := r.httpToken(req)
	return r.ParseToken(token)
}

// 从http中获取token
func (r *handle) httpToken(req *http.Request) string {
	token := ""
	cookie, err := req.Cookie(header)
	if err != nil || len(cookie.Value) == 0 {
		req.ParseForm()
		token = req.FormValue(header)
	} else {
		token = cookie.Value
	}
	return token
}

func (r *handle) SetKey(key []byte) *handle {
	r.lock.Lock()
	r.key = key
	r.lock.Unlock()
	return r
}

// 生成token
func (r *handle) GenToken(private *Private) (string, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	customClaims := newClaims(r.jwtConfig, private)

	c := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	token, err := c.SignedString(r.key)
	if err != nil {
		return "", err
	}
	return token, nil
}

// 创建jwtHandle实例对象
func NewHandle(conf *jwtConfig) *handle {
	if conf == nil {
		handleLock.RLock()
		conf = defaultConfig
		conf.key = []byte("https://github.com/hzwy23")
		handleLock.RUnlock()
	}
	return &handle{
		jwtConfig: conf,
		lock:      new(sync.RWMutex),
	}
}
