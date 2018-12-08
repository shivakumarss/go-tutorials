package main

import "fmt"

func main() {

	const (
		Hello  = iota
		Hello2 = iota
	)

	for i := 0; i < 3; i++ {
		fmt.Println(Hello)
		fmt.Println(Hello2)
	}
}
