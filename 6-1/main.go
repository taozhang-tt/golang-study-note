package main

import "fmt"

func main() {
	main1()
}

type N int

func (n N) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("p: %p, %v\n", n, *n)
}

func main1() {
	var a N = 25

	a.value()
	a.pointer()

	fmt.Printf("a: %p, %v\n", &a, a)
}
