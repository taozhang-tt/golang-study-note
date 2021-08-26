package main

import "fmt"

func main() {
	a := struct {
		name string
		age  int
	}{"TT", 18}
	fmt.Printf("%T\n", a) // struct { name string; age int }

	type Person struct {
		name string
		age  int
	}
	b := Person{"TT", 18}
	fmt.Printf("%T\n", b) // main.Person

}
