package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var str string = "raj"
	fmt.Println(utf8.ValidString(str))

	data2 := "A\xfeC"
	fmt.Println(utf8.ValidString(data2))

	data := "♥"
	fmt.Println(len(data)) //prints: 3

}
