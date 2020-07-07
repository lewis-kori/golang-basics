package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got: %v ", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of spades but got: %v", d[0])
	}

	if d[len(d)-1] != "Four of Clover" {
		t.Errorf("Expected Four of Clubs but got: %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// do cleanup to remove the file
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got: %v", len(loadedDeck))
	}
}
