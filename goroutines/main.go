package main

import (
	"fmt"
	"runtime"
	"time"
)

func printNums(num int, done chan bool) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
	done <- true
}

func main() {
	fmt.Println("Channels and goroutines")
	runtime.GOMAXPROCS(4)
	done := make(chan bool, 3)
	start := time.Now()
	go printNums(3, done)
	go printNums(3, done)
	elapsedTime := time.Since(start)
	fmt.Printf("Total time for execution is %v\n", elapsedTime.String())
	// time.Sleep(time.Second)
	fmt.Println(<-done)
	fmt.Println(<-done)
}
