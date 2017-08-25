package hret

import "fmt"

type httpPanicFunc func()

// HttpPanic user for stop panic up.
func HttpPanic(f ...httpPanicFunc) {
	if r := recover(); r != nil {
		fmt.Println("system generator panic.", r)
		for _, val := range f {
			val()
		}
	}
}
