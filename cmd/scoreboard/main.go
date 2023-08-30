package main

import (
	"scoreboard/data"
	"scoreboard/ui"
	"os"
)

func main() {
	app := ui.CreateApp()

	if os.Getenv("MODE") == "test" {
		// mode test, add a game, 3 players
		game := data.NewGame("Scrabble")
		game.AddPlayers(
			data.NewPlayer("Alice"),
			data.NewPlayer("Bob"),
			data.NewPlayer("Carol"),
			data.NewPlayer("Dave"),
			data.NewPlayer("Eve"),
		)

		app.SetGame(game)
		app.ScorePage()
	} else {
		app.IndexPage()
	}
	app.Run()
}
