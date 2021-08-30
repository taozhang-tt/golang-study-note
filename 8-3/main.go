package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var d data

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}

type data struct {
	sync.Mutex
}

// test() 的 receiver 为 data 时的输出
// write 0
// read 0
// write 1
// read 1
// write 2
// read 2
// read 3
// write 3
// write 4
// read 4

// test() 的 receiver 为 *data 时的输出
// write 0
// write 1
// write 2
// write 3
// write 4
// read 0
// read 1
// read 2
// read 3
// read 4
func (d data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		println(s, i)
		time.Sleep(time.Second)
	}
}
