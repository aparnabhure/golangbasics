package main

import "fmt"

func main() {
	fmt.Println("Wlecome to Differ test")
	defer fmt.Println("Hello")
	fmt.Println("world")
	defer fmt.Println("aparna") //Printed in reverse order LIFO so aparna will get rpint first and then Hello
	fmt.Println("welcome")
	defer myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
