package main

import (
	"fmt"
	"unsafe"
)

func main() {
}

func main1() {
	fmt.Println("vim-go")
	m := make(map[string]int)
	fmt.Printf("m: %p, %d\n", m, unsafe.Sizeof(m))
}

type User struct {
	Name string
	Age  int
}

func main2() {
	for i := 0; i < 10000; i++ {
		user := User{
			Name: "wuchang",
			Age:  27,
		}
		_ = user
	}
}

func main3() {
	users := make(map[int]User, 10000)
	for i := 0; i < 10000; i++ {
		users[i] = User{
			Name: "wuchang",
			Age:  27,
		}
		_ = users
	}
}
