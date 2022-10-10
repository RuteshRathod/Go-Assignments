package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Request Demo")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response Type: %T", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println("Content:", content)

	// mandatory to close response, as by default it does not closes
}
