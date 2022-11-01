package main

import (
	"fmt"
	"net/http"
)

// file means it should be in different files
// Model For course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var coursesDB []Course

// helpers are used for validations, aka middleware but not exactly they are middleware
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

// Controllers  - file
// serve home route-homepage type

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Serve Home")
	w.Write([]byte("<h1> Welcome to API Route</h1>"))
}
