package main

import (
	"runtime"
)

func main() {
	main2()
}

func main1() {
	exit := make(chan struct{})

	go func() {
		defer close(exit)
		defer println("a")

		func() {
			defer func() {
				println("b", recover() == nil)
			}()

			func() {
				println("c")
				runtime.Goexit()
				// return
				println("c done")
			}()
		}()

		println("b done") // 调用栈都被退出了，这里执行不到了
	}()

	<-exit
	println("main exit.")

}

func main2() {
	exit := make(chan struct{})

	go func() {
		defer close(exit)
		defer println("a")

		func() {
			defer func() {
				println("b", recover() == nil)
			}()

			func() {
				println("c")
				// runtime.Goexit()
				return
				println("c done")
			}()
		}()

		println("b done") // return 是函数正常的返回，不会结束调用栈，这里可以被执行
	}()

	<-exit
	println("main exit.")

}
