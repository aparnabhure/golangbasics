package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods")
	aparna := User{"Aparna", "email", true, 40, 20}
	fmt.Println("Name ", aparna.Name)
	fmt.Println("OneAge ", aparna.oneAge)
	aparna.GetStatus()
	aparna.NewMail()
	fmt.Println("Email ", aparna.Email)
}

//Functions within Struct/class called as Methods

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int //Now it is private
}

// Now it is called as method
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

// Passing a copy hence not original object is not getting updated
func (u User) NewMail() {
	u.Email = "test@email.com"
	fmt.Println("Email of the user is ", u.Email)
}
