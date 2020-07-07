package main

func main() {
	cards := newDeck()
	// deal(cards,2)
	cards.shuffle()
	cards.print()
}
