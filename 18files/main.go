package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files")

	content := "Hello Aparna this is conent for the file you are creating"

	//Create in current directory
	file, err := os.Create("./myTestFile.txt")

	checkNilError(err)

	contentLength, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is ", contentLength)

	//Close the file, add defer is the recommended way
	defer file.Close()

	readFile("./myTestFile.txt")

}

func readFile(fileName string) {
	contentBytes, err := ioutil.ReadFile(fileName)

	checkNilError(err)

	fmt.Println("Content from file is \n", string(contentBytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err) //This will shutdown the system and show error
	}
}
