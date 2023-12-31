package ui

import (
	"fmt"
	"scoreboard/ui/resources"
	"scoreboard/version"

	. "scoreboard/i18n"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// IndexPage is the main page of the application
func (a *App) IndexPage() {

	gamepageButton := widget.NewButtonWithIcon(I("Create a new game"), theme.DocumentCreateIcon(), func() {
		a.GamePage()
	})
	gamepageButton.Importance = widget.HighImportance

	aboutButton := widget.NewButtonWithIcon(I("About"), theme.InfoIcon(), func() {
		icon := canvas.NewImageFromResource(resources.ResourceIconPng)
		icon.FillMode = canvas.ImageFillContain
		icon.SetMinSize(fyne.NewSize(64, 64))
		about := fmt.Sprintf(I("about.content"), version.Version())
		richtext := widget.NewRichTextFromMarkdown(about)
		richtext.Wrapping = fyne.TextWrapWord

		scroll := container.NewVScroll(richtext)

		content := container.NewBorder(
			icon, nil, nil, nil,
			scroll,
		)
		d := dialog.NewCustom(I("About"), I("Close"), content, a.mainWindow)
		d.Resize(fyne.NewSize(400, 400))
		d.Show()

	})
	aboutButton.Importance = widget.WarningImportance

	icon := canvas.NewImageFromResource(resources.ResourceIconPng)
	icon.FillMode = canvas.ImageFillContain
	icon.SetMinSize(fyne.NewSize(64, 64))

	doc := widget.NewRichTextFromMarkdown(I("documentation.text"))
	doc.Wrapping = fyne.TextWrapWord

	a.mainWindow.SetTitle(I("Scoreboard"))
	a.mainWindow.SetContent(container.NewVBox(
		icon,
		doc,
		gamepageButton,
		aboutButton,
	))
}
