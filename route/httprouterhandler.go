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

func GET(path string, handle httprouter.Handle) {
	lock.RLock()
	defaultHttprouter.GET(path,handle)
	lock.RUnlock()
}

func POST(path string ,handle httprouter.Handle){
	lock.RLock()
	defaultHttprouter.POST(path,handle)
	lock.RUnlock()
}

func DELETE(path string, handle httprouter.Handle){
	lock.RLock()
	defaultHttprouter.DELETE(path,handle)
	lock.RUnlock()
}

func PUT(path string, handle httprouter.Handle){
	lock.RLock()
	defaultHttprouter.PUT(path,handle)
	lock.RUnlock()
}

func HandlerFunc(method, path string, handler http.HandlerFunc){
	lock.RLock()
	defaultHttprouter.HandlerFunc(method,path,handler)
	lock.RUnlock()
}

// 获取httprouter默认路由
func GetHttpRouter() *httprouter.Router {
	lock.RLock()
	defer lock.RUnlock()
	return defaultHttprouter
}

func SetHttpRouter(r *httprouter.Router){
	lock.Lock()
	defaultHttprouter = r
	lock.Unlock()
}

func init(){
	defaultHttprouter = httprouter.New()
}

