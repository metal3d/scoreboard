package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"scoreboard/data"
	"scoreboard/i18n"
	"scoreboard/ui"
	"strings"
)

func getLocale() (string, error) {
	envlang, ok := os.LookupEnv("LANG")
	if ok {
		return strings.Split(envlang, ".")[0], nil
	}

	cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
	output, err := cmd.Output()
	if err == nil {
		return strings.Trim(string(output), "\r\n"), nil
	}

	return "", fmt.Errorf("cannot determine locale")
}

func init() {
	// detect the os lang
	lang, err := getLocale()
	if err != nil {
		log.Println(err)
		lang = "en_US"
	}
	i18n.SetLang(lang)
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
