package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*BgColoredLabel)(nil)

type BgColoredLabel struct {
	widget.BaseWidget
	text *canvas.Text
	rect *canvas.Rectangle
}

func NewBGColoredLabel(text string, col color.Color) *BgColoredLabel {

	c := &BgColoredLabel{
		rect: canvas.NewRectangle(col),
		text: canvas.NewText(text, color.Black),
	}
	c.ExtendBaseWidget(c)
	c.text.Alignment = fyne.TextAlignCenter
	c.SetBgColor(col)

	return c
}

func (c *BgColoredLabel) SetBgColor(col color.Color) {
	if c.rect == nil {
		return
	}
	c.rect.FillColor = col

	// is col dark or light?
	// if dark, set text to white
	// if light, set text to black
	r, g, b, _ := col.RGBA()
	luminance := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 255
	if luminance > 128 {
		c.text.Color = color.Black
	} else {
		c.text.Color = color.White
	}
	c.text.Refresh()
	c.rect.Refresh()
}

func (c *BgColoredLabel) Resize(size fyne.Size) {
	c.rect.Resize(size)
	c.text.Resize(size)
}

func (c *BgColoredLabel) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		container.NewMax(c.rect, c.text),
	)
}

func (c *BgColoredLabel) SetText(text string) {
	c.text.Text = text
}
