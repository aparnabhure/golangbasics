package main

import "fmt"

func main() {
	fmt.Println("Wlecome to Maps")

	//Map<String, String>
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println(languages)

	//Get from map
	fmt.Println("JS ", languages["JS"])

	//Delete from map
	delete(languages, "RB")
	fmt.Println(languages)

	//loops
	for key, value := range languages {
		fmt.Printf("Key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("value is %v\n", value)
	}

}
