package main

import "fmt"

type S struct{ i int }

type J struct{ i int }

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

type I interface {
	Get() int
	Put(int)
}

func fs(p I) {
	fmt.Println(p.Get())
	p.Put(1)
}

type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }

func main() {
	var s S
	var j S

	fmt.Println(j)
	fs(&s)
	fs(&s)
	fsss(&s)
	//t := s.(I)
	//someVariable.(SomeType)

}

func fsss(p I) {
	switch t := p.(type) {
	case *S:
		println("Hello", t)
	case *R:
	default:

	}

}
