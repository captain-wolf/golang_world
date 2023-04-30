package main

import "fmt"

var card string

func main() {
	cards := deck{ newCard()}

	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	cards.print()
}

func newCard() {
	return "Five of Diamonds"
}