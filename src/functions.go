package main

import (
	"fmt"
	"os"
)

func main() {

	type foo int

	count := 1
	closure := func(msg string) {
		noOfArgs, ErroMsg := printf("%d %s\n", count, msg)
		fmt.Println(noOfArgs)
		fmt.Println(ErroMsg)
		count++
	}

	closure("A Message")

	closure("Another Message")
	m := thisisFunction()
	fmt.Println(m)

}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	os.NewFile()
}

func thisisFunction() int {

	_, k := "shiva", 123
	return k
}

func printf(str string, args ...interface{}) (int, error) {

	// println returns number of bytes written and error  2 params.
	// Here we ignore
	_, err := fmt.Printf(str, args...)
	return len(args), err
}
