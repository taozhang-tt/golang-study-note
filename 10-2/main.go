package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	main2()
}

// 修改目标对象
func main1() {
	a := 100

	va, vp := reflect.ValueOf(a), reflect.ValueOf(&a).Elem()

	fmt.Println(va.CanAddr(), va.CanSet()) // false, false
	fmt.Println(vp.CanAddr(), vp.CanSet()) // true, true
}

// 修改目标字段
func main2() {
	p := new(User)
	v := reflect.ValueOf(p).Elem()

	name := v.FieldByName("Name")
	code := v.FieldByName("code")

	fmt.Printf("name: canaddr= %v, canset= %v\n", name.CanAddr(), name.CanSet()) //name: canaddr= true, canset= true
	fmt.Printf("code: canaddr= %v, canset= %v\n", code.CanAddr(), code.CanSet()) //code: canaddr= true, canset= false

	if name.CanSet() {
		name.SetString("Tom")
	}

	if code.CanSet() {
		*(*int)(unsafe.Pointer(code.UnsafeAddr())) = 100
	}

	fmt.Printf("%+v\n", *p) //{Name:Tom code:0}
}

// 类型推断
func main3() {
	type user struct {
		Name string
		Age  int
	}
	u := user{
		"wuchang",
		27,
	}
	v := reflect.ValueOf(&u)

	if !v.CanInterface() {
		println("CanInterface:fial.")
		return
	}

	p, ok := v.Interface().(*user)

	if !ok {
		println("Interface:fail.")
		return
	}
	p.Age++
	fmt.Printf("%+v\n", u)
}

// 通过 IsNil 判断接口是否为nil
func main4() {
	var a interface{} = nil
	var b interface{} = (*int)(nil)

	fmt.Println(a == nil)                             // true
	fmt.Println(b == nil, reflect.ValueOf(b).IsNil()) // false true
}

type User struct {
	Name string
	code int
}
