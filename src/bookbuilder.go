package main

import (
	"example"
	"fmt"
)

func main() {

	/***
	Create a book object
	*/
	bookObject := example.Book{
		Title:  " A journey of go lang",
		Author: "Shiva",
		Isbn:   12388774,
	}

	fmt.Println("This is book object ", bookObject)
	/***
	look up isbn using GetBookrequest
	*/
	request := example.GetBookRequest{
		Isbn: 1234,
	}

	isExists, responseObj := getBookbyBookRequest(request)
	fmt.Println("Exists ", isExists, " book object ", responseObj)

	/***
	Now trying to lookup which has the value.
	*/
	request2 := example.GetBookRequest{
		Isbn: 12345,
	}

	isExists2, responseObj2 := getBookbyBookRequest(request2)
	fmt.Println("Exists ", isExists2, " book object ", responseObj2)

	value, ok := example.EnumSample_name[1] // 1 is for Started.
	if !ok {
		panic("invalid enum value")
	}

	fmt.Println(value)

	/***
	vice versa lookup by name
	*/

	value2, ok := example.EnumSample_value[value] // 1 is for Started.
	if !ok {
		panic("invalid enum value")
	}

	fmt.Println(value2)

}

func getBookbyBookRequest(request example.GetBookRequest) (bool, example.Book) {

	mapper := make(map[int64]example.Book, 10)
	mapper[12345] = example.Book{
		Title:  " A journey of go lang",
		Author: "Shiva",
		Isbn:   12388774,
	}

	mapper[45678] = example.Book{
		Title:  " A journey of java",
		Author: "Akilesh yadav",
		Isbn:   123887876,
	}

	/***
	parse the request and return from map
	*/

	fmt.Println("Trying to look up ", request.GetIsbn())

	value, isExists := mapper[request.GetIsbn()]
	fmt.Println(value)
	fmt.Println(isExists)

	return isExists, value

}
