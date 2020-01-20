package main

import (
	"fmt"
)

func main() {
	cards := []Card{Card{2, Diamonds}, Card{3, Diamonds}, Card{4, Clubs}}
	deck := Deck{cards}

	deck.Shuffle()

	player1 := Player{[2]Card{Card{4, Diamonds}, Card{5, Diamonds}}, 0}
	player2 := Player{[2]Card{Card{6, Diamonds}, Card{7, Diamonds}}, 1}

	flop := [3]Card{Card{8, Diamonds}, Card{9, Diamonds}, Card{10, Diamonds}}

	game := Game{deck, []Player{player1, player2}, flop, Card{11, Diamonds}, Card{12, Diamonds}}

	fmt.Println(deck)
	fmt.Println(game)
	fmt.Println("hello")
}
