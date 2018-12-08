package main

import "fmt"

func main() {

	var str = "shiva"
	//str[0] = 'T'
	xstr := []byte(str)
	xstr[0] = 'T'

	var str2 = string(xstr)

	fmt.Println(xstr)
	fmt.Println(str2)

}
