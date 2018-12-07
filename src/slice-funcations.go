package main

import "fmt"

func main() {

	var arr []int = make([]int, 2)

	arr[0] = 1
	arr[1] = 2

	printArrayElements(arr)

	fmt.Println("final elements of array ", arr)

}

func printArrayElements(arrs []int) {
	fmt.Println("elements of array ", arrs)

	arrs[0] = 3
	arrs[1] = 4

	fmt.Println("after modification ", arrs)

}
