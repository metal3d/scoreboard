package components

import (
	"image/color"
	"scoreboard/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*PlayerElement)(nil)
var _ fyne.Draggable = (*PlayerElement)(nil)

// PlayerElement represents a player in the "add players" page. It is
// composed of a label, and utiltiy buttons to delete, move up or move down the
// player.
type PlayerElement struct {
	widget.BaseWidget
	score *data.Score
	list  *PlayerList

	OnDrag       func(*fyne.DragEvent)
	OnDragEnd    func()
	deleteButton *widget.Button
	upButton     *widget.Button
	downButton   *widget.Button
	colorButton  *widget.Button
	label        *BgColoredLabel
}

// NewPlayerElement creates a new PlayerElement.
func NewPlayerElement(player *data.Score, list *PlayerList) *PlayerElement {
	p := &PlayerElement{
		label: NewBGColoredLabel("template", color.RGBA{0, 0, 0, 0}),
		list:  list,
	}

	p.colorButton = widget.NewButtonWithIcon("", theme.ColorPaletteIcon(), p.changeColor)
	p.deleteButton = widget.NewButtonWithIcon("", theme.DeleteIcon(), p.delete)
	p.upButton = widget.NewButtonWithIcon("", theme.MoveUpIcon(), p.up)
	p.downButton = widget.NewButtonWithIcon("", theme.MoveDownIcon(), p.down)

	p.SetPlayer(player)

	p.ExtendBaseWidget(p)
	return p
}

// CreateRenderer creates a new renderer for the PlayerElement.
//
// Implements fyne.Widget.
func (p *PlayerElement) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		container.NewBorder(
			nil, nil, nil,
			container.NewHBox(
				p.deleteButton,
				p.upButton,
				p.downButton,
				p.colorButton,
			),
			p.label,
		),
	)
}

func (p *PlayerElement) SetPlayer(player *data.Score) {
	p.score = player
	if player == nil {
		return
	}
	p.label.SetText(player.Player.Name)
	p.label.SetBgColor(player.Player.Color)
}

func (p *PlayerElement) GetPlayer() *data.Player {
	return p.score.Player
}

func (p *PlayerElement) changeColor() {
	if len(fyne.CurrentApp().Driver().AllWindows()) == 0 {
		return
	}
	currentWindow := fyne.CurrentApp().Driver().AllWindows()[0]
	picker := dialog.NewColorPicker("Choose a color", "Choose a color", func(c color.Color) {
		p.score.Player.Color = c
		p.label.SetBgColor(c)
	}, currentWindow)
	picker.Advanced = true
	picker.SetColor(p.score.Player.Color)
	picker.Show()

}

func (p *PlayerElement) up() {
	scores := p.list.Scores()
	for i, score := range scores {
		if score == p.score {
			if i == 0 {
				return
			}
			scores[i], scores[i-1] = scores[i-1], scores[i]
			p.list.SetPlayers(scores)
			return
		}
	}
}

func (p *PlayerElement) down() {
	scores := p.list.Scores()
	for i, score := range scores {
		if score == p.score {
			if i == len(scores)-1 {
				return
			}
			scores[i], scores[i+1] = scores[i+1], scores[i]
			p.list.SetPlayers(scores)
			return
		}
	}
}

func (p *PlayerElement) delete() {
	scores := p.list.Scores()
	for i, score := range scores {
		if score.Player.Name == p.score.Player.Name {
			scores = append(scores[:i], scores[i+1:]...)
			p.list.SetPlayers(scores)
			return
		}
	}
}

func (p *PlayerElement) DragEnd() {
	scores := p.list.Scores()
	p.list.SetPlayers(scores)
	if p.OnDragEnd != nil {
		p.OnDragEnd()
		p.list.Refresh()
	}
}

func (p *PlayerElement) Dragged(event *fyne.DragEvent) {
	if p.OnDrag != nil {
		p.OnDrag(event)
	}
}
