package main

import "fmt"

func main() {

	k := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(k)
	i := 0
	i++

	fmt.Println(k[i])

	fmt.Println(^1) //error

	var d uint8 = 7
	fmt.Printf("%08b\n", d)

	var d1 uint8 = 7
	fmt.Printf("%08b\n", ^d1)

}
