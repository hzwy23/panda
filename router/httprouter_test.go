package router_test

import (
	"testing"
	"github.com/hzwy23/panda/router"
	"net/http"
	"fmt"
	"github.com/hzwy23/httprouter"
)


func Index(w http.ResponseWriter,r *http.Request){
	fmt.Println("hello world",w,r)
}

func Index2(w http.ResponseWriter,r *http.Request,ps httprouter.Params){
	fmt.Println(ps,w,r)
}

func TestGET(t *testing.T) {
	mux:=router.GetHttpMux()
	router.GetNative("/",Index)
	router.GET("/:httprouter",Index2)
	middle:= router.NewMiddleware(router.Wrap(mux))
	http.ListenAndServe(":8080",middle)
}
