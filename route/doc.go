// Package route 提供http路由相关服务.
//
// http路由服务包，主要分为两块：
// 	1. httprouter，管理路由注册与分发
// 	2. middleware，管理中间件
//
// 示例代码：
//		package main
//
//		import (
//			"testing"
//			"github.com/hzwy23/panda/route"
//			"net/http"
//			"fmt"
//		)
//
//		func Index(w http.ResponseWriter,r *http.Request){
//			fmt.Println("hello world",w,r)
//		}
//
//		func Index2(w http.ResponseWriter,r *http.Request,ps route.Params){
//			fmt.Println(ps,w,r)
//		}
//
//		func main() {
//			mux := route.GetRouter()
//			route.HandlerFunc("GET","/",Index)
//			route.GET("/:httprouter",Index2)
//			http.ListenAndServe(":8080",mux)
//		}
package route
