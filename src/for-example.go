package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 40}

	for j, i := range arr {
		fmt.Println(i, j)
	}
}
