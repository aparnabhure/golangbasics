package main

import "fmt"

func main() {
	fmt.Println("Wlecome to IF ELSE")

	loginCount := 23

	var result string

	// { brach should be at the same line
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount < 30 {
		result = "< 30"
	} else {
		result = "Something else"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is < 10")
	} else {
		fmt.Println("Num is >=10")
	}

}
