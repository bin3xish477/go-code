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
	cards = newDeckFromFile("my_cards.txt")
	cards.shuffle()
	cards.print()
	fmt.Println("Completed!")
}
