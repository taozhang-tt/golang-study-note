package main

import (
	"reflect"
)

func main() {

}

type X int

func main1() {
	var x X = 10
	t := reflect.TypeOf(x)
	if t.Kind() == reflect.Ptr {

	}
}
