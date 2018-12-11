package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	One int
	Two string
}

func main() {

	var myData = MyData{1, "Two"}
	fmt.Println(myData)

	b, _ := json.Marshal(myData)
	fmt.Println(string(b)) // 1st variable now marshalled because identifier declared as small case.

	var out MyData
	json.Unmarshal(b, &out)

	fmt.Println(out)

	fmt.Printf("%#v\n", out) //prints main.MyData{One:1, two:""}

}
