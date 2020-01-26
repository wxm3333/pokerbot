package main

type PlayerId string

type Player struct {
	Id PlayerId // MUST be unique
}

func (p *Player) Bet() {

}

func (p *Player) Fold() {

}

func (p *Player) Raise() {

}
