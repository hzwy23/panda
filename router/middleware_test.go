package router

import (
	"fmt"
	"net/http"
	"testing"
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
