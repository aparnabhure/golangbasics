package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Wlecome to Slices")

	//Slices, dynamic array if using this syntex then {} is required iw initialization is required
	var fruitList = []string{}

	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	var vegList = []string{"potato", "beans", "gilki"}

	fmt.Printf("Type of vegList is %T\n", vegList)

	vegList = append(vegList, "cabbage", "onion")
	fmt.Println("veg list is ", vegList)

	//1: is to slice the array in 2 list so the list got sliced from 1st
	// vegList = append(vegList[1:])
	// fmt.Println("veg list is ", vegList)

	// Slice from 1 to 3, last is exclusive
	// vegList = append(vegList[1:3])
	// fmt.Println("veg list is ", vegList)

	vegList = append(vegList[:3])
	fmt.Println("veg list is ", vegList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 635
	highScores[2] = 136
	highScores[3] = 437
	//highScores[4] = 444 //Crash as the size has given as 4

	fmt.Println(highScores)
	//Add values at run time 4 is the initial size
	highScores = append(highScores, 555, 777, 333)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	//Remove value from slice based on index
	var courses = []string{"java", "react", "js", "sql", "python"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
