package main

import (
	"testing"
	"os"
)

/*****Testing NewDeck Function*****
 -Check that length of deck is 52
 -Check that first card is Ace of Clubs
 -Check that last card is King of Spades
*/

func TestNewDeck(t *testing.T)  {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card of Ace of Clubs, but got %s", d[0])
	}

	if d[len(d) - 1] != "King of Spades" {
		t.Errorf("Expected last card of King of Spades, but got %s", d[len(d) - 1])
	}
}

/*****Testing saveToFile and newDeckFromFile Function******/
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
  os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %d", len(loadedDeck))
	}
}