package main

type Game struct {
	D       Deck
	Players []Player
	Flop    [3]Card
	Turn    Card
	River   Card
}

func (g *Game) PlayerJoin(buyin int) {
    g.Players = append(g.Players, NewPlayer(len(g.Players), buyin))
}

func (g *Game) PlayerLeave(position int) {
    g.Players = append(g.Players[:position], g.Players[position+1:]...)
}

func (g *Game) DealToPlayer(position int) {
    hand := g.D.Deal(2)
    g.Players[position].Hand = hand
}
