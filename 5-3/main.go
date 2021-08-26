package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	main4()
}

func main1() {
	s := []int{0, 1, 2, 3, 4}

	p := &s
	p0 := &s[0]
	p1 := &s[1]

	println(p, p0, p1)

	(*p)[0] += 100
	*p1 += 100

	fmt.Println(s)
}

func main2() {
	x := [][]int{
		{1, 2},
		{10, 20, 30},
		{100},
	}

	fmt.Println(x[1])

	x[2] = append(x[2], 200, 300)
	fmt.Println(x[2])
}

func main3() {
	s := make([]int, 0, 5)

	s1 := append(s, 10)
	s2 := append(s1, 20, 30)
	s3 := append(s2, 40, 50)

	fmt.Println(s, len(s), cap(s), (*reflect.SliceHeader)(unsafe.Pointer(&s)))
	fmt.Println(s1, len(s1), cap(s1), (*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	fmt.Println(s2, len(s2), cap(s2), (*reflect.SliceHeader)(unsafe.Pointer(&s2)))
	fmt.Println(s3, len(s3), cap(s3), (*reflect.SliceHeader)(unsafe.Pointer(&s3)))

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	data := sliceHeader.Data
	dataArr := (*[5]int)(unsafe.Pointer(data)) // 查看底层数组数据
	fmt.Printf("dataArr %v", *dataArr)
}

func main4() {
	s := []int{1, 2, 3, 4, 5}
	s1 := s[1:3]
	fmt.Println(s1, len(s1), cap(s1), (*reflect.SliceHeader)(unsafe.Pointer(&s1)))
}
