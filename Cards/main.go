package main

import (
	"fmt"
	"time"
)

func main() {
	cards := newDeck()
	saveToFile(cards, "my_cards.txt")
	// Using the Sleep function in the time package
	// for halting the programs execution for three seconds.
	time.Sleep(time.Second)
	data := newDeckFromFile("my_cards.txt")
	data.print()
	fmt.Println("Completed!")
}