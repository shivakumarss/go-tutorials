package main

import "fmt"

type Xii []int

func (p Xii) Len() int               { return len(p) }
func (p Xii) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xii) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

func SortInts(x Xii) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}

func main() {
	ints := Xii{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	SortInts(ints)
	fmt.Println("Pointer of ints ", &ints)
	fmt.Println(ints)

}
