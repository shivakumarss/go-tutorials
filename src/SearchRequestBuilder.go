package main

import (
	"example"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {

	/***
	building the Model
	*/
	query := example.SearchRequestMessage{}
	query.Query = "sample query"
	query.PageNumber = 100

	fmt.Println(query.String())

	/**
	building another model
	*/
	query2 := &example.SearchRequestMessage{
		Query:      "another query",
		PageNumber: 120,
	}

	fmt.Println(query2)

	data, _ := proto.Marshal(query2)
	fmt.Println("Data after Marshal ", data)

	query3 := &example.SearchRequestMessage{}

	e := proto.Unmarshal(data, query3)
	if e != nil {
		fmt.Println("Error in unmarshalling bytes ")
		return
	}

	fmt.Println("After unmarshalling this is data ")
	fmt.Println(query3)

}
