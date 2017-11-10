package route

import (
	"github.com/hzwy23/httprouter"
	"net/http"
)

// 正则路由匹配到的参数列表，
// 只有在注册路由服务的时候，使用了正则路由，Params中才会有值，
// 并且正则路由只能使用Handler类型的函数一起使用。
type Params = httprouter.Params

// NewRouter函数新建一个路由服务实例,
// 路由服务使用httprouter包实现。
func NewRouter() *httprouter.Router{
	return httprouter.New()
}

var (
	// 定义httprouter的默认路由
	defaultHttprouter = NewRouter()
)

// 获取默认路由实例对象
func GetDefaultRouter() *httprouter.Router {
	return defaultHttprouter
}

// 注册路由，
// 使用http.HandlerFunc类型的函数注册路由，
// method是路由请求方法，
// path 是路由地址，
// handler 响应路由的函数。
func Handler(method, path string, handler http.HandlerFunc){
	defaultHttprouter.Handler(method,path,handler)
}

// 向默认路由实例对象中注册路由，路由请求方式是GET，
// path 是路由地址，
// handle 是响应路由的服务函数
func GET(path string, handle httprouter.Handle) {
	defaultHttprouter.GET(path,handle)
}

// 向默认路由实例对象中注册路由，路由请求方式是POST，
// path 是路由地址，
// handle 是响应路由的服务函数
func POST(path string ,handle httprouter.Handle){
	defaultHttprouter.POST(path,handle)
}

// 向默认路由实例对象中注册路由，路由请求方式是DELETE，
// path 是路由地址，
// handle 是响应路由的服务函数
func DELETE(path string, handle httprouter.Handle){
	defaultHttprouter.DELETE(path,handle)
}

// 向默认路由实例对象中注册路由，路由请求方式是PUT，
// path 是路由地址，
// handle 是响应路由的服务函数
func PUT(path string, handle httprouter.Handle){
	defaultHttprouter.PUT(path,handle)
}