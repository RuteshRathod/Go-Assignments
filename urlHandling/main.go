package main

import (
	"fmt"
	"net/url"
)

const mUrl string = "https://lco.dev:3000/learn?coursename=react.js&paymentid=jhjkjska"

func main() {
	fmt.Println("Handling mUrl")
	fmt.Println(mUrl)

	//Parsing mUrl

	result, _ := url.Parse(mUrl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Println("qPAram", qParams["coursename"])

	for index, val := range qParams {
		fmt.Println("Param is", index, " ", val)
	}

	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherURL := partsOfURL.String()
	fmt.Printf("Type of Another URL:- %T \n", anotherURL)
	fmt.Println("another URL", anotherURL)

}
