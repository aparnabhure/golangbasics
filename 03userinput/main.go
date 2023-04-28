package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	//pkg.go.dev
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter some number")

	//Comma Ok | Error Ok syntax | Try/Catch
	input, _ := reader.ReadString('\n')
	fmt.Println("You have entered ", input)
	fmt.Printf("Type of the input is %T ", input)

}
