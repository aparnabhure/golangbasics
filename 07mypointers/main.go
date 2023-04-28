package main

import "fmt"

func main() {
	fmt.Println("Wlecome to Pointers")
	//* represents Pointer
	var ptr *int
	fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	//Pointer referencing myNumber address
	var ptrNum = &myNumber
	fmt.Println("Value of MyNumber is ", ptrNum)
	fmt.Println("Value of MyNumber is ", *ptrNum) //Actual Value
	fmt.Println("Value of MyNumber is ", &ptrNum)

	*ptrNum = *ptrNum * 2
	fmt.Println("New number is ", myNumber)
}
