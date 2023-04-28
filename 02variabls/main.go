package main

import "fmt"

// Capital First letter makes it PUBLIC, accesible to anywhere

const LoginToken string = "some token"

func main() {

	var username string = "aparna"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type %T \n", smallval)

	var smallfloat float64 = 255.1234232987437368368
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type %T \n", smallfloat)

	//Default values
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type %T \n", anotherVariable)

	//Implicite Data type
	var website = "learncodeline.in"
	fmt.Println(website)
	fmt.Printf("Variable is of type %T \n", website)

	//No var style := volorous operators are allowed inside the method, not outside
	numberOfUsers := 2485543.23232
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type %T \n", numberOfUsers)

	//Using constants
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T \n", LoginToken)

}
