package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*ColoredLabel)(nil)

// ColoredLabel is a widget that displays a colored text.
type ColoredLabel struct {
	widget.BaseWidget
	text *canvas.Text
}

// NewColoredLabel creates a new ColoredLabel widget.
func NewColoredLabel(text string, color color.Color) *ColoredLabel {
	label := &ColoredLabel{
		text: canvas.NewText(text, color),
	}
	label.ExtendBaseWidget(label)
	label.text.Alignment = fyne.TextAlignCenter
	label.text.TextSize = 16

	return label
}

// CreateRenderer implements fyne.Widget. It creates a new renderer for the widget.
func (c *ColoredLabel) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		c.text,
	)
}

// SetText sets the text of the widget.
func (c *ColoredLabel) SetText(text string) {
	c.text.Text = text
	c.Refresh()
}
