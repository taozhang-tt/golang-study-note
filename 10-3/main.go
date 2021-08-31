package main

import (
	"fmt"
	"reflect"
)

func main() {
	main2()
}

// 动态调用方法
func main1() {
	var a X

	v := reflect.ValueOf(&a)
	m := v.MethodByName("Test")

	in := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	}

	out := m.Call(in)

	for _, v := range out {
		fmt.Println(v)
	}

	var ans int
	var err error
	ans, ok := out[0].Interface().(int)
	if !ok {
		return
	}
	err, ok = out[1].Interface().(error)
	if !ok {
		return
	}
	fmt.Println(ans, err)
}

type X struct{}

func (X) Test(x, y int) (int, error) {
	return x + y, fmt.Errorf("err: %d", x+y)
}

// 动态调用变参方法
func main2() {
	var a X

	v := reflect.ValueOf(&a)
	m := v.MethodByName("Format")

	out := m.Call([]reflect.Value{
		reflect.ValueOf("%s= %d"),
		reflect.ValueOf("x"),
		reflect.ValueOf(100),
	})
	fmt.Println(out)

	out = m.CallSlice([]reflect.Value{
		reflect.ValueOf("%s= %d"),
		reflect.ValueOf([]interface{}{"x", 100}),
	})
	fmt.Println(out)

}

func (X) Format(s string, a ...interface{}) string {
	return fmt.Sprintf(s, a...)
}
