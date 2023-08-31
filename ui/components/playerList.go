package components

import (
	"log"
	"scoreboard/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*PlayerList)(nil)

// PlayerList is a widget that displays a list of players. It displays a BgColoredLabel and buttons to
// move the player up and down in the list, change the color and delete the player.
type PlayerList struct {
	widget.BaseWidget
	players   []*data.Score
	container *fyne.Container
}

// NewPlayerList creates a new PlayerList widget.
func NewPlayerList(list []*data.Score) *PlayerList {
	pl := &PlayerList{
		players:   list,
		container: container.NewVBox(),
	}
	pl.ExtendBaseWidget(pl)

	pl.SetPlayers(list)

	return pl
}

// SwapPlayers swaps the players at index i and j.
func (pl *PlayerList) SwapPlayers(i, j int) {
	pl.players[i], pl.players[j] = pl.players[j], pl.players[i]
	log.Println("swap", i, j, "order")
	for _, p := range pl.players {
		log.Println(p.Player.Name)
	}
	pl.Refresh()
}

// SetPlayers sets the players of the widget.
func (pl *PlayerList) SetPlayers(players []*data.Score) {
	pl.players = players
	pl.container.RemoveAll()
	for _, score := range pl.Scores() {
		element := NewPlayerElement(score, pl)
		pl.container.Add(element)
	}
}

// CreateRenderer implements fyne.Widget. It creates a new renderer for the widget.
func (pl *PlayerList) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		container.NewVScroll(
			pl.container,
		),
	)
}

// Scores returns the players of the widget.
func (pl *PlayerList) Scores() []*data.Score {
	return pl.players
}
