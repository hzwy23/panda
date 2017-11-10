package router

import (
	"github.com/hzwy23/httprouter"
	"net/http"
)

var defaultHttprouter *httprouter.Router

func GetNative(path string,handlerFunc http.HandlerFunc){
	defaultHttprouter.HandlerFunc("GET",path,handlerFunc)
}

func PostNative(path string,handlerFunc http.HandlerFunc){
	defaultHttprouter.HandlerFunc("POST",path,handlerFunc)
}

func DeleteNative(path string,handlerFunc http.HandlerFunc){
	defaultHttprouter.HandlerFunc("DELETE",path,handlerFunc)
}

func PutNative(path string,handlerFunc http.HandlerFunc){
	defaultHttprouter.HandlerFunc("PUT",path,handlerFunc)
}

type Handler = httprouter.Handle

func GET(path string, handle httprouter.Handle) {
	defaultHttprouter.GET(path,handle)
}

func POST(path string ,handle httprouter.Handle){
	defaultHttprouter.POST(path,handle)
}

func DELETE(path string, handle httprouter.Handle){
	defaultHttprouter.DELETE(path,handle)
}

func PUT(path string, handle httprouter.Handle){
	defaultHttprouter.PUT(path,handle)
}

func GetHttpMux() *httprouter.Router {
	return defaultHttprouter
}

func init(){
	defaultHttprouter = httprouter.New()
}