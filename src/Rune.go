package main

import (
	"fmt"
	"os"
)

func main() {
	var array [10]int

	for i := 0; i < 10; i++ {
		array[i] = 100 + i
	}

	fmt.Println(array)

	slice := array[0:10]

	fmt.Println(slice)
	//slice[8] = 'a'
	//slice[9] = 'a'

	os.Exit(-1)

	sl := make([]int, 10)

	fmt.Println(len(sl))
	fmt.Println(cap(sl))
	sl = append(sl, 0, 1, 2, 3, 4, 5, 6)
	fmt.Println(sl)

	//fmt.Println("start of program")
	//myfunc()
	i := 1

	switch i {
	case 1:
		fallthrough
	case 2:
		fmt.Println("this is case 2")

	case 3:
		fmt.Println("this is case 3")
	}

	fmt.Println("end of the program")
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+':
		return true
	}
	return false
}

func unhex(c byte, k byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	case k <= c && c <= k:
		return c - 'A' + 10
	}
	return 0
}

func myfunc() {
	i := 0
Here:

	if i > 5 {
		return
	}

	println(i)
	i++
	goto Here

}
