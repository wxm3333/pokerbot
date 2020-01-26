package main

import (
	"math/rand"
	"time"
)

// A deck is a list of cards.
// There are at most 52 cards in a deck.
// Cards get dealt from the lowest index.
type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	cards := []Card{}
	for _, suit := range AllCardSuits {
		for _, value := range AllCardValues {
			card := Card{value, suit}
			cards = append(cards, card)
		}
	}
	return Deck{cards}
}

// Randomize the order of cards in the deck
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

// Remove and return the top cards from the deck
func (d *Deck) Deal(numCards int) []Card {
	out := d.Cards[:numCards]
	d.Cards = d.Cards[numCards:]
	return out
}
