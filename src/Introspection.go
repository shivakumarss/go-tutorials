package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string "namestr"
	age  int
}

func main() {

	s := Person{"raj", 12}
	//fmt.Println(ShowTag(&s))

	show(&s)
}

func ShowTag(i interface{}) reflect.StructTag {
	//← Called with *Person
	switch t := reflect.TypeOf(i); t.Kind() {
	case reflect.Ptr:
		//← A pointer, thus reflect.Ptr
		tag := t.Elem().Field(0).Tag
		return tag
	}

	return ""

}

func show(i interface{}) {

	switch t := i.(type) {

	case *Person:

		t2 := reflect.TypeOf(i)
		v := reflect.ValueOf(i)

		//← Type meta data
		//← Actual values
		tag := t2.Elem().Field(0).Tag
		name := v.Elem().Field(0).String()
		fmt.Println(tag)
		fmt.Println(name)
		fmt.Println(t)

	}
}
