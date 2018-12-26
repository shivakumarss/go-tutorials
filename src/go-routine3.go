package main

import (
	"fmt"
	"runtime"
	"time"
)

var c1 chan int

func main() {

	c1 = make(chan int)
	previousValue := runtime.GOMAXPROCS(5)
	fmt.Println("rPrevious value ", previousValue)
	go isReady2("Coffee", 2)
	go isReady2("Tea", 1)

	var i int

L:
	for {
		select {

		case <-c1:
			i++
			if i > 1 {
				break L
			}
		}

	}

	//time.Sleep(time.Duration(3) * time.Second)
	//<-c1
	//<-c1

}

func isReady2(str string, sec int) {

	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(str, " is ready ")
	c1 <- 1

}
