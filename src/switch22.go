package main

import "fmt"

func Less(l, r interface{}) bool {
	switch l.(type) {
	case int:
		if _, ok := r.(int); ok {
			return l.(int) < r.(int)
		}
	case float32:
		if _, ok := r.(float32); ok {
			return l.(float32) < r.(float32)
		}
	}
	return false
}
func main() {
	var a, b, c int = 5, 15, 0
	var x, y, z float32 = 5.4, 29.3, 0.0
	if c = a; Less(a, b) {

		c = b
		// 15
	}
	if z = x; Less(x, y) {

		z = y
		// 29.3
	}
	//println(c, z) // this is sync call , there could be a chance that it would print later than fmt.
	fmt.Println(c)
	fmt.Println(z)
}
