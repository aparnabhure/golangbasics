package main

import (
	"fmt"
	"sync"
)

// https://www.youtube.com/watch?v=i5HEE5Ikv3w&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=56
// Channels help to connect multiple goroutines to talk to each other
// Channels cannot be utilize unless some goroutine is not listening to me
func main() {
	fmt.Println("Welcome to Channels")

	//Buffered channel myChannel := make(chan int, 5) //Capacity
	myChannel := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	//RECEIVE Only goroutine, we should not close the channel in Receiver
	go func(ch <-chan int, w *sync.WaitGroup) {
		_, isOpen := <-myChannel
		fmt.Println("Is Channel Open ", isOpen)
		fmt.Println(<-myChannel)
		fmt.Println(<-myChannel)
		wg.Done()

	}(myChannel, wg)

	//SEND ONLY goroutine
	go func(ch chan<- int, w *sync.WaitGroup) {
		myChannel <- 5
		myChannel <- 6
		close(myChannel)
		_, isOpen := <-myChannel
		fmt.Println("Is Channel Open ", isOpen)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
