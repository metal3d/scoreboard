package ui

import (
	"fmt"
	"scoreboard/data"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// GamePage is a form to create a new Game.
func (a *App) GamePage() {

	var (
		form  *widget.Form
		entry *widget.Entry
	)

	label := widget.NewRichTextFromMarkdown(gameCreateText)
	label.Wrapping = fyne.TextWrapWord

	entry = widget.NewEntry()
	entry.OnChanged = func(text string) {
		text = strings.TrimSpace(text)
		if text == "" {
			return
		}
	}
	entry.Validator = func(text string) error {
		text = strings.TrimSpace(text)
		if text == "" {
			return fmt.Errorf("Game name cannot be empty")
		}
		return nil
	}
	entry.OnSubmitted = func(text string) {
		if err := form.Validate(); err != nil {
			return
		}
		form.OnSubmit()
	}

	form = widget.NewForm(
		widget.NewFormItem("Game Name", entry),
	)
	form.SubmitText = "Create"
	form.OnSubmit = func() {
		a.SetGame(data.NewGame(entry.Text))
		doCapture(a.mainWindow, "game")
		a.AddPlayersPage()
	}

	a.mainWindow.SetContent(container.NewBorder(
		label, nil, nil, nil,
		form,
	))
}

func (a *App) SetGame(game *data.Game) {
	a.currentGame = game
}
