package main

import "fmt"

func main() {

	//var i = nil // explicit type is required for declaration
	var i interface{} = nil // this works
	fmt.Println(i)

}
