package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck'

type deck []string

// new deck function which is of the 'deck' type
func newDeck() deck {
	cards := deck{}

	//Slices of card suits and values
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// for index suit within range of cardSuits
	// instead of using for loop with a variable like "for i, blah", we use _ because variables are not needed for this particular function
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			//in this instance +" of "+ is tying suit and value into 1 string
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// the goal of the deal function is to pass in the deck variable as d and the index as handSize
// addtionally we want to tell go to expect 2 returns of type "deck" notated in (deck,deck)
// upToIncluding:notIncluding <-- range inside of slice syntax
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// we will need to turn decks into strings for type conversion. function has a receiver of deck and returns a string
func (d deck) toString() string {
	//convert deck back to a slice of strings
	return strings.Join([]string(d), ",")
}

// save to file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	//we need the byte slice and error out
	bs, err := os.ReadFile(filename)
	if err != nil {
		//Option #1 - log the error and return to newDeck()
		//Option #2 - log error and quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//we will do the reverse of save to file essentially. We will use the Split func
	s := strings.Split(string(bs), ",")

	return deck(s)

}

func (d deck) shuffle() {
	for i := range d {
		// create new position of card using rand package
		// len is short for length function which returns the length according to type
		newPosition := rand.Intn(len(d) - 1)

		//swap the elements --
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
