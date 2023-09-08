package main

import (
	"fmt"
	"os"
	"os/exec"
	"scoreboard/data"
	"scoreboard/i18n"
	"scoreboard/ui"
	"strings"

	"fyne.io/fyne/v2"
)

func getLocale() (string, error) {
	envlang, ok := os.LookupEnv("LANG")
	if ok {
		lang := strings.Split(envlang, ".")[0]
		if len(lang) == 0 {
			return "", fmt.Errorf("cannot determine locale")
		}
		return lang, nil
	}

	cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
	output, err := cmd.Output()
	if err == nil {
		lang := strings.Trim(string(output), "\r\n ")
		if len(lang) == 0 {
			return "", fmt.Errorf("cannot determine locale")
		}
		return lang, nil
	}
	return "", fmt.Errorf("cannot determine locale")
}

func init() {
	// detect the os lang
	lang, err := getLocale()
	if err != nil {
		fyne.LogError("Cannot determine lang", err)
	} else {
		i18n.SetLang(lang)
	}
}

func main() {
	app := ui.CreateApp()

	if os.Getenv("MODE") == "test" {
		// mode test, add a game, 3 players
		game := data.NewGame("Scrabble")
		game.AddPlayers(
			data.NewPlayer("Alice"),
			data.NewPlayer("Bob"),
			data.NewPlayer("Carol"),
			data.NewPlayer("Dave"),
			data.NewPlayer("Eve"),
		)

		app.SetGame(game)
		app.ScorePage()
	} else {
		app.IndexPage()
	}
	app.Run()
}
