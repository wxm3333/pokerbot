package main

type Game struct {
	D       Deck
	Players []Player
	Flop    [3]Card
	Turn    Card
	River   Card
}
