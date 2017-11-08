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

const (
	owner    string = "hzwy23"
	duration int64  = 3600
)

func newClaims() *customClaims {
	c := &customClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + duration,
			Issuer:    owner,
		},
	}
	return c
}



//func CheckToken(req *http.Request) bool {
//	cookie, err := req.Cookie("Authorization")
//	token := ""
//	if err != nil || len(cookie.Value) == 0 {
//		token = req.FormValue("Authorization")
//	} else {
//		token = cookie.Value
//	}
//
//	_, err = ParseJwt(token)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	return true
//}
//

//
//func Getclaims(request *http.Request) (*claims, error) {
//	cookie, err := request.Cookie("Authorization")
//	if err != nil || cookie == nil || len(cookie.Value) == 0 {
//		jwt := request.Header.Get("Authorization")
//		return ParseJwt(jwt)
//	}
//	return ParseJwt(cookie.Value)
//}
