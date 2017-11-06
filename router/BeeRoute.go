package router

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

const (
	BeforeStatic = iota
	BeforeRouter
	BeforeExec
	AfterExec
	FinishRouter
)

type Context = *context.Context
type FilterFunc = beego.FilterFunc
type Controller = beego.Controller

func Post(pattern string, handle FilterFunc) {
	beego.Post(pattern, handle)
}

func Put(pattern string, handle FilterFunc) {
	beego.Put(pattern, handle)
}

func Delete(pattern string, handle FilterFunc) {
	beego.Delete(pattern, handle)
}

func Get(pattern string, handle FilterFunc) {
	beego.Get(pattern, handle)
}

func Any(rootpath string, f FilterFunc) {
	beego.BeeApp.Handlers.Any(rootpath, f)
}

func RESTful(pattern string, c beego.ControllerInterface) {
	beego.Router(pattern, c)
}

func InsertFilter(pattern string, pos int, filter FilterFunc, params ...bool) {
	beego.InsertFilter(pattern,pos,filter,params...)
}

func FindRouter(ctx Context) bool {
	_, yes := beego.BeeApp.Handlers.FindRouter(ctx)
	return yes
}

func SetStaticPath(url string,path string)  {
	beego.SetStaticPath(url, path)
}

func Run()  {
	beego.Run()
}