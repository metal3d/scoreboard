package ui

import (
	"scoreboard/ui/resources"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (a *App) IndexPage() {

	gamepageButton := widget.NewButtonWithIcon("Create a new game", theme.DocumentCreateIcon(), func() {
		a.GamePage()
	})
	gamepageButton.Importance = widget.HighImportance

	aboutButton := widget.NewButtonWithIcon("About", theme.InfoIcon(), func() {
		icon := canvas.NewImageFromResource(resources.ResourceIconPng)
		icon.FillMode = canvas.ImageFillContain
		icon.SetMinSize(fyne.NewSize(64, 64))
		richtext := widget.NewRichTextFromMarkdown(aboutText)
		richtext.Wrapping = fyne.TextWrapWord

		scroll := container.NewVScroll(richtext)

		content := container.NewBorder(
			icon, nil, nil, nil,
			scroll,
		)
		d := dialog.NewCustom("About", "Close", content, a.mainWindow)
		d.Resize(fyne.NewSize(400, 400))
		d.Show()

	})
	aboutButton.Importance = widget.WarningImportance

	icon := canvas.NewImageFromResource(resources.ResourceIconPng)
	icon.FillMode = canvas.ImageFillContain
	icon.SetMinSize(fyne.NewSize(64, 64))

	a.mainWindow.SetContent(container.NewVBox(
		widget.NewLabel("Scoreboard!"),
		icon,
		widget.NewLabel("A simple scoreboard application. Selection an action:"),
		gamepageButton,
		aboutButton,
	))
}
