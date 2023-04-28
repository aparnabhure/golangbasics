package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcom to web requests")

	res, err := http.Get(url)

	checkNilError(err)

	fmt.Printf("Type of the response is %T\n", res)

	fmt.Println("Status is ", res.Status)

	//You have to close the response as it is not gets closed automatically it is caller's responsibility to close, it is very important
	defer res.Body.Close()

	contentBytes, err := ioutil.ReadAll(res.Body)

	checkNilError(err)

	fmt.Println("Content is \n", string(contentBytes))

}

func checkNilError(err error) {
	if err != nil {
		panic(err) //This will shutdown the system and show error
	}
}
