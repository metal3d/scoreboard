package ui

import (
	"fmt"
	"scoreboard/data"
	. "scoreboard/i18n"
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

	label := widget.NewRichTextFromMarkdown(fmt.Sprintf(I("add.player"), minplayer))
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
				dialog.ShowError(fmt.Errorf(I("Player %s already exists"), name), a.mainWindow)
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
	entry.SetPlaceHolder(I("Player Name"))

	addPlayerButton := widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {
		addPlayer(entry.Text)
		entry.SetText("")
	})
	addPlayerButton.Disable()

	var okbutton *widget.Button
	okbutton = widget.NewButtonWithIcon(I("Start!"), theme.ConfirmIcon(), func() {
		if len(players.Scores()) < minplayer {
			okbutton.Disable()
			dialog.ShowError(fmt.Errorf(I("You need at lease %d players"), minplayer), a.mainWindow)
			return
		}
		a.currentGame.PlayerScores = players.Scores()
		a.ScorePage()
	})
	okbutton.Importance = widget.HighImportance
	okbutton.Disable()

	entry.OnChanged = func(text string) {
		if text == "" {
			addPlayerButton.Disable()
		} else {
			addPlayerButton.Enable()
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

	backbutton := widget.NewButtonWithIcon("", theme.MediaSkipPreviousIcon(), func() {
		a.GamePage()
	})

	addPlayerView := container.NewBorder(
		label, nil, nil, addPlayerButton, entry,
	)

	topView := container.NewBorder(
		nil, addPlayerView, backbutton, okbutton,
	)

	a.mainWindow.SetContent(container.NewBorder(
		topView,
		nil, nil, nil,
		players,
	))
}
