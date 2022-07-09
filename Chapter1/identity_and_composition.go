package main

import (
	"fmt"
)

type F func(x interface{}) interface{}

func main() {
	myFunc := func(x interface{}) interface{} { return "abcd" }

	// Expect that id . f == f . id == f
	// Output should be abcd, not input
	fmt.Println(compose(id, myFunc)("input"))
	fmt.Println(compose(myFunc, id)("input"))
}

func id(x interface{}) interface{} {
	return x
}

func compose(f, g F) F {
	return func(x interface{}) interface{} { return f(g(x)) }
}