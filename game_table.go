package main

import "errors"

type Position int

const (
	Button = iota
	SmallBlind
	BigBlind
	UnderTheGun
	UnderTheGun1
	UnderTheGun2
	MiddlePosition
	Hijack
	Cutoff
)

var AllPositions = [9]Position{Button, SmallBlind, BigBlind, UnderTheGun, UnderTheGun1, UnderTheGun2, MiddlePosition, Hijack, Cutoff}

type GameTable struct {
	Seats  map[Position]Player
	Stacks map[PlayerId]int
	Cards  map[PlayerId][]Card
}

func NewGameTable() GameTable {
	seats := make(map[Position]Player)
	stacks := make(map[PlayerId]int)
	cards := make(map[PlayerId][]Card)
	return GameTable{seats, stacks, cards}
}

func (table *GameTable) PlayerJoin(player Player, seat Position, buyin int) error {
	if _, exists := table.Seats[seat]; exists {
		return errors.New("A player is already seated there!")
	} else {
		table.Seats[seat] = player
		table.Stacks[player.Id] = buyin
		return nil
	}
}

func (table *GameTable) PlayerLeave(player Player) error {
	if _, exists := table.Stacks[player.Id]; !exists {
		return errors.New("The player is not seated!")
	}
	for position, seatedPlayer := range table.Seats {
		if seatedPlayer.Id == player.Id {
			delete(table.Seats, position)
		}
	}
	delete(table.Stacks, player.Id)
	delete(table.Cards, player.Id)
	return nil
}

func (table GameTable) GetOpenPosition() (Position, error) {
	positionsToSeat := []Position{
		Button,
		SmallBlind,
		BigBlind,
		UnderTheGun,
		MiddlePosition,
		Cutoff,
		Hijack,
		UnderTheGun1,
		UnderTheGun2,
	}
	for _, position := range positionsToSeat {
		if _, exists := table.Seats[position]; !exists {
			return position, nil
		}
	}
	return Button, errors.New("There is no open seat!")
}

func (table *GameTable) DealToPlayer(player Player, cards []Card) error {
	if _, exists := table.Stacks[player.Id]; !exists {
		return errors.New("Player is not seated at the table!")
	}
	table.Cards[player.Id] = cards
	return nil
}
