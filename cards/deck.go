package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{} // initialize an empty deck - similar to []string{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits { // underscore is used to ignore the index, else it throws an unused variable error
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

func (d deck) print() { // receiver function - d is of type deck
	for i, card := range d {
		fmt.Println(i, card)
	}
}
