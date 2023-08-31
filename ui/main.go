package ui

import (
	"log"
	"math"
	"scoreboard/data"
	"scoreboard/ui/resources"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/x/fyne/theme"
)

// App represents the application.
type App struct {
	app         fyne.App
	mainWindow  fyne.Window
	currentGame *data.Game
}

// CreateApp creates a new App.
func CreateApp() *App {
	app := app.NewWithID("org.metal3d.scoreboard")
	app.SetIcon(resources.ResourceIconPng)
	app.Settings().SetTheme(theme.AdwaitaTheme())
	window := app.NewWindow("Scoreboard")
	var w, h, scale float32
	scale = 1
	w = 1024 * scale
	h = w / 16 * 9
	w = float32(math.Ceil(float64(w)))
	h = float32(math.Floor(float64(h)))
	log.Printf("w: %f, h: %f", w, h)
	window.Resize(fyne.NewSize(w, h))

	return &App{
		app:         app,
		currentGame: nil,
		mainWindow:  window,
	}
}

// Run runs the application.
func (a *App) Run() {
	a.mainWindow.ShowAndRun()
}
