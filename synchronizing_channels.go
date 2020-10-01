package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Start Worker")
	time.Sleep(time.Second)
	fmt.Println("Worker completed task")
	done <- true
}
func main() {
	// Channel Synchronization
	done := make(chan bool, 1)
	go worker(done)
	fmt.Println("Function completed:", <-done)
}
