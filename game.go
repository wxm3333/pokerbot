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

func (game *Game) DealCards() {
	for _, position := range game.Table.SeatedPositions() {
		cards := game.D.Deal(2)
		game.Table.DealCardsToPosition(position, cards)
	}
}
