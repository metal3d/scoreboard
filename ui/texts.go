package ui

import "fmt"

var (
	gameCreateText = `# Create a new game

This is the name of the game you'll play. For example, Scrabble, Boggle...
`
	addPlayersText = `# Add players

Create new players for the game. You need at least ` + fmt.Sprintf("%d", minplayer) + ` players!
`
	presentationText = `# Scoreboard

Scoreboard is a simple app to easily count points for a given game.
`

	aboutText = `# ScoreBooard - A simple Scroring Board

Scoreboard is a simple application to easily count points for a given game. It is developped in Go with [Fyne.io](https://fyne.io).

Author: [Patrice Ferlet](https://github.com/metal3d) aka Metal3d

All rights reserved. This software is distributed without any warranty.

`

	documentationText = `# Scoreboard

Simply press the "Create a new game" button to create a new game. You'll be asked for a name for the game. Then you will add players and start to add scores.

`
)
