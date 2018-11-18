package main

import (
	"fmt"
)

func main() {

	loops := 1

	for loops > 0 {
		fmt.Println("Enter number of loops required ? ")
		fmt.Scanf("%d", &loops)

		fmt.Println("No of loops entered as %d ", loops)

		for k := 0; k < loops; k++ {

			fmt.Println("Loop counter %d", k)
		}
	}
	fmt.Println("While loop ended ")

	/***
	This is infinate loop, getting breaked here.
	*/
	for {

		break
	}

	fmt.Println("All done thank you.")

}
