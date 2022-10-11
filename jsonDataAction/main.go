package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"CourseName"` // alias in ``
	Price    int
	Platform string   `json:"Website"`
	Password string   `json:"-"`              // "-" means it wont be visible whoever will consume this api
	Tags     []string `json:"tags,omitempty"` // this tag will not show if value is nil
}

// we will be encoding json in struct or slices
func main() {
	fmt.Println("JSON Handling")
	EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	courses := []course{

		{"Golanng", 2966, "go.dev.com", "abc123", []string{"web dev", "backend"}},
		{"Python", 1221, "go.dev.com", "asas123", []string{"web dev", "backend"}},
		{"Vue", 3312, "go.dev.com", "xcs123", nil},
	}

	// packing data into json
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)

}

func DecodeJson() {
	jsonDataFromSource := []byte(`
        {
                "CourseName": "GOlanng",
                "Price": 2966,
                "Website": "go.dev.com",
                "tags": [
                        "web dev",
                        "backend"
                ]
        }
	`)
	var courseGo course
	checkValid := json.Valid(jsonDataFromSource)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromSource, &courseGo) // here in second param of unmarshall we are using
		//reference for actual variale instad of copy of it which is mentioninng vairable without "&"
		fmt.Printf("%#v \n", courseGo) // used %# to print interface
	} else {
		fmt.Println("Something went wrong for unmarshalling")
	}

	// case: when you want to add data to key value pair instead of straight decoding
	var myData map[string]interface{}
	json.Unmarshal(jsonDataFromSource, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("key is %v: and value is %v: with its type is: %T \n", k, v, v)
	}

}
