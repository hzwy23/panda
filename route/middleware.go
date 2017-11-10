package route

import (
	"net/http"
)



// 中间件类，使用链表实现，每一个元素都实现了ServeHTTP方法
type Middleware struct {
	handler MiddlewareHandle
	next    *Middleware
}

// 创建并初始化中间件实例对象,中间件类实现了http.Handler接口，
// 所以中间件实例对象可以直接当成golang默认http服务的路由使用
func NewMiddleware(handles ...MiddlewareHandle) *Middleware {
	middle := &Middleware{}
	for _, handle := range handles {
		middle.Add(handle)
	}
	return middle
}

// 中间件实现了net/http包中http.Handler接口，当应用使用golang默认的http程序时，
// 中间件实例对象可以直接与golang默认http对接
func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m != nil {
		m.handler.ServeHTTP(w, r, m.next.ServeHTTP)
	}
}

// 向中间件中追加实例对象，
// 这个实例对象必须实现了MiddlewareHandle接口，中间件对象采用链表来存储，
// 新增加的对象将会追加到中间件最尾部
func (m *Middleware) Add(handle MiddlewareHandle) *Middleware {
	if handle == nil {
		panic("the handle is nil, forbid add nil handle into the Middleware")
	}
	if m.handler == nil {
		m.handler = handle
		m.next = nil
		return m
	}
	h := &Middleware{
		handler: handle,
		next:    nil,
	}
	t := m
	for t.next != nil {
		t = t.next
	}
	t.next = h
	return m
}