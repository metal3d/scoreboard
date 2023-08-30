package ui

import (
	"image/jpeg"
	"os"
	"sync"

	"fyne.io/fyne/v2"
)

var capture = false
var locker = &sync.Mutex{}

func doCapture(w fyne.Window, name string) {
	if capture {
		locker.Lock()
		defer locker.Unlock()
		i := w.Canvas().Capture()
		f, _ := os.Create("captures/" + name + ".jpg")
		defer f.Close()
		jpeg.Encode(f, i, &jpeg.Options{
			Quality: 90,
		})
	}
}
