package main

import (
	"fmt"
	"math/rand"
	"time"
)

type F func(x interface{}) interface{}

type Memoizer interface {
	Memoize(f F) F
}

type memo struct {
	cache map[interface{}]interface{}
}

var _ Memoizer = (*memo)(nil)

/*
 * This function takes a pure function f as
 * an argument and returns a function that behaves almost exactly
 * like f, except that it only calls the original function once for every
 * argument, stores the result internally, and subsequently returns
 * this stored result every time it’s called with the same argument.
*/
func (m *memo) Memoize(f F) F {
	return func(x interface{}) interface{} {
		val, exists := m.cache[x]
		if exists {
			return val
		}
		m.cache[x] = f(x)
		return m.cache[x]
	}
}

func New() Memoizer {
	return &memo{
		cache: make(map[interface{}]interface{}),
	}
}

func testMemo() {
	input1 := "world!"
	input2 := "is it me you're looking for?"

	myFunc := func(x interface{}) interface{} {
		// simulates a long running function
		time.Sleep(5 * time.Second)
		return fmt.Sprintf("Hello %+v", x)
	}

	fmt.Println("Sanity Check:")
	fmt.Println(myFunc(input1))
	
	memoizer := New()
	memFunc := memoizer.Memoize(myFunc)

	fmt.Printf("One Slow:\n%s\n\n", memFunc(input1))

	fmt.Printf("Three Fast:\n%s\n%s\n%s\n\n", memFunc(input1), memFunc(input1), memFunc(input1))

	fmt.Printf("Another Slow:\n%s\n\n", memFunc(input2))

	fmt.Printf("Three More Fast:\n%s\n%s\n%s\n\n", memFunc(input2), memFunc(input2), memFunc(input2))
}

// insecure rand is necessary here
func testRandomMemo() {
	myFunc := func(x interface{}) interface{} {
		i, ok := x.(int64)
		if !ok {
			fmt.Printf("Input of type %T was not an int64.", x)
		}

		rand.Seed(i)

		// simulates a long running function
		time.Sleep(2 * time.Second)
		return rand.Intn(10)
	}

	memoizer := New()
	memFunc := memoizer.Memoize(myFunc)

	for i := 0; i < 100; i++ {
		val := memFunc(int64(i % 10))
		fmt.Println("Pseudo-Random Number #", (i+1), ":", val)
	}
}

func main() {
	fmt.Println("Testing Memoize Function:")
	testMemo()

	fmt.Println("Testing Memoize Random Number Generator:")
	testRandomMemo()
}
