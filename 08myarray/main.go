package main

import "fmt"

func main() {
	fmt.Println("Wlecome to Array")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Mango"

	fmt.Println("Fruits are ", fruitList)
	fmt.Println("Fruits are ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "cabbage"}
	fmt.Println("Vegs are ", vegList)
	fmt.Println("Vegs are ", len(vegList))

}
