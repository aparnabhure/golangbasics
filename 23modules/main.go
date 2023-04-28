package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Couple of commands
//goland.org/ref/mod
// https://github.com/gorilla/mux
// go get -u github.com/gorilla/mux To add Gorilla MUX library it is HTTP router lib,
// default it  will add an entry in go.mod with //indirect comment
// the indirect will go away when you use the lib in the code to do that you need to run below command
// go mod tidy (cleanup dependencies if not in use)
// you need to copy below code in order to get gorilla mux inported
/*
r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	http.Handle("/", r)
*/
// To verify all modules are valid run below command
// go mod verify
// To view all dependencies
// go list
// go list all
// go list -m all
// To view all list of gorilla mux versions
// go list -m -versions github.com/gorilla/mux

// To know why module dependencies are required
// go mod why

// Graph list
// go mod graph

//To edit go.mod file, this might help while changes are required via Pipeline or script
// go mod edit -go 1.16 (to change version)

//go mod vendor

func main() {
	fmt.Println("Welcome to MOD")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")
	http.Handle("/", r)

	err := http.ListenAndServe(":4000", r)
	log.Fatal(err)
}

func greeter() {
	fmt.Println("Hello there")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcom To Golang</h1>"))
}
