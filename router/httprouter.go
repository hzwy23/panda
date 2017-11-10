package router

import (
	"github.com/hzwy23/httprouter"
)

// httprouter路由服务函数接口
type Handle = httprouter.Handle

// 正则路由匹配到的参数列表
// 只有在注册路由服务的时候，使用了正则路由，Params中才会有值
// 并且正则路由只能使用Handler类型的函数一起使用
type Params = httprouter.Params

// httprouter路由
type Router = httprouter.Router

func NewRouter() *Router{
	return httprouter.New()
}