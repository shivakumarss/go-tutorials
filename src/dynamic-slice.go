package main

import "fmt"

func main() {
	x := 5
	y := 5

	table := make([][]int, x)

	for i := range table {
		table[i] = make([]int, y)

		for i, _ := range table[i] {
			table[i] = []int{1, 2, 3, 4, 5}

		}
	}

	fmt.Println(table)
}
