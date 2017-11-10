package route

import (
	"github.com/hzwy23/httprouter"
	"net/http"
	"sync"
)


var (
	// 定义httprouter的默认路由
	defaultHttprouter *httprouter.Router
	// 默认路由对象服务所
	lock = new(sync.RWMutex)
)

// 向默认路由实例对象中注册路由，路由请求方式是GET，
// path 是路由地址，
// handle 是响应路由的服务函数
func GET(path string, handle httprouter.Handle) {
	lock.RLock()
	defaultHttprouter.GET(path,handle)
	lock.RUnlock()
}

// 向默认路由实例对象中注册路由，路由请求方式是POST，
// path 是路由地址，
// handle 是响应路由的服务函数
func POST(path string ,handle httprouter.Handle){
	lock.RLock()
	defaultHttprouter.POST(path,handle)
	lock.RUnlock()
}

// 向默认路由实例对象中注册路由，路由请求方式是DELETE，
// path 是路由地址，
// handle 是响应路由的服务函数
func DELETE(path string, handle httprouter.Handle){
	lock.RLock()
	defaultHttprouter.DELETE(path,handle)
	lock.RUnlock()
}

// 向默认路由实例对象中注册路由，路由请求方式是PUT，
// path 是路由地址，
// handle 是响应路由的服务函数
func PUT(path string, handle httprouter.Handle){
	lock.RLock()
	defaultHttprouter.PUT(path,handle)
	lock.RUnlock()
}

// 注册路由，
// 使用http.HandlerFunc类型的函数注册路由，
// method是路由请求方法，
// path 是路由地址，
// handler 响应路由的函数。
func Handler(method, path string, handler http.HandlerFunc){
	lock.RLock()
	defaultHttprouter.Handler(method,path,handler)
	lock.RUnlock()
}

// 获取默认路由实例对象
func GetRouter() *httprouter.Router {
	lock.RLock()
	defer lock.RUnlock()
	return defaultHttprouter
}

// 修改默认路由实例对象
func SetRouter(r *httprouter.Router){
	lock.Lock()
	defaultHttprouter = r
	lock.Unlock()
}

func init(){
	defaultHttprouter = httprouter.New()
}
