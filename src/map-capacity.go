package main

func main() {
	m := make(map[string]int, 99)
	//cap(m) //error cannot use cap on map.
	println(m)

}
