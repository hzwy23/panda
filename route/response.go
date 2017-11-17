package route

import (
	"net/http"
	"net"
	"bufio"
	"errors"
)

// Response封装http请求返回信息，
type Response struct{
	http.ResponseWriter
	// 返沪状态吗
	Status int
	// 是否已经写入返回信息，如果已经写入返沪信息，值为true，没有写入返沪信息，值为false
	Written bool
}

func NewResponse(w http.ResponseWriter) *Response {
	nw:=&Response{
		w,0,false,
	}
	return nw
}

func (t *Response)Write(p []byte)(int,error){
	t.Written = true
	return t.ResponseWriter.Write(p)
}

func (t *Response)WriteHeader(code int){
	if t.Status > 0 {
		return
	}
	t.Status = code
	t.Written = true
	t.ResponseWriter.WriteHeader(code)
}

func (t *Response)CloseNotify()<-chan bool{
	if cn,ok:=t.ResponseWriter.(http.CloseNotifier);ok{
		cn.CloseNotify()
	}
	return nil
}

func (t *Response)Flush(){
	if f,ok:=t.ResponseWriter.(http.Flusher);ok{
		f.Flush()
	}
}

func (t *Response)Hijack()(net.Conn,*bufio.ReadWriter,error){
	if hj,ok:=t.ResponseWriter.(http.Hijacker);ok{
		return hj.Hijack()
	}
	return nil,nil,errors.New("webserver doesn't support hijacking")
}