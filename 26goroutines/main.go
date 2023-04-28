package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

var mutex sync.Mutex
var waitgrp sync.WaitGroup

func main() {
	fmt.Println("Welcome to goroutine")
	// go fireup a thread but it is waiting to come back
	//go greeter("Hello")
	//greeter("World")

	websites := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websites {
		go getStatusCode(web)
		waitgrp.Add(1)
	}

	//Always be the end of Main method, because all other thread or routines might not have been completed
	waitgrp.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer waitgrp.Done()

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Opps an error")
	} else {
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()
		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
	}

}
