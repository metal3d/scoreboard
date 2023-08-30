package data

import (
	"image/color"
	"math/rand"
)

// Player represents a player in a game. It contains a name and a color.
type Player struct {
	Name  string
	Color color.Color
}

// NewPlayer creates a new player with the given name and a random color.
func NewPlayer(name string) *Player {
	randColor := func() color.Color {

		// generate a random paster color
		col := color.NRGBA{
			R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)),
			B: uint8(rand.Intn(255)),
			A: 190,
		}

		return col

	}

	return &Player{
		Name:  name,
		Color: randColor(),
	}
}
