package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	params := mux.Vars(r)
	fmt.Println(params)

	for _, course := range coursesDB {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(params["id"])
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Single Course")
	w.Header().Set("Content-Type", "application/json")

	// validation/ helper : what if body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Empty Data is not Accepted")
	}

	// what if body has empty paranthesis"{}", as it would be consider something
	var course Course // refernece of struct as object
	_ = json.NewDecoder(r.Body).Decode(&course)
	fmt.Println("Inserted, :", course.CourseName)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Empty Data{} is not Accepted")
		return
	}
	// Check only if title is duplicate
	// loop, titte matches with course. coursename, JSON
	for _, courseX := range coursesDB {
		fmt.Println(courseX)
		if courseX.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Duplicate CourseName is not Accepted")
			return
		}
	}

	// generate random unique id, which would be string according to struct
	// then append that input of cousre into slice of main Course

	rand.Seed(time.Now().Unix())
	//converted int into string with random id
	course.CourseId = strconv.Itoa(rand.Intn(1000))
	coursesDB = append(coursesDB, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Single Course")
	w.Header().Set("Content-Type", "application/json")

	// grab the id from request
	params := mux.Vars(r)

	// loop, id, remove, add/update with Id
	// 1. Loop
	for index, course := range coursesDB {
		//2. Match Id
		if course.CourseId == params["id"] {
			// removed the course basis on the id given in request, with varidic function of append()
			coursesDB = append(coursesDB[:index], coursesDB[index+1:]...)
			// add
			var course Course // refernece of Course struct as object
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			coursesDB = append(coursesDB, course)
			json.NewEncoder(w).Encode(course)
			return
		} else {
			json.NewEncoder(w).Encode(params["id"])
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete single course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// loop, id, remove

	for index, course := range coursesDB {
		if course.CourseId == params["id"] {
			coursesDB = append(coursesDB[:index], coursesDB[index+1:]...)
			json.NewEncoder(w).Encode(params["id"])
			return
			// As deletion happened we have to break the loop as conditon has been satiesfied
			break

		}
	}

}

func main() {
	fmt.Println("API-Courses")
	r := mux.NewRouter()

	// seeding data for courses
	coursesDB = append(coursesDB, Course{CourseId: "2", CourseName: "Java", CoursePrice: 5000, Author: &Author{FullName: "Rutesh", Website: "lco.dev"}})
	coursesDB = append(coursesDB, Course{CourseId: "23", CourseName: "Go", CoursePrice: 1000, Author: &Author{FullName: "Rutesh", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}
