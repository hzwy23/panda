package route

import "net/http"

// 定义中间件MiddlewareHandle接口，
// 实现这个接口的类，其实例对象可以被追加到中间件中。
type MiddlewareHandle interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

// 将http.Handler类型的函数包装成 middlewarehandlefunc类型的函数，
// middlewareHandlefunc类型函数实现了MiddlerwareHandler接口，
// http.Handler只有转换成实现了MiddlewareHandler接口类型的实例，才能被添加到中间件中。
func Wrap(handler http.Handler) MiddlewareHandle{
	return middlewareHandlefunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		handler.ServeHTTP(rw, r)
		next(rw, r)
	})
}

// Handler默认实现，将函数转换成实现了MiddlewareHandle接口的实例对象，
// 转换之后，使其能够被追加到中间件中。
type middlewareHandlefunc func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (fc middlewareHandlefunc) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fc(w, r, next)
}
