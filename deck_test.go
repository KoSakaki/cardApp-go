package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expect := 4 * 4
	if len(d) != expect {
		t.Errorf("Expected deck length of %v, but got %v", expect, len(d))
	}

	firstExpected := "Ace of Spades"
	if d[0] != firstExpected {
		t.Errorf("Expected first card of %s, but got %v", firstExpected, d[0])
	}

	lastExpected := "Four of Clubs"
	if d[len(d)-1] != lastExpected {
		t.Errorf("Expected last card of %s, but got %v", lastExpected, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	expected := 160
	if len(loadedDeck) != expected {
		t.Errorf("Expected %v cards in deck, got %v",expected, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
