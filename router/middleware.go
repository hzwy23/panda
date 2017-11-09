// middleware 参考了negroni项目的中间件设计思路
package router

import (
	"net/http"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (fc HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fc(w, r, next)
}

// 中间件类，使用链表实现，每一个元素都实现了ServeHTTP方法
type middleware struct {
	handler Handler
	next    *middleware
}

// 创建并初始化中间件实例
func NewMiddleware(handles ...Handler) *middleware {
	middle := &middleware{}
	for _, handle := range handles {
		middle.add(handle)
	}
	return middle
}

// 实现net/http默认接口
func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m != nil {
		m.handler.ServeHTTP(w, r, m.next.ServeHTTP)
	}
}

// 在middleware的最尾部追加新的Handle，形成一条链表
func (m *middleware) add(handle Handler) {
	if handle == nil {
		panic("the handle is nil, forbid add nil handle into the middleware")
	}

	if m.handler == nil {
		m.handler = handle
		m.next = nil
		return
	}
	h := &middleware{
		handler: handle,
		next:    nil,
	}
	t := m
	for t.next != nil {
		t = t.next
	}
	t.next = h
}

// 将实现了http.HandlerFunc的实例对象进行包装，转换成实现Handler接口，然后添加到中间件
func Wrap(handler http.Handler) Handler {
	return HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		handler.ServeHTTP(rw, r)
		next(rw, r)
	})
}
