package main

type CardValue int
type CardSuit string

const (
	Two = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

const (
	Diamonds CardSuit = "diamonds"
	Clubs    CardSuit = "clubs"
	Hearts   CardSuit = "hearts"
	Spades   CardSuit = "spades"
)

var (
	AllCardValues = [14]CardValue{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	AllCardSuits  = [4]CardSuit{Diamonds, Clubs, Hearts, Spades}
)

/*
  Value is an int between 1 and 13 inclusive.
  S is one of the 4 suits.
*/
type Card struct {
	Value CardValue
	Suit  CardSuit
}
