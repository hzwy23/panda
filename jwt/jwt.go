package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	*jwt.StandardClaims
	UserId string
	//DomainId    string
	OrgUnitId   string
	Authorities string `json:"authorities"`
	ClientIp    string
}

var (
	key []byte = []byte("hzwy23@163.com-jwt")
)

func GenToken(user_id, org_id string, dt int64, clientIp string) string {
	claims := JwtClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + dt,
			Issuer:    "hzwy23",
		},
		UserId: user_id,
		//DomainId:    domain_id,
		OrgUnitId:   org_id,
		Authorities: "ROLE_ADMIN,AUTH_WRITE,ACTUATOR",
		ClientIp:    clientIp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return ss
}

func Identify(token string) bool {
	_, err := ParseJwt(token)
	if err != nil {
		return false
	}
	return true
}

func CheckToken(req *http.Request) bool {
	cookie, err := req.Cookie("Authorization")
	token := ""
	if err != nil || len(cookie.Value) == 0 {
		token = req.FormValue("Authorization")
	} else {
		token = cookie.Value
	}

	jclaim, err := ParseJwt(token)
	if err != nil {
		fmt.Println(err)
		return false
	}
	reqIp := getIp(req)
	return jclaim.ClientIp == reqIp
}

func ParseJwt(token string) (*JwtClaims, error) {
	var jclaim = &JwtClaims{}
	_, err := jwt.ParseWithClaims(token, jclaim, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err, token)
		return nil, errors.New("parase with claims failed.")
	}
	return jclaim, nil
}

func GetJwtClaims(request *http.Request) (*JwtClaims, error) {
	cookie, err := request.Cookie("Authorization")
	if err != nil || cookie == nil || len(cookie.Value) == 0 {
		jwt := request.Header.Get("Authorization")
		return ParseJwt(jwt)
	}
	return ParseJwt(cookie.Value)
}

func getIp(req *http.Request) string {
	reg, err := regexp.Compile(`:[\d]+$`)
	if err != nil {
		return req.RemoteAddr
	}
	return reg.ReplaceAllString(req.RemoteAddr, "")
}
