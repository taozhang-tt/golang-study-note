package main

import (
	"fmt"
	"reflect"
)

func main() {
	var t T
	methodSet(t)
	fmt.Println("---------------")
	methodSet(&t)
}

type S struct{}

type T struct {
	S
}

func (s S) SVal() {}
func (*S) SPtr()  {}
func (T) TVal()   {}
func (*T) TPtr()  {}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)

	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
