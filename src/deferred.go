package main

import (
	"fmt"
	"testing"
)

func main() {

	testing.Coverage()

	for i := 0; i < 5; i++ {

		defer fmt.Println(i) // last in first out . LIFO
	}

	defer func(x int) {
		/* ... */
		fmt.Println(x)
	}(99)

	fmt.Println("Hello ")

	k := f()

	println(k)

	//a := func() {
	//	println("Hello")
	//}
	//a()

	//var xs = map[int]func() int{
	//	1: func() int { return 10 },
	//	2: func() int { return 20 },
	//	3: func() int { return 30 },
	//}
	//
	//k := xs[1]()
	//
	//fmt.Println(k)

}

func f() (ret int) {
	//← ret is initialized with zero
	defer func() {
		ret++
		//	← Increment ret with 1
	}()
	return 0
	//← 1 not 0 will be returned!
}
