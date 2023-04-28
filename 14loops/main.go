package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("******")

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("******")
	for index, day := range days {
		fmt.Printf("index is %v, day is %v \n", index, day)
	}

	fmt.Println("******")
	for _, day := range days {
		fmt.Printf("day is %v \n", day)
	}

	fmt.Println("******")
	roughValue := 1
	for roughValue < 10 {
		if roughValue == 8 {
			break
		}

		if roughValue == 2 {
			goto lbl
		}

		if roughValue == 3 {
			roughValue++
			continue
		}
		fmt.Println("Value is ", roughValue)
		roughValue++
	}

lbl:
	fmt.Println("You reached to GOTO Labesl")

}
