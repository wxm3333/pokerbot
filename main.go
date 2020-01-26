package main

import (
	"fmt"
)

func main() {
	player1 := Player{"xiaomin"}
	player2 := Player{"daniel negreanu"}
	player3 := Player{"phil ivey"}

	game := NewGame()

	game.PlayerJoin(player1, 100)
	game.PlayerJoin(player2, 120)
	game.PlayerJoin(player3, 200)

	game.DealCards()

	fmt.Println(game)
	// game.PlayerLeave(0)
	// fmt.Println(game)
	// fmt.Println("hello")
}
