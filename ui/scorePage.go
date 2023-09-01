package ui

import (
	"scoreboard/data"
	. "scoreboard/i18n"
	"scoreboard/ui/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// ScorePage creates the score page. It displays the scores of the players. All the scores are editable.
func (a *App) ScorePage() {

	var currentScore *data.Score
	elements := []fyne.CanvasObject{}

	for _, score := range a.currentGame.PlayerScores {
		element := components.NewPlayerScore(score)
		element.OnSubmitted = func(score float64, edited bool) {

			// find the next element score
			var nextScore *components.PlayerScore
			for i, element := range elements {
				if currentScore == element.(*components.PlayerScore).Score() {
					if i+1 < len(elements) {
						nextScore = elements[i+1].(*components.PlayerScore)
					} else {
						nextScore = elements[0].(*components.PlayerScore)
					}
				}
			}
			if nextScore == nil {
				return
			}
			if !edited {
				if v := nextScore.AddScoreLine(); v != nil {
					currentScore = v.Score()
				}
			}
		}
		elements = append(elements, element)
	}
	currentScore = a.currentGame.PlayerScores[0]
	elements[0].(*components.PlayerScore).AddScoreLine()

	top := topBar(a)

	a.mainWindow.SetContent(
		container.NewBorder(
			top, nil, nil, nil,
			container.NewHScroll(
				container.NewGridWithColumns(len(a.currentGame.PlayerScores),
					elements...,
				),
			),
		),
	)
}

// topBar creates the top bar of the page. It contains a button to go back to the game page and a button to reset the game.
func topBar(app *App) fyne.CanvasObject {

	resetButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		dialog.NewConfirm(I("Sure?"), I("Are you sure you want to reset the game?"), func(ok bool) {
			if !ok {
				return
			}

			for _, score := range app.currentGame.PlayerScores {
				score.Scores = []float32{}
			}
			app.ScorePage()
		}, app.mainWindow).Show()
	})

	backButton := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		dialog.NewConfirm(I("Sure?"), I("Are you sure you want to go back?"), func(ok bool) {
			if !ok {
				return
			}
			app.IndexPage()
		}, app.mainWindow).Show()
	})

	title := widget.NewLabel(app.currentGame.Name)
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}
	return container.NewBorder(
		nil, nil,
		backButton,
		resetButton,
		title,
	)
}
