package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

// Declaring a custom type 'deck'
// This is basically Golang's way of creating
// a class. There is no OOP in Golang!
type deck []string

// newDeck will create a deck of cards which is of type deck
// and will return the deck created.
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

// This "receiver" function is essentially a property of
// any instance of type deck.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal will return two slices of type deck.
// d = type deck
// handSize = int
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// toString: A receiver function that will convert a
// type deck value into a string.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saveToFile: A  function that will save a deck
// type value into a file.
func saveToFile(d deck, fileName string) error {
	// ioutil.WriteFile excepts a filename, a byte slice of the data
	// that will be written to the file, and the permissions that the file
	// will have.
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// newDeckFromFile: A function that will read data from a 
// file.
func newDeckFromFile(fileName string) (deck) {
	var receivedDeck deck
	byteslice, err := ioutil.ReadFile(fileName)
	// Handling error
	// nil is the default placeholder in Go.
	// If the error returned is not nil, then
	// an error occured.
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {
		receivedDeck = strings.Split(string(byteslice), ",")
		// Type conversion to our custom type, deck.
		return deck(receivedDeck)
	}
	return receivedDeck
}
