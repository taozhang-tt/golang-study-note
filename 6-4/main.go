package main

import "fmt"

func main() {
	main1()
	fmt.Println("---------------------------")

	main2()
}

type N int

func (n N) test1() {
	fmt.Printf("test1.n: %p, %v\n", &n, n)
}

func main1() {
	var n N = 100
	p := &n

	n++
	f1 := n.test // 因为test方法的receiver是 N 类型，所以复制n，等于 101

	n++
	f2 := n.test // 还是因为test方法的receiver是 N 类型，这里虽然会复制一个 *p，但绑定的还是 p，而不是 *p，所以是 102

	n++
	fmt.Printf("main.n: %p, %v\n", p, n)

	f1()
	f2()
}

// main.n: 0xc000096010, 103
// test.n: 0xc000096028, 101
// test.n: 0xc000096038, 102

func (n *N) test() {
	fmt.Printf("test2.n: %p, %v\n", n, *n)
}

func main2() {
	var n N = 100
	p := &n

	n++
	f1 := n.test // 因为test方法的receiver是 N 类型，所以复制n，等于 101

	n++
	f2 := n.test // 还是因为test方法的receiver是 N 类型，这里虽然会复制一个 *p，但绑定的还是 p，而不是 *p，所以是 102

	n++
	fmt.Printf("main.n: %p, %v\n", p, n)

	f1()
	f2()
}
