package main

import "fmt"

func main() {

	type Person struct {
		name string
		age  int
	}

	var p Person
	p.age = 12
	p.name = "shiva"
	fmt.Println(p)

	p2 := new(Person)
	p2.age = 15
	p2.name = "ssshiva"
	fmt.Println(p2)

}
