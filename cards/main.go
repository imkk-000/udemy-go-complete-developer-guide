package main

func main() {
	oldCards := newDeck()
	oldCards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.print()
}
