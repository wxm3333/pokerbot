package main

type Game struct {
	D              Deck
	CommunityCards []Card
	Table          GameTable
}

func NewGame() Game {
	deck := NewDeck()
	deck.Shuffle()

	communityCards := []Card{}
	table := NewGameTable()

	return Game{deck, communityCards, table}
}

func (game *Game) PlayerJoin(player Player, buyin int) error {
	emptyPosition, err := game.Table.GetOpenPosition()
	if err != nil {
		return err
	}
	game.Table.PlayerJoin(player, emptyPosition, buyin)
	return nil
}

func (game *Game) PlayerLeave(player Player) error {
	return game.Table.PlayerLeave(player)
}

func (game *Game) DealToPlayer(player Player) error {
	cards := []Card{game.D.Deal(), game.D.Deal()}
	return game.Table.DealToPlayer(player, cards)
}
