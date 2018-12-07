package main

import "fmt"

func main() {
	x := [3]int{1, 2, 3}
	print(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x) //prints [1 2 3] (not ok if you need [7 2 3])

	/**
	if you want then use array pointers
	*/
	y := [3]int{1, 2, 3}
	print_pointers(&y)
	fmt.Println(y) //prints [1 2 3] (not ok if you need [7 2 3])

}

func print(arr [3]int) {
	arr[0] = 7
	fmt.Println(arr) //prints [7 2 3]
}

func print_pointers(arr *[3]int) {
	(*arr)[0] = 7
	fmt.Println(arr) //prints [7 2 3]
}
