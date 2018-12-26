package main

import (
	"fmt"
	"time"
)

var c chan int

func main() {

	c = make(chan int)

	go isReady("Coffee", 2)
	go isReady("Tea", 1)
	//time.Sleep(time.Duration(3) * time.Second)
	<-c
	<-c

}

func isReady(str string, sec int) {

	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(str, " is ready ")
	c <- 1

}
