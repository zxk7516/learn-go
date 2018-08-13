package main

func main() {
	// cards := []Card{Card{"red", 1}}
	cards := newDeck()
	cards.shuffle()
	hand, cards := deal(cards, 5)

	hand.print()
	hand.saveToFile("hand_cards")
	hadnFromFile := newDeckFromFile("hand_cards")
	hadnFromFile.print()
}
