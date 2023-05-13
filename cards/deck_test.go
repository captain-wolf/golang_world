package main

import "testing"

func TestNewDeck(t *testing.T) {
	// 1. Create a deck to test
	// 2. Count the cards to ensure 52 cards total
	// If test fails, present error
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card Ace of Spades but got %v", d[0])
	}
}
