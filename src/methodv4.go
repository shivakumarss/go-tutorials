package main

import (
	"fmt"
)

type Foo int

func (self Foo) Emit() {
	fmt.Printf("%v", self)
}

type Emitter interface {
	Emit()
}

func main() {

}
