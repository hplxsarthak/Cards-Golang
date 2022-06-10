package main

func main() {

	// cards := [] string {"Ace of spades", newCard()}

	cards := deck{"Ace of spades", newCard()}

	cards = append(cards, "Six of hearts")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
