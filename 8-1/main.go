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

		println("b done")
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

		println("b done")
	}()

	<-exit
	println("main exit.")

}
