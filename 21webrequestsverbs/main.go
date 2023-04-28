package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const myurl = "http://localhost:8080/v1/home"

func main() {
	fmt.Println("Welcome to web verb")
	GetRequest(myurl)
	PostJson(myurl)
	PostForm(myurl)
}

func GetRequest(uri string) {
	res, err := http.Get(uri)
	checkError(err)

	defer res.Body.Close()

	fmt.Println("Status code ", res.StatusCode)
	fmt.Println("Content length ", res.ContentLength)

	content, err := ioutil.ReadAll(res.Body)
	checkError(err)

	fmt.Println("Content ", string(content))

	var responseString strings.Builder
	byteCounts, err := responseString.Write(content)

	fmt.Println("ByteCount is ", byteCounts)
	fmt.Println("Response is ", responseString.String())
}

func PostJson(uri string) {
	//Json payload
	requestBody := strings.NewReader(`
	{
		"name":"aparna",
		"age":40,
		"email":"myemail@email.com"
	}
	`)

	res, err := http.Post(uri, "application/json", requestBody)
	checkError(err)

	//Close the connection
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	checkError(err)

	fmt.Println(string(content))

}

func PostForm(uri string) {
	//Form data

	data := url.Values{}
	data.Add("Country", "India")
	data.Add("Name", "Aparna")
	data.Add("Email", "test@email.com")

	//Respons is a pointer
	res, err := http.PostForm(uri, data)
	checkError(err)

	//Close request
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	checkError(err)

	fmt.Println(string(content))

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
