package hret

type recvPanicFunc func()

// 捕获panic信息，阻止程序因为panic而推出。
func RecvPanic(f ...recvPanicFunc) {
	if r := recover(); r != nil {
		for _, val := range f {
			val()
		}
	}
}
