package main

import (
	"fmt"
	"sync"
	"time"
)

// just basic unbuffered channel practice, easy stuff
func main() {
	channel := make(chan int)
	wg := &sync.WaitGroup{}
	odd := [5]int{1, 2, 3, 4, 5}
	wg.Add(5)
	for _, value := range odd {
		go rand(value, wg, channel)
		time.Sleep(time.Millisecond * 100)
		fmt.Println(<-channel)
	}

	wg.Wait()
}
func rand(n int, wg *sync.WaitGroup, channel chan int) {
	defer wg.Done()
	channel <- n * 2
}
