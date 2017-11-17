package route

import (
	"net/http"
	"regexp"
	"strings"
)

func RequestIP(r *http.Request) string {
	ip := r.Header.Get("Remote_addr")
	if len(ip) == 0 {
		ip =  r.RemoteAddr
	}
	if strings.HasPrefix(ip,"[::1]"){
		ip =  "127.0.0.1"
	}
	reg,_ := regexp.Compile(`^\d+\.\d+\.\d+\.\d*`)
	return reg.FindString(ip)
}
