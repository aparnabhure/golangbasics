package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	fmt.Println("Welcome to Maths")

	var numOne int = 2
	var numTwo float64 = 4
	fmt.Println("Sum is ", numOne+int(numTwo)) //loosing precision

	//Random numbers
	//Math.Rand can generate duplicate where as Crypto.Rand gives more precise numbers
	// rand.Seed(time.Now().UnixNano()) //This is deprecated
	// fmt.Println(rand.Intn(5) + 1)

	//Random from Chrpto
	myRand, err := rand.Int(rand.Reader, big.NewInt(5))
	if err != nil {
		panic(err)
	}

	fmt.Println("***********")
	fmt.Println(myRand)

}
