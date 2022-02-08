package main

import (
	"fmt"
	"runtime"
	"time"
)

// Very basic implementation of monitoring resources with runtime package
func f(i interface{}) {
	fmt.Printf("%v\n", i)
}

func monitor() {
	mem := &runtime.MemStats{}
	for {
		cpu := runtime.NumCPU()
		fmt.Println("CPU:", cpu)

		rot := runtime.NumGoroutine()
		fmt.Println("Goroutine:", rot)

		// Byte
		runtime.ReadMemStats(mem)
		fmt.Println("Memory:", mem.Alloc)
		time.Sleep(time.Second * 2)
	}
}
func main() {
	go monitor()
	var a interface{} = "a string"
	var c int = 42

	var numCpu int
	numCpu = runtime.NumCPU()
	runtime.GOMAXPROCS(numCpu)
	fmt.Printf("Total Cpu: %d\n", numCpu)
	f(a) // string
	f(c) // int
	time.Sleep(time.Second * 3)
}
