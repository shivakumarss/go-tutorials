package main

import (
	"fmt"
	"time"
)

func main() {

	var i int = 2

	for j := 0; j < i; j++ {
		fmt.Println("this is ", j)
		doIt(j)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("all done!")

}

func doIt(workerId int) {
	fmt.Printf("[%v] is running\n", workerId)
	time.Sleep(3 * time.Second)
	fmt.Printf("[%v] is done\n", workerId)
}
