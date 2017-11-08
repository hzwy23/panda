package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// 自定义Claims类
type customClaims struct {
	*jwt.StandardClaims
	*Private
}

func newClaims(conf *jwtConfig, private *Private) *customClaims {
	c := &customClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + conf.duration,
			Issuer:    conf.owner,
		},
		Private: private,
	}
	return c
}

type Private struct {
	// 用户账号
	UserId string
	// 用户机构号
	OrgUnitId string
	// 用户角色
	Authorities string `json:"authorities"`
}

func NewPrivate() *Private {
	return &Private{
		Authorities: "",
	}
}

func (r *Private) SetUserId(userId string) *Private {
	r.UserId = userId
	return r
}

func (r *Private) SetOrgunitId(orgId string) *Private {
	r.OrgUnitId = orgId
	return r
}
func (r *Private) SetAuthorities(a string) *Private {
	r.Authorities = a
	return r
}
