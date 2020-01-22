package main

// Position 0 is dealer.
type Player struct {
	Hand     []Card
	Position int
    Stack int
}

func NewPlayer (position int, buyin int) Player {
    return Player{[]Card{}, position, buyin}
 }

func (p *Player) Bet() {

}

func (p *Player) Fold() {

}

func (p *Player) Raise() {

}
