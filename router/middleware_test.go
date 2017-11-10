package router

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/hzwy23/httprouter"
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

	handle := Wrap(mux)
	mid := NewMiddleware(handle, &b{}, &a{})
	http.ListenAndServe(":8080", mid)
}

func Index(w http.ResponseWriter,r *http.Request, ps httprouter.Params){
	fmt.Println(ps)
	fmt.Println(w,r)
}

func Index2(w http.ResponseWriter,r *http.Request,ps httprouter.Params){
	fmt.Println(ps)
}

func TestNewMiddleware(t *testing.T) {
	mux:=httprouter.New()

	mux.POST("/:name/:bcd",Index)

	handle := Wrap(mux)
	mid := NewMiddleware(handle, &b{}, &a{})
	http.ListenAndServe(":8080", mid)
}
