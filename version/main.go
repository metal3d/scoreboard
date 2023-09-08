package version

import (
	"fmt"

	"fyne.io/fyne/v2"
)

var appVersion = ""

func Version() string {
	if appVersion != "" {
		return appVersion
	}

	version := fyne.CurrentApp().Metadata().Version
	release := fyne.CurrentApp().Metadata().Build
	isreleased := fyne.CurrentApp().Metadata().Release
	if isreleased {
		return fmt.Sprintf("%s-%d released", version, release)
	} else {
		return fmt.Sprintf("%s-%d not released", version, release)
	}
}
