package main

import "fmt"

func main() {

	/**
	works on slice
	*/

	var s []int = nil
	s = append(s, 2, 2)
	fmt.Println(s)

	/**
	doesnt work on map
	*/
	//var m map[string]int = nil// doesnt work
	var m = make(map[string]int)
	m["one"] = 1
	fmt.Println(m)

}
