package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*ColoredLabel)(nil)

type ColoredLabel struct {
	widget.BaseWidget
	text *canvas.Text
}

func NewColoredLabel(text string, color color.Color) *ColoredLabel {
	label := &ColoredLabel{
		text: canvas.NewText(text, color),
	}
	label.ExtendBaseWidget(label)
	label.text.Alignment = fyne.TextAlignCenter
	label.text.TextSize = 16

	return label
}

func (c *ColoredLabel) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		c.text,
	)
}

func (c *ColoredLabel) SetText(text string) {
	c.text.Text = text
	c.Refresh()
}
