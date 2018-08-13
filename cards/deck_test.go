package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("expected deck length of 52, but got %d", len(d))
	}
	first := Card{suit: "Spade", value: "Ace"}
	if d[0] != first {
		t.Errorf("expected first card to be %v bug got %v", first, d[0])
	}

	last := Card{suit: "Diamond", value: "King"}
	if d[len(d)-1] != last {
		t.Errorf("expected first card to be %v bug got %v", last, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("expected deck length of 52, but got %d", len(d))
	}
	os.Remove("_decktesting")
}
