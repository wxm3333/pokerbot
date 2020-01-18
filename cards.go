package main

type Suit string
const (
  Diamonds Suit = "diamonds"
  Clubs Suit = "clubs"
  Hearts Suit = "hearts"
  Spades Suit = "spades"
)

/*
  Value is an int between 1 and 13 inclusive.
  S is one of the 4 suits.
*/
type Card struct {
    Value int
    S Suit
}

/*
    A deck is a list of cards.
    There are at most 52 cards in a deck.
    Cards get dealt from the lowest index.
    */
type Deck struct {
    Cards []Card
}

// Randomize the order of cards in the deck
func (d *Deck) Shuffle() {
  // TODO
}

// Remove and return the top numCards cards from d.Cards
func (d *Deck) Deal(numCards int) []Card {
    out := d.Cards[:numCards]
    d.Cards = d.Cards[numCards:]
    return out
}
