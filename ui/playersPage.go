package ui

import (
	"fmt"
	"scoreboard/data"
	"scoreboard/ui/components"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const minplayer = 2

// AddPlayersPage creates the add players page. It prepares the game to play.
// This page gives the possibility to add players to the game, and to modify their color or remove the player.
func (a *App) AddPlayersPage() {

	label := widget.NewRichTextFromMarkdown(addPlayersText)
	label.Wrapping = fyne.TextWrapWord

	players := components.NewPlayerList(a.currentGame.PlayerScores)
	addPlayer := func(name string) {
		name = strings.TrimSpace(name)
		if name == "" {
			return
		}
		// check if player already exists
		for _, player := range players.Scores() {
			if player.Player.Name == name {
				dialog.ShowError(fmt.Errorf("Player %s already exists", name), a.mainWindow)
				return
			}
		}

		player := data.NewPlayer(name)
		list := players.Scores()
		list = append(list, &data.Score{
			Player: player,
			Scores: []float32{},
		})

		players.SetPlayers(list)
		players.Refresh()
	}

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Player Name")

	button := widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {
		addPlayer(entry.Text)
		entry.SetText("")
	})
	button.Disable()

	var okbutton *widget.Button
	okbutton = widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {
		if len(players.Scores()) < minplayer {
			okbutton.Disable()
			dialog.ShowError(fmt.Errorf("You need at lease %d players", minplayer), a.mainWindow)
			return
		}
		a.currentGame.PlayerScores = players.Scores()
		a.ScorePage()
	})
	okbutton.Disable()

	entry.OnChanged = func(text string) {
		if text == "" {
			button.Disable()
		} else {
			button.Enable()
		}
	}

	entry.OnSubmitted = func(text string) {
		addPlayer(text)
		if len(players.Scores()) >= minplayer {
			okbutton.Enable()
		} else {
			okbutton.Disable()
		}
		entry.SetText("")
	}

	view := container.NewBorder(
		label, nil, nil, button, entry,
	)

	a.mainWindow.SetContent(container.NewBorder(
		view,
		okbutton, nil, nil,
		players,
	))
}
