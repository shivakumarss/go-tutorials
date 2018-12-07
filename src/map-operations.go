package main

import "fmt"

func main() {
	x := map[string]string{"one": "a", "two": "", "three": "c"}

	v := x["two"]

	if v == "" { //incorrect
		fmt.Println("no entry")
	}

	value, isExists := x["one"]
	fmt.Println(value)
	fmt.Println(isExists)

	//if _, ok := x["two"]; !ok {
	//	fmt.Println("no entry")
	//}

}
