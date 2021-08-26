package main

import "fmt"

func main() {
	main1()
}

type data int

func (d data) String() string {
	return fmt.Sprintf("data: %d", d)
}

func main1() {
	var d data = 15
	var x interface{} = d

	if n, ok := x.(fmt.Stringer); ok { // 转换为更具体的接口类型
		fmt.Println(n)
	}

	if d2, ok := x.(data); ok { // 转换会原始类型
		fmt.Println(d2)
	}
}
