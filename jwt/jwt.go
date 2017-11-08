package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// 自定义Claims类
type customClaims struct {
	*jwt.StandardClaims
	// 用户账号
	UserId string
	// 用户机构号
	OrgUnitId string
	// 用户角色
	Authorities string `json:"authorities"`
}

func newClaims(conf *jwtConfig) *customClaims {
	c := &customClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + conf.duration,
			Issuer:    conf.owner,
		},
	}
	return c
}
