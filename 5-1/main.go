package main

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "无常"
	println(len(s), utf8.RuneCountInString(s))

}

func test1() string {
	var s string
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	return s
}

func test2() string {
	s := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}
	return strings.Join(s, "")
}

func test3() string {
	var b bytes.Buffer
	b.Grow(1000)

	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}
	return b.String()
}

func test4() string {
	var b bytes.Buffer

	//b.Grow(500)
	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}
	return b.String()
}
