package main

import (
	"fmt"
)

const (
	Diamonds = '\u25c6' // Karo
	Spades   = '\u2660' // Pik
	Clubs    = '\u2663' // Kreuz
	Hearts   = '\u2665' // Herz

	Six   = '\u2465'
	Seven = '\u2466'
	Eight = '\u2467'
	Nine  = '\u2468'
	Ten   = '\u2469'
	Jack  = 'J'
	Queen = 'Q'
	King  = 'K'
	Ace   = 'A'
)

func main() {
	suits := []rune{Diamonds, Spades, Clubs, Hearts}
	ranks := []rune{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	for _, rank := range ranks {
		for _, suit := range suits {
			fmt.Printf("%c%c\t", suit, rank)
		}
		fmt.Println()
	}

	fmt.Println("\nZusatzaufgabe:")

	cards := make(map[rune][]string)
	for _, suit := range suits {
		var cardList []string
		for _, rank := range ranks {
			cardList = append(cardList, fmt.Sprintf("%c%c", suit, rank))
		}
		cards[suit] = cardList
	}

	for suit, cardList := range cards {
		fmt.Printf("%c: %v\n", suit, cardList)
	}
}
