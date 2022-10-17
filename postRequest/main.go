package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("POST REQUEST")
	PerformPostRequest()
	PerformPostFormRequest()
}
func PerformPostRequest() {
	const myURL = "http://localhost:8000/post"

	//fake json payload which is being treated as string
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

// using a post form
func PerformPostFormRequest() {
	const myURL = "http://localhost:8000/post"

	//formData
	data := url.Values{}
	data.Add("firstName", "Rutesh")
	data.Add("email", "ruteshr@google.com")

	responseF, err := http.PostForm(myURL, data)
	if err != nil {
		panic(err)
	}
	defer responseF.Body.Close()
	content, _ := ioutil.ReadAll(responseF.Body)
	fmt.Println(string(content))

}
