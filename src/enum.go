package main

import "fmt"

func main() {

	const (
		Red             = 1 << iota
		Green           = 1 << iota
		Blue, ColorMask = 1 << iota, (1 << (iota + 1)) - 1
	)

	fmt.Println("I am printing constant values now ")
	fmt.Println(Red)
	fmt.Println(Green)
	fmt.Println(Blue)
	fmt.Println(ColorMask)

}
