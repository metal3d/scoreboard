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

type PlayerScore struct {
	widget.BaseWidget
	playerNameLabel  fyne.CanvasObject
	playerScoreLabel fyne.CanvasObject
	score            *data.Score
	scoreBox         *fyne.Container
	OnSubmitted      func(score float64, edited bool)
}

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

func (ps *PlayerScore) Refresh() {
	ps.playerNameLabel.Refresh()
	ps.playerScoreLabel.Refresh()
}

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

func (ps *PlayerScore) AddScoreLine() *PlayerScore {
	element := NewEditableLabel()
	element.SetOnSubmitted(ps.Submitted)
	ps.scoreBox.Add(element)
	element.SetFocus()
	return ps
}

func (ps *PlayerScore) Submitted(score float64, editing bool) {
	ps.CalculateTotalScore()
	if ps.OnSubmitted != nil {
		ps.OnSubmitted(score, editing)
	}
}

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

func (ps *PlayerScore) Score() *data.Score {
	return ps.score
}
