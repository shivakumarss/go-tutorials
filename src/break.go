package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
	L:
		for {
			for {
				break L
				/**
				This never gets here.
				*/
				fmt.Println("hello")
			}
		}
		fmt.Printf("%d\n", i)
	}

}
