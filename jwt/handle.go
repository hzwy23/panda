package jwt

import (
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	handleLock    = new(sync.RWMutex)
	defaultHandle *handle
)

const (
	// 默认情况下从http请求的中读取的key名称
	header string = "Authorization"
)

// jwt操作
type handle struct {
	*jwtConfig
	*customClaims
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
	r.Valid()
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

func (r *handle) SetUserId(userId string) *handle {
	r.lock.Lock()
	r.UserId = userId
	r.lock.Unlock()
	return r
}

func (r *handle) SetOrgUnitId(orgUnitId string) *handle {
	r.lock.Lock()
	r.OrgUnitId = orgUnitId
	r.lock.Unlock()
	return r
}

func (r *handle) SetAuthorities(author string) *handle {
	r.lock.Lock()
	r.Authorities = author
	r.lock.Unlock()
	return r
}

func (r *handle) SetExpiresAt(dt int64) *handle {
	r.lock.Lock()
	r.ExpiresAt = time.Now().Unix() + dt
	r.lock.Unlock()
	return r
}

func (r *handle) SetIssuer(issuer string) *handle {
	r.lock.Lock()
	r.Issuer = issuer
	r.lock.Unlock()
	return r
}

func (r *handle) SetKey(key []byte) *handle {
	r.lock.Lock()
	r.key = key
	r.lock.Unlock()
	return r
}

// 生成token
func (r *handle) Build() (string, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	c := jwt.NewWithClaims(jwt.SigningMethodHS256, r)
	token, err := c.SignedString(r.key)
	if err != nil {
		return "", err
	}
	return token, nil
}

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

// 创建jwtHandle实例对象
func NewHandle(conf *jwtConfig) *handle {
	if conf == nil {
		handleLock.RLock()
		conf = &jwtConfig{
			key:     []byte("https://github.com/hzwy23"),
			ipValid: false,
		}
		handleLock.RUnlock()
	}
	return &handle{
		jwtConfig:    conf,
		customClaims: newClaims(conf),
		lock:         new(sync.RWMutex),
	}
}

// 修改默认的Handle方法
func SetHandle(jwtHandle *handle) {
	handleLock.Lock()
	defaultHandle = jwtHandle
	handleLock.Unlock()
}

func init() {
	defaultHandle = NewHandle(defaultConfig)
}
