package route_test

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/hzwy23/panda/route"
)

type a struct {
}

func (a) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("hello func1")
	next(w, r)
}

type b struct {
}

func (b) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("hello func2")
	next(w, r)
}

func TestNewMiddle(t *testing.T) {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("index")
	})

	handle := route.Wrap(mux)
	mid := route.NewMiddleware(handle, &b{}, &a{})
	http.ListenAndServe(":8080", mid)
}


func TestNewMiddleware(t *testing.T) {
	mux:= route.NewRouter()

	mux.POST("/:name/:bcd",Index2)
	mux.HandlerFunc("POST","/",func(w http.ResponseWriter,r *http.Request){
		fmt.Println("hi HandleFunc")
	})

	handle := route.Wrap(mux)
	mid := route.NewMiddleware(handle, &b{}, &a{})
	http.ListenAndServe(":8080", mid)
}
