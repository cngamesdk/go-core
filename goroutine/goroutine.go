package goroutine

// CreateGoroutine create goroutine and recover the panic
func CreateGoroutine(f func(), panicCallbackFun func(any2 any)) {
	go func() {
		defer func() {
			if err := recover(); err != any(nil) {
				panicCallbackFun(err)
			}
		}()

		f()
	}()
}
