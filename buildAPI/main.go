package main

import (
	"encoding/json"
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

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coursesDB)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Single Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from incoming request
	params := mux.Vars()
	fmt.Println(params)

	for _, course := range coursesDB {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found for id: ", params["id"])
	return
}
