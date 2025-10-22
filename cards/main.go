package main

import "fmt"

func main() {
	cards := newDeck()

	hands, remainingCards := deal(cards, 5)
	fmt.Println("Hand dealt:")
	hands.print()
	fmt.Println("Remaining cards:")
	remainingCards.print()

	// Saving the deck to a file
	cards.print()
	cards.saveToFile("my_cards")

	// Loading the deck from a file
	loadedDeck := newDeckFromFile("my_cards")
	fmt.Println("Loaded deck from file:")
	loadedDeck.print()

	fmt.Println("Shuffling the deck:")
	cards.shuffle()
	cards.print()
}
