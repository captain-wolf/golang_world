package main

func main() {
	cards := newDeckFromFile("my_cards")
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	//we are assigning 2 new variables based off the return from the deal() function. Both are of type deck
	// hand, remainingCards := deal(cards, 5)

	//print output
	cards.print()
	// hand.print()
	// remainingCards.print()
}
