package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")
	//no ingeritence in golang, no super or parent concept

	aparna := User{"aparna", "aparna@test.com", true, 40}
	fmt.Println(aparna)

	//For struct use+v to get in detailed information ie is giving Attribute names
	fmt.Printf("Details are %+v \n", aparna)

	fmt.Printf("Name is %v, Email is %v, Age is %v\n", aparna.Name, aparna.Email, aparna.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
