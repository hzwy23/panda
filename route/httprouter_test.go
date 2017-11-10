package route_test

import (
	"testing"
	"github.com/hzwy23/panda/route"
	"net/http"
	"fmt"
)


func Index(w http.ResponseWriter,r *http.Request){
	fmt.Println("hello world",w,r)
}

func Index2(w http.ResponseWriter,r *http.Request,ps route.Params){
	fmt.Println(ps,w,r)
}

func TestGET(t *testing.T) {
	mux := route.GetDefaultRouter()
	route.Handler("GET","/",Index)
	route.GET("/:httprouter",Index2)
	http.ListenAndServe(":8080",mux)
}

func TestNewRouter(t *testing.T) {
	mux := route.NewRouter()
	mux.HandlerFunc("GET","/hello",Index)
	mux.GET("/",Index2)
	http.ListenAndServe(":8090",mux)
}