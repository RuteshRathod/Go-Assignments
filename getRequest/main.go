package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("GET REQUEST")
	PerformGetRequest()
}
func PerformGetRequest() {
	const myURL = "http://localhost:8000/get"
	response, err := http.Get(myURL)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status Codde: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	// Used Strings package for decoding and many more methods
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	// using string library we decoded byte to string and it has many more methods
	fmt.Println(string(content))

	byteContent, _ := responseString.Write(content)
	fmt.Println("Byte Content", byteContent)
	fmt.Println(responseString.String())

	// strings vs string package use differs according to use case and also cons, pros of it

}
