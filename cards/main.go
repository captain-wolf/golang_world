package main

func main() {
	// cards := newDeckFromFile("my_cards")
	// cards.saveToFile("my_cards")
	// we're assigning 2 new variables based off the return from the deal() function. Both are of type deck
	// hand, remainingCards := deal(cards, 5)
	// print output
	// hand.print()
	// remainingCards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
