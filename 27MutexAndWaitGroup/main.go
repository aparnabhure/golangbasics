package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race conditions handling")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	//Go ahead and read about RWMutex

	var score = []int{0}

	//Number of goroutines to wait for
	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Go 1")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Go 2")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Go 3")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	//Notify the main method for all wait groups to finish the jobs
	wg.Wait()
	fmt.Println(score)
	// to analyse what is happening with goroutine use command
	// go run --race .
	// to fix the race conditions use mutex
}
