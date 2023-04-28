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

type Course struct {
	Id     string  `json:"courseid"`
	Name   string  `json:"coursename"`
	Price  int     `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Name    string `json:"authorname"`
	Website string `json:"website"`
}

// Fake DB
var courses []Course

// Helpers
func (c *Course) IsEmpty() bool {
	//return c.Id == "" && c.Name == ""
	return c.Name == ""
}

func (c *Course) IsIdEmpty() bool {
	return c.Id == ""
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course")
	w.Header().Set("Content-Type", "application/json")

	//Grab ID from request
	params := mux.Vars(r)

	for _, course := range courses {
		if course.Id == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Found with given id ")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send valid data")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsIdEmpty() {
		json.NewEncoder(w).Encode("Please add valid ID")
		return
	}

	//Add course in the courses slice
	//Generate unique ID string
	rand.Seed(time.Now().UnixNano())
	course.Id = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send valid data")
		return
	}

	var newCourse Course
	_ = json.NewDecoder(r.Body).Decode(&newCourse)
	if newCourse.IsEmpty() {
		json.NewEncoder(w).Encode("Please add valid ID & name")
		return
	}

	//Grab ID from request
	params := mux.Vars(r)

	for index, course := range courses {
		if course.Id == params["id"] {
			//Remove old from courses
			courses = append(courses[:index], courses[index+1:]...)
			//Add as new
			newCourse.Id = params["id"]
			fmt.Println("NewCourse Name ", newCourse.Name)
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode(newCourse)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Found with given id ")
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send valid data")
		return
	}

	//Grab ID from request
	params := mux.Vars(r)

	for index, course := range courses {
		if course.Id == params["id"] {
			//Remove old from courses
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Success")
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Found with given id ")
	return
}

func main() {
	fmt.Println("Welcome to Rest APIs Course")
	r := mux.NewRouter()
	courses = append(courses, Course{"1", "JS", 299, &Author{Name: "aparna", Website: "test.com"}})
	courses = append(courses, Course{"2", "Java", 399, &Author{Name: "aparna", Website: "test.com"}})

	//Routing
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")

	//Listen to the port
	log.Fatal(http.ListenAndServe(":4000", r))
}
