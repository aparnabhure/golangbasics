package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=shdjksdsj"

func main() {
	fmt.Println("Welcome to URL")
	fmt.Println(myurl)

	result, err := url.Parse(myurl)

	checkNilError(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("Type of parmas is %T\n", queryParams)

	fmt.Println(queryParams["coursename"])
	fmt.Println(queryParams["paymentid"])

	for _, val := range queryParams {
		fmt.Println(val)
	}

	//Don't forgot to use & to have the reference not the copy
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=aparna",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println("Another URL: ", anotherUrl)

}

func checkNilError(err error) {
	if err != nil {
		panic(err) //This will shutdown the system and show error
	}
}
