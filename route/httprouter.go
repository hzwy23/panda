package route

import (
	"github.com/hzwy23/httprouter"
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