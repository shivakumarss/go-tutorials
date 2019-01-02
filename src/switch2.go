package main

import "fmt"

func main() {

	isSpace := func(ch byte) bool {

		switch ch {

		case 'i':
			fmt.Println("This is i")

			return true

		}

		return false

	}

	fmt.Println(isSpace('i'))
	fmt.Println(isSpace('k'))
	fmt.Println(isSpace(' '))

}
