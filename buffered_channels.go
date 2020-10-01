package main

import (
	"fmt"
)

func main() {
	// Channel Buffering Example
	msgs = make(chan string, 2)
	// Sending strings into the channel
	msgs <- "Hello,"
	msgs <- "World!"
	// Retrieving and printing strings from the channel
	fmt.Println(<-msgs)
	fmt.Println(<-msgs)	
}
