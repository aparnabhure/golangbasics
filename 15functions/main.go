package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is ", result)

	proResult := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Result is ", proResult)

	proAdderResult, prodAdderMessage := proAdderResult(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Result is ", proAdderResult)
	fmt.Println("Message is ", prodAdderMessage)

}

// Returned value is at the end
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}

	return total
}

func proAdderResult(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "You got the result"
}

func greeter() {
	fmt.Println("Hello from golang")
}
