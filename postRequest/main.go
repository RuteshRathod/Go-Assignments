package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("POST REQUEST")
	PerformPostRequest()
}
func PerformPostRequest() {
	const myURL = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
	{
		"Name":"Rutesh",
		"Gender": "Male",
		"Age": 24 
	}
	`)
	response, err := http.Post(myURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	contentResponse, _ := ioutil.ReadAll(response.Body)
	// echo's back whatever we send in request body
	fmt.Println(string(contentResponse))
}
