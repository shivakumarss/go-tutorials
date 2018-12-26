package main

import "fmt"

func main() {

	type NameAge struct {
		Name string
		Age  int
	}

	ins := new(NameAge)
	ins.Age = 10
	ins.Name = "raja"

	fmt.Printf("Name is %s and age is %d  \n", ins.Name, ins.Age)

	fmt.Printf("%s\n", ins)
}
