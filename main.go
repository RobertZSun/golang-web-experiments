package main

func main() {
	// cards := newDeckFromFile("my")
	// cards.print()

	// cards := newDeckFromFile("my_cards")

	// fmt.Println(cards.toString())
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
