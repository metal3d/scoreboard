package components

import (
	"fmt"
	"scoreboard/data"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Textable interface {
	SetText(string)
}

var _ fyne.Widget = (*PlayerScore)(nil)

// PlayerScore is a widget that displays a player's name, score and a list of scores.
type PlayerScore struct {
	widget.BaseWidget
	playerNameLabel  fyne.CanvasObject
	playerScoreLabel fyne.CanvasObject
	score            *data.Score
	scoreBox         *fyne.Container
	OnSubmitted      func(score float64, edited bool)
}

// NewPlayerScore creates a new PlayerScore widget.
func NewPlayerScore(score *data.Score) *PlayerScore {
	ps := &PlayerScore{
		playerNameLabel:  NewBGColoredLabel(score.Player.Name, score.Player.Color),
		playerScoreLabel: NewColoredLabel("0", score.Player.Color),
		score:            score,
		scoreBox:         container.NewVBox(),
	}
	ps.ExtendBaseWidget(ps)

	return ps
}

// Refresh implements fyne.Widget. It refreshes the widget.
func (ps *PlayerScore) Refresh() {
	ps.playerNameLabel.Refresh()
	ps.playerScoreLabel.Refresh()
}

// CreateRenderer implements fyne.Widget. It creates a new renderer for the widget.
func (ps *PlayerScore) CreateRenderer() fyne.WidgetRenderer {
	renderer := widget.NewSimpleRenderer(
		container.NewBorder(
			ps.playerNameLabel,
			container.NewVBox(
				canvas.NewRectangle(ps.score.Player.Color),
				ps.playerScoreLabel,
			),
			nil,
			nil,
			container.NewVScroll(
				ps.scoreBox,
			),
		),
	)
	return renderer
}

// AddScoreLine adds a new score line to the widget.
func (ps *PlayerScore) AddScoreLine() *PlayerScore {
	element := NewEditableLabel()
	element.SetOnSubmitted(ps.Submitted)
	ps.scoreBox.Add(element)
	element.SetFocus()
	return ps
}

// Submitted is called when a score line is submitted.
func (ps *PlayerScore) Submitted(score float64, editing bool) {
	ps.CalculateTotalScore()
	if ps.OnSubmitted != nil {
		ps.OnSubmitted(score, editing)
	}
}

// CalculateTotalScore calculates the total score of the player.
func (ps *PlayerScore) CalculateTotalScore() {
	scores := []float32{}
	var total float32

	for _, s := range ps.scoreBox.Objects {
		v := s.(*EditableLabel).Text()

		score, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fyne.LogError("Error parsing score <"+v+">", err)
			continue
		}
		scores = append(scores, float32(score))
	}
	ps.score.Scores = scores

	for _, score := range ps.score.Scores {
		total += score
	}
	label := ps.playerScoreLabel.(Textable)
	label.SetText(fmt.Sprintf("%g", total))

	ps.Refresh()
}

// Score returns the score object of the player.
func (ps *PlayerScore) Score() *data.Score {
	return ps.score
}
