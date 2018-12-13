package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close() //ok, most of the time :-)

	_, err = io.Copy(ioutil.Discard, resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
