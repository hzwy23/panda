package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// 自定义Claims类
type customClaims struct {
	*jwt.StandardClaims
	Userdata *Userdata
}

func newClaims(conf *Config, userData *Userdata) *customClaims {
	c := &customClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + conf.duration,
			Issuer:    conf.owner,
		},
		Userdata: userData,
	}
	return c
}

type Userdata struct {
	// 用户账号
	UserId string
	// 用户机构号
	OrgUnitId string
	// 用户角色
	Authorities string `json:"authorities"`
}

func NewUserdata() *Userdata {
	return &Userdata{
		Authorities: "",
	}
}

func (r *Userdata) SetUserId(userId string) *Userdata {
	r.UserId = userId
	return r
}

func (r *Userdata) SetOrgunitId(orgId string) *Userdata {
	r.OrgUnitId = orgId
	return r
}
func (r *Userdata) SetAuthorities(a string) *Userdata {
	r.Authorities = a
	return r
}
