package main

import "fmt"

func main() {

	var i int
	var Θ float32
	var explicitly, typed, pointers *complex128
	int_pointer := &i
	another_int_pointer := new(int)
	generic_channel_pointer := make(chan interface{})

	fmt.Println(i)
	fmt.Println(Θ)
	fmt.Println(explicitly)
	fmt.Println(typed)
	fmt.Println(pointers)
	fmt.Println(int_pointer)
	fmt.Println(another_int_pointer)
	fmt.Println(generic_channel_pointer)


}
